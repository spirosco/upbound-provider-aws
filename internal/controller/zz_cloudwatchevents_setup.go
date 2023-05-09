/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	apidestination "github.com/spirosco/upbound-provider-aws/internal/controller/cloudwatchevents/apidestination"
	archive "github.com/spirosco/upbound-provider-aws/internal/controller/cloudwatchevents/archive"
	bus "github.com/spirosco/upbound-provider-aws/internal/controller/cloudwatchevents/bus"
	buspolicy "github.com/spirosco/upbound-provider-aws/internal/controller/cloudwatchevents/buspolicy"
	connection "github.com/spirosco/upbound-provider-aws/internal/controller/cloudwatchevents/connection"
	permission "github.com/spirosco/upbound-provider-aws/internal/controller/cloudwatchevents/permission"
	rule "github.com/spirosco/upbound-provider-aws/internal/controller/cloudwatchevents/rule"
	target "github.com/spirosco/upbound-provider-aws/internal/controller/cloudwatchevents/target"
)

// Setup_cloudwatchevents creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_cloudwatchevents(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		apidestination.Setup,
		archive.Setup,
		bus.Setup,
		buspolicy.Setup,
		connection.Setup,
		permission.Setup,
		rule.Setup,
		target.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
