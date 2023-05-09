/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	channel "github.com/spirosco/upbound-provider-aws/internal/controller/ivs/channel"
	recordingconfiguration "github.com/spirosco/upbound-provider-aws/internal/controller/ivs/recordingconfiguration"
)

// Setup_ivs creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_ivs(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		channel.Setup,
		recordingconfiguration.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
