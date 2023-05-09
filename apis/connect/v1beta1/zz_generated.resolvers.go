/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1beta1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	errors "github.com/pkg/errors"
	v1beta11 "github.com/spirosco/upbound-provider-aws/apis/ds/v1beta1"
	v1beta12 "github.com/spirosco/upbound-provider-aws/apis/firehose/v1beta1"
	v1beta13 "github.com/spirosco/upbound-provider-aws/apis/kinesis/v1beta1"
	v1beta14 "github.com/spirosco/upbound-provider-aws/apis/kms/v1beta1"
	v1beta16 "github.com/spirosco/upbound-provider-aws/apis/lambda/v1beta1"
	v1beta1 "github.com/spirosco/upbound-provider-aws/apis/lexmodels/v1beta1"
	v1beta15 "github.com/spirosco/upbound-provider-aws/apis/s3/v1beta1"
	resource "github.com/upbound/upjet/pkg/resource"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this BotAssociation.
func (mg *BotAssociation) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.InstanceID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.InstanceIDRef,
		Selector:     mg.Spec.ForProvider.InstanceIDSelector,
		To: reference.To{
			List:    &InstanceList{},
			Managed: &Instance{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.InstanceID")
	}
	mg.Spec.ForProvider.InstanceID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.InstanceIDRef = rsp.ResolvedReference

	for i3 := 0; i3 < len(mg.Spec.ForProvider.LexBot); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.LexBot[i3].Name),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.LexBot[i3].NameRef,
			Selector:     mg.Spec.ForProvider.LexBot[i3].NameSelector,
			To: reference.To{
				List:    &v1beta1.BotList{},
				Managed: &v1beta1.Bot{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.LexBot[i3].Name")
		}
		mg.Spec.ForProvider.LexBot[i3].Name = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.LexBot[i3].NameRef = rsp.ResolvedReference

	}

	return nil
}

