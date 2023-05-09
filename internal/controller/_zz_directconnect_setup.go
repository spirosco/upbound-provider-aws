/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	bgppeer "github.com/spirosco/upbound-provider-aws/internal/controller/directconnect/bgppeer"
	connection "github.com/spirosco/upbound-provider-aws/internal/controller/directconnect/connection"
	connectionassociation "github.com/spirosco/upbound-provider-aws/internal/controller/directconnect/connectionassociation"
	gateway "github.com/spirosco/upbound-provider-aws/internal/controller/directconnect/gateway"
	gatewayassociation "github.com/spirosco/upbound-provider-aws/internal/controller/directconnect/gatewayassociation"
	gatewayassociationproposal "github.com/spirosco/upbound-provider-aws/internal/controller/directconnect/gatewayassociationproposal"
	hostedprivatevirtualinterface "github.com/spirosco/upbound-provider-aws/internal/controller/directconnect/hostedprivatevirtualinterface"
	hostedprivatevirtualinterfaceaccepter "github.com/spirosco/upbound-provider-aws/internal/controller/directconnect/hostedprivatevirtualinterfaceaccepter"
	hostedpublicvirtualinterface "github.com/spirosco/upbound-provider-aws/internal/controller/directconnect/hostedpublicvirtualinterface"
	hostedpublicvirtualinterfaceaccepter "github.com/spirosco/upbound-provider-aws/internal/controller/directconnect/hostedpublicvirtualinterfaceaccepter"
	hostedtransitvirtualinterface "github.com/spirosco/upbound-provider-aws/internal/controller/directconnect/hostedtransitvirtualinterface"
	hostedtransitvirtualinterfaceaccepter "github.com/spirosco/upbound-provider-aws/internal/controller/directconnect/hostedtransitvirtualinterfaceaccepter"
	lag "github.com/spirosco/upbound-provider-aws/internal/controller/directconnect/lag"
	privatevirtualinterface "github.com/spirosco/upbound-provider-aws/internal/controller/directconnect/privatevirtualinterface"
	publicvirtualinterface "github.com/spirosco/upbound-provider-aws/internal/controller/directconnect/publicvirtualinterface"
	transitvirtualinterface "github.com/spirosco/upbound-provider-aws/internal/controller/directconnect/transitvirtualinterface"
)

// Setup_directconnect creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_directconnect(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		bgppeer.Setup,
		connection.Setup,
		connectionassociation.Setup,
		gateway.Setup,
		gatewayassociation.Setup,
		gatewayassociationproposal.Setup,
		hostedprivatevirtualinterface.Setup,
		hostedprivatevirtualinterfaceaccepter.Setup,
		hostedpublicvirtualinterface.Setup,
		hostedpublicvirtualinterfaceaccepter.Setup,
		hostedtransitvirtualinterface.Setup,
		hostedtransitvirtualinterfaceaccepter.Setup,
		lag.Setup,
		privatevirtualinterface.Setup,
		publicvirtualinterface.Setup,
		transitvirtualinterface.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
