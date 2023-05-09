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

type VPCConnectorObservation struct {

	// ARN of VPC connector.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// List of IDs of security groups that App Runner should use for access to AWS resources under the specified subnets. If not specified, App Runner uses the default security group of the Amazon VPC. The default security group allows all outbound traffic.
	SecurityGroups []*string `json:"securityGroups,omitempty" tf:"security_groups,omitempty"`

	// Current state of the VPC connector. If the status of a connector revision is INACTIVE, it was deleted and can't be used. Inactive connector revisions are permanently removed some time after they are deleted.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// List of IDs of subnets that App Runner should use when it associates your service with a custom Amazon VPC. Specify IDs of subnets of a single Amazon VPC. App Runner determines the Amazon VPC from the subnets you specify.
	Subnets []*string `json:"subnets,omitempty" tf:"subnets,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`

	// Name for the VPC connector.
	VPCConnectorName *string `json:"vpcConnectorName,omitempty" tf:"vpc_connector_name,omitempty"`

	// The revision of VPC connector. It's unique among all the active connectors ("Status": "ACTIVE") that share the same Name.
	VPCConnectorRevision *float64 `json:"vpcConnectorRevision,omitempty" tf:"vpc_connector_revision,omitempty"`
}

type VPCConnectorParameters struct {

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// References to SecurityGroup in ec2 to populate securityGroups.
	// +kubebuilder:validation:Optional
	SecurityGroupRefs []v1.Reference `json:"securityGroupRefs,omitempty" tf:"-"`

	// Selector for a list of SecurityGroup in ec2 to populate securityGroups.
	// +kubebuilder:validation:Optional
	SecurityGroupSelector *v1.Selector `json:"securityGroupSelector,omitempty" tf:"-"`

	// List of IDs of security groups that App Runner should use for access to AWS resources under the specified subnets. If not specified, App Runner uses the default security group of the Amazon VPC. The default security group allows all outbound traffic.
	// +crossplane:generate:reference:type=github.com/spirosco/upbound-provider-aws/apis/ec2/v1beta1.SecurityGroup
	// +crossplane:generate:reference:refFieldName=SecurityGroupRefs
	// +crossplane:generate:reference:selectorFieldName=SecurityGroupSelector
	// +kubebuilder:validation:Optional
	SecurityGroups []*string `json:"securityGroups,omitempty" tf:"security_groups,omitempty"`

	// References to Subnet in ec2 to populate subnets.
	// +kubebuilder:validation:Optional
	SubnetRefs []v1.Reference `json:"subnetRefs,omitempty" tf:"-"`

	// Selector for a list of Subnet in ec2 to populate subnets.
	// +kubebuilder:validation:Optional
	SubnetSelector *v1.Selector `json:"subnetSelector,omitempty" tf:"-"`

	// List of IDs of subnets that App Runner should use when it associates your service with a custom Amazon VPC. Specify IDs of subnets of a single Amazon VPC. App Runner determines the Amazon VPC from the subnets you specify.
	// +crossplane:generate:reference:type=github.com/spirosco/upbound-provider-aws/apis/ec2/v1beta1.Subnet
	// +crossplane:generate:reference:refFieldName=SubnetRefs
	// +crossplane:generate:reference:selectorFieldName=SubnetSelector
	// +kubebuilder:validation:Optional
	Subnets []*string `json:"subnets,omitempty" tf:"subnets,omitempty"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Name for the VPC connector.
	// +kubebuilder:validation:Optional
	VPCConnectorName *string `json:"vpcConnectorName,omitempty" tf:"vpc_connector_name,omitempty"`
}

// VPCConnectorSpec defines the desired state of VPCConnector
type VPCConnectorSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     VPCConnectorParameters `json:"forProvider"`
}

// VPCConnectorStatus defines the observed state of VPCConnector.
type VPCConnectorStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        VPCConnectorObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// VPCConnector is the Schema for the VPCConnectors API. Manages an App Runner VPC Connector.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type VPCConnector struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.vpcConnectorName)",message="vpcConnectorName is a required parameter"
	Spec   VPCConnectorSpec   `json:"spec"`
	Status VPCConnectorStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VPCConnectorList contains a list of VPCConnectors
type VPCConnectorList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VPCConnector `json:"items"`
}

// Repository type metadata.
var (
	VPCConnector_Kind             = "VPCConnector"
	VPCConnector_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: VPCConnector_Kind}.String()
	VPCConnector_KindAPIVersion   = VPCConnector_Kind + "." + CRDGroupVersion.String()
	VPCConnector_GroupVersionKind = CRDGroupVersion.WithKind(VPCConnector_Kind)
)

func init() {
	SchemeBuilder.Register(&VPCConnector{}, &VPCConnectorList{})
}
