/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	certificate "github.com/spirosco/upbound-provider-aws/internal/controller/acm/certificate"
	certificatevalidation "github.com/spirosco/upbound-provider-aws/internal/controller/acm/certificatevalidation"
)

// Setup_acm creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_acm(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		certificate.Setup,
		certificatevalidation.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
