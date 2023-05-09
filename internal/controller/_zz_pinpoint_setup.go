/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	app "github.com/spirosco/upbound-provider-aws/internal/controller/pinpoint/app"
	smschannel "github.com/spirosco/upbound-provider-aws/internal/controller/pinpoint/smschannel"
)

// Setup_pinpoint creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_pinpoint(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		app.Setup,
		smschannel.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
