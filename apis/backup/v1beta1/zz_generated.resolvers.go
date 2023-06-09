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
	v1beta12 "github.com/spirosco/upbound-provider-aws/apis/sns/v1beta1"
	common "github.com/spirosco/upbound-provider-aws/config/common"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this Plan.
func (mg *Plan) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	for i3 := 0; i3 < len(mg.Spec.ForProvider.Rule); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Rule[i3].TargetVaultName),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.Rule[i3].TargetVaultNameRef,
			Selector:     mg.Spec.ForProvider.Rule[i3].TargetVaultNameSelector,
			To: reference.To{
				List:    &VaultList{},
				Managed: &Vault{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.Rule[i3].TargetVaultName")
		}
		mg.Spec.ForProvider.Rule[i3].TargetVaultName = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.Rule[i3].TargetVaultNameRef = rsp.ResolvedReference

	}

	return nil
}

// ResolveReferences of this Selection.
func (mg *Selection) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.IAMRoleArn),
		Extract:      common.ARNExtractor(),
		Reference:    mg.Spec.ForProvider.IAMRoleArnRef,
		Selector:     mg.Spec.ForProvider.IAMRoleArnSelector,
		To: reference.To{
			List:    &v1beta1.RoleList{},
			Managed: &v1beta1.Role{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.IAMRoleArn")
	}
	mg.Spec.ForProvider.IAMRoleArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.IAMRoleArnRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.PlanID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.PlanIDRef,
		Selector:     mg.Spec.ForProvider.PlanIDSelector,
		To: reference.To{
			List:    &PlanList{},
			Managed: &Plan{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.PlanID")
	}
	mg.Spec.ForProvider.PlanID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.PlanIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Vault.
func (mg *Vault) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.KMSKeyArn),
		Extract:      common.ARNExtractor(),
		Reference:    mg.Spec.ForProvider.KMSKeyArnRef,
		Selector:     mg.Spec.ForProvider.KMSKeyArnSelector,
		To: reference.To{
			List:    &v1beta11.KeyList{},
			Managed: &v1beta11.Key{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.KMSKeyArn")
	}
	mg.Spec.ForProvider.KMSKeyArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.KMSKeyArnRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this VaultLockConfiguration.
func (mg *VaultLockConfiguration) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.BackupVaultName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.BackupVaultNameRef,
		Selector:     mg.Spec.ForProvider.BackupVaultNameSelector,
		To: reference.To{
			List:    &VaultList{},
			Managed: &Vault{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.BackupVaultName")
	}
	mg.Spec.ForProvider.BackupVaultName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.BackupVaultNameRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this VaultNotifications.
func (mg *VaultNotifications) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.BackupVaultName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.BackupVaultNameRef,
		Selector:     mg.Spec.ForProvider.BackupVaultNameSelector,
		To: reference.To{
			List:    &VaultList{},
			Managed: &Vault{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.BackupVaultName")
	}
	mg.Spec.ForProvider.BackupVaultName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.BackupVaultNameRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.SnsTopicArn),
		Extract:      common.ARNExtractor(),
		Reference:    mg.Spec.ForProvider.SnsTopicArnRef,
		Selector:     mg.Spec.ForProvider.SnsTopicArnSelector,
		To: reference.To{
			List:    &v1beta12.TopicList{},
			Managed: &v1beta12.Topic{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.SnsTopicArn")
	}
	mg.Spec.ForProvider.SnsTopicArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.SnsTopicArnRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this VaultPolicy.
func (mg *VaultPolicy) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.BackupVaultName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.BackupVaultNameRef,
		Selector:     mg.Spec.ForProvider.BackupVaultNameSelector,
		To: reference.To{
			List:    &VaultList{},
			Managed: &Vault{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.BackupVaultName")
	}
	mg.Spec.ForProvider.BackupVaultName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.BackupVaultNameRef = rsp.ResolvedReference

	return nil
}
