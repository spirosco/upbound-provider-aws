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

type DomainNameConfigurationObservation struct {

	// ARN of an AWS-managed certificate that will be used by the endpoint for the domain name. AWS Certificate Manager is the only supported source. Use the aws_acm_certificate resource to configure an ACM certificate.
	CertificateArn *string `json:"certificateArn,omitempty" tf:"certificate_arn,omitempty"`

	// Endpoint type. Valid values: REGIONAL.
	EndpointType *string `json:"endpointType,omitempty" tf:"endpoint_type,omitempty"`

	// (Computed) Amazon Route 53 Hosted Zone ID of the endpoint.
	HostedZoneID *string `json:"hostedZoneId,omitempty" tf:"hosted_zone_id,omitempty"`

	// ARN of the AWS-issued certificate used to validate custom domain ownership (when certificate_arn is issued via an ACM Private CA or mutual_tls_authentication is configured with an ACM-imported certificate.)
	OwnershipVerificationCertificateArn *string `json:"ownershipVerificationCertificateArn,omitempty" tf:"ownership_verification_certificate_arn,omitempty"`

	// Transport Layer Security (TLS) version of the security policy for the domain name. Valid values: TLS_1_2.
	SecurityPolicy *string `json:"securityPolicy,omitempty" tf:"security_policy,omitempty"`

	// (Computed) Target domain name.
	TargetDomainName *string `json:"targetDomainName,omitempty" tf:"target_domain_name,omitempty"`
}

type DomainNameConfigurationParameters struct {

	// ARN of an AWS-managed certificate that will be used by the endpoint for the domain name. AWS Certificate Manager is the only supported source. Use the aws_acm_certificate resource to configure an ACM certificate.
	// +crossplane:generate:reference:type=github.com/spirosco/upbound-provider-aws/apis/acm/v1beta1.Certificate
	// +crossplane:generate:reference:extractor=github.com/spirosco/upbound-provider-aws/config/common.ARNExtractor()
	// +kubebuilder:validation:Optional
	CertificateArn *string `json:"certificateArn,omitempty" tf:"certificate_arn,omitempty"`

	// Reference to a Certificate in acm to populate certificateArn.
	// +kubebuilder:validation:Optional
	CertificateArnRef *v1.Reference `json:"certificateArnRef,omitempty" tf:"-"`

	// Selector for a Certificate in acm to populate certificateArn.
	// +kubebuilder:validation:Optional
	CertificateArnSelector *v1.Selector `json:"certificateArnSelector,omitempty" tf:"-"`

	// Endpoint type. Valid values: REGIONAL.
	// +kubebuilder:validation:Required
	EndpointType *string `json:"endpointType" tf:"endpoint_type,omitempty"`

	// ARN of the AWS-issued certificate used to validate custom domain ownership (when certificate_arn is issued via an ACM Private CA or mutual_tls_authentication is configured with an ACM-imported certificate.)
	// +kubebuilder:validation:Optional
	OwnershipVerificationCertificateArn *string `json:"ownershipVerificationCertificateArn,omitempty" tf:"ownership_verification_certificate_arn,omitempty"`

	// Transport Layer Security (TLS) version of the security policy for the domain name. Valid values: TLS_1_2.
	// +kubebuilder:validation:Required
	SecurityPolicy *string `json:"securityPolicy" tf:"security_policy,omitempty"`
}

type DomainNameObservation struct {

	// API mapping selection expression for the domain name.
	APIMappingSelectionExpression *string `json:"apiMappingSelectionExpression,omitempty" tf:"api_mapping_selection_expression,omitempty"`

	// ARN of the domain name.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// Domain name configuration. See below.
	DomainNameConfiguration []DomainNameConfigurationObservation `json:"domainNameConfiguration,omitempty" tf:"domain_name_configuration,omitempty"`

	// Domain name identifier.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Mutual TLS authentication configuration for the domain name.
	MutualTLSAuthentication []MutualTLSAuthenticationObservation `json:"mutualTlsAuthentication,omitempty" tf:"mutual_tls_authentication,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type DomainNameParameters struct {

	// Domain name configuration. See below.
	// +kubebuilder:validation:Optional
	DomainNameConfiguration []DomainNameConfigurationParameters `json:"domainNameConfiguration,omitempty" tf:"domain_name_configuration,omitempty"`

	// Mutual TLS authentication configuration for the domain name.
	// +kubebuilder:validation:Optional
	MutualTLSAuthentication []MutualTLSAuthenticationParameters `json:"mutualTlsAuthentication,omitempty" tf:"mutual_tls_authentication,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type MutualTLSAuthenticationObservation struct {

	// Amazon S3 URL that specifies the truststore for mutual TLS authentication, for example, s3://bucket-name/key-name. The truststore can contain certificates from public or private certificate authorities. To update the truststore, upload a new version to S3, and then update your custom domain name to use the new version.
	TruststoreURI *string `json:"truststoreUri,omitempty" tf:"truststore_uri,omitempty"`

	// Version of the S3 object that contains the truststore. To specify a version, you must have versioning enabled for the S3 bucket.
	TruststoreVersion *string `json:"truststoreVersion,omitempty" tf:"truststore_version,omitempty"`
}

type MutualTLSAuthenticationParameters struct {

	// Amazon S3 URL that specifies the truststore for mutual TLS authentication, for example, s3://bucket-name/key-name. The truststore can contain certificates from public or private certificate authorities. To update the truststore, upload a new version to S3, and then update your custom domain name to use the new version.
	// +kubebuilder:validation:Required
	TruststoreURI *string `json:"truststoreUri" tf:"truststore_uri,omitempty"`

	// Version of the S3 object that contains the truststore. To specify a version, you must have versioning enabled for the S3 bucket.
	// +kubebuilder:validation:Optional
	TruststoreVersion *string `json:"truststoreVersion,omitempty" tf:"truststore_version,omitempty"`
}

// DomainNameSpec defines the desired state of DomainName
type DomainNameSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DomainNameParameters `json:"forProvider"`
}

// DomainNameStatus defines the observed state of DomainName.
type DomainNameStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DomainNameObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// DomainName is the Schema for the DomainNames API. Manages an Amazon API Gateway Version 2 domain name.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type DomainName struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.domainNameConfiguration)",message="domainNameConfiguration is a required parameter"
	Spec   DomainNameSpec   `json:"spec"`
	Status DomainNameStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DomainNameList contains a list of DomainNames
type DomainNameList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DomainName `json:"items"`
}

// Repository type metadata.
var (
	DomainName_Kind             = "DomainName"
	DomainName_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: DomainName_Kind}.String()
	DomainName_KindAPIVersion   = DomainName_Kind + "." + CRDGroupVersion.String()
	DomainName_GroupVersionKind = CRDGroupVersion.WithKind(DomainName_Kind)
)

func init() {
	SchemeBuilder.Register(&DomainName{}, &DomainNameList{})
}
