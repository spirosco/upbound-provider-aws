/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	bytematchset "github.com/spirosco/upbound-provider-aws/internal/controller/waf/bytematchset"
	geomatchset "github.com/spirosco/upbound-provider-aws/internal/controller/waf/geomatchset"
	ipset "github.com/spirosco/upbound-provider-aws/internal/controller/waf/ipset"
	ratebasedrule "github.com/spirosco/upbound-provider-aws/internal/controller/waf/ratebasedrule"
	regexmatchset "github.com/spirosco/upbound-provider-aws/internal/controller/waf/regexmatchset"
	regexpatternset "github.com/spirosco/upbound-provider-aws/internal/controller/waf/regexpatternset"
	rule "github.com/spirosco/upbound-provider-aws/internal/controller/waf/rule"
	sizeconstraintset "github.com/spirosco/upbound-provider-aws/internal/controller/waf/sizeconstraintset"
	sqlinjectionmatchset "github.com/spirosco/upbound-provider-aws/internal/controller/waf/sqlinjectionmatchset"
	webacl "github.com/spirosco/upbound-provider-aws/internal/controller/waf/webacl"
	xssmatchset "github.com/spirosco/upbound-provider-aws/internal/controller/waf/xssmatchset"
)

// Setup_waf creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_waf(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		bytematchset.Setup,
		geomatchset.Setup,
		ipset.Setup,
		ratebasedrule.Setup,
		regexmatchset.Setup,
		regexpatternset.Setup,
		rule.Setup,
		sizeconstraintset.Setup,
		sqlinjectionmatchset.Setup,
		webacl.Setup,
		xssmatchset.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
