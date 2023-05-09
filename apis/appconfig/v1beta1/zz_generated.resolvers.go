/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1beta1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	errors "github.com/pkg/errors"
	v1beta11 "github.com/spirosco/upbound-provider-aws/apis/cloudwatch/v1beta1"
	v1beta1 "github.com/spirosco/upbound-provider-aws/apis/iam/v1beta1"
	v1beta12 "github.com/spirosco/upbound-provider-aws/apis/sns/v1beta1"
	common "github.com/spirosco/upbound-provider-aws/config/common"
	resource "github.com/upbound/upjet/pkg/resource"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this ConfigurationProfile.
func (mg *ConfigurationProfile) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ApplicationID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.ApplicationIDRef,
		Selector:     mg.Spec.ForProvider.ApplicationIDSelector,
		To: reference.To{
			List:    &ApplicationList{},
			Managed: &Application{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ApplicationID")
	}
	mg.Spec.ForProvider.ApplicationID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ApplicationIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.RetrievalRoleArn),
		Extract:      common.ARNExtractor(),
		Reference:    mg.Spec.ForProvider.RetrievalRoleArnRef,
		Selector:     mg.Spec.ForProvider.RetrievalRoleArnSelector,
		To: reference.To{
			List:    &v1beta1.RoleList{},
			Managed: &v1beta1.Role{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.RetrievalRoleArn")
	}
	mg.Spec.ForProvider.RetrievalRoleArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.RetrievalRoleArnRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Deployment.
func (mg *Deployment) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ApplicationID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.ApplicationIDRef,
		Selector:     mg.Spec.ForProvider.ApplicationIDSelector,
		To: reference.To{
			List:    &ApplicationList{},
			Managed: &Application{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ApplicationID")
	}
	mg.Spec.ForProvider.ApplicationID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ApplicationIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ConfigurationProfileID),
		Extract:      resource.ExtractParamPath("configuration_profile_id", true),
		Reference:    mg.Spec.ForProvider.ConfigurationProfileIDRef,
		Selector:     mg.Spec.ForProvider.ConfigurationProfileIDSelector,
		To: reference.To{
			List:    &ConfigurationProfileList{},
			Managed: &ConfigurationProfile{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ConfigurationProfileID")
	}
	mg.Spec.ForProvider.ConfigurationProfileID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ConfigurationProfileIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ConfigurationVersion),
		Extract:      resource.ExtractParamPath("version_number", true),
		Reference:    mg.Spec.ForProvider.ConfigurationVersionRef,
		Selector:     mg.Spec.ForProvider.ConfigurationVersionSelector,
		To: reference.To{
			List:    &HostedConfigurationVersionList{},
			Managed: &HostedConfigurationVersion{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ConfigurationVersion")
	}
	mg.Spec.ForProvider.ConfigurationVersion = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ConfigurationVersionRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.DeploymentStrategyID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.DeploymentStrategyIDRef,
		Selector:     mg.Spec.ForProvider.DeploymentStrategyIDSelector,
		To: reference.To{
			List:    &DeploymentStrategyList{},
			Managed: &DeploymentStrategy{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.DeploymentStrategyID")
	}
	mg.Spec.ForProvider.DeploymentStrategyID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.DeploymentStrategyIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.EnvironmentID),
		Extract:      resource.ExtractParamPath("environment_id", true),
		Reference:    mg.Spec.ForProvider.EnvironmentIDRef,
		Selector:     mg.Spec.ForProvider.EnvironmentIDSelector,
		To: reference.To{
			List:    &EnvironmentList{},
			Managed: &Environment{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.EnvironmentID")
	}
	mg.Spec.ForProvider.EnvironmentID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.EnvironmentIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Environment.
func (mg *Environment) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ApplicationID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.ApplicationIDRef,
		Selector:     mg.Spec.ForProvider.ApplicationIDSelector,
		To: reference.To{
			List:    &ApplicationList{},
			Managed: &Application{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ApplicationID")
	}
	mg.Spec.ForProvider.ApplicationID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ApplicationIDRef = rsp.ResolvedReference

	for i3 := 0; i3 < len(mg.Spec.ForProvider.Monitor); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Monitor[i3].AlarmArn),
			Extract:      resource.ExtractParamPath("arn", true),
			Reference:    mg.Spec.ForProvider.Monitor[i3].AlarmArnRef,
			Selector:     mg.Spec.ForProvider.Monitor[i3].AlarmArnSelector,
			To: reference.To{
				List:    &v1beta11.MetricAlarmList{},
				Managed: &v1beta11.MetricAlarm{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.Monitor[i3].AlarmArn")
		}
		mg.Spec.ForProvider.Monitor[i3].AlarmArn = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.Monitor[i3].AlarmArnRef = rsp.ResolvedReference

	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.Monitor); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Monitor[i3].AlarmRoleArn),
			Extract:      resource.ExtractParamPath("arn", true),
			Reference:    mg.Spec.ForProvider.Monitor[i3].AlarmRoleArnRef,
			Selector:     mg.Spec.ForProvider.Monitor[i3].AlarmRoleArnSelector,
			To: reference.To{
				List:    &v1beta1.RoleList{},
				Managed: &v1beta1.Role{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.Monitor[i3].AlarmRoleArn")
		}
		mg.Spec.ForProvider.Monitor[i3].AlarmRoleArn = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.Monitor[i3].AlarmRoleArnRef = rsp.ResolvedReference

	}

	return nil
}

// ResolveReferences of this Extension.
func (mg *Extension) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	for i3 := 0; i3 < len(mg.Spec.ForProvider.ActionPoint); i3++ {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.ActionPoint[i3].Action); i4++ {
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ActionPoint[i3].Action[i4].RoleArn),
				Extract:      resource.ExtractParamPath("arn", true),
				Reference:    mg.Spec.ForProvider.ActionPoint[i3].Action[i4].RoleArnRef,
				Selector:     mg.Spec.ForProvider.ActionPoint[i3].Action[i4].RoleArnSelector,
				To: reference.To{
					List:    &v1beta1.RoleList{},
					Managed: &v1beta1.Role{},
				},
			})
			if err != nil {
				return errors.Wrap(err, "mg.Spec.ForProvider.ActionPoint[i3].Action[i4].RoleArn")
			}
			mg.Spec.ForProvider.ActionPoint[i3].Action[i4].RoleArn = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.ForProvider.ActionPoint[i3].Action[i4].RoleArnRef = rsp.ResolvedReference

		}
	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.ActionPoint); i3++ {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.ActionPoint[i3].Action); i4++ {
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ActionPoint[i3].Action[i4].URI),
				Extract:      resource.ExtractParamPath("arn", true),
				Reference:    mg.Spec.ForProvider.ActionPoint[i3].Action[i4].URIRef,
				Selector:     mg.Spec.ForProvider.ActionPoint[i3].Action[i4].URISelector,
				To: reference.To{
					List:    &v1beta12.TopicList{},
					Managed: &v1beta12.Topic{},
				},
			})
			if err != nil {
				return errors.Wrap(err, "mg.Spec.ForProvider.ActionPoint[i3].Action[i4].URI")
			}
			mg.Spec.ForProvider.ActionPoint[i3].Action[i4].URI = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.ForProvider.ActionPoint[i3].Action[i4].URIRef = rsp.ResolvedReference

		}
	}

	return nil
}

