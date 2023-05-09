/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	experimenttemplate "github.com/spirosco/upbound-provider-aws/internal/controller/fis/experimenttemplate"
)

// Setup_fis creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_fis(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		experimenttemplate.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
