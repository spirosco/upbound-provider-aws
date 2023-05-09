/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1beta1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	errors "github.com/pkg/errors"
	v1beta1 "github.com/spirosco/upbound-provider-aws/apis/iam/v1beta1"
	v1beta11 "github.com/spirosco/upbound-provider-aws/apis/kms/v1beta1"
	common "github.com/spirosco/upbound-provider-aws/config/common"
	resource "github.com/upbound/upjet/pkg/resource"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this LifecyclePolicy.
func (mg *LifecyclePolicy) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ExecutionRoleArn),
		Extract:      common.ARNExtractor(),
		Reference:    mg.Spec.ForProvider.ExecutionRoleArnRef,
		Selector:     mg.Spec.ForProvider.ExecutionRoleArnSelector,
		To: reference.To{
			List:    &v1beta1.RoleList{},
			Managed: &v1beta1.Role{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ExecutionRoleArn")
	}
	mg.Spec.ForProvider.ExecutionRoleArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ExecutionRoleArnRef = rsp.ResolvedReference

	for i3 := 0; i3 < len(mg.Spec.ForProvider.PolicyDetails); i3++ {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.PolicyDetails[i3].Schedule); i4++ {
			for i5 := 0; i5 < len(mg.Spec.ForProvider.PolicyDetails[i3].Schedule[i4].CrossRegionCopyRule); i5++ {
				rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
					CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.PolicyDetails[i3].Schedule[i4].CrossRegionCopyRule[i5].CmkArn),
					Extract:      resource.ExtractParamPath("arn", true),
					Reference:    mg.Spec.ForProvider.PolicyDetails[i3].Schedule[i4].CrossRegionCopyRule[i5].CmkArnRef,
					Selector:     mg.Spec.ForProvider.PolicyDetails[i3].Schedule[i4].CrossRegionCopyRule[i5].CmkArnSelector,
					To: reference.To{
						List:    &v1beta11.KeyList{},
						Managed: &v1beta11.Key{},
					},
				})
				if err != nil {
					return errors.Wrap(err, "mg.Spec.ForProvider.PolicyDetails[i3].Schedule[i4].CrossRegionCopyRule[i5].CmkArn")
				}
				mg.Spec.ForProvider.PolicyDetails[i3].Schedule[i4].CrossRegionCopyRule[i5].CmkArn = reference.ToPtrValue(rsp.ResolvedValue)
				mg.Spec.ForProvider.PolicyDetails[i3].Schedule[i4].CrossRegionCopyRule[i5].CmkArnRef = rsp.ResolvedReference

			}
		}
	}

	return nil
}
