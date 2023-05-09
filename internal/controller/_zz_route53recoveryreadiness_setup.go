/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	cell "github.com/spirosco/upbound-provider-aws/internal/controller/route53recoveryreadiness/cell"
	readinesscheck "github.com/spirosco/upbound-provider-aws/internal/controller/route53recoveryreadiness/readinesscheck"
	recoverygroup "github.com/spirosco/upbound-provider-aws/internal/controller/route53recoveryreadiness/recoverygroup"
	resourceset "github.com/spirosco/upbound-provider-aws/internal/controller/route53recoveryreadiness/resourceset"
)

// Setup_route53recoveryreadiness creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_route53recoveryreadiness(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		cell.Setup,
		readinesscheck.Setup,
		recoverygroup.Setup,
		resourceset.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
