/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	devicepool "github.com/spirosco/upbound-provider-aws/internal/controller/devicefarm/devicepool"
	instanceprofile "github.com/spirosco/upbound-provider-aws/internal/controller/devicefarm/instanceprofile"
	networkprofile "github.com/spirosco/upbound-provider-aws/internal/controller/devicefarm/networkprofile"
	project "github.com/spirosco/upbound-provider-aws/internal/controller/devicefarm/project"
	testgridproject "github.com/spirosco/upbound-provider-aws/internal/controller/devicefarm/testgridproject"
	upload "github.com/spirosco/upbound-provider-aws/internal/controller/devicefarm/upload"
)

// Setup_devicefarm creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_devicefarm(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		devicepool.Setup,
		instanceprofile.Setup,
		networkprofile.Setup,
		project.Setup,
		testgridproject.Setup,
		upload.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
