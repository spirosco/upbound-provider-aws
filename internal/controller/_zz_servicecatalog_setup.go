/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	budgetresourceassociation "github.com/spirosco/upbound-provider-aws/internal/controller/servicecatalog/budgetresourceassociation"
	constraint "github.com/spirosco/upbound-provider-aws/internal/controller/servicecatalog/constraint"
	portfolio "github.com/spirosco/upbound-provider-aws/internal/controller/servicecatalog/portfolio"
	portfolioshare "github.com/spirosco/upbound-provider-aws/internal/controller/servicecatalog/portfolioshare"
	principalportfolioassociation "github.com/spirosco/upbound-provider-aws/internal/controller/servicecatalog/principalportfolioassociation"
	product "github.com/spirosco/upbound-provider-aws/internal/controller/servicecatalog/product"
	productportfolioassociation "github.com/spirosco/upbound-provider-aws/internal/controller/servicecatalog/productportfolioassociation"
	provisioningartifact "github.com/spirosco/upbound-provider-aws/internal/controller/servicecatalog/provisioningartifact"
	serviceaction "github.com/spirosco/upbound-provider-aws/internal/controller/servicecatalog/serviceaction"
	tagoption "github.com/spirosco/upbound-provider-aws/internal/controller/servicecatalog/tagoption"
	tagoptionresourceassociation "github.com/spirosco/upbound-provider-aws/internal/controller/servicecatalog/tagoptionresourceassociation"
)

// Setup_servicecatalog creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_servicecatalog(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		budgetresourceassociation.Setup,
		constraint.Setup,
		portfolio.Setup,
		portfolioshare.Setup,
		principalportfolioassociation.Setup,
		product.Setup,
		productportfolioassociation.Setup,
		provisioningartifact.Setup,
		serviceaction.Setup,
		tagoption.Setup,
		tagoptionresourceassociation.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
