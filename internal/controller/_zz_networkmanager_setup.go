/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	attachmentaccepter "github.com/spirosco/upbound-provider-aws/internal/controller/networkmanager/attachmentaccepter"
	connectattachment "github.com/spirosco/upbound-provider-aws/internal/controller/networkmanager/connectattachment"
	connection "github.com/spirosco/upbound-provider-aws/internal/controller/networkmanager/connection"
	corenetwork "github.com/spirosco/upbound-provider-aws/internal/controller/networkmanager/corenetwork"
	customergatewayassociation "github.com/spirosco/upbound-provider-aws/internal/controller/networkmanager/customergatewayassociation"
	device "github.com/spirosco/upbound-provider-aws/internal/controller/networkmanager/device"
	globalnetwork "github.com/spirosco/upbound-provider-aws/internal/controller/networkmanager/globalnetwork"
	link "github.com/spirosco/upbound-provider-aws/internal/controller/networkmanager/link"
	linkassociation "github.com/spirosco/upbound-provider-aws/internal/controller/networkmanager/linkassociation"
	site "github.com/spirosco/upbound-provider-aws/internal/controller/networkmanager/site"
	transitgatewayconnectpeerassociation "github.com/spirosco/upbound-provider-aws/internal/controller/networkmanager/transitgatewayconnectpeerassociation"
	transitgatewayregistration "github.com/spirosco/upbound-provider-aws/internal/controller/networkmanager/transitgatewayregistration"
	vpcattachment "github.com/spirosco/upbound-provider-aws/internal/controller/networkmanager/vpcattachment"
)

// Setup_networkmanager creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_networkmanager(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		attachmentaccepter.Setup,
		connectattachment.Setup,
		connection.Setup,
		corenetwork.Setup,
		customergatewayassociation.Setup,
		device.Setup,
		globalnetwork.Setup,
		link.Setup,
		linkassociation.Setup,
		site.Setup,
		transitgatewayconnectpeerassociation.Setup,
		transitgatewayregistration.Setup,
		vpcattachment.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
