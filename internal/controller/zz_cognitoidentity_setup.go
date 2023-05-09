/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	cognitoidentitypoolproviderprincipaltag "github.com/spirosco/upbound-provider-aws/internal/controller/cognitoidentity/cognitoidentitypoolproviderprincipaltag"
	pool "github.com/spirosco/upbound-provider-aws/internal/controller/cognitoidentity/pool"
	poolrolesattachment "github.com/spirosco/upbound-provider-aws/internal/controller/cognitoidentity/poolrolesattachment"
)

// Setup_cognitoidentity creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_cognitoidentity(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		cognitoidentitypoolproviderprincipaltag.Setup,
		pool.Setup,
		poolrolesattachment.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
