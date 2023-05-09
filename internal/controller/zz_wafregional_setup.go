/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	bytematchset "github.com/spirosco/upbound-provider-aws/internal/controller/wafregional/bytematchset"
	geomatchset "github.com/spirosco/upbound-provider-aws/internal/controller/wafregional/geomatchset"
	ipset "github.com/spirosco/upbound-provider-aws/internal/controller/wafregional/ipset"
	ratebasedrule "github.com/spirosco/upbound-provider-aws/internal/controller/wafregional/ratebasedrule"
	regexmatchset "github.com/spirosco/upbound-provider-aws/internal/controller/wafregional/regexmatchset"
	regexpatternset "github.com/spirosco/upbound-provider-aws/internal/controller/wafregional/regexpatternset"
	rule "github.com/spirosco/upbound-provider-aws/internal/controller/wafregional/rule"
	sizeconstraintset "github.com/spirosco/upbound-provider-aws/internal/controller/wafregional/sizeconstraintset"
	sqlinjectionmatchset "github.com/spirosco/upbound-provider-aws/internal/controller/wafregional/sqlinjectionmatchset"
	webacl "github.com/spirosco/upbound-provider-aws/internal/controller/wafregional/webacl"
	xssmatchset "github.com/spirosco/upbound-provider-aws/internal/controller/wafregional/xssmatchset"
)

// Setup_wafregional creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_wafregional(mgr ctrl.Manager, o controller.Options) error {
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
