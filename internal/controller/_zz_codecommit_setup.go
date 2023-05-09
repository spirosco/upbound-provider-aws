/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	approvalruletemplate "github.com/spirosco/upbound-provider-aws/internal/controller/codecommit/approvalruletemplate"
	approvalruletemplateassociation "github.com/spirosco/upbound-provider-aws/internal/controller/codecommit/approvalruletemplateassociation"
	repository "github.com/spirosco/upbound-provider-aws/internal/controller/codecommit/repository"
	trigger "github.com/spirosco/upbound-provider-aws/internal/controller/codecommit/trigger"
)

// Setup_codecommit creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_codecommit(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		approvalruletemplate.Setup,
		approvalruletemplateassociation.Setup,
		repository.Setup,
		trigger.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
