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

type VPNGatewayRoutePropagationObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The id of the aws_route_table to propagate routes into.
	RouteTableID *string `json:"routeTableId,omitempty" tf:"route_table_id,omitempty"`

	// The id of the aws_vpn_gateway to propagate routes from.
	VPNGatewayID *string `json:"vpnGatewayId,omitempty" tf:"vpn_gateway_id,omitempty"`
}

type VPNGatewayRoutePropagationParameters struct {

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// The id of the aws_route_table to propagate routes into.
	// +crossplane:generate:reference:type=github.com/spirosco/upbound-provider-aws/apis/ec2/v1beta1.RouteTable
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	RouteTableID *string `json:"routeTableId,omitempty" tf:"route_table_id,omitempty"`

	// Reference to a RouteTable in ec2 to populate routeTableId.
	// +kubebuilder:validation:Optional
	RouteTableIDRef *v1.Reference `json:"routeTableIdRef,omitempty" tf:"-"`

	// Selector for a RouteTable in ec2 to populate routeTableId.
	// +kubebuilder:validation:Optional
	RouteTableIDSelector *v1.Selector `json:"routeTableIdSelector,omitempty" tf:"-"`

	// The id of the aws_vpn_gateway to propagate routes from.
	// +crossplane:generate:reference:type=github.com/spirosco/upbound-provider-aws/apis/ec2/v1beta1.VPNGateway
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	VPNGatewayID *string `json:"vpnGatewayId,omitempty" tf:"vpn_gateway_id,omitempty"`

	// Reference to a VPNGateway in ec2 to populate vpnGatewayId.
	// +kubebuilder:validation:Optional
	VPNGatewayIDRef *v1.Reference `json:"vpnGatewayIdRef,omitempty" tf:"-"`

	// Selector for a VPNGateway in ec2 to populate vpnGatewayId.
	// +kubebuilder:validation:Optional
	VPNGatewayIDSelector *v1.Selector `json:"vpnGatewayIdSelector,omitempty" tf:"-"`
}

// VPNGatewayRoutePropagationSpec defines the desired state of VPNGatewayRoutePropagation
type VPNGatewayRoutePropagationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     VPNGatewayRoutePropagationParameters `json:"forProvider"`
}

// VPNGatewayRoutePropagationStatus defines the observed state of VPNGatewayRoutePropagation.
type VPNGatewayRoutePropagationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        VPNGatewayRoutePropagationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// VPNGatewayRoutePropagation is the Schema for the VPNGatewayRoutePropagations API. Requests automatic route propagation between a VPN gateway and a route table.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type VPNGatewayRoutePropagation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VPNGatewayRoutePropagationSpec   `json:"spec"`
	Status            VPNGatewayRoutePropagationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VPNGatewayRoutePropagationList contains a list of VPNGatewayRoutePropagations
type VPNGatewayRoutePropagationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VPNGatewayRoutePropagation `json:"items"`
}

// Repository type metadata.
var (
	VPNGatewayRoutePropagation_Kind             = "VPNGatewayRoutePropagation"
	VPNGatewayRoutePropagation_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: VPNGatewayRoutePropagation_Kind}.String()
	VPNGatewayRoutePropagation_KindAPIVersion   = VPNGatewayRoutePropagation_Kind + "." + CRDGroupVersion.String()
	VPNGatewayRoutePropagation_GroupVersionKind = CRDGroupVersion.WithKind(VPNGatewayRoutePropagation_Kind)
)

func init() {
	SchemeBuilder.Register(&VPNGatewayRoutePropagation{}, &VPNGatewayRoutePropagationList{})
}
