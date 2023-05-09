/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	lifecyclepolicy "github.com/spirosco/upbound-provider-aws/internal/controller/ecr/lifecyclepolicy"
	pullthroughcacherule "github.com/spirosco/upbound-provider-aws/internal/controller/ecr/pullthroughcacherule"
	registrypolicy "github.com/spirosco/upbound-provider-aws/internal/controller/ecr/registrypolicy"
	registryscanningconfiguration "github.com/spirosco/upbound-provider-aws/internal/controller/ecr/registryscanningconfiguration"
	replicationconfiguration "github.com/spirosco/upbound-provider-aws/internal/controller/ecr/replicationconfiguration"
	repository "github.com/spirosco/upbound-provider-aws/internal/controller/ecr/repository"
	repositorypolicy "github.com/spirosco/upbound-provider-aws/internal/controller/ecr/repositorypolicy"
)

// Setup_ecr creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_ecr(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		lifecyclepolicy.Setup,
		pullthroughcacherule.Setup,
		registrypolicy.Setup,
		registryscanningconfiguration.Setup,
		replicationconfiguration.Setup,
		repository.Setup,
		repositorypolicy.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
