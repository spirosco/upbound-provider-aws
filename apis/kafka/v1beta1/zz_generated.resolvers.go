/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1beta1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	errors "github.com/pkg/errors"
	v1beta12 "github.com/spirosco/upbound-provider-aws/apis/cloudwatchlogs/v1beta1"
	v1beta1 "github.com/spirosco/upbound-provider-aws/apis/ec2/v1beta1"
	v1beta13 "github.com/spirosco/upbound-provider-aws/apis/firehose/v1beta1"
	v1beta11 "github.com/spirosco/upbound-provider-aws/apis/kms/v1beta1"
	v1beta14 "github.com/spirosco/upbound-provider-aws/apis/s3/v1beta1"
	common "github.com/spirosco/upbound-provider-aws/config/common"
	resource "github.com/upbound/upjet/pkg/resource"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this Cluster.
func (mg *Cluster) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var mrsp reference.MultiResolutionResponse
	var err error

	for i3 := 0; i3 < len(mg.Spec.ForProvider.BrokerNodeGroupInfo); i3++ {
		mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
			CurrentValues: reference.FromPtrValues(mg.Spec.ForProvider.BrokerNodeGroupInfo[i3].ClientSubnets),
			Extract:       reference.ExternalName(),
			References:    mg.Spec.ForProvider.BrokerNodeGroupInfo[i3].ClientSubnetsRefs,
			Selector:      mg.Spec.ForProvider.BrokerNodeGroupInfo[i3].ClientSubnetsSelector,
			To: reference.To{
				List:    &v1beta1.SubnetList{},
				Managed: &v1beta1.Subnet{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.BrokerNodeGroupInfo[i3].ClientSubnets")
		}
		mg.Spec.ForProvider.BrokerNodeGroupInfo[i3].ClientSubnets = reference.ToPtrValues(mrsp.ResolvedValues)
		mg.Spec.ForProvider.BrokerNodeGroupInfo[i3].ClientSubnetsRefs = mrsp.ResolvedReferences

	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.BrokerNodeGroupInfo); i3++ {
		mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
			CurrentValues: reference.FromPtrValues(mg.Spec.ForProvider.BrokerNodeGroupInfo[i3].SecurityGroups),
			Extract:       reference.ExternalName(),
			References:    mg.Spec.ForProvider.BrokerNodeGroupInfo[i3].SecurityGroupsRefs,
			Selector:      mg.Spec.ForProvider.BrokerNodeGroupInfo[i3].SecurityGroupsSelector,
			To: reference.To{
				List:    &v1beta1.SecurityGroupList{},
				Managed: &v1beta1.SecurityGroup{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.BrokerNodeGroupInfo[i3].SecurityGroups")
		}
		mg.Spec.ForProvider.BrokerNodeGroupInfo[i3].SecurityGroups = reference.ToPtrValues(mrsp.ResolvedValues)
		mg.Spec.ForProvider.BrokerNodeGroupInfo[i3].SecurityGroupsRefs = mrsp.ResolvedReferences

	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.EncryptionInfo); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.EncryptionInfo[i3].EncryptionAtRestKMSKeyArn),
			Extract:      common.ARNExtractor(),
			Reference:    mg.Spec.ForProvider.EncryptionInfo[i3].EncryptionAtRestKMSKeyArnRef,
			Selector:     mg.Spec.ForProvider.EncryptionInfo[i3].EncryptionAtRestKMSKeyArnSelector,
			To: reference.To{
				List:    &v1beta11.KeyList{},
				Managed: &v1beta11.Key{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.EncryptionInfo[i3].EncryptionAtRestKMSKeyArn")
		}
		mg.Spec.ForProvider.EncryptionInfo[i3].EncryptionAtRestKMSKeyArn = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.EncryptionInfo[i3].EncryptionAtRestKMSKeyArnRef = rsp.ResolvedReference

	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.LoggingInfo); i3++ {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.LoggingInfo[i3].BrokerLogs); i4++ {
			for i5 := 0; i5 < len(mg.Spec.ForProvider.LoggingInfo[i3].BrokerLogs[i4].CloudwatchLogs); i5++ {
				rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
					CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.LoggingInfo[i3].BrokerLogs[i4].CloudwatchLogs[i5].LogGroup),
					Extract:      reference.ExternalName(),
					Reference:    mg.Spec.ForProvider.LoggingInfo[i3].BrokerLogs[i4].CloudwatchLogs[i5].LogGroupRef,
					Selector:     mg.Spec.ForProvider.LoggingInfo[i3].BrokerLogs[i4].CloudwatchLogs[i5].LogGroupSelector,
					To: reference.To{
						List:    &v1beta12.GroupList{},
						Managed: &v1beta12.Group{},
					},
				})
				if err != nil {
					return errors.Wrap(err, "mg.Spec.ForProvider.LoggingInfo[i3].BrokerLogs[i4].CloudwatchLogs[i5].LogGroup")
				}
				mg.Spec.ForProvider.LoggingInfo[i3].BrokerLogs[i4].CloudwatchLogs[i5].LogGroup = reference.ToPtrValue(rsp.ResolvedValue)
				mg.Spec.ForProvider.LoggingInfo[i3].BrokerLogs[i4].CloudwatchLogs[i5].LogGroupRef = rsp.ResolvedReference

			}
		}
	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.LoggingInfo); i3++ {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.LoggingInfo[i3].BrokerLogs); i4++ {
			for i5 := 0; i5 < len(mg.Spec.ForProvider.LoggingInfo[i3].BrokerLogs[i4].Firehose); i5++ {
				rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
					CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.LoggingInfo[i3].BrokerLogs[i4].Firehose[i5].DeliveryStream),
					Extract:      resource.ExtractParamPath("name", false),
					Reference:    mg.Spec.ForProvider.LoggingInfo[i3].BrokerLogs[i4].Firehose[i5].DeliveryStreamRef,
					Selector:     mg.Spec.ForProvider.LoggingInfo[i3].BrokerLogs[i4].Firehose[i5].DeliveryStreamSelector,
					To: reference.To{
						List:    &v1beta13.DeliveryStreamList{},
						Managed: &v1beta13.DeliveryStream{},
					},
				})
				if err != nil {
					return errors.Wrap(err, "mg.Spec.ForProvider.LoggingInfo[i3].BrokerLogs[i4].Firehose[i5].DeliveryStream")
				}
				mg.Spec.ForProvider.LoggingInfo[i3].BrokerLogs[i4].Firehose[i5].DeliveryStream = reference.ToPtrValue(rsp.ResolvedValue)
				mg.Spec.ForProvider.LoggingInfo[i3].BrokerLogs[i4].Firehose[i5].DeliveryStreamRef = rsp.ResolvedReference

			}
		}
	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.LoggingInfo); i3++ {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.LoggingInfo[i3].BrokerLogs); i4++ {
			for i5 := 0; i5 < len(mg.Spec.ForProvider.LoggingInfo[i3].BrokerLogs[i4].S3); i5++ {
				rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
					CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.LoggingInfo[i3].BrokerLogs[i4].S3[i5].Bucket),
					Extract:      reference.ExternalName(),
					Reference:    mg.Spec.ForProvider.LoggingInfo[i3].BrokerLogs[i4].S3[i5].BucketRef,
					Selector:     mg.Spec.ForProvider.LoggingInfo[i3].BrokerLogs[i4].S3[i5].BucketSelector,
					To: reference.To{
						List:    &v1beta14.BucketList{},
						Managed: &v1beta14.Bucket{},
					},
				})
				if err != nil {
					return errors.Wrap(err, "mg.Spec.ForProvider.LoggingInfo[i3].BrokerLogs[i4].S3[i5].Bucket")
				}
				mg.Spec.ForProvider.LoggingInfo[i3].BrokerLogs[i4].S3[i5].Bucket = reference.ToPtrValue(rsp.ResolvedValue)
				mg.Spec.ForProvider.LoggingInfo[i3].BrokerLogs[i4].S3[i5].BucketRef = rsp.ResolvedReference

			}
		}
	}

	return nil
}