// ResolveReferences of this ExtensionAssociation.
func (mg *ExtensionAssociation) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ExtensionArn),
		Extract:      resource.ExtractParamPath("arn", true),
		Reference:    mg.Spec.ForProvider.ExtensionArnRef,
		Selector:     mg.Spec.ForProvider.ExtensionArnSelector,
		To: reference.To{
			List:    &ExtensionList{},
			Managed: &Extension{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ExtensionArn")
	}
	mg.Spec.ForProvider.ExtensionArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ExtensionArnRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ResourceArn),
		Extract:      resource.ExtractParamPath("arn", true),
		Reference:    mg.Spec.ForProvider.ResourceArnRef,
		Selector:     mg.Spec.ForProvider.ResourceArnSelector,
		To: reference.To{
			List:    &ApplicationList{},
			Managed: &Application{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ResourceArn")
	}
	mg.Spec.ForProvider.ResourceArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ResourceArnRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this HostedConfigurationVersion.
func (mg *HostedConfigurationVersion) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ApplicationID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.ApplicationIDRef,
		Selector:     mg.Spec.ForProvider.ApplicationIDSelector,
		To: reference.To{
			List:    &ApplicationList{},
			Managed: &Application{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ApplicationID")
	}
	mg.Spec.ForProvider.ApplicationID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ApplicationIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ConfigurationProfileID),
		Extract:      resource.ExtractParamPath("configuration_profile_id", true),
		Reference:    mg.Spec.ForProvider.ConfigurationProfileIDRef,
		Selector:     mg.Spec.ForProvider.ConfigurationProfileIDSelector,
		To: reference.To{
			List:    &ConfigurationProfileList{},
			Managed: &ConfigurationProfile{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ConfigurationProfileID")
	}
	mg.Spec.ForProvider.ConfigurationProfileID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ConfigurationProfileIDRef = rsp.ResolvedReference

	return nil
}
