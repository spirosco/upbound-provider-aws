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

type EBSEncryptionByDefaultObservation struct {

	// Whether or not default EBS encryption is enabled. Valid values are true or false. Defaults to true.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type EBSEncryptionByDefaultParameters struct {

	// Whether or not default EBS encryption is enabled. Valid values are true or false. Defaults to true.
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`
}

// EBSEncryptionByDefaultSpec defines the desired state of EBSEncryptionByDefault
type EBSEncryptionByDefaultSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     EBSEncryptionByDefaultParameters `json:"forProvider"`
}

// EBSEncryptionByDefaultStatus defines the observed state of EBSEncryptionByDefault.
type EBSEncryptionByDefaultStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        EBSEncryptionByDefaultObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// EBSEncryptionByDefault is the Schema for the EBSEncryptionByDefaults API. Manages whether default EBS encryption is enabled for your AWS account in the current AWS region.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type EBSEncryptionByDefault struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              EBSEncryptionByDefaultSpec   `json:"spec"`
	Status            EBSEncryptionByDefaultStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// EBSEncryptionByDefaultList contains a list of EBSEncryptionByDefaults
type EBSEncryptionByDefaultList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []EBSEncryptionByDefault `json:"items"`
}

// Repository type metadata.
var (
	EBSEncryptionByDefault_Kind             = "EBSEncryptionByDefault"
	EBSEncryptionByDefault_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: EBSEncryptionByDefault_Kind}.String()
	EBSEncryptionByDefault_KindAPIVersion   = EBSEncryptionByDefault_Kind + "." + CRDGroupVersion.String()
	EBSEncryptionByDefault_GroupVersionKind = CRDGroupVersion.WithKind(EBSEncryptionByDefault_Kind)
)

func init() {
	SchemeBuilder.Register(&EBSEncryptionByDefault{}, &EBSEncryptionByDefaultList{})
}
