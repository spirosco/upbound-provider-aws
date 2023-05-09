/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	policy "github.com/spirosco/upbound-provider-aws/internal/controller/appautoscaling/policy"
	scheduledaction "github.com/spirosco/upbound-provider-aws/internal/controller/appautoscaling/scheduledaction"
	target "github.com/spirosco/upbound-provider-aws/internal/controller/appautoscaling/target"
)

// Setup_appautoscaling creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_appautoscaling(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		policy.Setup,
		scheduledaction.Setup,
		target.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
