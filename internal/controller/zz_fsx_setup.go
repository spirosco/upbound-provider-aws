/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	backup "github.com/spirosco/upbound-provider-aws/internal/controller/fsx/backup"
	datarepositoryassociation "github.com/spirosco/upbound-provider-aws/internal/controller/fsx/datarepositoryassociation"
	lustrefilesystem "github.com/spirosco/upbound-provider-aws/internal/controller/fsx/lustrefilesystem"
	ontapfilesystem "github.com/spirosco/upbound-provider-aws/internal/controller/fsx/ontapfilesystem"
	ontapstoragevirtualmachine "github.com/spirosco/upbound-provider-aws/internal/controller/fsx/ontapstoragevirtualmachine"
	windowsfilesystem "github.com/spirosco/upbound-provider-aws/internal/controller/fsx/windowsfilesystem"
)

// Setup_fsx creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_fsx(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		backup.Setup,
		datarepositoryassociation.Setup,
		lustrefilesystem.Setup,
		ontapfilesystem.Setup,
		ontapstoragevirtualmachine.Setup,
		windowsfilesystem.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
