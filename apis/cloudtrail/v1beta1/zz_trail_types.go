/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type AdvancedEventSelectorObservation struct {

	// Specifies the selector statements in an advanced event selector. Fields documented below.
	FieldSelector []FieldSelectorObservation `json:"fieldSelector,omitempty" tf:"field_selector,omitempty"`

	// Name of the trail.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type AdvancedEventSelectorParameters struct {

	// Specifies the selector statements in an advanced event selector. Fields documented below.
	// +kubebuilder:validation:Required
	FieldSelector []FieldSelectorParameters `json:"fieldSelector" tf:"field_selector,omitempty"`

	// Name of the trail.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type DataResourceObservation struct {

	// Resource type in which you want to log data events. You can specify only the following value: "AWS::S3::Object", "AWS::Lambda::Function" and "AWS::DynamoDB::Table".
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// List of ARN strings or partial ARN strings to specify selectors for data audit events over data resources. ARN list is specific to single-valued type. For example, arn:aws:s3:::<bucket name>/ for all objects in a bucket, arn:aws:s3:::<bucket name>/key for specific objects, arn:aws:lambda for all lambda events within an account, arn:aws:lambda:<region>:<account number>:function:<function name> for a specific Lambda function, arn:aws:dynamodb for all DDB events for all tables within an account, or arn:aws:dynamodb:<region>:<account number>:table/<table name> for a specific DynamoDB table.
	Values []*string `json:"values,omitempty" tf:"values,omitempty"`
}

