/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	directory "github.com/spirosco/upbound-provider-aws/internal/controller/workspaces/directory"
	ipgroup "github.com/spirosco/upbound-provider-aws/internal/controller/workspaces/ipgroup"
)

// Setup_workspaces creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_workspaces(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		directory.Setup,
		ipgroup.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
