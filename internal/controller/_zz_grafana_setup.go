/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	licenseassociation "github.com/spirosco/upbound-provider-aws/internal/controller/grafana/licenseassociation"
	roleassociation "github.com/spirosco/upbound-provider-aws/internal/controller/grafana/roleassociation"
	workspace "github.com/spirosco/upbound-provider-aws/internal/controller/grafana/workspace"
	workspaceapikey "github.com/spirosco/upbound-provider-aws/internal/controller/grafana/workspaceapikey"
	workspacesamlconfiguration "github.com/spirosco/upbound-provider-aws/internal/controller/grafana/workspacesamlconfiguration"
)

// Setup_grafana creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_grafana(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		licenseassociation.Setup,
		roleassociation.Setup,
		workspace.Setup,
		workspaceapikey.Setup,
		workspacesamlconfiguration.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