type DataResourceParameters struct {

	// Resource type in which you want to log data events. You can specify only the following value: "AWS::S3::Object", "AWS::Lambda::Function" and "AWS::DynamoDB::Table".
	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`

	// List of ARN strings or partial ARN strings to specify selectors for data audit events over data resources. ARN list is specific to single-valued type. For example, arn:aws:s3:::<bucket name>/ for all objects in a bucket, arn:aws:s3:::<bucket name>/key for specific objects, arn:aws:lambda for all lambda events within an account, arn:aws:lambda:<region>:<account number>:function:<function name> for a specific Lambda function, arn:aws:dynamodb for all DDB events for all tables within an account, or arn:aws:dynamodb:<region>:<account number>:table/<table name> for a specific DynamoDB table.
	// +kubebuilder:validation:Required
	Values []*string `json:"values" tf:"values,omitempty"`
}

type EventSelectorObservation struct {

	// Configuration block for data events. See details below.
	DataResource []DataResourceObservation `json:"dataResource,omitempty" tf:"data_resource,omitempty"`

	// A set of event sources to exclude. Valid values include: kms.amazonaws.com and rdsdata.amazonaws.com. include_management_events must be set totrue to allow this.
	ExcludeManagementEventSources []*string `json:"excludeManagementEventSources,omitempty" tf:"exclude_management_event_sources,omitempty"`

	// Whether to include management events for your trail. Defaults to true.
	IncludeManagementEvents *bool `json:"includeManagementEvents,omitempty" tf:"include_management_events,omitempty"`

	// Type of events to log. Valid values are ReadOnly, WriteOnly, All. Default value is All.
	ReadWriteType *string `json:"readWriteType,omitempty" tf:"read_write_type,omitempty"`
}

type EventSelectorParameters struct {

	// Configuration block for data events. See details below.
	// +kubebuilder:validation:Optional
	DataResource []DataResourceParameters `json:"dataResource,omitempty" tf:"data_resource,omitempty"`

	// A set of event sources to exclude. Valid values include: kms.amazonaws.com and rdsdata.amazonaws.com. include_management_events must be set totrue to allow this.
	// +kubebuilder:validation:Optional
	ExcludeManagementEventSources []*string `json:"excludeManagementEventSources,omitempty" tf:"exclude_management_event_sources,omitempty"`

	// Whether to include management events for your trail. Defaults to true.
	// +kubebuilder:validation:Optional
	IncludeManagementEvents *bool `json:"includeManagementEvents,omitempty" tf:"include_management_events,omitempty"`

	// Type of events to log. Valid values are ReadOnly, WriteOnly, All. Default value is All.
	// +kubebuilder:validation:Optional
	ReadWriteType *string `json:"readWriteType,omitempty" tf:"read_write_type,omitempty"`
}

type FieldSelectorObservation struct {

	// A list of values that includes events that match the last few characters of the event record field specified as the value of field.
	EndsWith []*string `json:"endsWith,omitempty" tf:"ends_with,omitempty"`

	// A list of values that includes events that match the exact value of the event record field specified as the value of field. This is the only valid operator that you can use with the readOnly, eventCategory, and resources.type fields.
	Equals []*string `json:"equals,omitempty" tf:"equals,omitempty"`

	// Field in an event record on which to filter events to be logged. You can specify only the following values: readOnly, eventSource, eventName, eventCategory, resources.type, resources.ARN.
	Field *string `json:"field,omitempty" tf:"field,omitempty"`

	// A list of values that excludes events that match the last few characters of the event record field specified as the value of field.
	NotEndsWith []*string `json:"notEndsWith,omitempty" tf:"not_ends_with,omitempty"`

	// A list of values that excludes events that match the exact value of the event record field specified as the value of field.
	NotEquals []*string `json:"notEquals,omitempty" tf:"not_equals,omitempty"`

	// A list of values that excludes events that match the first few characters of the event record field specified as the value of field.
	NotStartsWith []*string `json:"notStartsWith,omitempty" tf:"not_starts_with,omitempty"`

	// A list of values that includes events that match the first few characters of the event record field specified as the value of field.
	StartsWith []*string `json:"startsWith,omitempty" tf:"starts_with,omitempty"`
}

type FieldSelectorParameters struct {

	// A list of values that includes events that match the last few characters of the event record field specified as the value of field.
	// +kubebuilder:validation:Optional
	EndsWith []*string `json:"endsWith,omitempty" tf:"ends_with,omitempty"`

	// A list of values that includes events that match the exact value of the event record field specified as the value of field. This is the only valid operator that you can use with the readOnly, eventCategory, and resources.type fields.
	// +kubebuilder:validation:Optional
	Equals []*string `json:"equals,omitempty" tf:"equals,omitempty"`

	// Field in an event record on which to filter events to be logged. You can specify only the following values: readOnly, eventSource, eventName, eventCategory, resources.type, resources.ARN.
	// +kubebuilder:validation:Required
	Field *string `json:"field" tf:"field,omitempty"`

	// A list of values that excludes events that match the last few characters of the event record field specified as the value of field.
	// +kubebuilder:validation:Optional
	NotEndsWith []*string `json:"notEndsWith,omitempty" tf:"not_ends_with,omitempty"`

	// A list of values that excludes events that match the exact value of the event record field specified as the value of field.
	// +kubebuilder:validation:Optional
	NotEquals []*string `json:"notEquals,omitempty" tf:"not_equals,omitempty"`

	// A list of values that excludes events that match the first few characters of the event record field specified as the value of field.
	// +kubebuilder:validation:Optional
	NotStartsWith []*string `json:"notStartsWith,omitempty" tf:"not_starts_with,omitempty"`

	// A list of values that includes events that match the first few characters of the event record field specified as the value of field.
	// +kubebuilder:validation:Optional
	StartsWith []*string `json:"startsWith,omitempty" tf:"starts_with,omitempty"`
}

type InsightSelectorObservation struct {

	// Type of insights to log on a trail. Valid values are: ApiCallRateInsight and ApiErrorRateInsight.
	InsightType *string `json:"insightType,omitempty" tf:"insight_type,omitempty"`
}

type InsightSelectorParameters struct {

	// Type of insights to log on a trail. Valid values are: ApiCallRateInsight and ApiErrorRateInsight.
	// +kubebuilder:validation:Required
	InsightType *string `json:"insightType" tf:"insight_type,omitempty"`
}

type TrailObservation struct {

	// Specifies an advanced event selector for enabling data event logging. Fields documented below. Conflicts with event_selector.
	AdvancedEventSelector []AdvancedEventSelectorObservation `json:"advancedEventSelector,omitempty" tf:"advanced_event_selector,omitempty"`

	// ARN of the trail.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// Log group name using an ARN that represents the log group to which CloudTrail logs will be delivered. Note that CloudTrail requires the Log Stream wildcard.
	CloudWatchLogsGroupArn *string `json:"cloudWatchLogsGroupArn,omitempty" tf:"cloud_watch_logs_group_arn,omitempty"`

	// Role for the CloudWatch Logs endpoint to assume to write to a user’s log group.
	CloudWatchLogsRoleArn *string `json:"cloudWatchLogsRoleArn,omitempty" tf:"cloud_watch_logs_role_arn,omitempty"`

	// Whether log file integrity validation is enabled. Defaults to false.
	EnableLogFileValidation *bool `json:"enableLogFileValidation,omitempty" tf:"enable_log_file_validation,omitempty"`

	// Enables logging for the trail. Defaults to true. Setting this to false will pause logging.
	EnableLogging *bool `json:"enableLogging,omitempty" tf:"enable_logging,omitempty"`

	// Specifies an event selector for enabling data event logging. Fields documented below. Please note the CloudTrail limits when configuring these. Conflicts with advanced_event_selector.
	EventSelector []EventSelectorObservation `json:"eventSelector,omitempty" tf:"event_selector,omitempty"`

	// Region in which the trail was created.
	HomeRegion *string `json:"homeRegion,omitempty" tf:"home_region,omitempty"`

	// Name of the trail.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Whether the trail is publishing events from global services such as IAM to the log files. Defaults to true.
	IncludeGlobalServiceEvents *bool `json:"includeGlobalServiceEvents,omitempty" tf:"include_global_service_events,omitempty"`

	// Configuration block for identifying unusual operational activity. See details below.
	InsightSelector []InsightSelectorObservation `json:"insightSelector,omitempty" tf:"insight_selector,omitempty"`

	// Whether the trail is created in the current region or in all regions. Defaults to false.
	IsMultiRegionTrail *bool `json:"isMultiRegionTrail,omitempty" tf:"is_multi_region_trail,omitempty"`

	// Whether the trail is an AWS Organizations trail. Organization trails log events for the master account and all member accounts. Can only be created in the organization master account. Defaults to false.
	IsOrganizationTrail *bool `json:"isOrganizationTrail,omitempty" tf:"is_organization_trail,omitempty"`

	// KMS key ARN to use to encrypt the logs delivered by CloudTrail.
	KMSKeyID *string `json:"kmsKeyId,omitempty" tf:"kms_key_id,omitempty"`

	// Name of the S3 bucket designated for publishing log files.
	S3BucketName *string `json:"s3BucketName,omitempty" tf:"s3_bucket_name,omitempty"`

	// S3 key prefix that follows the name of the bucket you have designated for log file delivery.
	S3KeyPrefix *string `json:"s3KeyPrefix,omitempty" tf:"s3_key_prefix,omitempty"`

	// Name of the Amazon SNS topic defined for notification of log file delivery.
	SnsTopicName *string `json:"snsTopicName,omitempty" tf:"sns_topic_name,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type TrailParameters struct {

	// Specifies an advanced event selector for enabling data event logging. Fields documented below. Conflicts with event_selector.
	// +kubebuilder:validation:Optional
	AdvancedEventSelector []AdvancedEventSelectorParameters `json:"advancedEventSelector,omitempty" tf:"advanced_event_selector,omitempty"`

	// Log group name using an ARN that represents the log group to which CloudTrail logs will be delivered. Note that CloudTrail requires the Log Stream wildcard.
	// +kubebuilder:validation:Optional
	CloudWatchLogsGroupArn *string `json:"cloudWatchLogsGroupArn,omitempty" tf:"cloud_watch_logs_group_arn,omitempty"`

	// Role for the CloudWatch Logs endpoint to assume to write to a user’s log group.
	// +crossplane:generate:reference:type=github.com/spirosco/upbound-provider-aws/apis/iam/v1beta1.Role
	// +crossplane:generate:reference:extractor=github.com/spirosco/upbound-provider-aws/config/common.ARNExtractor()
	// +kubebuilder:validation:Optional
	CloudWatchLogsRoleArn *string `json:"cloudWatchLogsRoleArn,omitempty" tf:"cloud_watch_logs_role_arn,omitempty"`

	// Reference to a Role in iam to populate cloudWatchLogsRoleArn.
	// +kubebuilder:validation:Optional
	CloudWatchLogsRoleArnRef *v1.Reference `json:"cloudWatchLogsRoleArnRef,omitempty" tf:"-"`

	// Selector for a Role in iam to populate cloudWatchLogsRoleArn.
	// +kubebuilder:validation:Optional
	CloudWatchLogsRoleArnSelector *v1.Selector `json:"cloudWatchLogsRoleArnSelector,omitempty" tf:"-"`

	// Whether log file integrity validation is enabled. Defaults to false.
	// +kubebuilder:validation:Optional
	EnableLogFileValidation *bool `json:"enableLogFileValidation,omitempty" tf:"enable_log_file_validation,omitempty"`

	// Enables logging for the trail. Defaults to true. Setting this to false will pause logging.
	// +kubebuilder:validation:Optional
	EnableLogging *bool `json:"enableLogging,omitempty" tf:"enable_logging,omitempty"`

	// Specifies an event selector for enabling data event logging. Fields documented below. Please note the CloudTrail limits when configuring these. Conflicts with advanced_event_selector.
	// +kubebuilder:validation:Optional
	EventSelector []EventSelectorParameters `json:"eventSelector,omitempty" tf:"event_selector,omitempty"`

	// Whether the trail is publishing events from global services such as IAM to the log files. Defaults to true.
	// +kubebuilder:validation:Optional
	IncludeGlobalServiceEvents *bool `json:"includeGlobalServiceEvents,omitempty" tf:"include_global_service_events,omitempty"`

	// Configuration block for identifying unusual operational activity. See details below.
	// +kubebuilder:validation:Optional
	InsightSelector []InsightSelectorParameters `json:"insightSelector,omitempty" tf:"insight_selector,omitempty"`

	// Whether the trail is created in the current region or in all regions. Defaults to false.
	// +kubebuilder:validation:Optional
	IsMultiRegionTrail *bool `json:"isMultiRegionTrail,omitempty" tf:"is_multi_region_trail,omitempty"`

	// Whether the trail is an AWS Organizations trail. Organization trails log events for the master account and all member accounts. Can only be created in the organization master account. Defaults to false.
	// +kubebuilder:validation:Optional
	IsOrganizationTrail *bool `json:"isOrganizationTrail,omitempty" tf:"is_organization_trail,omitempty"`

	// KMS key ARN to use to encrypt the logs delivered by CloudTrail.
	// +crossplane:generate:reference:type=github.com/spirosco/upbound-provider-aws/apis/kms/v1beta1.Key
	// +kubebuilder:validation:Optional
	KMSKeyID *string `json:"kmsKeyId,omitempty" tf:"kms_key_id,omitempty"`

	// Reference to a Key in kms to populate kmsKeyId.
	// +kubebuilder:validation:Optional
	KMSKeyIDRef *v1.Reference `json:"kmsKeyIdRef,omitempty" tf:"-"`

	// Selector for a Key in kms to populate kmsKeyId.
	// +kubebuilder:validation:Optional
	KMSKeyIDSelector *v1.Selector `json:"kmsKeyIdSelector,omitempty" tf:"-"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Name of the S3 bucket designated for publishing log files.
	// +crossplane:generate:reference:type=github.com/spirosco/upbound-provider-aws/apis/s3/v1beta1.Bucket
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	S3BucketName *string `json:"s3BucketName,omitempty" tf:"s3_bucket_name,omitempty"`

	// Reference to a Bucket in s3 to populate s3BucketName.
	// +kubebuilder:validation:Optional
	S3BucketNameRef *v1.Reference `json:"s3BucketNameRef,omitempty" tf:"-"`

	// Selector for a Bucket in s3 to populate s3BucketName.
	// +kubebuilder:validation:Optional
	S3BucketNameSelector *v1.Selector `json:"s3BucketNameSelector,omitempty" tf:"-"`

	// S3 key prefix that follows the name of the bucket you have designated for log file delivery.
	// +kubebuilder:validation:Optional
	S3KeyPrefix *string `json:"s3KeyPrefix,omitempty" tf:"s3_key_prefix,omitempty"`

	// Name of the Amazon SNS topic defined for notification of log file delivery.
	// +kubebuilder:validation:Optional
	SnsTopicName *string `json:"snsTopicName,omitempty" tf:"sns_topic_name,omitempty"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// TrailSpec defines the desired state of Trail
type TrailSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     TrailParameters `json:"forProvider"`
}

// TrailStatus defines the observed state of Trail.
type TrailStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        TrailObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Trail is the Schema for the Trails API. Provides a CloudTrail resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Trail struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              TrailSpec   `json:"spec"`
	Status            TrailStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TrailList contains a list of Trails
type TrailList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Trail `json:"items"`
}

// Repository type metadata.
var (
	Trail_Kind             = "Trail"
	Trail_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Trail_Kind}.String()
	Trail_KindAPIVersion   = Trail_Kind + "." + CRDGroupVersion.String()
	Trail_GroupVersionKind = CRDGroupVersion.WithKind(Trail_Kind)
)

func init() {
	SchemeBuilder.Register(&Trail{}, &TrailList{})
}
