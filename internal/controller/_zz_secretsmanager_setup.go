/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	secret "github.com/spirosco/upbound-provider-aws/internal/controller/secretsmanager/secret"
	secretpolicy "github.com/spirosco/upbound-provider-aws/internal/controller/secretsmanager/secretpolicy"
	secretrotation "github.com/spirosco/upbound-provider-aws/internal/controller/secretsmanager/secretrotation"
	secretversion "github.com/spirosco/upbound-provider-aws/internal/controller/secretsmanager/secretversion"
)

// Setup_secretsmanager creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_secretsmanager(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		secret.Setup,
		secretpolicy.Setup,
		secretrotation.Setup,
		secretversion.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
