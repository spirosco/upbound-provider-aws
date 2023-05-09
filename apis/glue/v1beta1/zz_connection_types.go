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

type ConnectionObservation struct {

	// The ARN of the Glue Connection.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// –  The ID of the Data Catalog in which to create the connection. If none is supplied, the AWS account ID is used by default.
	CatalogID *string `json:"catalogId,omitempty" tf:"catalog_id,omitempty"`

	// –  The type of the connection. Supported are: CUSTOM, JDBC, KAFKA, MARKETPLACE, MONGODB, and NETWORK. Defaults to JBDC.
	ConnectionType *string `json:"connectionType,omitempty" tf:"connection_type,omitempty"`

	// –  Description of the connection.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Catalog ID and name of the connection
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// –  A list of criteria that can be used in selecting this connection.
	MatchCriteria []*string `json:"matchCriteria,omitempty" tf:"match_criteria,omitempty"`

	// A map of physical connection requirements, such as VPC and SecurityGroup. Defined below.
	PhysicalConnectionRequirements []PhysicalConnectionRequirementsObservation `json:"physicalConnectionRequirements,omitempty" tf:"physical_connection_requirements,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type ConnectionParameters struct {

	// –  The ID of the Data Catalog in which to create the connection. If none is supplied, the AWS account ID is used by default.
	// +kubebuilder:validation:Required
	CatalogID *string `json:"catalogId" tf:"catalog_id,omitempty"`

	// value pairs used as parameters for this connection.
	// +kubebuilder:validation:Optional
	ConnectionPropertiesSecretRef *v1.SecretReference `json:"connectionPropertiesSecretRef,omitempty" tf:"-"`

	// –  The type of the connection. Supported are: CUSTOM, JDBC, KAFKA, MARKETPLACE, MONGODB, and NETWORK. Defaults to JBDC.
	// +kubebuilder:validation:Optional
	ConnectionType *string `json:"connectionType,omitempty" tf:"connection_type,omitempty"`

	// –  Description of the connection.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// –  A list of criteria that can be used in selecting this connection.
	// +kubebuilder:validation:Optional
	MatchCriteria []*string `json:"matchCriteria,omitempty" tf:"match_criteria,omitempty"`

	// A map of physical connection requirements, such as VPC and SecurityGroup. Defined below.
	// +kubebuilder:validation:Optional
	PhysicalConnectionRequirements []PhysicalConnectionRequirementsParameters `json:"physicalConnectionRequirements,omitempty" tf:"physical_connection_requirements,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type PhysicalConnectionRequirementsObservation struct {

	// The availability zone of the connection. This field is redundant and implied by subnet_id, but is currently an api requirement.
	AvailabilityZone *string `json:"availabilityZone,omitempty" tf:"availability_zone,omitempty"`

	// The security group ID list used by the connection.
	SecurityGroupIDList []*string `json:"securityGroupIdList,omitempty" tf:"security_group_id_list,omitempty"`

	// The subnet ID used by the connection.
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`
}

type PhysicalConnectionRequirementsParameters struct {

	// The availability zone of the connection. This field is redundant and implied by subnet_id, but is currently an api requirement.
	// +crossplane:generate:reference:type=github.com/spirosco/upbound-provider-aws/apis/ec2/v1beta1.Subnet
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractParamPath("availability_zone",false)
	// +kubebuilder:validation:Optional
	AvailabilityZone *string `json:"availabilityZone,omitempty" tf:"availability_zone,omitempty"`

	// Reference to a Subnet in ec2 to populate availabilityZone.
	// +kubebuilder:validation:Optional
	AvailabilityZoneRef *v1.Reference `json:"availabilityZoneRef,omitempty" tf:"-"`

	// Selector for a Subnet in ec2 to populate availabilityZone.
	// +kubebuilder:validation:Optional
	AvailabilityZoneSelector *v1.Selector `json:"availabilityZoneSelector,omitempty" tf:"-"`

	// The security group ID list used by the connection.
	// +kubebuilder:validation:Optional
	SecurityGroupIDList []*string `json:"securityGroupIdList,omitempty" tf:"security_group_id_list,omitempty"`

	// The subnet ID used by the connection.
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

// ConnectionSpec defines the desired state of Connection
type ConnectionSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ConnectionParameters `json:"forProvider"`
}

// ConnectionStatus defines the observed state of Connection.
type ConnectionStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ConnectionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Connection is the Schema for the Connections API. Provides an Glue Connection resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Connection struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ConnectionSpec   `json:"spec"`
	Status            ConnectionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ConnectionList contains a list of Connections
type ConnectionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Connection `json:"items"`
}

// Repository type metadata.
var (
	Connection_Kind             = "Connection"
	Connection_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Connection_Kind}.String()
	Connection_KindAPIVersion   = Connection_Kind + "." + CRDGroupVersion.String()
	Connection_GroupVersionKind = CRDGroupVersion.WithKind(Connection_Kind)
)

func init() {
	SchemeBuilder.Register(&Connection{}, &ConnectionList{})
}
