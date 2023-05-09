/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	datalakesettings "github.com/spirosco/upbound-provider-aws/internal/controller/lakeformation/datalakesettings"
	permissions "github.com/spirosco/upbound-provider-aws/internal/controller/lakeformation/permissions"
	resource "github.com/spirosco/upbound-provider-aws/internal/controller/lakeformation/resource"
)

// Setup_lakeformation creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_lakeformation(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		datalakesettings.Setup,
		permissions.Setup,
		resource.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
