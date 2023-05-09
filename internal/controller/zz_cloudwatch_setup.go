/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	compositealarm "github.com/spirosco/upbound-provider-aws/internal/controller/cloudwatch/compositealarm"
	dashboard "github.com/spirosco/upbound-provider-aws/internal/controller/cloudwatch/dashboard"
	metricalarm "github.com/spirosco/upbound-provider-aws/internal/controller/cloudwatch/metricalarm"
	metricstream "github.com/spirosco/upbound-provider-aws/internal/controller/cloudwatch/metricstream"
)

// Setup_cloudwatch creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_cloudwatch(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		compositealarm.Setup,
		dashboard.Setup,
		metricalarm.Setup,
		metricstream.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
