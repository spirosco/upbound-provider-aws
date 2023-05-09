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

type EnvironmentMembershipObservation struct {

	// The ID of the environment that contains the environment member you want to add.
	EnvironmentID *string `json:"environmentId,omitempty" tf:"environment_id,omitempty"`

	// The ID of the environment membership.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The type of environment member permissions you want to associate with this environment member. Allowed values are read-only and read-write .
	Permissions *string `json:"permissions,omitempty" tf:"permissions,omitempty"`

	// The Amazon Resource Name (ARN) of the environment member you want to add.
	UserArn *string `json:"userArn,omitempty" tf:"user_arn,omitempty"`

	// he user ID in AWS Identity and Access Management (AWS IAM) of the environment member.
	UserID *string `json:"userId,omitempty" tf:"user_id,omitempty"`
}

type EnvironmentMembershipParameters struct {

	// The ID of the environment that contains the environment member you want to add.
	// +crossplane:generate:reference:type=github.com/spirosco/upbound-provider-aws/apis/cloud9/v1beta1.EnvironmentEC2
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	EnvironmentID *string `json:"environmentId,omitempty" tf:"environment_id,omitempty"`

	// Reference to a EnvironmentEC2 in cloud9 to populate environmentId.
	// +kubebuilder:validation:Optional
	EnvironmentIDRef *v1.Reference `json:"environmentIdRef,omitempty" tf:"-"`

	// Selector for a EnvironmentEC2 in cloud9 to populate environmentId.
	// +kubebuilder:validation:Optional
	EnvironmentIDSelector *v1.Selector `json:"environmentIdSelector,omitempty" tf:"-"`

	// The type of environment member permissions you want to associate with this environment member. Allowed values are read-only and read-write .
	// +kubebuilder:validation:Optional
	Permissions *string `json:"permissions,omitempty" tf:"permissions,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// The Amazon Resource Name (ARN) of the environment member you want to add.
	// +crossplane:generate:reference:type=github.com/spirosco/upbound-provider-aws/apis/iam/v1beta1.User
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractParamPath("arn",true)
	// +kubebuilder:validation:Optional
	UserArn *string `json:"userArn,omitempty" tf:"user_arn,omitempty"`

	// Reference to a User in iam to populate userArn.
	// +kubebuilder:validation:Optional
	UserArnRef *v1.Reference `json:"userArnRef,omitempty" tf:"-"`

	// Selector for a User in iam to populate userArn.
	// +kubebuilder:validation:Optional
	UserArnSelector *v1.Selector `json:"userArnSelector,omitempty" tf:"-"`
}

// EnvironmentMembershipSpec defines the desired state of EnvironmentMembership
type EnvironmentMembershipSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     EnvironmentMembershipParameters `json:"forProvider"`
}

// EnvironmentMembershipStatus defines the observed state of EnvironmentMembership.
type EnvironmentMembershipStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        EnvironmentMembershipObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// EnvironmentMembership is the Schema for the EnvironmentMemberships API. Provides an environment member to an AWS Cloud9 development environment.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type EnvironmentMembership struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.permissions)",message="permissions is a required parameter"
	Spec   EnvironmentMembershipSpec   `json:"spec"`
	Status EnvironmentMembershipStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// EnvironmentMembershipList contains a list of EnvironmentMemberships
type EnvironmentMembershipList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []EnvironmentMembership `json:"items"`
}

// Repository type metadata.
var (
	EnvironmentMembership_Kind             = "EnvironmentMembership"
	EnvironmentMembership_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: EnvironmentMembership_Kind}.String()
	EnvironmentMembership_KindAPIVersion   = EnvironmentMembership_Kind + "." + CRDGroupVersion.String()
	EnvironmentMembership_GroupVersionKind = CRDGroupVersion.WithKind(EnvironmentMembership_Kind)
)

func init() {
	SchemeBuilder.Register(&EnvironmentMembership{}, &EnvironmentMembershipList{})
}
