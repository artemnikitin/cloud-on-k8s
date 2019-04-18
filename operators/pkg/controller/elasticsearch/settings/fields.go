// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

package settings

const (
	ClusterName = "cluster.name"

	DiscoveryZenMinimumMasterNodes = "discovery.zen.minimum_master_nodes"
	DiscoveryZenPingUnicastHosts   = "discovery.zen.ping.unicast.hosts"
	ClusterInitialMasterNodes      = "cluster.initial_master_nodes"

	NetworkHost        = "network.host"
	NetworkPublishHost = "network.publish_host"

	NodeName = "node.name"

	PathData = "path.data"
	PathLogs = "path.logs"

	XPackLicenseSelfGeneratedType                   = "xpack.license.self_generated.type"
	XPackSecurityAuthcReservedRealmEnabled          = "xpack.security.authc.reserved_realm.enabled"
	XPackSecurityEnabled                            = "xpack.security.enabled"
	XPackSecurityHttpSslCertificate                 = "xpack.security.http.ssl.certificate"
	XPackSecurityHttpSslCertificateAuthorities      = "xpack.security.http.ssl.certificate_authorities"
	XPackSecurityHttpSslEnabled                     = "xpack.security.http.ssl.enabled"
	XPackSecurityHttpSslKey                         = "xpack.security.http.ssl.key"
	XPackSecurityTransportSslCertificate            = "xpack.security.transport.ssl.certificate"
	XPackSecurityTransportSslCertificateAuthorities = "xpack.security.transport.ssl.certificate_authorities"
	XPackSecurityTransportSslEnabled                = "xpack.security.transport.ssl.enabled"
	XPackSecurityTransportSslKey                    = "xpack.security.transport.ssl.key"
	XPackSecurityTransportSslTrustRestrictionsPath  = "xpack.security.transport.ssl.trust_restrictions.path"
	XPackSecurityTransportSslVerificationMode       = "xpack.security.transport.ssl.verification_mode"
)

var Blacklist = []string{
	ClusterName,
	DiscoveryZenMinimumMasterNodes,
	DiscoveryZenPingUnicastHosts,
	ClusterInitialMasterNodes,
	NetworkHost,
	NetworkPublishHost,
	NodeName,
	PathData,
	PathLogs,
	XPackLicenseSelfGeneratedType,
	XPackSecurityAuthcReservedRealmEnabled,
	XPackSecurityEnabled,
	XPackSecurityHttpSslCertificate,
	XPackSecurityHttpSslCertificateAuthorities,
	XPackSecurityHttpSslEnabled,
	XPackSecurityHttpSslKey,
	XPackSecurityTransportSslCertificate,
	XPackSecurityTransportSslCertificateAuthorities,
	XPackSecurityTransportSslEnabled,
	XPackSecurityTransportSslKey,
	XPackSecurityTransportSslTrustRestrictionsPath,
	XPackSecurityTransportSslVerificationMode,
}