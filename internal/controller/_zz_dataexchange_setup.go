/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	dataset "github.com/spirosco/upbound-provider-aws/internal/controller/dataexchange/dataset"
	revision "github.com/spirosco/upbound-provider-aws/internal/controller/dataexchange/revision"
)

// Setup_dataexchange creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_dataexchange(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		dataset.Setup,
		revision.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
