/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	directoryconfig "github.com/spirosco/upbound-provider-aws/internal/controller/appstream/directoryconfig"
	fleet "github.com/spirosco/upbound-provider-aws/internal/controller/appstream/fleet"
	fleetstackassociation "github.com/spirosco/upbound-provider-aws/internal/controller/appstream/fleetstackassociation"
	imagebuilder "github.com/spirosco/upbound-provider-aws/internal/controller/appstream/imagebuilder"
	stack "github.com/spirosco/upbound-provider-aws/internal/controller/appstream/stack"
	user "github.com/spirosco/upbound-provider-aws/internal/controller/appstream/user"
	userstackassociation "github.com/spirosco/upbound-provider-aws/internal/controller/appstream/userstackassociation"
)

// Setup_appstream creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_appstream(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		directoryconfig.Setup,
		fleet.Setup,
		fleetstackassociation.Setup,
		imagebuilder.Setup,
		stack.Setup,
		user.Setup,
		userstackassociation.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
