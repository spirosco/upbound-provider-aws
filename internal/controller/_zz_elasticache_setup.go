/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	cluster "github.com/spirosco/upbound-provider-aws/internal/controller/elasticache/cluster"
	parametergroup "github.com/spirosco/upbound-provider-aws/internal/controller/elasticache/parametergroup"
	replicationgroup "github.com/spirosco/upbound-provider-aws/internal/controller/elasticache/replicationgroup"
	subnetgroup "github.com/spirosco/upbound-provider-aws/internal/controller/elasticache/subnetgroup"
	user "github.com/spirosco/upbound-provider-aws/internal/controller/elasticache/user"
	usergroup "github.com/spirosco/upbound-provider-aws/internal/controller/elasticache/usergroup"
)

// Setup_elasticache creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_elasticache(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		cluster.Setup,
		parametergroup.Setup,
		replicationgroup.Setup,
		subnetgroup.Setup,
		user.Setup,
		usergroup.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
