/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	delegationset "github.com/spirosco/upbound-provider-aws/internal/controller/route53/delegationset"
	healthcheck "github.com/spirosco/upbound-provider-aws/internal/controller/route53/healthcheck"
	hostedzonednssec "github.com/spirosco/upbound-provider-aws/internal/controller/route53/hostedzonednssec"
	record "github.com/spirosco/upbound-provider-aws/internal/controller/route53/record"
	resolverconfig "github.com/spirosco/upbound-provider-aws/internal/controller/route53/resolverconfig"
	trafficpolicy "github.com/spirosco/upbound-provider-aws/internal/controller/route53/trafficpolicy"
	trafficpolicyinstance "github.com/spirosco/upbound-provider-aws/internal/controller/route53/trafficpolicyinstance"
	vpcassociationauthorization "github.com/spirosco/upbound-provider-aws/internal/controller/route53/vpcassociationauthorization"
	zone "github.com/spirosco/upbound-provider-aws/internal/controller/route53/zone"
)

// Setup_route53 creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_route53(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		delegationset.Setup,
		healthcheck.Setup,
		hostedzonednssec.Setup,
		record.Setup,
		resolverconfig.Setup,
		trafficpolicy.Setup,
		trafficpolicyinstance.Setup,
		vpcassociationauthorization.Setup,
		zone.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
