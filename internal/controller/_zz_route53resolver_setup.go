/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	endpoint "github.com/spirosco/upbound-provider-aws/internal/controller/route53resolver/endpoint"
	rule "github.com/spirosco/upbound-provider-aws/internal/controller/route53resolver/rule"
	ruleassociation "github.com/spirosco/upbound-provider-aws/internal/controller/route53resolver/ruleassociation"
)

// Setup_route53resolver creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_route53resolver(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		endpoint.Setup,
		rule.Setup,
		ruleassociation.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
