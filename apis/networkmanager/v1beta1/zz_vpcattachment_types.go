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

type VPCAttachmentObservation struct {

	// The ARN of the attachment.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// The policy rule number associated with the attachment.
	AttachmentPolicyRuleNumber *float64 `json:"attachmentPolicyRuleNumber,omitempty" tf:"attachment_policy_rule_number,omitempty"`

	// The type of attachment.
	AttachmentType *string `json:"attachmentType,omitempty" tf:"attachment_type,omitempty"`

	// The ARN of a core network.
	CoreNetworkArn *string `json:"coreNetworkArn,omitempty" tf:"core_network_arn,omitempty"`

	// The ID of a core network for the VPC attachment.
	CoreNetworkID *string `json:"coreNetworkId,omitempty" tf:"core_network_id,omitempty"`

	// The Region where the edge is located.
	EdgeLocation *string `json:"edgeLocation,omitempty" tf:"edge_location,omitempty"`

	// The ID of the attachment.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Options for the VPC attachment.
	Options []VPCAttachmentOptionsObservation `json:"options,omitempty" tf:"options,omitempty"`

	// The ID of the attachment account owner.
	OwnerAccountID *string `json:"ownerAccountId,omitempty" tf:"owner_account_id,omitempty"`

	// The attachment resource ARN.
	ResourceArn *string `json:"resourceArn,omitempty" tf:"resource_arn,omitempty"`

	// The name of the segment attachment.
	SegmentName *string `json:"segmentName,omitempty" tf:"segment_name,omitempty"`

	// The state of the attachment.
	State *string `json:"state,omitempty" tf:"state,omitempty"`

	// The subnet ARN of the VPC attachment.
	SubnetArns []*string `json:"subnetArns,omitempty" tf:"subnet_arns,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`

	// The ARN of the VPC.
	VPCArn *string `json:"vpcArn,omitempty" tf:"vpc_arn,omitempty"`
}

type VPCAttachmentOptionsObservation struct {

	// Indicates whether appliance mode is supported. If enabled, traffic flow between a source and destination use the same Availability Zone for the VPC attachment for the lifetime of that flow.
	ApplianceModeSupport *bool `json:"applianceModeSupport,omitempty" tf:"appliance_mode_support,omitempty"`

	// Indicates whether IPv6 is supported.
	IPv6Support *bool `json:"ipv6Support,omitempty" tf:"ipv6_support,omitempty"`
}

type VPCAttachmentOptionsParameters struct {

	// Indicates whether appliance mode is supported. If enabled, traffic flow between a source and destination use the same Availability Zone for the VPC attachment for the lifetime of that flow.
	// +kubebuilder:validation:Optional
	ApplianceModeSupport *bool `json:"applianceModeSupport,omitempty" tf:"appliance_mode_support,omitempty"`

	// Indicates whether IPv6 is supported.
	// +kubebuilder:validation:Optional
	IPv6Support *bool `json:"ipv6Support,omitempty" tf:"ipv6_support,omitempty"`
}

type VPCAttachmentParameters struct {

	// The ID of a core network for the VPC attachment.
	// +crossplane:generate:reference:type=CoreNetwork
	// +kubebuilder:validation:Optional
	CoreNetworkID *string `json:"coreNetworkId,omitempty" tf:"core_network_id,omitempty"`

	// Reference to a CoreNetwork to populate coreNetworkId.
	// +kubebuilder:validation:Optional
	CoreNetworkIDRef *v1.Reference `json:"coreNetworkIdRef,omitempty" tf:"-"`

	// Selector for a CoreNetwork to populate coreNetworkId.
	// +kubebuilder:validation:Optional
	CoreNetworkIDSelector *v1.Selector `json:"coreNetworkIdSelector,omitempty" tf:"-"`

	// Options for the VPC attachment.
	// +kubebuilder:validation:Optional
	Options []VPCAttachmentOptionsParameters `json:"options,omitempty" tf:"options,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// The subnet ARN of the VPC attachment.
	// +crossplane:generate:reference:type=github.com/spirosco/upbound-provider-aws/apis/ec2/v1beta1.Subnet
	// +crossplane:generate:reference:extractor=github.com/spirosco/upbound-provider-aws/config/common.ARNExtractor()
	// +kubebuilder:validation:Optional
	SubnetArns []*string `json:"subnetArns,omitempty" tf:"subnet_arns,omitempty"`

	// References to Subnet in ec2 to populate subnetArns.
	// +kubebuilder:validation:Optional
	SubnetArnsRefs []v1.Reference `json:"subnetArnsRefs,omitempty" tf:"-"`

	// Selector for a list of Subnet in ec2 to populate subnetArns.
	// +kubebuilder:validation:Optional
	SubnetArnsSelector *v1.Selector `json:"subnetArnsSelector,omitempty" tf:"-"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The ARN of the VPC.
	// +crossplane:generate:reference:type=github.com/spirosco/upbound-provider-aws/apis/ec2/v1beta1.VPC
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractParamPath("arn",true)
	// +kubebuilder:validation:Optional
	VPCArn *string `json:"vpcArn,omitempty" tf:"vpc_arn,omitempty"`

	// Reference to a VPC in ec2 to populate vpcArn.
	// +kubebuilder:validation:Optional
	VPCArnRef *v1.Reference `json:"vpcArnRef,omitempty" tf:"-"`

	// Selector for a VPC in ec2 to populate vpcArn.
	// +kubebuilder:validation:Optional
	VPCArnSelector *v1.Selector `json:"vpcArnSelector,omitempty" tf:"-"`
}

// VPCAttachmentSpec defines the desired state of VPCAttachment
type VPCAttachmentSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     VPCAttachmentParameters `json:"forProvider"`
}

// VPCAttachmentStatus defines the observed state of VPCAttachment.
type VPCAttachmentStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        VPCAttachmentObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// VPCAttachment is the Schema for the VPCAttachments API.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type VPCAttachment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VPCAttachmentSpec   `json:"spec"`
	Status            VPCAttachmentStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VPCAttachmentList contains a list of VPCAttachments
type VPCAttachmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VPCAttachment `json:"items"`
}

// Repository type metadata.
var (
	VPCAttachment_Kind             = "VPCAttachment"
	VPCAttachment_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: VPCAttachment_Kind}.String()
	VPCAttachment_KindAPIVersion   = VPCAttachment_Kind + "." + CRDGroupVersion.String()
	VPCAttachment_GroupVersionKind = CRDGroupVersion.WithKind(VPCAttachment_Kind)
)

func init() {
	SchemeBuilder.Register(&VPCAttachment{}, &VPCAttachmentList{})
}
