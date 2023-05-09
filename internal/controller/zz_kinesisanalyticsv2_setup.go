/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	application "github.com/spirosco/upbound-provider-aws/internal/controller/kinesisanalyticsv2/application"
	applicationsnapshot "github.com/spirosco/upbound-provider-aws/internal/controller/kinesisanalyticsv2/applicationsnapshot"
)

// Setup_kinesisanalyticsv2 creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_kinesisanalyticsv2(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		application.Setup,
		applicationsnapshot.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
