/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	definition "github.com/spirosco/upbound-provider-aws/internal/controller/cloudwatchlogs/definition"
	destination "github.com/spirosco/upbound-provider-aws/internal/controller/cloudwatchlogs/destination"
	destinationpolicy "github.com/spirosco/upbound-provider-aws/internal/controller/cloudwatchlogs/destinationpolicy"
	group "github.com/spirosco/upbound-provider-aws/internal/controller/cloudwatchlogs/group"
	metricfilter "github.com/spirosco/upbound-provider-aws/internal/controller/cloudwatchlogs/metricfilter"
	resourcepolicy "github.com/spirosco/upbound-provider-aws/internal/controller/cloudwatchlogs/resourcepolicy"
	stream "github.com/spirosco/upbound-provider-aws/internal/controller/cloudwatchlogs/stream"
	subscriptionfilter "github.com/spirosco/upbound-provider-aws/internal/controller/cloudwatchlogs/subscriptionfilter"
)

// Setup_cloudwatchlogs creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_cloudwatchlogs(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		definition.Setup,
		destination.Setup,
		destinationpolicy.Setup,
		group.Setup,
		metricfilter.Setup,
		resourcepolicy.Setup,
		stream.Setup,
		subscriptionfilter.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
