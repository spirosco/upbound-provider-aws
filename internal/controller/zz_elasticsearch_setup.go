/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	domain "github.com/spirosco/upbound-provider-aws/internal/controller/elasticsearch/domain"
	domainpolicy "github.com/spirosco/upbound-provider-aws/internal/controller/elasticsearch/domainpolicy"
	domainsamloptions "github.com/spirosco/upbound-provider-aws/internal/controller/elasticsearch/domainsamloptions"
)

// Setup_elasticsearch creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_elasticsearch(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		domain.Setup,
		domainpolicy.Setup,
		domainsamloptions.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
