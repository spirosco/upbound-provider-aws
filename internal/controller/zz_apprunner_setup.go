/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	autoscalingconfigurationversion "github.com/spirosco/upbound-provider-aws/internal/controller/apprunner/autoscalingconfigurationversion"
	connection "github.com/spirosco/upbound-provider-aws/internal/controller/apprunner/connection"
	observabilityconfiguration "github.com/spirosco/upbound-provider-aws/internal/controller/apprunner/observabilityconfiguration"
	service "github.com/spirosco/upbound-provider-aws/internal/controller/apprunner/service"
	vpcconnector "github.com/spirosco/upbound-provider-aws/internal/controller/apprunner/vpcconnector"
)

// Setup_apprunner creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_apprunner(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		autoscalingconfigurationversion.Setup,
		connection.Setup,
		observabilityconfiguration.Setup,
		service.Setup,
		vpcconnector.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
