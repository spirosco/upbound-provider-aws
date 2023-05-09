/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1beta1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	errors "github.com/pkg/errors"
	v1beta1 "github.com/spirosco/upbound-provider-aws/apis/kms/v1beta1"
	v1beta11 "github.com/spirosco/upbound-provider-aws/apis/lambda/v1beta1"
	resource "github.com/upbound/upjet/pkg/resource"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this Secret.
func (mg *Secret) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.KMSKeyID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.KMSKeyIDRef,
		Selector:     mg.Spec.ForProvider.KMSKeyIDSelector,
		To: reference.To{
			List:    &v1beta1.KeyList{},
			Managed: &v1beta1.Key{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.KMSKeyID")
	}
	mg.Spec.ForProvider.KMSKeyID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.KMSKeyIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this SecretPolicy.
func (mg *SecretPolicy) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.SecretArn),
		Extract:      resource.ExtractParamPath("arn", true),
		Reference:    mg.Spec.ForProvider.SecretArnRef,
		Selector:     mg.Spec.ForProvider.SecretArnSelector,
		To: reference.To{
			List:    &SecretList{},
			Managed: &Secret{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.SecretArn")
	}
	mg.Spec.ForProvider.SecretArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.SecretArnRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this SecretRotation.
func (mg *SecretRotation) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.RotationLambdaArn),
		Extract:      resource.ExtractParamPath("arn", true),
		Reference:    mg.Spec.ForProvider.RotationLambdaArnRef,
		Selector:     mg.Spec.ForProvider.RotationLambdaArnSelector,
		To: reference.To{
			List:    &v1beta11.FunctionList{},
			Managed: &v1beta11.Function{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.RotationLambdaArn")
	}
	mg.Spec.ForProvider.RotationLambdaArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.RotationLambdaArnRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.SecretID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.SecretIDRef,
		Selector:     mg.Spec.ForProvider.SecretIDSelector,
		To: reference.To{
			List:    &SecretList{},
			Managed: &Secret{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.SecretID")
	}
	mg.Spec.ForProvider.SecretID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.SecretIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this SecretVersion.
func (mg *SecretVersion) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.SecretID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.SecretIDRef,
		Selector:     mg.Spec.ForProvider.SecretIDSelector,
		To: reference.To{
			List:    &SecretList{},
			Managed: &Secret{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.SecretID")
	}
	mg.Spec.ForProvider.SecretID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.SecretIDRef = rsp.ResolvedReference

	return nil
}
