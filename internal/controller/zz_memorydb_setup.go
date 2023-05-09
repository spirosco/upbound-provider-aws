/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	acl "github.com/spirosco/upbound-provider-aws/internal/controller/memorydb/acl"
	cluster "github.com/spirosco/upbound-provider-aws/internal/controller/memorydb/cluster"
	parametergroup "github.com/spirosco/upbound-provider-aws/internal/controller/memorydb/parametergroup"
	snapshot "github.com/spirosco/upbound-provider-aws/internal/controller/memorydb/snapshot"
	subnetgroup "github.com/spirosco/upbound-provider-aws/internal/controller/memorydb/subnetgroup"
)

// Setup_memorydb creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_memorydb(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		acl.Setup,
		cluster.Setup,
		parametergroup.Setup,
		snapshot.Setup,
		subnetgroup.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
