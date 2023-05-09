/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	analyzer "github.com/spirosco/upbound-provider-aws/internal/controller/accessanalyzer/analyzer"
	archiverule "github.com/spirosco/upbound-provider-aws/internal/controller/accessanalyzer/archiverule"
)

// Setup_accessanalyzer creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_accessanalyzer(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		analyzer.Setup,
		archiverule.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
