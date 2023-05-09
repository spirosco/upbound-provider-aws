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

type EndpointObservation struct {

	// The ARN of the Route 53 Resolver endpoint.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// The direction of DNS queries to or from the Route 53 Resolver endpoint.
	// Valid values are INBOUND (resolver forwards DNS queries to the DNS service for a VPC from your network or another VPC)
	// or OUTBOUND (resolver forwards DNS queries from the DNS service for a VPC to your network or another VPC).
	Direction *string `json:"direction,omitempty" tf:"direction,omitempty"`

	// The ID of the VPC that you want to create the resolver endpoint in.
	HostVPCID *string `json:"hostVpcId,omitempty" tf:"host_vpc_id,omitempty"`

	// The ID of the Route 53 Resolver endpoint.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The subnets and IP addresses in your VPC that you want DNS queries to pass through on the way from your VPCs
	// to your network (for outbound endpoints) or on the way from your network to your VPCs (for inbound endpoints). Described below.
	IPAddress []IPAddressObservation `json:"ipAddress,omitempty" tf:"ip_address,omitempty"`

	// The friendly name of the Route 53 Resolver endpoint.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The ID of one or more security groups that you want to use to control access to this VPC.
	SecurityGroupIds []*string `json:"securityGroupIds,omitempty" tf:"security_group_ids,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type EndpointParameters struct {

	// The direction of DNS queries to or from the Route 53 Resolver endpoint.
	// Valid values are INBOUND (resolver forwards DNS queries to the DNS service for a VPC from your network or another VPC)
	// or OUTBOUND (resolver forwards DNS queries from the DNS service for a VPC to your network or another VPC).
	// +kubebuilder:validation:Optional
	Direction *string `json:"direction,omitempty" tf:"direction,omitempty"`

	// The subnets and IP addresses in your VPC that you want DNS queries to pass through on the way from your VPCs
	// to your network (for outbound endpoints) or on the way from your network to your VPCs (for inbound endpoints). Described below.
	// +kubebuilder:validation:Optional
	IPAddress []IPAddressParameters `json:"ipAddress,omitempty" tf:"ip_address,omitempty"`

	// The friendly name of the Route 53 Resolver endpoint.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// References to SecurityGroup in ec2 to populate securityGroupIds.
	// +kubebuilder:validation:Optional
	SecurityGroupIDRefs []v1.Reference `json:"securityGroupIdRefs,omitempty" tf:"-"`

	// Selector for a list of SecurityGroup in ec2 to populate securityGroupIds.
	// +kubebuilder:validation:Optional
	SecurityGroupIDSelector *v1.Selector `json:"securityGroupIdSelector,omitempty" tf:"-"`

	// The ID of one or more security groups that you want to use to control access to this VPC.
	// +crossplane:generate:reference:type=github.com/spirosco/upbound-provider-aws/apis/ec2/v1beta1.SecurityGroup
	// +crossplane:generate:reference:refFieldName=SecurityGroupIDRefs
	// +crossplane:generate:reference:selectorFieldName=SecurityGroupIDSelector
	// +kubebuilder:validation:Optional
	SecurityGroupIds []*string `json:"securityGroupIds,omitempty" tf:"security_group_ids,omitempty"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type IPAddressObservation struct {

	// The IP address in the subnet that you want to use for DNS queries.
	IP *string `json:"ip,omitempty" tf:"ip,omitempty"`

	// The ID of the Route 53 Resolver endpoint.
	IPID *string `json:"ipId,omitempty" tf:"ip_id,omitempty"`

	// The ID of the subnet that contains the IP address.
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`
}

type IPAddressParameters struct {

	// The IP address in the subnet that you want to use for DNS queries.
	// +kubebuilder:validation:Optional
	IP *string `json:"ip,omitempty" tf:"ip,omitempty"`

	// The ID of the subnet that contains the IP address.
	// +crossplane:generate:reference:type=github.com/spirosco/upbound-provider-aws/apis/ec2/v1beta1.Subnet
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// Reference to a Subnet in ec2 to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDRef *v1.Reference `json:"subnetIdRef,omitempty" tf:"-"`

	// Selector for a Subnet in ec2 to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDSelector *v1.Selector `json:"subnetIdSelector,omitempty" tf:"-"`
}

// EndpointSpec defines the desired state of Endpoint
type EndpointSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     EndpointParameters `json:"forProvider"`
}

// EndpointStatus defines the observed state of Endpoint.
type EndpointStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        EndpointObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Endpoint is the Schema for the Endpoints API. Provides a Route 53 Resolver endpoint resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Endpoint struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.direction)",message="direction is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.ipAddress)",message="ipAddress is a required parameter"
	Spec   EndpointSpec   `json:"spec"`
	Status EndpointStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// EndpointList contains a list of Endpoints
type EndpointList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Endpoint `json:"items"`
}

// Repository type metadata.
var (
	Endpoint_Kind             = "Endpoint"
	Endpoint_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Endpoint_Kind}.String()
	Endpoint_KindAPIVersion   = Endpoint_Kind + "." + CRDGroupVersion.String()
	Endpoint_GroupVersionKind = CRDGroupVersion.WithKind(Endpoint_Kind)
)

func init() {
	SchemeBuilder.Register(&Endpoint{}, &EndpointList{})
}
