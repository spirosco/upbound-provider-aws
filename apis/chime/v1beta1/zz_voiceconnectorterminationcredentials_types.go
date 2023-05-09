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

type CredentialsObservation struct {

	// RFC2617 compliant username associated with the SIP credentials.
	Username *string `json:"username,omitempty" tf:"username,omitempty"`
}

type CredentialsParameters struct {

	// RFC2617 compliant password associated with the SIP credentials.
	// +kubebuilder:validation:Required
	PasswordSecretRef v1.SecretKeySelector `json:"passwordSecretRef" tf:"-"`

	// RFC2617 compliant username associated with the SIP credentials.
	// +kubebuilder:validation:Required
	Username *string `json:"username" tf:"username,omitempty"`
}

type VoiceConnectorTerminationCredentialsObservation struct {

	// List of termination SIP credentials.
	Credentials []CredentialsObservation `json:"credentials,omitempty" tf:"credentials,omitempty"`

	// Amazon Chime Voice Connector ID.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Amazon Chime Voice Connector ID.
	VoiceConnectorID *string `json:"voiceConnectorId,omitempty" tf:"voice_connector_id,omitempty"`
}

type VoiceConnectorTerminationCredentialsParameters struct {

	// List of termination SIP credentials.
	// +kubebuilder:validation:Optional
	Credentials []CredentialsParameters `json:"credentials,omitempty" tf:"credentials,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Amazon Chime Voice Connector ID.
	// +crossplane:generate:reference:type=github.com/spirosco/upbound-provider-aws/apis/chime/v1beta1.VoiceConnector
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	VoiceConnectorID *string `json:"voiceConnectorId,omitempty" tf:"voice_connector_id,omitempty"`

	// Reference to a VoiceConnector in chime to populate voiceConnectorId.
	// +kubebuilder:validation:Optional
	VoiceConnectorIDRef *v1.Reference `json:"voiceConnectorIdRef,omitempty" tf:"-"`

	// Selector for a VoiceConnector in chime to populate voiceConnectorId.
	// +kubebuilder:validation:Optional
	VoiceConnectorIDSelector *v1.Selector `json:"voiceConnectorIdSelector,omitempty" tf:"-"`
}

// VoiceConnectorTerminationCredentialsSpec defines the desired state of VoiceConnectorTerminationCredentials
type VoiceConnectorTerminationCredentialsSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     VoiceConnectorTerminationCredentialsParameters `json:"forProvider"`
}

// VoiceConnectorTerminationCredentialsStatus defines the observed state of VoiceConnectorTerminationCredentials.
type VoiceConnectorTerminationCredentialsStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        VoiceConnectorTerminationCredentialsObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// VoiceConnectorTerminationCredentials is the Schema for the VoiceConnectorTerminationCredentialss API. Adds termination SIP credentials for the specified Amazon Chime Voice Connector.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type VoiceConnectorTerminationCredentials struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.credentials)",message="credentials is a required parameter"
	Spec   VoiceConnectorTerminationCredentialsSpec   `json:"spec"`
	Status VoiceConnectorTerminationCredentialsStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VoiceConnectorTerminationCredentialsList contains a list of VoiceConnectorTerminationCredentialss
type VoiceConnectorTerminationCredentialsList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VoiceConnectorTerminationCredentials `json:"items"`
}

// Repository type metadata.
var (
	VoiceConnectorTerminationCredentials_Kind             = "VoiceConnectorTerminationCredentials"
	VoiceConnectorTerminationCredentials_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: VoiceConnectorTerminationCredentials_Kind}.String()
	VoiceConnectorTerminationCredentials_KindAPIVersion   = VoiceConnectorTerminationCredentials_Kind + "." + CRDGroupVersion.String()
	VoiceConnectorTerminationCredentials_GroupVersionKind = CRDGroupVersion.WithKind(VoiceConnectorTerminationCredentials_Kind)
)

func init() {
	SchemeBuilder.Register(&VoiceConnectorTerminationCredentials{}, &VoiceConnectorTerminationCredentialsList{})
}
