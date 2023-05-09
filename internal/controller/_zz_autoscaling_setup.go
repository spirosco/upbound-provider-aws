/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	attachment "github.com/spirosco/upbound-provider-aws/internal/controller/autoscaling/attachment"
	autoscalinggroup "github.com/spirosco/upbound-provider-aws/internal/controller/autoscaling/autoscalinggroup"
	grouptag "github.com/spirosco/upbound-provider-aws/internal/controller/autoscaling/grouptag"
	launchconfiguration "github.com/spirosco/upbound-provider-aws/internal/controller/autoscaling/launchconfiguration"
	lifecyclehook "github.com/spirosco/upbound-provider-aws/internal/controller/autoscaling/lifecyclehook"
	notification "github.com/spirosco/upbound-provider-aws/internal/controller/autoscaling/notification"
	policy "github.com/spirosco/upbound-provider-aws/internal/controller/autoscaling/policy"
	schedule "github.com/spirosco/upbound-provider-aws/internal/controller/autoscaling/schedule"
)

// Setup_autoscaling creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_autoscaling(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		attachment.Setup,
		autoscalinggroup.Setup,
		grouptag.Setup,
		launchconfiguration.Setup,
		lifecyclehook.Setup,
		notification.Setup,
		policy.Setup,
		schedule.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
