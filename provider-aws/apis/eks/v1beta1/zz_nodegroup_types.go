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

type AutoscalingGroupsObservation struct {
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type AutoscalingGroupsParameters struct {
}

type LaunchTemplateObservation struct {
}

type LaunchTemplateParameters struct {

	// +kubebuilder:validation:Optional
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	Version *string `json:"version" tf:"version,omitempty"`
}

type NodeGroupObservation struct {

	// Amazon Resource Name (ARN) of the EKS Node Group.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// List of objects containing information about underlying resources.
	Resources []ResourcesObservation `json:"resources,omitempty" tf:"resources,omitempty"`

	// Status of the EKS Node Group.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type NodeGroupParameters struct {

	// Type of Amazon Machine Image (AMI) associated with the EKS Node Group. See the AWS documentation for valid values.
	// +kubebuilder:validation:Optional
	AMIType *string `json:"amiType,omitempty" tf:"ami_type,omitempty"`

	// Type of capacity associated with the EKS Node Group. Valid values: ON_DEMAND, SPOT.
	// +kubebuilder:validation:Optional
	CapacityType *string `json:"capacityType,omitempty" tf:"capacity_type,omitempty"`

	// 100 characters in length. Must begin with an alphanumeric character, and must only contain alphanumeric characters, dashes and underscores (^[0-9A-Za-z][A-Za-z0-9\-_]+$).
	// +crossplane:generate:reference:type=Cluster
	// +crossplane:generate:reference:extractor=ExternalNameIfClusterActive()
	// +kubebuilder:validation:Optional
	ClusterName *string `json:"clusterName,omitempty" tf:"cluster_name,omitempty"`

	// Reference to a Cluster to populate clusterName.
	// +kubebuilder:validation:Optional
	ClusterNameRef *v1.Reference `json:"clusterNameRef,omitempty" tf:"-"`

	// Selector for a Cluster to populate clusterName.
	// +kubebuilder:validation:Optional
	ClusterNameSelector *v1.Selector `json:"clusterNameSelector,omitempty" tf:"-"`

	// Disk size in GiB for worker nodes. Defaults to 20.
	// +kubebuilder:validation:Optional
	DiskSize *float64 `json:"diskSize,omitempty" tf:"disk_size,omitempty"`

	// Force version update if existing pods are unable to be drained due to a pod disruption budget issue.
	// +kubebuilder:validation:Optional
	ForceUpdateVersion *bool `json:"forceUpdateVersion,omitempty" tf:"force_update_version,omitempty"`

	// List of instance types associated with the EKS Node Group. Defaults to ["t3.medium"].
	// +kubebuilder:validation:Optional
	InstanceTypes []*string `json:"instanceTypes,omitempty" tf:"instance_types,omitempty"`

	// Key-value map of Kubernetes labels. Only labels that are applied with the EKS API are managed by this argument. Other Kubernetes labels applied to the EKS Node Group will not be managed.
	// +kubebuilder:validation:Optional
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// Configuration block with Launch Template settings. Detailed below.
	// +kubebuilder:validation:Optional
	LaunchTemplate []LaunchTemplateParameters `json:"launchTemplate,omitempty" tf:"launch_template,omitempty"`

	// –  Amazon Resource Name (ARN) of the IAM Role that provides permissions for the EKS Node Group.
	// +crossplane:generate:reference:type=github.com/upbound/official-providers/provider-aws/apis/iam/v1beta1.Role
	// +crossplane:generate:reference:extractor=github.com/upbound/official-providers/provider-aws/config/common.ARNExtractor()
	// +kubebuilder:validation:Optional
	NodeRoleArn *string `json:"nodeRoleArn,omitempty" tf:"node_role_arn,omitempty"`

	// Reference to a Role in iam to populate nodeRoleArn.
	// +kubebuilder:validation:Optional
	NodeRoleArnRef *v1.Reference `json:"nodeRoleArnRef,omitempty" tf:"-"`

	// Selector for a Role in iam to populate nodeRoleArn.
	// +kubebuilder:validation:Optional
	NodeRoleArnSelector *v1.Selector `json:"nodeRoleArnSelector,omitempty" tf:"-"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// –  AMI version of the EKS Node Group. Defaults to latest version for Kubernetes version.
	// +kubebuilder:validation:Optional
	ReleaseVersion *string `json:"releaseVersion,omitempty" tf:"release_version,omitempty"`

	// Configuration block with remote access settings. Detailed below.
	// +kubebuilder:validation:Optional
	RemoteAccess []RemoteAccessParameters `json:"remoteAccess,omitempty" tf:"remote_access,omitempty"`

	// Configuration block with scaling settings. Detailed below.
	// +kubebuilder:validation:Required
	ScalingConfig []ScalingConfigParameters `json:"scalingConfig" tf:"scaling_config,omitempty"`

	// References to Subnet in ec2 to populate subnetIds.
	// +kubebuilder:validation:Optional
	SubnetIDRefs []v1.Reference `json:"subnetIdRefs,omitempty" tf:"-"`

	// Selector for a list of Subnet in ec2 to populate subnetIds.
	// +kubebuilder:validation:Optional
	SubnetIDSelector *v1.Selector `json:"subnetIdSelector,omitempty" tf:"-"`

	// –  Identifiers of EC2 Subnets to associate with the EKS Node Group. These subnets must have the following resource tag: kubernetes.io/cluster/CLUSTER_NAME (where CLUSTER_NAME is replaced with the name of the EKS Cluster).
	// +crossplane:generate:reference:type=github.com/upbound/official-providers/provider-aws/apis/ec2/v1beta1.Subnet
	// +crossplane:generate:reference:refFieldName=SubnetIDRefs
	// +crossplane:generate:reference:selectorFieldName=SubnetIDSelector
	// +kubebuilder:validation:Optional
	SubnetIds []*string `json:"subnetIds,omitempty" tf:"subnet_ids,omitempty"`

	// Key-value map of resource tags. If configured with a provider default_tags configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The Kubernetes taints to be applied to the nodes in the node group. Maximum of 50 taints per node group. Detailed below.
	// +kubebuilder:validation:Optional
	Taint []TaintParameters `json:"taint,omitempty" tf:"taint,omitempty"`

	// +kubebuilder:validation:Optional
	UpdateConfig []UpdateConfigParameters `json:"updateConfig,omitempty" tf:"update_config,omitempty"`

	// +kubebuilder:validation:Optional
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

type RemoteAccessObservation struct {
}

type RemoteAccessParameters struct {

	// EC2 Key Pair name that provides access for SSH communication with the worker nodes in the EKS Node Group. If you specify this configuration, but do not specify source_security_group_ids when you create an EKS Node Group, port 22 on the worker nodes is opened to the Internet (0.0.0.0/0).
	// +kubebuilder:validation:Optional
	EC2SSHKey *string `json:"ec2SshKey,omitempty" tf:"ec2_ssh_key,omitempty"`

	// References to SecurityGroup in ec2 to populate sourceSecurityGroupIds.
	// +kubebuilder:validation:Optional
	SourceSecurityGroupIDRefs []v1.Reference `json:"sourceSecurityGroupIdRefs,omitempty" tf:"-"`

	// Selector for a list of SecurityGroup in ec2 to populate sourceSecurityGroupIds.
	// +kubebuilder:validation:Optional
	SourceSecurityGroupIDSelector *v1.Selector `json:"sourceSecurityGroupIdSelector,omitempty" tf:"-"`

	// Set of EC2 Security Group IDs to allow SSH access (port 22) from on the worker nodes. If you specify ec2_ssh_key, but do not specify this configuration when you create an EKS Node Group, port 22 on the worker nodes is opened to the Internet (0.0.0.0/0).
	// +crossplane:generate:reference:type=github.com/upbound/official-providers/provider-aws/apis/ec2/v1beta1.SecurityGroup
	// +crossplane:generate:reference:refFieldName=SourceSecurityGroupIDRefs
	// +crossplane:generate:reference:selectorFieldName=SourceSecurityGroupIDSelector
	// +kubebuilder:validation:Optional
	SourceSecurityGroupIds []*string `json:"sourceSecurityGroupIds,omitempty" tf:"source_security_group_ids,omitempty"`
}

type ResourcesObservation struct {

	// List of objects containing information about AutoScaling Groups.
	AutoscalingGroups []AutoscalingGroupsObservation `json:"autoscalingGroups,omitempty" tf:"autoscaling_groups,omitempty"`

	// Identifier of the remote access EC2 Security Group.
	RemoteAccessSecurityGroupID *string `json:"remoteAccessSecurityGroupId,omitempty" tf:"remote_access_security_group_id,omitempty"`
}

type ResourcesParameters struct {
}

type ScalingConfigObservation struct {
}

type ScalingConfigParameters struct {

	// Desired number of worker nodes.
	// +kubebuilder:validation:Required
	DesiredSize *float64 `json:"desiredSize" tf:"desired_size,omitempty"`

	// Maximum number of worker nodes.
	// +kubebuilder:validation:Required
	MaxSize *float64 `json:"maxSize" tf:"max_size,omitempty"`

	// Minimum number of worker nodes.
	// +kubebuilder:validation:Required
	MinSize *float64 `json:"minSize" tf:"min_size,omitempty"`
}

type TaintObservation struct {
}

type TaintParameters struct {

	// The effect of the taint. Valid values: NO_SCHEDULE, NO_EXECUTE, PREFER_NO_SCHEDULE.
	// +kubebuilder:validation:Required
	Effect *string `json:"effect" tf:"effect,omitempty"`

	// The key of the taint. Maximum length of 63.
	// +kubebuilder:validation:Required
	Key *string `json:"key" tf:"key,omitempty"`

	// The value of the taint. Maximum length of 63.
	// +kubebuilder:validation:Optional
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type UpdateConfigObservation struct {
}

type UpdateConfigParameters struct {

	// Desired max number of unavailable worker nodes during node group update.
	// +kubebuilder:validation:Optional
	MaxUnavailable *float64 `json:"maxUnavailable,omitempty" tf:"max_unavailable,omitempty"`

	// Desired max percentage of unavailable worker nodes during node group update.
	// +kubebuilder:validation:Optional
	MaxUnavailablePercentage *float64 `json:"maxUnavailablePercentage,omitempty" tf:"max_unavailable_percentage,omitempty"`
}

// NodeGroupSpec defines the desired state of NodeGroup
type NodeGroupSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     NodeGroupParameters `json:"forProvider"`
}

// NodeGroupStatus defines the observed state of NodeGroup.
type NodeGroupStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        NodeGroupObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// NodeGroup is the Schema for the NodeGroups API. Manages an EKS Node Group
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type NodeGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              NodeGroupSpec   `json:"spec"`
	Status            NodeGroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// NodeGroupList contains a list of NodeGroups
type NodeGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NodeGroup `json:"items"`
}

// Repository type metadata.
var (
	NodeGroup_Kind             = "NodeGroup"
	NodeGroup_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: NodeGroup_Kind}.String()
	NodeGroup_KindAPIVersion   = NodeGroup_Kind + "." + CRDGroupVersion.String()
	NodeGroup_GroupVersionKind = CRDGroupVersion.WithKind(NodeGroup_Kind)
)

func init() {
	SchemeBuilder.Register(&NodeGroup{}, &NodeGroupList{})
}
