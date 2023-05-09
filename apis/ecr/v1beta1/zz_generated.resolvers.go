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
	common "github.com/spirosco/upbound-provider-aws/config/common"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this LifecyclePolicy.
func (mg *LifecyclePolicy) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Repository),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.RepositoryRef,
		Selector:     mg.Spec.ForProvider.RepositorySelector,
		To: reference.To{
			List:    &RepositoryList{},
			Managed: &Repository{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Repository")
	}
	mg.Spec.ForProvider.Repository = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.RepositoryRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Repository.
func (mg *Repository) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	for i3 := 0; i3 < len(mg.Spec.ForProvider.EncryptionConfiguration); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.EncryptionConfiguration[i3].KMSKey),
			Extract:      common.ARNExtractor(),
			Reference:    mg.Spec.ForProvider.EncryptionConfiguration[i3].KMSKeyRef,
			Selector:     mg.Spec.ForProvider.EncryptionConfiguration[i3].KMSKeySelector,
			To: reference.To{
				List:    &v1beta1.KeyList{},
				Managed: &v1beta1.Key{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.EncryptionConfiguration[i3].KMSKey")
		}
		mg.Spec.ForProvider.EncryptionConfiguration[i3].KMSKey = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.EncryptionConfiguration[i3].KMSKeyRef = rsp.ResolvedReference

	}

	return nil
}

// ResolveReferences of this RepositoryPolicy.
func (mg *RepositoryPolicy) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Repository),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.RepositoryRef,
		Selector:     mg.Spec.ForProvider.RepositorySelector,
		To: reference.To{
			List:    &RepositoryList{},
			Managed: &Repository{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Repository")
	}
	mg.Spec.ForProvider.Repository = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.RepositoryRef = rsp.ResolvedReference

	return nil
}
