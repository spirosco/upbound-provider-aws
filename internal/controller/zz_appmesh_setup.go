/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	gatewayroute "github.com/spirosco/upbound-provider-aws/internal/controller/appmesh/gatewayroute"
	mesh "github.com/spirosco/upbound-provider-aws/internal/controller/appmesh/mesh"
	route "github.com/spirosco/upbound-provider-aws/internal/controller/appmesh/route"
	virtualgateway "github.com/spirosco/upbound-provider-aws/internal/controller/appmesh/virtualgateway"
	virtualnode "github.com/spirosco/upbound-provider-aws/internal/controller/appmesh/virtualnode"
	virtualrouter "github.com/spirosco/upbound-provider-aws/internal/controller/appmesh/virtualrouter"
	virtualservice "github.com/spirosco/upbound-provider-aws/internal/controller/appmesh/virtualservice"
)

// Setup_appmesh creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_appmesh(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		gatewayroute.Setup,
		mesh.Setup,
		route.Setup,
		virtualgateway.Setup,
		virtualnode.Setup,
		virtualrouter.Setup,
		virtualservice.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
