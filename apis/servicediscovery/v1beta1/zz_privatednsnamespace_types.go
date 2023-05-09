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

type PrivateDNSNamespaceObservation struct {

	// The ARN that Amazon Route 53 assigns to the namespace when you create it.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// The description that you specify for the namespace when you create it.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The ID for the hosted zone that Amazon Route 53 creates when you create a namespace.
	HostedZone *string `json:"hostedZone,omitempty" tf:"hosted_zone,omitempty"`

	// The ID of a namespace.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The name of the namespace.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`

	// The ID of VPC that you want to associate the namespace with.
	VPC *string `json:"vpc,omitempty" tf:"vpc,omitempty"`
}

type PrivateDNSNamespaceParameters struct {

	// The description that you specify for the namespace when you create it.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The name of the namespace.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The ID of VPC that you want to associate the namespace with.
	// +crossplane:generate:reference:type=github.com/spirosco/upbound-provider-aws/apis/ec2/v1beta1.VPC
	// +kubebuilder:validation:Optional
	VPC *string `json:"vpc,omitempty" tf:"vpc,omitempty"`

	// Reference to a VPC in ec2 to populate vpc.
	// +kubebuilder:validation:Optional
	VPCRef *v1.Reference `json:"vpcRef,omitempty" tf:"-"`

	// Selector for a VPC in ec2 to populate vpc.
	// +kubebuilder:validation:Optional
	VPCSelector *v1.Selector `json:"vpcSelector,omitempty" tf:"-"`
}

// PrivateDNSNamespaceSpec defines the desired state of PrivateDNSNamespace
type PrivateDNSNamespaceSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     PrivateDNSNamespaceParameters `json:"forProvider"`
}

// PrivateDNSNamespaceStatus defines the observed state of PrivateDNSNamespace.
type PrivateDNSNamespaceStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        PrivateDNSNamespaceObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// PrivateDNSNamespace is the Schema for the PrivateDNSNamespaces API. Provides a Service Discovery Private DNS Namespace resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type PrivateDNSNamespace struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.name)",message="name is a required parameter"
	Spec   PrivateDNSNamespaceSpec   `json:"spec"`
	Status PrivateDNSNamespaceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PrivateDNSNamespaceList contains a list of PrivateDNSNamespaces
type PrivateDNSNamespaceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PrivateDNSNamespace `json:"items"`
}

// Repository type metadata.
var (
	PrivateDNSNamespace_Kind             = "PrivateDNSNamespace"
	PrivateDNSNamespace_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: PrivateDNSNamespace_Kind}.String()
	PrivateDNSNamespace_KindAPIVersion   = PrivateDNSNamespace_Kind + "." + CRDGroupVersion.String()
	PrivateDNSNamespace_GroupVersionKind = CRDGroupVersion.WithKind(PrivateDNSNamespace_Kind)
)

func init() {
	SchemeBuilder.Register(&PrivateDNSNamespace{}, &PrivateDNSNamespaceList{})
}
