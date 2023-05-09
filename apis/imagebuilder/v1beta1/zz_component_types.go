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

type ComponentObservation struct {

	// Amazon Resource Name (ARN) of the component.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// Change description of the component.
	ChangeDescription *string `json:"changeDescription,omitempty" tf:"change_description,omitempty"`

	// Inline YAML string with data of the component. Exactly one of data and uri can be specified.
	Data *string `json:"data,omitempty" tf:"data,omitempty"`

	// Date the component was created.
	DateCreated *string `json:"dateCreated,omitempty" tf:"date_created,omitempty"`

	// Description of the component.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Encryption status of the component.
	Encrypted *bool `json:"encrypted,omitempty" tf:"encrypted,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Amazon Resource Name (ARN) of the Key Management Service (KMS) Key used to encrypt the component.
	KMSKeyID *string `json:"kmsKeyId,omitempty" tf:"kms_key_id,omitempty"`

	// Name of the component.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Owner of the component.
	Owner *string `json:"owner,omitempty" tf:"owner,omitempty"`

	// Platform of the component.
	Platform *string `json:"platform,omitempty" tf:"platform,omitempty"`

	// Whether to retain the old version when the resource is destroyed or replacement is necessary. Defaults to false.
	SkipDestroy *bool `json:"skipDestroy,omitempty" tf:"skip_destroy,omitempty"`

	// Set of Operating Systems (OS) supported by the component.
	SupportedOsVersions []*string `json:"supportedOsVersions,omitempty" tf:"supported_os_versions,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`

	// Type of the component.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// S3 URI with data of the component. Exactly one of data and uri can be specified.
	URI *string `json:"uri,omitempty" tf:"uri,omitempty"`

	// Version of the component.
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

type ComponentParameters struct {

	// Change description of the component.
	// +kubebuilder:validation:Optional
	ChangeDescription *string `json:"changeDescription,omitempty" tf:"change_description,omitempty"`

	// Inline YAML string with data of the component. Exactly one of data and uri can be specified.
	// +kubebuilder:validation:Optional
	Data *string `json:"data,omitempty" tf:"data,omitempty"`

	// Description of the component.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Amazon Resource Name (ARN) of the Key Management Service (KMS) Key used to encrypt the component.
	// +crossplane:generate:reference:type=github.com/spirosco/upbound-provider-aws/apis/kms/v1beta1.Key
	// +kubebuilder:validation:Optional
	KMSKeyID *string `json:"kmsKeyId,omitempty" tf:"kms_key_id,omitempty"`

	// Reference to a Key in kms to populate kmsKeyId.
	// +kubebuilder:validation:Optional
	KMSKeyIDRef *v1.Reference `json:"kmsKeyIdRef,omitempty" tf:"-"`

	// Selector for a Key in kms to populate kmsKeyId.
	// +kubebuilder:validation:Optional
	KMSKeyIDSelector *v1.Selector `json:"kmsKeyIdSelector,omitempty" tf:"-"`

	// Name of the component.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Platform of the component.
	// +kubebuilder:validation:Optional
	Platform *string `json:"platform,omitempty" tf:"platform,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Whether to retain the old version when the resource is destroyed or replacement is necessary. Defaults to false.
	// +kubebuilder:validation:Optional
	SkipDestroy *bool `json:"skipDestroy,omitempty" tf:"skip_destroy,omitempty"`

	// Set of Operating Systems (OS) supported by the component.
	// +kubebuilder:validation:Optional
	SupportedOsVersions []*string `json:"supportedOsVersions,omitempty" tf:"supported_os_versions,omitempty"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// S3 URI with data of the component. Exactly one of data and uri can be specified.
	// +kubebuilder:validation:Optional
	URI *string `json:"uri,omitempty" tf:"uri,omitempty"`

	// Version of the component.
	// +kubebuilder:validation:Optional
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

// ComponentSpec defines the desired state of Component
type ComponentSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ComponentParameters `json:"forProvider"`
}

// ComponentStatus defines the observed state of Component.
type ComponentStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ComponentObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Component is the Schema for the Components API. Manage an Image Builder Component
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Component struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.name)",message="name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.platform)",message="platform is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.version)",message="version is a required parameter"
	Spec   ComponentSpec   `json:"spec"`
	Status ComponentStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ComponentList contains a list of Components
type ComponentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Component `json:"items"`
}

// Repository type metadata.
var (
	Component_Kind             = "Component"
	Component_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Component_Kind}.String()
	Component_KindAPIVersion   = Component_Kind + "." + CRDGroupVersion.String()
	Component_GroupVersionKind = CRDGroupVersion.WithKind(Component_Kind)
)

func init() {
	SchemeBuilder.Register(&Component{}, &ComponentList{})
}
