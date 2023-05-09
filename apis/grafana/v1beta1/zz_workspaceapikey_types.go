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

type WorkspaceAPIKeyObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The key token in JSON format. Use this value as a bearer token to authenticate HTTP requests to the workspace.
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	// Specifies the name of the API key. Key names must be unique to the workspace.
	KeyName *string `json:"keyName,omitempty" tf:"key_name,omitempty"`

	// Specifies the permission level of the API key. Valid values are VIEWER, EDITOR, or ADMIN.
	KeyRole *string `json:"keyRole,omitempty" tf:"key_role,omitempty"`

	// Specifies the time in seconds until the API key expires. Keys can be valid for up to 30 days.
	SecondsToLive *float64 `json:"secondsToLive,omitempty" tf:"seconds_to_live,omitempty"`

	// The ID of the workspace that the API key is valid for.
	WorkspaceID *string `json:"workspaceId,omitempty" tf:"workspace_id,omitempty"`
}

type WorkspaceAPIKeyParameters struct {

	// Specifies the name of the API key. Key names must be unique to the workspace.
	// +kubebuilder:validation:Optional
	KeyName *string `json:"keyName,omitempty" tf:"key_name,omitempty"`

	// Specifies the permission level of the API key. Valid values are VIEWER, EDITOR, or ADMIN.
	// +kubebuilder:validation:Optional
	KeyRole *string `json:"keyRole,omitempty" tf:"key_role,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Specifies the time in seconds until the API key expires. Keys can be valid for up to 30 days.
	// +kubebuilder:validation:Optional
	SecondsToLive *float64 `json:"secondsToLive,omitempty" tf:"seconds_to_live,omitempty"`

	// The ID of the workspace that the API key is valid for.
	// +crossplane:generate:reference:type=github.com/spirosco/upbound-provider-aws/apis/grafana/v1beta1.Workspace
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	WorkspaceID *string `json:"workspaceId,omitempty" tf:"workspace_id,omitempty"`

	// Reference to a Workspace in grafana to populate workspaceId.
	// +kubebuilder:validation:Optional
	WorkspaceIDRef *v1.Reference `json:"workspaceIdRef,omitempty" tf:"-"`

	// Selector for a Workspace in grafana to populate workspaceId.
	// +kubebuilder:validation:Optional
	WorkspaceIDSelector *v1.Selector `json:"workspaceIdSelector,omitempty" tf:"-"`
}

// WorkspaceAPIKeySpec defines the desired state of WorkspaceAPIKey
type WorkspaceAPIKeySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     WorkspaceAPIKeyParameters `json:"forProvider"`
}

// WorkspaceAPIKeyStatus defines the observed state of WorkspaceAPIKey.
type WorkspaceAPIKeyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        WorkspaceAPIKeyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// WorkspaceAPIKey is the Schema for the WorkspaceAPIKeys API. Creates a Grafana API key for the workspace. This key can be used to authenticate requests sent to the workspace's HTTP API.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type WorkspaceAPIKey struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.keyName)",message="keyName is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.keyRole)",message="keyRole is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.secondsToLive)",message="secondsToLive is a required parameter"
	Spec   WorkspaceAPIKeySpec   `json:"spec"`
	Status WorkspaceAPIKeyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// WorkspaceAPIKeyList contains a list of WorkspaceAPIKeys
type WorkspaceAPIKeyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []WorkspaceAPIKey `json:"items"`
}

// Repository type metadata.
var (
	WorkspaceAPIKey_Kind             = "WorkspaceAPIKey"
	WorkspaceAPIKey_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: WorkspaceAPIKey_Kind}.String()
	WorkspaceAPIKey_KindAPIVersion   = WorkspaceAPIKey_Kind + "." + CRDGroupVersion.String()
	WorkspaceAPIKey_GroupVersionKind = CRDGroupVersion.WithKind(WorkspaceAPIKey_Kind)
)

func init() {
	SchemeBuilder.Register(&WorkspaceAPIKey{}, &WorkspaceAPIKeyList{})
}
