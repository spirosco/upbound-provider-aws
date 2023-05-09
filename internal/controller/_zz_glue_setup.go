/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	catalogdatabase "github.com/spirosco/upbound-provider-aws/internal/controller/glue/catalogdatabase"
	catalogtable "github.com/spirosco/upbound-provider-aws/internal/controller/glue/catalogtable"
	classifier "github.com/spirosco/upbound-provider-aws/internal/controller/glue/classifier"
	connection "github.com/spirosco/upbound-provider-aws/internal/controller/glue/connection"
	crawler "github.com/spirosco/upbound-provider-aws/internal/controller/glue/crawler"
	datacatalogencryptionsettings "github.com/spirosco/upbound-provider-aws/internal/controller/glue/datacatalogencryptionsettings"
	job "github.com/spirosco/upbound-provider-aws/internal/controller/glue/job"
	registry "github.com/spirosco/upbound-provider-aws/internal/controller/glue/registry"
	resourcepolicy "github.com/spirosco/upbound-provider-aws/internal/controller/glue/resourcepolicy"
	schema "github.com/spirosco/upbound-provider-aws/internal/controller/glue/schema"
	securityconfiguration "github.com/spirosco/upbound-provider-aws/internal/controller/glue/securityconfiguration"
	trigger "github.com/spirosco/upbound-provider-aws/internal/controller/glue/trigger"
	userdefinedfunction "github.com/spirosco/upbound-provider-aws/internal/controller/glue/userdefinedfunction"
	workflow "github.com/spirosco/upbound-provider-aws/internal/controller/glue/workflow"
)

// Setup_glue creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_glue(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		catalogdatabase.Setup,
		catalogtable.Setup,
		classifier.Setup,
		connection.Setup,
		crawler.Setup,
		datacatalogencryptionsettings.Setup,
		job.Setup,
		registry.Setup,
		resourcepolicy.Setup,
		schema.Setup,
		securityconfiguration.Setup,
		trigger.Setup,
		userdefinedfunction.Setup,
		workflow.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
