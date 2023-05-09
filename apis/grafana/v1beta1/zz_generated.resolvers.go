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
	common "github.com/spirosco/upbound-provider-aws/config/common"
	resource "github.com/upbound/upjet/pkg/resource"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this LicenseAssociation.
func (mg *LicenseAssociation) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.WorkspaceID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.WorkspaceIDRef,
		Selector:     mg.Spec.ForProvider.WorkspaceIDSelector,
		To: reference.To{
			List:    &WorkspaceList{},
			Managed: &Workspace{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.WorkspaceID")
	}
	mg.Spec.ForProvider.WorkspaceID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.WorkspaceIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this RoleAssociation.
func (mg *RoleAssociation) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.WorkspaceID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.WorkspaceIDRef,
		Selector:     mg.Spec.ForProvider.WorkspaceIDSelector,
		To: reference.To{
			List:    &WorkspaceList{},
			Managed: &Workspace{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.WorkspaceID")
	}
	mg.Spec.ForProvider.WorkspaceID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.WorkspaceIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Workspace.
func (mg *Workspace) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.RoleArn),
		Extract:      common.ARNExtractor(),
		Reference:    mg.Spec.ForProvider.RoleArnRef,
		Selector:     mg.Spec.ForProvider.RoleArnSelector,
		To: reference.To{
			List:    &v1beta1.RoleList{},
			Managed: &v1beta1.Role{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.RoleArn")
	}
	mg.Spec.ForProvider.RoleArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.RoleArnRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this WorkspaceAPIKey.
func (mg *WorkspaceAPIKey) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.WorkspaceID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.WorkspaceIDRef,
		Selector:     mg.Spec.ForProvider.WorkspaceIDSelector,
		To: reference.To{
			List:    &WorkspaceList{},
			Managed: &Workspace{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.WorkspaceID")
	}
	mg.Spec.ForProvider.WorkspaceID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.WorkspaceIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this WorkspaceSAMLConfiguration.
func (mg *WorkspaceSAMLConfiguration) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.WorkspaceID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.WorkspaceIDRef,
		Selector:     mg.Spec.ForProvider.WorkspaceIDSelector,
		To: reference.To{
			List:    &WorkspaceList{},
			Managed: &Workspace{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.WorkspaceID")
	}
	mg.Spec.ForProvider.WorkspaceID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.WorkspaceIDRef = rsp.ResolvedReference

	return nil
}
