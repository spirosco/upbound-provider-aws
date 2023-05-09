/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	contributorinsights "github.com/spirosco/upbound-provider-aws/internal/controller/dynamodb/contributorinsights"
	globaltable "github.com/spirosco/upbound-provider-aws/internal/controller/dynamodb/globaltable"
	kinesisstreamingdestination "github.com/spirosco/upbound-provider-aws/internal/controller/dynamodb/kinesisstreamingdestination"
	table "github.com/spirosco/upbound-provider-aws/internal/controller/dynamodb/table"
	tableitem "github.com/spirosco/upbound-provider-aws/internal/controller/dynamodb/tableitem"
	tablereplica "github.com/spirosco/upbound-provider-aws/internal/controller/dynamodb/tablereplica"
	tag "github.com/spirosco/upbound-provider-aws/internal/controller/dynamodb/tag"
)

// Setup_dynamodb creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_dynamodb(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		contributorinsights.Setup,
		globaltable.Setup,
		kinesisstreamingdestination.Setup,
		table.Setup,
		tableitem.Setup,
		tablereplica.Setup,
		tag.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
