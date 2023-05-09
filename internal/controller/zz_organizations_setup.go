/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	account "github.com/spirosco/upbound-provider-aws/internal/controller/organizations/account"
	delegatedadministrator "github.com/spirosco/upbound-provider-aws/internal/controller/organizations/delegatedadministrator"
	organization "github.com/spirosco/upbound-provider-aws/internal/controller/organizations/organization"
	organizationalunit "github.com/spirosco/upbound-provider-aws/internal/controller/organizations/organizationalunit"
	policy "github.com/spirosco/upbound-provider-aws/internal/controller/organizations/policy"
	policyattachment "github.com/spirosco/upbound-provider-aws/internal/controller/organizations/policyattachment"
)

// Setup_organizations creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_organizations(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		account.Setup,
		delegatedadministrator.Setup,
		organization.Setup,
		organizationalunit.Setup,
		policy.Setup,
		policyattachment.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
