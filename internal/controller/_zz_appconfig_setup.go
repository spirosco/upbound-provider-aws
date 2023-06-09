/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	application "github.com/spirosco/upbound-provider-aws/internal/controller/appconfig/application"
	configurationprofile "github.com/spirosco/upbound-provider-aws/internal/controller/appconfig/configurationprofile"
	deployment "github.com/spirosco/upbound-provider-aws/internal/controller/appconfig/deployment"
	deploymentstrategy "github.com/spirosco/upbound-provider-aws/internal/controller/appconfig/deploymentstrategy"
	environment "github.com/spirosco/upbound-provider-aws/internal/controller/appconfig/environment"
	extension "github.com/spirosco/upbound-provider-aws/internal/controller/appconfig/extension"
	extensionassociation "github.com/spirosco/upbound-provider-aws/internal/controller/appconfig/extensionassociation"
	hostedconfigurationversion "github.com/spirosco/upbound-provider-aws/internal/controller/appconfig/hostedconfigurationversion"
)

// Setup_appconfig creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_appconfig(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		application.Setup,
		configurationprofile.Setup,
		deployment.Setup,
		deploymentstrategy.Setup,
		environment.Setup,
		extension.Setup,
		extensionassociation.Setup,
		hostedconfigurationversion.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
