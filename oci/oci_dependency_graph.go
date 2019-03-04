// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

package provider

var DependencyGraph map[string][]string

func initDependencyGraph() {
	DependencyGraph = make(map[string][]string)

	DependencyGraph["asset"] = append(DependencyGraph["asset"], "CoreVolumeBackupPolicyAssignment")
	DependencyGraph["autonomousDataWarehouse"] = append(DependencyGraph["autonomousDataWarehouse"], "DatabaseAutonomousDataWarehouseBackup")
	DependencyGraph["autonomousDataWarehouse"] = append(DependencyGraph["autonomousDataWarehouse"], "DatabaseAutonomousDataWarehouseWallet")
	DependencyGraph["autonomousDatabase"] = append(DependencyGraph["autonomousDatabase"], "DatabaseAutonomousDatabaseBackup")
	DependencyGraph["autonomousDatabase"] = append(DependencyGraph["autonomousDatabase"], "DatabaseAutonomousDatabaseWallet")
	DependencyGraph["backupPolicy"] = append(DependencyGraph["backupPolicy"], "CoreBootVolume")
	DependencyGraph["backupPolicy"] = append(DependencyGraph["backupPolicy"], "CoreVolume")
	DependencyGraph["backupSubnet"] = append(DependencyGraph["backupSubnet"], "DatabaseDbSystem")
	DependencyGraph["bootVolume"] = append(DependencyGraph["bootVolume"], "CoreBootVolumeAttachment")
	DependencyGraph["bootVolume"] = append(DependencyGraph["bootVolume"], "CoreBootVolumeBackup")
	DependencyGraph["cluster"] = append(DependencyGraph["cluster"], "ContainerengineClusterKubeConfig")
	DependencyGraph["cluster"] = append(DependencyGraph["cluster"], "ContainerengineNodePool")
	DependencyGraph["cpe"] = append(DependencyGraph["cpe"], "CoreIpSecConnection")
	DependencyGraph["crossConnectGroup"] = append(DependencyGraph["crossConnectGroup"], "CoreCrossConnect")
	DependencyGraph["database"] = append(DependencyGraph["database"], "DatabaseBackup")
	DependencyGraph["dbSystem"] = append(DependencyGraph["dbSystem"], "DatabaseDbHome")
	DependencyGraph["dhcpOptions"] = append(DependencyGraph["dhcpOptions"], "CoreSubnet")
	DependencyGraph["drg"] = append(DependencyGraph["drg"], "CoreDrgAttachment")
	DependencyGraph["drg"] = append(DependencyGraph["drg"], "CoreIpSecConnection")
	DependencyGraph["drg"] = append(DependencyGraph["drg"], "CoreRemotePeeringConnection")
	DependencyGraph["exportSet"] = append(DependencyGraph["exportSet"], "FileStorageExport")
	DependencyGraph["farCrossConnectOrCrossConnectGroup"] = append(DependencyGraph["farCrossConnectOrCrossConnectGroup"], "CoreCrossConnect")
	DependencyGraph["fileSystem"] = append(DependencyGraph["fileSystem"], "FileStorageExport")
	DependencyGraph["fileSystem"] = append(DependencyGraph["fileSystem"], "FileStorageSnapshot")
	DependencyGraph["gateway"] = append(DependencyGraph["gateway"], "CoreVirtualCircuit")
	DependencyGraph["group"] = append(DependencyGraph["group"], "IdentityIdpGroupMapping")
	DependencyGraph["group"] = append(DependencyGraph["group"], "IdentityUserGroupMembership")
	DependencyGraph["healthCheckMonitor"] = append(DependencyGraph["healthCheckMonitor"], "DnsSteeringPolicy")
	DependencyGraph["identityProvider"] = append(DependencyGraph["identityProvider"], "IdentityIdpGroupMapping")
	DependencyGraph["instance"] = append(DependencyGraph["instance"], "CoreBootVolumeAttachment")
	DependencyGraph["instance"] = append(DependencyGraph["instance"], "CoreConsoleHistory")
	DependencyGraph["instance"] = append(DependencyGraph["instance"], "CoreInstanceConsoleConnection")
	DependencyGraph["instance"] = append(DependencyGraph["instance"], "CoreVolumeAttachment")
	DependencyGraph["instanceConfiguration"] = append(DependencyGraph["instanceConfiguration"], "CoreInstancePool")
	DependencyGraph["internetGateway"] = append(DependencyGraph["internetGateway"], "CoreRouteTable")
	DependencyGraph["key"] = append(DependencyGraph["key"], "KmsDecryptedData")
	DependencyGraph["key"] = append(DependencyGraph["key"], "KmsEncryptedData")
	DependencyGraph["key"] = append(DependencyGraph["key"], "KmsGeneratedKey")
	DependencyGraph["key"] = append(DependencyGraph["key"], "KmsKeyVersion")
	DependencyGraph["kmsKey"] = append(DependencyGraph["kmsKey"], "CoreBootVolume")
	DependencyGraph["kmsKey"] = append(DependencyGraph["kmsKey"], "CoreVolume")
	DependencyGraph["kmsKey"] = append(DependencyGraph["kmsKey"], "ObjectStorageBucket")
	DependencyGraph["listing"] = append(DependencyGraph["listing"], "CoreAppCatalogSubscription")
	DependencyGraph["loadBalancer"] = append(DependencyGraph["loadBalancer"], "LoadBalancerBackend")
	DependencyGraph["loadBalancer"] = append(DependencyGraph["loadBalancer"], "LoadBalancerBackendSet")
	DependencyGraph["loadBalancer"] = append(DependencyGraph["loadBalancer"], "LoadBalancerCertificate")
	DependencyGraph["loadBalancer"] = append(DependencyGraph["loadBalancer"], "LoadBalancerHostname")
	DependencyGraph["loadBalancer"] = append(DependencyGraph["loadBalancer"], "LoadBalancerListener")
	DependencyGraph["loadBalancer"] = append(DependencyGraph["loadBalancer"], "LoadBalancerPathRouteSet")
	DependencyGraph["loadBalancer"] = append(DependencyGraph["loadBalancer"], "LoadBalancerRuleSet")
	DependencyGraph["metricCompartment"] = append(DependencyGraph["metricCompartment"], "MonitoringAlarm")
	DependencyGraph["nearCrossConnectOrCrossConnectGroup"] = append(DependencyGraph["nearCrossConnectOrCrossConnectGroup"], "CoreCrossConnect")
	DependencyGraph["policy"] = append(DependencyGraph["policy"], "CoreVolumeBackupPolicyAssignment")
	DependencyGraph["privateIp"] = append(DependencyGraph["privateIp"], "CorePublicIp")
	DependencyGraph["providerService"] = append(DependencyGraph["providerService"], "CoreVirtualCircuit")
	DependencyGraph["routeTable"] = append(DependencyGraph["routeTable"], "CoreDrgAttachment")
	DependencyGraph["routeTable"] = append(DependencyGraph["routeTable"], "CoreLocalPeeringGateway")
	DependencyGraph["routeTable"] = append(DependencyGraph["routeTable"], "CoreSubnet")
	DependencyGraph["steeringPolicy"] = append(DependencyGraph["steeringPolicy"], "DnsSteeringPolicyAttachment")
	DependencyGraph["subnet"] = append(DependencyGraph["subnet"], "CoreInstance")
	DependencyGraph["tagNamespace"] = append(DependencyGraph["tagNamespace"], "IdentityTag")
	DependencyGraph["tenancy"] = append(DependencyGraph["tenancy"], "IdentityRegionSubscription")
	DependencyGraph["topic"] = append(DependencyGraph["topic"], "OnsSubscription")
	DependencyGraph["user"] = append(DependencyGraph["user"], "IdentityApiKey")
	DependencyGraph["user"] = append(DependencyGraph["user"], "IdentityAuthToken")
	DependencyGraph["user"] = append(DependencyGraph["user"], "IdentityCustomerSecretKey")
	DependencyGraph["user"] = append(DependencyGraph["user"], "IdentitySmtpCredential")
	DependencyGraph["user"] = append(DependencyGraph["user"], "IdentitySwiftPassword")
	DependencyGraph["user"] = append(DependencyGraph["user"], "IdentityUiPassword")
	DependencyGraph["user"] = append(DependencyGraph["user"], "IdentityUserGroupMembership")
	DependencyGraph["vcn"] = append(DependencyGraph["vcn"], "ContainerengineCluster")
	DependencyGraph["vcn"] = append(DependencyGraph["vcn"], "CoreDhcpOptions")
	DependencyGraph["vcn"] = append(DependencyGraph["vcn"], "CoreDrgAttachment")
	DependencyGraph["vcn"] = append(DependencyGraph["vcn"], "CoreInternetGateway")
	DependencyGraph["vcn"] = append(DependencyGraph["vcn"], "CoreLocalPeeringGateway")
	DependencyGraph["vcn"] = append(DependencyGraph["vcn"], "CoreNatGateway")
	DependencyGraph["vcn"] = append(DependencyGraph["vcn"], "CoreRouteTable")
	DependencyGraph["vcn"] = append(DependencyGraph["vcn"], "CoreSecurityList")
	DependencyGraph["vcn"] = append(DependencyGraph["vcn"], "CoreServiceGateway")
	DependencyGraph["vcn"] = append(DependencyGraph["vcn"], "CoreSubnet")
	DependencyGraph["vnic"] = append(DependencyGraph["vnic"], "CorePrivateIp")
	DependencyGraph["volume"] = append(DependencyGraph["volume"], "CoreVolumeAttachment")
	DependencyGraph["volume"] = append(DependencyGraph["volume"], "CoreVolumeBackup")
	DependencyGraph["volumeGroup"] = append(DependencyGraph["volumeGroup"], "CoreVolumeGroupBackup")
	DependencyGraph["zone"] = append(DependencyGraph["zone"], "DnsSteeringPolicyAttachment")
}
