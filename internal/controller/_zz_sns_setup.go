/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	platformapplication "github.com/spirosco/upbound-provider-aws/internal/controller/sns/platformapplication"
	smspreferences "github.com/spirosco/upbound-provider-aws/internal/controller/sns/smspreferences"
	topic "github.com/spirosco/upbound-provider-aws/internal/controller/sns/topic"
	topicpolicy "github.com/spirosco/upbound-provider-aws/internal/controller/sns/topicpolicy"
	topicsubscription "github.com/spirosco/upbound-provider-aws/internal/controller/sns/topicsubscription"
)

// Setup_sns creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_sns(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		platformapplication.Setup,
		smspreferences.Setup,
		topic.Setup,
		topicpolicy.Setup,
		topicsubscription.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
