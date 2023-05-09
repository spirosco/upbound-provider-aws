/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	database "github.com/spirosco/upbound-provider-aws/internal/controller/athena/database"
	datacatalog "github.com/spirosco/upbound-provider-aws/internal/controller/athena/datacatalog"
	namedquery "github.com/spirosco/upbound-provider-aws/internal/controller/athena/namedquery"
	workgroup "github.com/spirosco/upbound-provider-aws/internal/controller/athena/workgroup"
)

// Setup_athena creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_athena(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		database.Setup,
		datacatalog.Setup,
		namedquery.Setup,
		workgroup.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
