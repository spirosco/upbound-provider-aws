/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	graph "github.com/spirosco/upbound-provider-aws/internal/controller/detective/graph"
	invitationaccepter "github.com/spirosco/upbound-provider-aws/internal/controller/detective/invitationaccepter"
	member "github.com/spirosco/upbound-provider-aws/internal/controller/detective/member"
)

// Setup_detective creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_detective(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		graph.Setup,
		invitationaccepter.Setup,
		member.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
