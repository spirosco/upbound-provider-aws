/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1beta1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	errors "github.com/pkg/errors"
	v1beta1 "github.com/spirosco/upbound-provider-aws/apis/cloudwatch/v1beta1"
	v1beta11 "github.com/spirosco/upbound-provider-aws/apis/ec2/v1beta1"
	resource "github.com/upbound/upjet/pkg/resource"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this HealthCheck.
func (mg *HealthCheck) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.CloudwatchAlarmName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.CloudwatchAlarmNameRef,
		Selector:     mg.Spec.ForProvider.CloudwatchAlarmNameSelector,
		To: reference.To{
			List:    &v1beta1.MetricAlarmList{},
			Managed: &v1beta1.MetricAlarm{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.CloudwatchAlarmName")
	}
	mg.Spec.ForProvider.CloudwatchAlarmName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.CloudwatchAlarmNameRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this HostedZoneDNSSEC.
func (mg *HostedZoneDNSSEC) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.HostedZoneID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.HostedZoneIDRef,
		Selector:     mg.Spec.ForProvider.HostedZoneIDSelector,
		To: reference.To{
			List:    &ZoneList{},
			Managed: &Zone{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.HostedZoneID")
	}
	mg.Spec.ForProvider.HostedZoneID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.HostedZoneIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Record.
func (mg *Record) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.HealthCheckID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.HealthCheckIDRef,
		Selector:     mg.Spec.ForProvider.HealthCheckIDSelector,
		To: reference.To{
			List:    &HealthCheckList{},
			Managed: &HealthCheck{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.HealthCheckID")
	}
	mg.Spec.ForProvider.HealthCheckID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.HealthCheckIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ZoneID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ZoneIDRef,
		Selector:     mg.Spec.ForProvider.ZoneIDSelector,
		To: reference.To{
			List:    &ZoneList{},
			Managed: &Zone{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ZoneID")
	}
	mg.Spec.ForProvider.ZoneID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ZoneIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this ResolverConfig.
func (mg *ResolverConfig) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ResourceID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.ResourceIDRef,
		Selector:     mg.Spec.ForProvider.ResourceIDSelector,
		To: reference.To{
			List:    &v1beta11.VPCList{},
			Managed: &v1beta11.VPC{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ResourceID")
	}
	mg.Spec.ForProvider.ResourceID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ResourceIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this TrafficPolicyInstance.
func (mg *TrafficPolicyInstance) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.HostedZoneID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.HostedZoneIDRef,
		Selector:     mg.Spec.ForProvider.HostedZoneIDSelector,
		To: reference.To{
			List:    &ZoneList{},
			Managed: &Zone{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.HostedZoneID")
	}
	mg.Spec.ForProvider.HostedZoneID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.HostedZoneIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.TrafficPolicyID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.TrafficPolicyIDRef,
		Selector:     mg.Spec.ForProvider.TrafficPolicyIDSelector,
		To: reference.To{
			List:    &TrafficPolicyList{},
			Managed: &TrafficPolicy{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.TrafficPolicyID")
	}
	mg.Spec.ForProvider.TrafficPolicyID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.TrafficPolicyIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this VPCAssociationAuthorization.
func (mg *VPCAssociationAuthorization) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.VPCID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.VPCIDRef,
		Selector:     mg.Spec.ForProvider.VPCIDSelector,
		To: reference.To{
			List:    &v1beta11.VPCList{},
			Managed: &v1beta11.VPC{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.VPCID")
	}
	mg.Spec.ForProvider.VPCID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.VPCIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ZoneID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ZoneIDRef,
		Selector:     mg.Spec.ForProvider.ZoneIDSelector,
		To: reference.To{
			List:    &ZoneList{},
			Managed: &Zone{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ZoneID")
	}
	mg.Spec.ForProvider.ZoneID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ZoneIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Zone.
func (mg *Zone) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.DelegationSetID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.DelegationSetIDRef,
		Selector:     mg.Spec.ForProvider.DelegationSetIDSelector,
		To: reference.To{
			List:    &DelegationSetList{},
			Managed: &DelegationSet{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.DelegationSetID")
	}
	mg.Spec.ForProvider.DelegationSetID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.DelegationSetIDRef = rsp.ResolvedReference

	for i3 := 0; i3 < len(mg.Spec.ForProvider.VPC); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.VPC[i3].VPCID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.VPC[i3].VPCIDRef,
			Selector:     mg.Spec.ForProvider.VPC[i3].VPCIDSelector,
			To: reference.To{
				List:    &v1beta11.VPCList{},
				Managed: &v1beta11.VPC{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.VPC[i3].VPCID")
		}
		mg.Spec.ForProvider.VPC[i3].VPCID = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.VPC[i3].VPCIDRef = rsp.ResolvedReference

	}

	return nil
}
