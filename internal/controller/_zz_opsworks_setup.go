/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	application "github.com/spirosco/upbound-provider-aws/internal/controller/opsworks/application"
	customlayer "github.com/spirosco/upbound-provider-aws/internal/controller/opsworks/customlayer"
	ecsclusterlayer "github.com/spirosco/upbound-provider-aws/internal/controller/opsworks/ecsclusterlayer"
	ganglialayer "github.com/spirosco/upbound-provider-aws/internal/controller/opsworks/ganglialayer"
	haproxylayer "github.com/spirosco/upbound-provider-aws/internal/controller/opsworks/haproxylayer"
	instance "github.com/spirosco/upbound-provider-aws/internal/controller/opsworks/instance"
	javaapplayer "github.com/spirosco/upbound-provider-aws/internal/controller/opsworks/javaapplayer"
	memcachedlayer "github.com/spirosco/upbound-provider-aws/internal/controller/opsworks/memcachedlayer"
	mysqllayer "github.com/spirosco/upbound-provider-aws/internal/controller/opsworks/mysqllayer"
	nodejsapplayer "github.com/spirosco/upbound-provider-aws/internal/controller/opsworks/nodejsapplayer"
	permission "github.com/spirosco/upbound-provider-aws/internal/controller/opsworks/permission"
	phpapplayer "github.com/spirosco/upbound-provider-aws/internal/controller/opsworks/phpapplayer"
	railsapplayer "github.com/spirosco/upbound-provider-aws/internal/controller/opsworks/railsapplayer"
	rdsdbinstance "github.com/spirosco/upbound-provider-aws/internal/controller/opsworks/rdsdbinstance"
	stack "github.com/spirosco/upbound-provider-aws/internal/controller/opsworks/stack"
	staticweblayer "github.com/spirosco/upbound-provider-aws/internal/controller/opsworks/staticweblayer"
	userprofile "github.com/spirosco/upbound-provider-aws/internal/controller/opsworks/userprofile"
)

// Setup_opsworks creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_opsworks(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		application.Setup,
		customlayer.Setup,
		ecsclusterlayer.Setup,
		ganglialayer.Setup,
		haproxylayer.Setup,
		instance.Setup,
		javaapplayer.Setup,
		memcachedlayer.Setup,
		mysqllayer.Setup,
		nodejsapplayer.Setup,
		permission.Setup,
		phpapplayer.Setup,
		railsapplayer.Setup,
		rdsdbinstance.Setup,
		stack.Setup,
		staticweblayer.Setup,
		userprofile.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
