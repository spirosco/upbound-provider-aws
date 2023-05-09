/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	accountsettingdefault "github.com/spirosco/upbound-provider-aws/internal/controller/ecs/accountsettingdefault"
	capacityprovider "github.com/spirosco/upbound-provider-aws/internal/controller/ecs/capacityprovider"
	cluster "github.com/spirosco/upbound-provider-aws/internal/controller/ecs/cluster"
	clustercapacityproviders "github.com/spirosco/upbound-provider-aws/internal/controller/ecs/clustercapacityproviders"
	service "github.com/spirosco/upbound-provider-aws/internal/controller/ecs/service"
	taskdefinition "github.com/spirosco/upbound-provider-aws/internal/controller/ecs/taskdefinition"
)

// Setup_ecs creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_ecs(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		accountsettingdefault.Setup,
		capacityprovider.Setup,
		cluster.Setup,
		clustercapacityproviders.Setup,
		service.Setup,
		taskdefinition.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
