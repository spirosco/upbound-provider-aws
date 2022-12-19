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

type FleetStackAssociationObservation struct {

	// Unique ID of the appstream stack fleet association, composed of the fleet_name and stack_name separated by a slash (/).
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type FleetStackAssociationParameters struct {

	// Name of the fleet.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/appstream/v1beta1.Fleet
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractParamPath("name",false)
	// +kubebuilder:validation:Optional
	FleetName *string `json:"fleetName,omitempty" tf:"fleet_name,omitempty"`

	// Reference to a Fleet in appstream to populate fleetName.
	// +kubebuilder:validation:Optional
	FleetNameRef *v1.Reference `json:"fleetNameRef,omitempty" tf:"-"`

	// Selector for a Fleet in appstream to populate fleetName.
	// +kubebuilder:validation:Optional
	FleetNameSelector *v1.Selector `json:"fleetNameSelector,omitempty" tf:"-"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Name of the stack.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/appstream/v1beta1.Stack
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractParamPath("name",false)
	// +kubebuilder:validation:Optional
	StackName *string `json:"stackName,omitempty" tf:"stack_name,omitempty"`

	// Reference to a Stack in appstream to populate stackName.
	// +kubebuilder:validation:Optional
	StackNameRef *v1.Reference `json:"stackNameRef,omitempty" tf:"-"`

	// Selector for a Stack in appstream to populate stackName.
	// +kubebuilder:validation:Optional
	StackNameSelector *v1.Selector `json:"stackNameSelector,omitempty" tf:"-"`
}

// FleetStackAssociationSpec defines the desired state of FleetStackAssociation
type FleetStackAssociationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     FleetStackAssociationParameters `json:"forProvider"`
}

// FleetStackAssociationStatus defines the observed state of FleetStackAssociation.
type FleetStackAssociationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        FleetStackAssociationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// FleetStackAssociation is the Schema for the FleetStackAssociations API. Manages an AppStream Fleet Stack association.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type FleetStackAssociation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              FleetStackAssociationSpec   `json:"spec"`
	Status            FleetStackAssociationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// FleetStackAssociationList contains a list of FleetStackAssociations
type FleetStackAssociationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FleetStackAssociation `json:"items"`
}

// Repository type metadata.
var (
	FleetStackAssociation_Kind             = "FleetStackAssociation"
	FleetStackAssociation_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: FleetStackAssociation_Kind}.String()
	FleetStackAssociation_KindAPIVersion   = FleetStackAssociation_Kind + "." + CRDGroupVersion.String()
	FleetStackAssociation_GroupVersionKind = CRDGroupVersion.WithKind(FleetStackAssociation_Kind)
)

func init() {
	SchemeBuilder.Register(&FleetStackAssociation{}, &FleetStackAssociationList{})
}