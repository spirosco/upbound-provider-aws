/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	stream "github.com/spirosco/upbound-provider-aws/internal/controller/kinesis/stream"
	streamconsumer "github.com/spirosco/upbound-provider-aws/internal/controller/kinesis/streamconsumer"
)

// Setup_kinesis creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_kinesis(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		stream.Setup,
		streamconsumer.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
