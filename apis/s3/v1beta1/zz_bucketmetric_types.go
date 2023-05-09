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

type BucketMetricFilterObservation struct {

	// Object prefix for filtering (singular).
	Prefix *string `json:"prefix,omitempty" tf:"prefix,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type BucketMetricFilterParameters struct {

	// Object prefix for filtering (singular).
	// +kubebuilder:validation:Optional
	Prefix *string `json:"prefix,omitempty" tf:"prefix,omitempty"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type BucketMetricObservation struct {

	// Name of the bucket to put metric configuration.
	Bucket *string `json:"bucket,omitempty" tf:"bucket,omitempty"`

	// Object filtering that accepts a prefix, tags, or a logical AND of prefix and tags (documented below).
	Filter []BucketMetricFilterObservation `json:"filter,omitempty" tf:"filter,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Unique identifier of the metrics configuration for the bucket. Must be less than or equal to 64 characters in length.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type BucketMetricParameters struct {

	// Name of the bucket to put metric configuration.
	// +crossplane:generate:reference:type=github.com/spirosco/upbound-provider-aws/apis/s3/v1beta1.Bucket
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	Bucket *string `json:"bucket,omitempty" tf:"bucket,omitempty"`

	// Reference to a Bucket in s3 to populate bucket.
	// +kubebuilder:validation:Optional
	BucketRef *v1.Reference `json:"bucketRef,omitempty" tf:"-"`

	// Selector for a Bucket in s3 to populate bucket.
	// +kubebuilder:validation:Optional
	BucketSelector *v1.Selector `json:"bucketSelector,omitempty" tf:"-"`

	// Object filtering that accepts a prefix, tags, or a logical AND of prefix and tags (documented below).
	// +kubebuilder:validation:Optional
	Filter []BucketMetricFilterParameters `json:"filter,omitempty" tf:"filter,omitempty"`

	// Unique identifier of the metrics configuration for the bucket. Must be less than or equal to 64 characters in length.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`
}

// BucketMetricSpec defines the desired state of BucketMetric
type BucketMetricSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     BucketMetricParameters `json:"forProvider"`
}

// BucketMetricStatus defines the observed state of BucketMetric.
type BucketMetricStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        BucketMetricObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// BucketMetric is the Schema for the BucketMetrics API. Provides a S3 bucket metrics configuration resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type BucketMetric struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.name)",message="name is a required parameter"
	Spec   BucketMetricSpec   `json:"spec"`
	Status BucketMetricStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// BucketMetricList contains a list of BucketMetrics
type BucketMetricList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BucketMetric `json:"items"`
}

// Repository type metadata.
var (
	BucketMetric_Kind             = "BucketMetric"
	BucketMetric_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: BucketMetric_Kind}.String()
	BucketMetric_KindAPIVersion   = BucketMetric_Kind + "." + CRDGroupVersion.String()
	BucketMetric_GroupVersionKind = CRDGroupVersion.WithKind(BucketMetric_Kind)
)

func init() {
	SchemeBuilder.Register(&BucketMetric{}, &BucketMetricList{})
}
