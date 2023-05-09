/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	cluster "github.com/spirosco/upbound-provider-aws/internal/controller/rds/cluster"
	clusteractivitystream "github.com/spirosco/upbound-provider-aws/internal/controller/rds/clusteractivitystream"
	clusterendpoint "github.com/spirosco/upbound-provider-aws/internal/controller/rds/clusterendpoint"
	clusterinstance "github.com/spirosco/upbound-provider-aws/internal/controller/rds/clusterinstance"
	clusterparametergroup "github.com/spirosco/upbound-provider-aws/internal/controller/rds/clusterparametergroup"
	clusterroleassociation "github.com/spirosco/upbound-provider-aws/internal/controller/rds/clusterroleassociation"
	clustersnapshot "github.com/spirosco/upbound-provider-aws/internal/controller/rds/clustersnapshot"
	dbinstanceautomatedbackupsreplication "github.com/spirosco/upbound-provider-aws/internal/controller/rds/dbinstanceautomatedbackupsreplication"
	dbsnapshotcopy "github.com/spirosco/upbound-provider-aws/internal/controller/rds/dbsnapshotcopy"
	eventsubscription "github.com/spirosco/upbound-provider-aws/internal/controller/rds/eventsubscription"
	globalcluster "github.com/spirosco/upbound-provider-aws/internal/controller/rds/globalcluster"
	instance "github.com/spirosco/upbound-provider-aws/internal/controller/rds/instance"
	instanceroleassociation "github.com/spirosco/upbound-provider-aws/internal/controller/rds/instanceroleassociation"
	optiongroup "github.com/spirosco/upbound-provider-aws/internal/controller/rds/optiongroup"
	parametergroup "github.com/spirosco/upbound-provider-aws/internal/controller/rds/parametergroup"
	proxy "github.com/spirosco/upbound-provider-aws/internal/controller/rds/proxy"
	proxydefaulttargetgroup "github.com/spirosco/upbound-provider-aws/internal/controller/rds/proxydefaulttargetgroup"
	proxyendpoint "github.com/spirosco/upbound-provider-aws/internal/controller/rds/proxyendpoint"
	proxytarget "github.com/spirosco/upbound-provider-aws/internal/controller/rds/proxytarget"
	securitygroup "github.com/spirosco/upbound-provider-aws/internal/controller/rds/securitygroup"
	snapshot "github.com/spirosco/upbound-provider-aws/internal/controller/rds/snapshot"
	subnetgroup "github.com/spirosco/upbound-provider-aws/internal/controller/rds/subnetgroup"
)

// Setup_rds creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_rds(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		cluster.Setup,
		clusteractivitystream.Setup,
		clusterendpoint.Setup,
		clusterinstance.Setup,
		clusterparametergroup.Setup,
		clusterroleassociation.Setup,
		clustersnapshot.Setup,
		dbinstanceautomatedbackupsreplication.Setup,
		dbsnapshotcopy.Setup,
		eventsubscription.Setup,
		globalcluster.Setup,
		instance.Setup,
		instanceroleassociation.Setup,
		optiongroup.Setup,
		parametergroup.Setup,
		proxy.Setup,
		proxydefaulttargetgroup.Setup,
		proxyendpoint.Setup,
		proxytarget.Setup,
		securitygroup.Setup,
		snapshot.Setup,
		subnetgroup.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