// ResolveReferences of this ContactFlow.
func (mg *ContactFlow) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.InstanceID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.InstanceIDRef,
		Selector:     mg.Spec.ForProvider.InstanceIDSelector,
		To: reference.To{
			List:    &InstanceList{},
			Managed: &Instance{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.InstanceID")
	}
	mg.Spec.ForProvider.InstanceID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.InstanceIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this ContactFlowModule.
func (mg *ContactFlowModule) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.InstanceID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.InstanceIDRef,
		Selector:     mg.Spec.ForProvider.InstanceIDSelector,
		To: reference.To{
			List:    &InstanceList{},
			Managed: &Instance{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.InstanceID")
	}
	mg.Spec.ForProvider.InstanceID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.InstanceIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this HoursOfOperation.
func (mg *HoursOfOperation) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.InstanceID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.InstanceIDRef,
		Selector:     mg.Spec.ForProvider.InstanceIDSelector,
		To: reference.To{
			List:    &InstanceList{},
			Managed: &Instance{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.InstanceID")
	}
	mg.Spec.ForProvider.InstanceID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.InstanceIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Instance.
func (mg *Instance) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.DirectoryID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.DirectoryIDRef,
		Selector:     mg.Spec.ForProvider.DirectoryIDSelector,
		To: reference.To{
			List:    &v1beta11.DirectoryList{},
			Managed: &v1beta11.Directory{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.DirectoryID")
	}
	mg.Spec.ForProvider.DirectoryID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.DirectoryIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this InstanceStorageConfig.
func (mg *InstanceStorageConfig) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.InstanceID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.InstanceIDRef,
		Selector:     mg.Spec.ForProvider.InstanceIDSelector,
		To: reference.To{
			List:    &InstanceList{},
			Managed: &Instance{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.InstanceID")
	}
	mg.Spec.ForProvider.InstanceID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.InstanceIDRef = rsp.ResolvedReference

	for i3 := 0; i3 < len(mg.Spec.ForProvider.StorageConfig); i3++ {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.StorageConfig[i3].KinesisFirehoseConfig); i4++ {
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.StorageConfig[i3].KinesisFirehoseConfig[i4].FirehoseArn),
				Extract:      resource.ExtractParamPath("arn", false),
				Reference:    mg.Spec.ForProvider.StorageConfig[i3].KinesisFirehoseConfig[i4].FirehoseArnRef,
				Selector:     mg.Spec.ForProvider.StorageConfig[i3].KinesisFirehoseConfig[i4].FirehoseArnSelector,
				To: reference.To{
					List:    &v1beta12.DeliveryStreamList{},
					Managed: &v1beta12.DeliveryStream{},
				},
			})
			if err != nil {
				return errors.Wrap(err, "mg.Spec.ForProvider.StorageConfig[i3].KinesisFirehoseConfig[i4].FirehoseArn")
			}
			mg.Spec.ForProvider.StorageConfig[i3].KinesisFirehoseConfig[i4].FirehoseArn = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.ForProvider.StorageConfig[i3].KinesisFirehoseConfig[i4].FirehoseArnRef = rsp.ResolvedReference

		}
	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.StorageConfig); i3++ {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.StorageConfig[i3].KinesisStreamConfig); i4++ {
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.StorageConfig[i3].KinesisStreamConfig[i4].StreamArn),
				Extract:      resource.ExtractParamPath("arn", false),
				Reference:    mg.Spec.ForProvider.StorageConfig[i3].KinesisStreamConfig[i4].StreamArnRef,
				Selector:     mg.Spec.ForProvider.StorageConfig[i3].KinesisStreamConfig[i4].StreamArnSelector,
				To: reference.To{
					List:    &v1beta13.StreamList{},
					Managed: &v1beta13.Stream{},
				},
			})
			if err != nil {
				return errors.Wrap(err, "mg.Spec.ForProvider.StorageConfig[i3].KinesisStreamConfig[i4].StreamArn")
			}
			mg.Spec.ForProvider.StorageConfig[i3].KinesisStreamConfig[i4].StreamArn = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.ForProvider.StorageConfig[i3].KinesisStreamConfig[i4].StreamArnRef = rsp.ResolvedReference

		}
	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.StorageConfig); i3++ {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.StorageConfig[i3].KinesisVideoStreamConfig); i4++ {
			for i5 := 0; i5 < len(mg.Spec.ForProvider.StorageConfig[i3].KinesisVideoStreamConfig[i4].EncryptionConfig); i5++ {
				rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
					CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.StorageConfig[i3].KinesisVideoStreamConfig[i4].EncryptionConfig[i5].KeyID),
					Extract:      resource.ExtractParamPath("arn", true),
					Reference:    mg.Spec.ForProvider.StorageConfig[i3].KinesisVideoStreamConfig[i4].EncryptionConfig[i5].KeyIDRef,
					Selector:     mg.Spec.ForProvider.StorageConfig[i3].KinesisVideoStreamConfig[i4].EncryptionConfig[i5].KeyIDSelector,
					To: reference.To{
						List:    &v1beta14.KeyList{},
						Managed: &v1beta14.Key{},
					},
				})
				if err != nil {
					return errors.Wrap(err, "mg.Spec.ForProvider.StorageConfig[i3].KinesisVideoStreamConfig[i4].EncryptionConfig[i5].KeyID")
				}
				mg.Spec.ForProvider.StorageConfig[i3].KinesisVideoStreamConfig[i4].EncryptionConfig[i5].KeyID = reference.ToPtrValue(rsp.ResolvedValue)
				mg.Spec.ForProvider.StorageConfig[i3].KinesisVideoStreamConfig[i4].EncryptionConfig[i5].KeyIDRef = rsp.ResolvedReference

			}
		}
	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.StorageConfig); i3++ {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.StorageConfig[i3].S3Config); i4++ {
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.StorageConfig[i3].S3Config[i4].BucketName),
				Extract:      resource.ExtractResourceID(),
				Reference:    mg.Spec.ForProvider.StorageConfig[i3].S3Config[i4].BucketNameRef,
				Selector:     mg.Spec.ForProvider.StorageConfig[i3].S3Config[i4].BucketNameSelector,
				To: reference.To{
					List:    &v1beta15.BucketList{},
					Managed: &v1beta15.Bucket{},
				},
			})
			if err != nil {
				return errors.Wrap(err, "mg.Spec.ForProvider.StorageConfig[i3].S3Config[i4].BucketName")
			}
			mg.Spec.ForProvider.StorageConfig[i3].S3Config[i4].BucketName = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.ForProvider.StorageConfig[i3].S3Config[i4].BucketNameRef = rsp.ResolvedReference

		}
	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.StorageConfig); i3++ {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.StorageConfig[i3].S3Config); i4++ {
			for i5 := 0; i5 < len(mg.Spec.ForProvider.StorageConfig[i3].S3Config[i4].EncryptionConfig); i5++ {
				rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
					CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.StorageConfig[i3].S3Config[i4].EncryptionConfig[i5].KeyID),
					Extract:      resource.ExtractParamPath("arn", true),
					Reference:    mg.Spec.ForProvider.StorageConfig[i3].S3Config[i4].EncryptionConfig[i5].KeyIDRef,
					Selector:     mg.Spec.ForProvider.StorageConfig[i3].S3Config[i4].EncryptionConfig[i5].KeyIDSelector,
					To: reference.To{
						List:    &v1beta14.KeyList{},
						Managed: &v1beta14.Key{},
					},
				})
				if err != nil {
					return errors.Wrap(err, "mg.Spec.ForProvider.StorageConfig[i3].S3Config[i4].EncryptionConfig[i5].KeyID")
				}
				mg.Spec.ForProvider.StorageConfig[i3].S3Config[i4].EncryptionConfig[i5].KeyID = reference.ToPtrValue(rsp.ResolvedValue)
				mg.Spec.ForProvider.StorageConfig[i3].S3Config[i4].EncryptionConfig[i5].KeyIDRef = rsp.ResolvedReference

			}
		}
	}

	return nil
}

