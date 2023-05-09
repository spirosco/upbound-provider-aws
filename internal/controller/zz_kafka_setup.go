/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	cluster "github.com/spirosco/upbound-provider-aws/internal/controller/kafka/cluster"
	configuration "github.com/spirosco/upbound-provider-aws/internal/controller/kafka/configuration"
)

// Setup_kafka creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_kafka(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		cluster.Setup,
		configuration.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
