package elasticsearch

import (
	"sort"
	"strings"
	"testing"

	"golang.org/x/crypto/bcrypt"

	deploymentsv1alpha1 "github.com/elastic/stack-operators/stack-operator/pkg/apis/deployments/v1alpha1"
	"github.com/elastic/stack-operators/stack-operator/pkg/controller/stack/elasticsearch/client"
	"github.com/stretchr/testify/assert"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var (
	mockStack = deploymentsv1alpha1.Stack{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "my-stack",
			Namespace: "default",
		}}
	testUser = []client.User{client.User{Name: "foo", Password: "bar"}}
)

func TestNewUserSecrets(t *testing.T) {

	elasticUsers, err := NewElasticUsersCredentials(mockStack, testUser)
	assert.NoError(t, err)

	tests := []struct {
		subject      UserCredentials
		expectedName string
		expectedKeys []string
	}{
		{
			subject:      NewInternalUserCredentials(mockStack),
			expectedName: "my-stack-internal-users",
			expectedKeys: []string{InternalControllerUserName, InternalKibanaServerUserName},
		},
		{
			subject:      NewExternalUserCredentials(mockStack),
			expectedName: "my-stack-elastic-user",
			expectedKeys: []string{ExternalUserName},
		},
		{
			subject:      elasticUsers,
			expectedName: "my-stack-users",
			expectedKeys: []string{ElasticUsersFile, ElasticUsersRolesFile},
		},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.expectedName, tt.subject.Secret().Name)
		var keys []string
		for k := range tt.subject.Secret().Data {
			keys = append(keys, k)
		}
		sort.Strings(keys)
		assert.EqualValues(t, tt.expectedKeys, keys)
	}

}

func TestNewElasticUsersSecret(t *testing.T) {
	creds, err := NewElasticUsersCredentials(mockStack, testUser)
	assert.NoError(t, err)
	assert.Equal(t, "superuser:foo", string(creds.Secret().Data[ElasticUsersRolesFile]))

	for _, user := range strings.Split(string(creds.Secret().Data[ElasticUsersFile]), "\n") {
		userPw := strings.Split(user, ":")
		assert.Equal(t, "foo", userPw[0])
		assert.NoError(t, bcrypt.CompareHashAndPassword([]byte(userPw[1]), []byte("bar")))
	}

}

func newTestCredentials(t *testing.T, users []client.User) UserCredentials {
	creds, err := NewElasticUsersCredentials(mockStack, users)
	assert.NoError(t, err)
	return creds

}

func TestNeedsUpdate(t *testing.T) {

	otherUser := client.User{Name: "baz", Password: "secret"}

	tests := []struct {
		desc        string
		subject1    UserCredentials
		subject2    UserCredentials
		needsUpdate bool
	}{
		{
			desc:        "internal clear text creds don't need update even if they contain different passwords (secret is source of truth)",
			subject1:    NewInternalUserCredentials(mockStack),
			subject2:    NewInternalUserCredentials(mockStack),
			needsUpdate: false,
		},
		{
			desc:        "external clear text creds don't need update even if they contain different passwords (secret is source of truth)",
			subject1:    NewExternalUserCredentials(mockStack),
			subject2:    NewExternalUserCredentials(mockStack),
			needsUpdate: false,
		},
		{
			desc:        "hashed creds: different hash but same password does not warrant an update of the secret",
			subject1:    newTestCredentials(t, testUser),
			subject2:    newTestCredentials(t, testUser),
			needsUpdate: false,
		},
		{
			desc:        "hashed creds: changed password warrants an update of the secret",
			subject1:    newTestCredentials(t, testUser),
			subject2:    newTestCredentials(t, []client.User{client.User{Name: "foo", Password: "changed!"}}),
			needsUpdate: true,
		},
		{
			desc:        "hashed creds: order of user credentials should not matter",
			subject1:    newTestCredentials(t, []client.User{testUser[0], otherUser}),
			subject2:    newTestCredentials(t, []client.User{otherUser, testUser[0]}),
			needsUpdate: false,
		},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.needsUpdate, tt.subject1.NeedsUpdate(tt.subject2.Secret()), tt.desc)
	}
}