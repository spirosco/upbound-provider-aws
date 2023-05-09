/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1beta1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	errors "github.com/pkg/errors"
	v1beta12 "github.com/spirosco/upbound-provider-aws/apis/ecs/v1beta1"
	v1beta11 "github.com/spirosco/upbound-provider-aws/apis/iam/v1beta1"
	v1beta1 "github.com/spirosco/upbound-provider-aws/apis/organizations/v1beta1"
	common "github.com/spirosco/upbound-provider-aws/config/common"
	resource "github.com/upbound/upjet/pkg/resource"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this APIDestination.
func (mg *APIDestination) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ConnectionArn),
		Extract:      resource.ExtractParamPath("arn", true),
		Reference:    mg.Spec.ForProvider.ConnectionArnRef,
		Selector:     mg.Spec.ForProvider.ConnectionArnSelector,
		To: reference.To{
			List:    &ConnectionList{},
			Managed: &Connection{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ConnectionArn")
	}
	mg.Spec.ForProvider.ConnectionArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ConnectionArnRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Archive.
func (mg *Archive) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.EventSourceArn),
		Extract:      resource.ExtractParamPath("arn", true),
		Reference:    mg.Spec.ForProvider.EventSourceArnRef,
		Selector:     mg.Spec.ForProvider.EventSourceArnSelector,
		To: reference.To{
			List:    &BusList{},
			Managed: &Bus{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.EventSourceArn")
	}
	mg.Spec.ForProvider.EventSourceArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.EventSourceArnRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this BusPolicy.
func (mg *BusPolicy) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.EventBusName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.EventBusNameRef,
		Selector:     mg.Spec.ForProvider.EventBusNameSelector,
		To: reference.To{
			List:    &BusList{},
			Managed: &Bus{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.EventBusName")
	}
	mg.Spec.ForProvider.EventBusName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.EventBusNameRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Permission.
func (mg *Permission) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	for i3 := 0; i3 < len(mg.Spec.ForProvider.Condition); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Condition[i3].Value),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.Condition[i3].ValueRef,
			Selector:     mg.Spec.ForProvider.Condition[i3].ValueSelector,
			To: reference.To{
				List:    &v1beta1.OrganizationList{},
				Managed: &v1beta1.Organization{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.Condition[i3].Value")
		}
		mg.Spec.ForProvider.Condition[i3].Value = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.Condition[i3].ValueRef = rsp.ResolvedReference

	}
	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.EventBusName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.EventBusNameRef,
		Selector:     mg.Spec.ForProvider.EventBusNameSelector,
		To: reference.To{
			List:    &BusList{},
			Managed: &Bus{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.EventBusName")
	}
	mg.Spec.ForProvider.EventBusName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.EventBusNameRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Rule.
func (mg *Rule) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.EventBusName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.EventBusNameRef,
		Selector:     mg.Spec.ForProvider.EventBusNameSelector,
		To: reference.To{
			List:    &BusList{},
			Managed: &Bus{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.EventBusName")
	}
	mg.Spec.ForProvider.EventBusName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.EventBusNameRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.RoleArn),
		Extract:      common.ARNExtractor(),
		Reference:    mg.Spec.ForProvider.RoleArnRef,
		Selector:     mg.Spec.ForProvider.RoleArnSelector,
		To: reference.To{
			List:    &v1beta11.RoleList{},
			Managed: &v1beta11.Role{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.RoleArn")
	}
	mg.Spec.ForProvider.RoleArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.RoleArnRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Target.
func (mg *Target) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	for i3 := 0; i3 < len(mg.Spec.ForProvider.EcsTarget); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.EcsTarget[i3].TaskDefinitionArn),
			Extract:      resource.ExtractParamPath("arn", true),
			Reference:    mg.Spec.ForProvider.EcsTarget[i3].TaskDefinitionArnRef,
			Selector:     mg.Spec.ForProvider.EcsTarget[i3].TaskDefinitionArnSelector,
			To: reference.To{
				List:    &v1beta12.TaskDefinitionList{},
				Managed: &v1beta12.TaskDefinition{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.EcsTarget[i3].TaskDefinitionArn")
		}
		mg.Spec.ForProvider.EcsTarget[i3].TaskDefinitionArn = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.EcsTarget[i3].TaskDefinitionArnRef = rsp.ResolvedReference

	}
	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.EventBusName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.EventBusNameRef,
		Selector:     mg.Spec.ForProvider.EventBusNameSelector,
		To: reference.To{
			List:    &BusList{},
			Managed: &Bus{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.EventBusName")
	}
	mg.Spec.ForProvider.EventBusName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.EventBusNameRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.RoleArn),
		Extract:      common.ARNExtractor(),
		Reference:    mg.Spec.ForProvider.RoleArnRef,
		Selector:     mg.Spec.ForProvider.RoleArnSelector,
		To: reference.To{
			List:    &v1beta11.RoleList{},
			Managed: &v1beta11.Role{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.RoleArn")
	}
	mg.Spec.ForProvider.RoleArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.RoleArnRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Rule),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.RuleRef,
		Selector:     mg.Spec.ForProvider.RuleSelector,
		To: reference.To{
			List:    &RuleList{},
			Managed: &Rule{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Rule")
	}
	mg.Spec.ForProvider.Rule = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.RuleRef = rsp.ResolvedReference

	return nil
}
