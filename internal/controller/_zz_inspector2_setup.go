/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	enabler "github.com/spirosco/upbound-provider-aws/internal/controller/inspector2/enabler"
)

// Setup_inspector2 creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_inspector2(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		enabler.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