// ResolveReferences of this LambdaFunctionAssociation.
func (mg *LambdaFunctionAssociation) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.FunctionArn),
		Extract:      resource.ExtractParamPath("arn", true),
		Reference:    mg.Spec.ForProvider.FunctionArnRef,
		Selector:     mg.Spec.ForProvider.FunctionArnSelector,
		To: reference.To{
			List:    &v1beta16.FunctionList{},
			Managed: &v1beta16.Function{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.FunctionArn")
	}
	mg.Spec.ForProvider.FunctionArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.FunctionArnRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.InstanceID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.InstanceIDRef,
		Selector:     mg.Spec.ForProvider.InstanceIDSelector,
		To: reference.To{
			List:    &InstanceList{},
			Managed: &Instance{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.InstanceID")
	}
	mg.Spec.ForProvider.InstanceID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.InstanceIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this PhoneNumber.
func (mg *PhoneNumber) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.TargetArn),
		Extract:      resource.ExtractParamPath("arn", true),
		Reference:    mg.Spec.ForProvider.TargetArnRef,
		Selector:     mg.Spec.ForProvider.TargetArnSelector,
		To: reference.To{
			List:    &InstanceList{},
			Managed: &Instance{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.TargetArn")
	}
	mg.Spec.ForProvider.TargetArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.TargetArnRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Queue.
func (mg *Queue) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.HoursOfOperationID),
		Extract:      resource.ExtractParamPath("hours_of_operation_id", true),
		Reference:    mg.Spec.ForProvider.HoursOfOperationIDRef,
		Selector:     mg.Spec.ForProvider.HoursOfOperationIDSelector,
		To: reference.To{
			List:    &HoursOfOperationList{},
			Managed: &HoursOfOperation{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.HoursOfOperationID")
	}
	mg.Spec.ForProvider.HoursOfOperationID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.HoursOfOperationIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.InstanceID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.InstanceIDRef,
		Selector:     mg.Spec.ForProvider.InstanceIDSelector,
		To: reference.To{
			List:    &InstanceList{},
			Managed: &Instance{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.InstanceID")
	}
	mg.Spec.ForProvider.InstanceID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.InstanceIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this QuickConnect.
func (mg *QuickConnect) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.InstanceID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.InstanceIDRef,
		Selector:     mg.Spec.ForProvider.InstanceIDSelector,
		To: reference.To{
			List:    &InstanceList{},
			Managed: &Instance{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.InstanceID")
	}
	mg.Spec.ForProvider.InstanceID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.InstanceIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this RoutingProfile.
func (mg *RoutingProfile) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.DefaultOutboundQueueID),
		Extract:      resource.ExtractParamPath("queue_id", true),
		Reference:    mg.Spec.ForProvider.DefaultOutboundQueueIDRef,
		Selector:     mg.Spec.ForProvider.DefaultOutboundQueueIDSelector,
		To: reference.To{
			List:    &QueueList{},
			Managed: &Queue{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.DefaultOutboundQueueID")
	}
	mg.Spec.ForProvider.DefaultOutboundQueueID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.DefaultOutboundQueueIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.InstanceID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.InstanceIDRef,
		Selector:     mg.Spec.ForProvider.InstanceIDSelector,
		To: reference.To{
			List:    &InstanceList{},
			Managed: &Instance{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.InstanceID")
	}
	mg.Spec.ForProvider.InstanceID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.InstanceIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this SecurityProfile.
func (mg *SecurityProfile) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.InstanceID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.InstanceIDRef,
		Selector:     mg.Spec.ForProvider.InstanceIDSelector,
		To: reference.To{
			List:    &InstanceList{},
			Managed: &Instance{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.InstanceID")
	}
	mg.Spec.ForProvider.InstanceID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.InstanceIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this User.
func (mg *User) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.InstanceID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.InstanceIDRef,
		Selector:     mg.Spec.ForProvider.InstanceIDSelector,
		To: reference.To{
			List:    &InstanceList{},
			Managed: &Instance{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.InstanceID")
	}
	mg.Spec.ForProvider.InstanceID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.InstanceIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.RoutingProfileID),
		Extract:      resource.ExtractParamPath("routing_profile_id", true),
		Reference:    mg.Spec.ForProvider.RoutingProfileIDRef,
		Selector:     mg.Spec.ForProvider.RoutingProfileIDSelector,
		To: reference.To{
			List:    &RoutingProfileList{},
			Managed: &RoutingProfile{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.RoutingProfileID")
	}
	mg.Spec.ForProvider.RoutingProfileID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.RoutingProfileIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this UserHierarchyStructure.
func (mg *UserHierarchyStructure) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.InstanceID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.InstanceIDRef,
		Selector:     mg.Spec.ForProvider.InstanceIDSelector,
		To: reference.To{
			List:    &InstanceList{},
			Managed: &Instance{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.InstanceID")
	}
	mg.Spec.ForProvider.InstanceID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.InstanceIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Vocabulary.
func (mg *Vocabulary) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.InstanceID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.InstanceIDRef,
		Selector:     mg.Spec.ForProvider.InstanceIDSelector,
		To: reference.To{
			List:    &InstanceList{},
			Managed: &Instance{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.InstanceID")
	}
	mg.Spec.ForProvider.InstanceID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.InstanceIDRef = rsp.ResolvedReference

	return nil
}
