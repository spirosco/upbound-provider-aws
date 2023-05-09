/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1beta1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	errors "github.com/pkg/errors"
	v1beta12 "github.com/spirosco/upbound-provider-aws/apis/iam/v1beta1"
	v1beta11 "github.com/spirosco/upbound-provider-aws/apis/kinesis/v1beta1"
	v1beta1 "github.com/spirosco/upbound-provider-aws/apis/kms/v1beta1"
	common "github.com/spirosco/upbound-provider-aws/config/common"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this Ledger.
func (mg *Ledger) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.KMSKey),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.KMSKeyRef,
		Selector:     mg.Spec.ForProvider.KMSKeySelector,
		To: reference.To{
			List:    &v1beta1.KeyList{},
			Managed: &v1beta1.Key{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.KMSKey")
	}
	mg.Spec.ForProvider.KMSKey = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.KMSKeyRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Stream.
func (mg *Stream) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	for i3 := 0; i3 < len(mg.Spec.ForProvider.KinesisConfiguration); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.KinesisConfiguration[i3].StreamArn),
			Extract:      common.TerraformID(),
			Reference:    mg.Spec.ForProvider.KinesisConfiguration[i3].StreamArnRef,
			Selector:     mg.Spec.ForProvider.KinesisConfiguration[i3].StreamArnSelector,
			To: reference.To{
				List:    &v1beta11.StreamList{},
				Managed: &v1beta11.Stream{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.KinesisConfiguration[i3].StreamArn")
		}
		mg.Spec.ForProvider.KinesisConfiguration[i3].StreamArn = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.KinesisConfiguration[i3].StreamArnRef = rsp.ResolvedReference

	}
	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.LedgerName),
		Extract:      common.TerraformID(),
		Reference:    mg.Spec.ForProvider.LedgerNameRef,
		Selector:     mg.Spec.ForProvider.LedgerNameSelector,
		To: reference.To{
			List:    &LedgerList{},
			Managed: &Ledger{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.LedgerName")
	}
	mg.Spec.ForProvider.LedgerName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.LedgerNameRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.RoleArn),
		Extract:      common.ARNExtractor(),
		Reference:    mg.Spec.ForProvider.RoleArnRef,
		Selector:     mg.Spec.ForProvider.RoleArnSelector,
		To: reference.To{
			List:    &v1beta12.RoleList{},
			Managed: &v1beta12.Role{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.RoleArn")
	}
	mg.Spec.ForProvider.RoleArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.RoleArnRef = rsp.ResolvedReference

	return nil
}
