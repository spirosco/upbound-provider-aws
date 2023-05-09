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

type ActionDefinitionPublishMetricActionObservation struct {

	// Set of configuration blocks containing the dimension settings to use for Amazon CloudWatch custom metrics. See Dimension below for details.
	Dimension []PublishMetricActionDimensionObservation `json:"dimension,omitempty" tf:"dimension,omitempty"`
}

type ActionDefinitionPublishMetricActionParameters struct {

	// Set of configuration blocks containing the dimension settings to use for Amazon CloudWatch custom metrics. See Dimension below for details.
	// +kubebuilder:validation:Required
	Dimension []PublishMetricActionDimensionParameters `json:"dimension" tf:"dimension,omitempty"`
}

type CustomActionActionDefinitionObservation struct {

	// A configuration block describing the stateless inspection criteria that publishes the specified metrics to Amazon CloudWatch for the matching packet. You can pair this custom action with any of the standard stateless rule actions. See Publish Metric Action below for details.
	PublishMetricAction []ActionDefinitionPublishMetricActionObservation `json:"publishMetricAction,omitempty" tf:"publish_metric_action,omitempty"`
}

type CustomActionActionDefinitionParameters struct {

	// A configuration block describing the stateless inspection criteria that publishes the specified metrics to Amazon CloudWatch for the matching packet. You can pair this custom action with any of the standard stateless rule actions. See Publish Metric Action below for details.
	// +kubebuilder:validation:Required
	PublishMetricAction []ActionDefinitionPublishMetricActionParameters `json:"publishMetricAction" tf:"publish_metric_action,omitempty"`
}

type CustomActionObservation struct {

	// A configuration block describing the custom action associated with the action_name. See Action Definition below for details.
	ActionDefinition []CustomActionActionDefinitionObservation `json:"actionDefinition,omitempty" tf:"action_definition,omitempty"`

	// A friendly name of the custom action.
	ActionName *string `json:"actionName,omitempty" tf:"action_name,omitempty"`
}

type CustomActionParameters struct {

	// A configuration block describing the custom action associated with the action_name. See Action Definition below for details.
	// +kubebuilder:validation:Required
	ActionDefinition []CustomActionActionDefinitionParameters `json:"actionDefinition" tf:"action_definition,omitempty"`

	// A friendly name of the custom action.
	// +kubebuilder:validation:Required
	ActionName *string `json:"actionName" tf:"action_name,omitempty"`
}

type DestinationObservation struct {

	// An IP address or a block of IP addresses in CIDR notation. AWS Network Firewall supports all address ranges for IPv4.
	AddressDefinition *string `json:"addressDefinition,omitempty" tf:"address_definition,omitempty"`
}

type DestinationParameters struct {

	// An IP address or a block of IP addresses in CIDR notation. AWS Network Firewall supports all address ranges for IPv4.
	// +kubebuilder:validation:Required
	AddressDefinition *string `json:"addressDefinition" tf:"address_definition,omitempty"`
}

type DestinationPortObservation struct {

	// The lower limit of the port range. This must be less than or equal to the to_port.
	FromPort *float64 `json:"fromPort,omitempty" tf:"from_port,omitempty"`

	// The upper limit of the port range. This must be greater than or equal to the from_port.
	ToPort *float64 `json:"toPort,omitempty" tf:"to_port,omitempty"`
}

type DestinationPortParameters struct {

	// The lower limit of the port range. This must be less than or equal to the to_port.
	// +kubebuilder:validation:Required
	FromPort *float64 `json:"fromPort" tf:"from_port,omitempty"`

	// The upper limit of the port range. This must be greater than or equal to the from_port.
	// +kubebuilder:validation:Optional
	ToPort *float64 `json:"toPort,omitempty" tf:"to_port,omitempty"`
}

type HeaderObservation struct {

	// Set of configuration blocks describing the destination IP address and address ranges to inspect for, in CIDR notation. If not specified, this matches with any destination address. See Destination below for details.
	Destination *string `json:"destination,omitempty" tf:"destination,omitempty"`

	// Set of configuration blocks describing the destination ports to inspect for. If not specified, this matches with any destination port. See Destination Port below for details.
	DestinationPort *string `json:"destinationPort,omitempty" tf:"destination_port,omitempty"`

	// The direction of traffic flow to inspect. Valid values: ANY or FORWARD.
	Direction *string `json:"direction,omitempty" tf:"direction,omitempty"`

	// The protocol to inspect. Valid values: IP, TCP, UDP, ICMP, HTTP, FTP, TLS, SMB, DNS, DCERPC, SSH, SMTP, IMAP, MSN, KRB5, IKEV2, TFTP, NTP, DHCP.
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`

	// Set of configuration blocks describing the source IP address and address ranges to inspect for, in CIDR notation. If not specified, this matches with any source address. See Source below for details.
	Source *string `json:"source,omitempty" tf:"source,omitempty"`

	// Set of configuration blocks describing the source ports to inspect for. If not specified, this matches with any source port. See Source Port below for details.
	SourcePort *string `json:"sourcePort,omitempty" tf:"source_port,omitempty"`
}

type HeaderParameters struct {

	// Set of configuration blocks describing the destination IP address and address ranges to inspect for, in CIDR notation. If not specified, this matches with any destination address. See Destination below for details.
	// +kubebuilder:validation:Required
	Destination *string `json:"destination" tf:"destination,omitempty"`

	// Set of configuration blocks describing the destination ports to inspect for. If not specified, this matches with any destination port. See Destination Port below for details.
	// +kubebuilder:validation:Required
	DestinationPort *string `json:"destinationPort" tf:"destination_port,omitempty"`

	// The direction of traffic flow to inspect. Valid values: ANY or FORWARD.
	// +kubebuilder:validation:Required
	Direction *string `json:"direction" tf:"direction,omitempty"`

	// The protocol to inspect. Valid values: IP, TCP, UDP, ICMP, HTTP, FTP, TLS, SMB, DNS, DCERPC, SSH, SMTP, IMAP, MSN, KRB5, IKEV2, TFTP, NTP, DHCP.
	// +kubebuilder:validation:Required
	Protocol *string `json:"protocol" tf:"protocol,omitempty"`

	// Set of configuration blocks describing the source IP address and address ranges to inspect for, in CIDR notation. If not specified, this matches with any source address. See Source below for details.
	// +kubebuilder:validation:Required
	Source *string `json:"source" tf:"source,omitempty"`

	// Set of configuration blocks describing the source ports to inspect for. If not specified, this matches with any source port. See Source Port below for details.
	// +kubebuilder:validation:Required
	SourcePort *string `json:"sourcePort" tf:"source_port,omitempty"`
}

type IPSetObservation struct {

	// Set of port ranges.
	Definition []*string `json:"definition,omitempty" tf:"definition,omitempty"`
}

type IPSetParameters struct {

	// Set of port ranges.
	// +kubebuilder:validation:Required
	Definition []*string `json:"definition" tf:"definition,omitempty"`
}

type IPSetReferenceObservation struct {

	// Set of Managed Prefix IP ARN(s)
	ReferenceArn *string `json:"referenceArn,omitempty" tf:"reference_arn,omitempty"`
}

type IPSetReferenceParameters struct {

	// Set of Managed Prefix IP ARN(s)
	// +crossplane:generate:reference:type=github.com/spirosco/upbound-provider-aws/apis/ec2/v1beta1.ManagedPrefixList
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractParamPath("arn",true)
	// +kubebuilder:validation:Optional
	ReferenceArn *string `json:"referenceArn,omitempty" tf:"reference_arn,omitempty"`

	// Reference to a ManagedPrefixList in ec2 to populate referenceArn.
	// +kubebuilder:validation:Optional
	ReferenceArnRef *v1.Reference `json:"referenceArnRef,omitempty" tf:"-"`

	// Selector for a ManagedPrefixList in ec2 to populate referenceArn.
	// +kubebuilder:validation:Optional
	ReferenceArnSelector *v1.Selector `json:"referenceArnSelector,omitempty" tf:"-"`
}

type IPSetReferencesObservation struct {

	// Set of configuration blocks that define the IP Reference information. See IP Set Reference below for details.
	IPSetReference []IPSetReferenceObservation `json:"ipSetReference,omitempty" tf:"ip_set_reference,omitempty"`

	// An unique alphanumeric string to identify the port_set.
	Key *string `json:"key,omitempty" tf:"key,omitempty"`
}

type IPSetReferencesParameters struct {

	// Set of configuration blocks that define the IP Reference information. See IP Set Reference below for details.
	// +kubebuilder:validation:Required
	IPSetReference []IPSetReferenceParameters `json:"ipSetReference" tf:"ip_set_reference,omitempty"`

	// An unique alphanumeric string to identify the port_set.
	// +kubebuilder:validation:Required
	Key *string `json:"key" tf:"key,omitempty"`
}

type IPSetsObservation struct {

	// A configuration block that defines a set of IP addresses. See IP Set below for details.
	IPSet []IPSetObservation `json:"ipSet,omitempty" tf:"ip_set,omitempty"`

	// An unique alphanumeric string to identify the port_set.
	Key *string `json:"key,omitempty" tf:"key,omitempty"`
}

type IPSetsParameters struct {

	// A configuration block that defines a set of IP addresses. See IP Set below for details.
	// +kubebuilder:validation:Required
	IPSet []IPSetParameters `json:"ipSet" tf:"ip_set,omitempty"`

	// An unique alphanumeric string to identify the port_set.
	// +kubebuilder:validation:Required
	Key *string `json:"key" tf:"key,omitempty"`
}

type MatchAttributesObservation struct {

	// Set of configuration blocks describing the destination IP address and address ranges to inspect for, in CIDR notation. If not specified, this matches with any destination address. See Destination below for details.
	Destination []DestinationObservation `json:"destination,omitempty" tf:"destination,omitempty"`

	// Set of configuration blocks describing the destination ports to inspect for. If not specified, this matches with any destination port. See Destination Port below for details.
	DestinationPort []DestinationPortObservation `json:"destinationPort,omitempty" tf:"destination_port,omitempty"`

	// Set of protocols to inspect for, specified using the protocol's assigned internet protocol number (IANA). If not specified, this matches with any protocol.
	Protocols []*float64 `json:"protocols,omitempty" tf:"protocols,omitempty"`

	// Set of configuration blocks describing the source IP address and address ranges to inspect for, in CIDR notation. If not specified, this matches with any source address. See Source below for details.
	Source []SourceObservation `json:"source,omitempty" tf:"source,omitempty"`

	// Set of configuration blocks describing the source ports to inspect for. If not specified, this matches with any source port. See Source Port below for details.
	SourcePort []SourcePortObservation `json:"sourcePort,omitempty" tf:"source_port,omitempty"`

	// Set of configuration blocks containing the TCP flags and masks to inspect for. If not specified, this matches with any settings.
	TCPFlag []TCPFlagObservation `json:"tcpFlag,omitempty" tf:"tcp_flag,omitempty"`
}

type MatchAttributesParameters struct {

	// Set of configuration blocks describing the destination IP address and address ranges to inspect for, in CIDR notation. If not specified, this matches with any destination address. See Destination below for details.
	// +kubebuilder:validation:Optional
	Destination []DestinationParameters `json:"destination,omitempty" tf:"destination,omitempty"`

	// Set of configuration blocks describing the destination ports to inspect for. If not specified, this matches with any destination port. See Destination Port below for details.
	// +kubebuilder:validation:Optional
	DestinationPort []DestinationPortParameters `json:"destinationPort,omitempty" tf:"destination_port,omitempty"`

	// Set of protocols to inspect for, specified using the protocol's assigned internet protocol number (IANA). If not specified, this matches with any protocol.
	// +kubebuilder:validation:Optional
	Protocols []*float64 `json:"protocols,omitempty" tf:"protocols,omitempty"`

	// Set of configuration blocks describing the source IP address and address ranges to inspect for, in CIDR notation. If not specified, this matches with any source address. See Source below for details.
	// +kubebuilder:validation:Optional
	Source []SourceParameters `json:"source,omitempty" tf:"source,omitempty"`

	// Set of configuration blocks describing the source ports to inspect for. If not specified, this matches with any source port. See Source Port below for details.
	// +kubebuilder:validation:Optional
	SourcePort []SourcePortParameters `json:"sourcePort,omitempty" tf:"source_port,omitempty"`

	// Set of configuration blocks containing the TCP flags and masks to inspect for. If not specified, this matches with any settings.
	// +kubebuilder:validation:Optional
	TCPFlag []TCPFlagParameters `json:"tcpFlag,omitempty" tf:"tcp_flag,omitempty"`
}

type PortSetObservation struct {

	// Set of port ranges.
	Definition []*string `json:"definition,omitempty" tf:"definition,omitempty"`
}

type PortSetParameters struct {

	// Set of port ranges.
	// +kubebuilder:validation:Required
	Definition []*string `json:"definition" tf:"definition,omitempty"`
}

type PortSetsObservation struct {

	// An unique alphanumeric string to identify the port_set.
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	// A configuration block that defines a set of port ranges. See Port Set below for details.
	PortSet []PortSetObservation `json:"portSet,omitempty" tf:"port_set,omitempty"`
}

type PortSetsParameters struct {

	// An unique alphanumeric string to identify the port_set.
	// +kubebuilder:validation:Required
	Key *string `json:"key" tf:"key,omitempty"`

	// A configuration block that defines a set of port ranges. See Port Set below for details.
	// +kubebuilder:validation:Required
	PortSet []PortSetParameters `json:"portSet" tf:"port_set,omitempty"`
}

type PublishMetricActionDimensionObservation struct {

	// The value to use in the custom metric dimension.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type PublishMetricActionDimensionParameters struct {

	// The value to use in the custom metric dimension.
	// +kubebuilder:validation:Required
	Value *string `json:"value" tf:"value,omitempty"`
}

type ReferenceSetsObservation struct {
	IPSetReferences []IPSetReferencesObservation `json:"ipSetReferences,omitempty" tf:"ip_set_references,omitempty"`
}

type ReferenceSetsParameters struct {

	// +kubebuilder:validation:Optional
	IPSetReferences []IPSetReferencesParameters `json:"ipSetReferences,omitempty" tf:"ip_set_references,omitempty"`
}

type RuleDefinitionObservation struct {

	// Set of actions to take on a packet that matches one of the stateless rule definition's match_attributes. For every rule you must specify 1 standard action, and you can add custom actions. Standard actions include: aws:pass, aws:drop, aws:forward_to_sfe.
	Actions []*string `json:"actions,omitempty" tf:"actions,omitempty"`

	// A configuration block containing criteria for AWS Network Firewall to use to inspect an individual packet in stateless rule inspection. See Match Attributes below for details.
	MatchAttributes []MatchAttributesObservation `json:"matchAttributes,omitempty" tf:"match_attributes,omitempty"`
}

type RuleDefinitionParameters struct {

	// Set of actions to take on a packet that matches one of the stateless rule definition's match_attributes. For every rule you must specify 1 standard action, and you can add custom actions. Standard actions include: aws:pass, aws:drop, aws:forward_to_sfe.
	// +kubebuilder:validation:Required
	Actions []*string `json:"actions" tf:"actions,omitempty"`

	// A configuration block containing criteria for AWS Network Firewall to use to inspect an individual packet in stateless rule inspection. See Match Attributes below for details.
	// +kubebuilder:validation:Required
	MatchAttributes []MatchAttributesParameters `json:"matchAttributes" tf:"match_attributes,omitempty"`
}

type RuleGroupEncryptionConfigurationObservation struct {

	// The ID of the customer managed key. You can use any of the key identifiers that KMS supports, unless you're using a key that's managed by another account. If you're using a key managed by another account, then specify the key ARN.
	KeyID *string `json:"keyId,omitempty" tf:"key_id,omitempty"`

	// The type of AWS KMS key to use for encryption of your Network Firewall resources. Valid values are CUSTOMER_KMS and AWS_OWNED_KMS_KEY.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type RuleGroupEncryptionConfigurationParameters struct {

	// The ID of the customer managed key. You can use any of the key identifiers that KMS supports, unless you're using a key that's managed by another account. If you're using a key managed by another account, then specify the key ARN.
	// +kubebuilder:validation:Optional
	KeyID *string `json:"keyId,omitempty" tf:"key_id,omitempty"`

	// The type of AWS KMS key to use for encryption of your Network Firewall resources. Valid values are CUSTOMER_KMS and AWS_OWNED_KMS_KEY.
	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`
}

type RuleGroupObservation struct {

	// The Amazon Resource Name (ARN) that identifies the rule group.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// The maximum number of operating resources that this rule group can use. For a stateless rule group, the capacity required is the sum of the capacity requirements of the individual rules. For a stateful rule group, the minimum capacity required is the number of individual rules.
	Capacity *float64 `json:"capacity,omitempty" tf:"capacity,omitempty"`

	// A friendly description of the rule group.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// KMS encryption configuration settings. See Encryption Configuration below for details.
	EncryptionConfiguration []RuleGroupEncryptionConfigurationObservation `json:"encryptionConfiguration,omitempty" tf:"encryption_configuration,omitempty"`

	// The Amazon Resource Name (ARN) that identifies the rule group.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// A friendly name of the rule group.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// A configuration block that defines the rule group rules. Required unless rules is specified. See Rule Group below for details.
	RuleGroup []RuleGroupRuleGroupObservation `json:"ruleGroup,omitempty" tf:"rule_group,omitempty"`

	// The stateful rule group rules specifications in Suricata file format, with one rule per line. Use this to import your existing Suricata compatible rule groups. Required unless rule_group is specified.
	Rules *string `json:"rules,omitempty" tf:"rules,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`

	// Whether the rule group is stateless (containing stateless rules) or stateful (containing stateful rules). Valid values include: STATEFUL or STATELESS.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// A string token used when updating the rule group.
	UpdateToken *string `json:"updateToken,omitempty" tf:"update_token,omitempty"`
}

type RuleGroupParameters struct {

	// The maximum number of operating resources that this rule group can use. For a stateless rule group, the capacity required is the sum of the capacity requirements of the individual rules. For a stateful rule group, the minimum capacity required is the number of individual rules.
	// +kubebuilder:validation:Optional
	Capacity *float64 `json:"capacity,omitempty" tf:"capacity,omitempty"`

	// A friendly description of the rule group.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// KMS encryption configuration settings. See Encryption Configuration below for details.
	// +kubebuilder:validation:Optional
	EncryptionConfiguration []RuleGroupEncryptionConfigurationParameters `json:"encryptionConfiguration,omitempty" tf:"encryption_configuration,omitempty"`

	// A friendly name of the rule group.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// A configuration block that defines the rule group rules. Required unless rules is specified. See Rule Group below for details.
	// +kubebuilder:validation:Optional
	RuleGroup []RuleGroupRuleGroupParameters `json:"ruleGroup,omitempty" tf:"rule_group,omitempty"`

	// The stateful rule group rules specifications in Suricata file format, with one rule per line. Use this to import your existing Suricata compatible rule groups. Required unless rule_group is specified.
	// +kubebuilder:validation:Optional
	Rules *string `json:"rules,omitempty" tf:"rules,omitempty"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Whether the rule group is stateless (containing stateless rules) or stateful (containing stateful rules). Valid values include: STATEFUL or STATELESS.
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type RuleGroupRuleGroupObservation struct {

	// A configuration block that defines the IP Set References for the rule group. See Reference Sets below for details.
	ReferenceSets []ReferenceSetsObservation `json:"referenceSets,omitempty" tf:"reference_sets,omitempty"`

	// A configuration block that defines additional settings available to use in the rules defined in the rule group. Can only be specified for stateful rule groups. See Rule Variables below for details.
	RuleVariables []RuleVariablesObservation `json:"ruleVariables,omitempty" tf:"rule_variables,omitempty"`

	// A configuration block that defines the stateful or stateless rules for the rule group. See Rules Source below for details.
	RulesSource []RulesSourceObservation `json:"rulesSource,omitempty" tf:"rules_source,omitempty"`

	// A configuration block that defines stateful rule options for the rule group. See Stateful Rule Options below for details.
	StatefulRuleOptions []StatefulRuleOptionsObservation `json:"statefulRuleOptions,omitempty" tf:"stateful_rule_options,omitempty"`
}

type RuleGroupRuleGroupParameters struct {

	// A configuration block that defines the IP Set References for the rule group. See Reference Sets below for details.
	// +kubebuilder:validation:Optional
	ReferenceSets []ReferenceSetsParameters `json:"referenceSets,omitempty" tf:"reference_sets,omitempty"`

	// A configuration block that defines additional settings available to use in the rules defined in the rule group. Can only be specified for stateful rule groups. See Rule Variables below for details.
	// +kubebuilder:validation:Optional
	RuleVariables []RuleVariablesParameters `json:"ruleVariables,omitempty" tf:"rule_variables,omitempty"`

	// A configuration block that defines the stateful or stateless rules for the rule group. See Rules Source below for details.
	// +kubebuilder:validation:Required
	RulesSource []RulesSourceParameters `json:"rulesSource" tf:"rules_source,omitempty"`

	// A configuration block that defines stateful rule options for the rule group. See Stateful Rule Options below for details.
	// +kubebuilder:validation:Optional
	StatefulRuleOptions []StatefulRuleOptionsParameters `json:"statefulRuleOptions,omitempty" tf:"stateful_rule_options,omitempty"`
}

type RuleOptionObservation struct {

	// Keyword defined by open source detection systems like Snort or Suricata for stateful rule inspection.
	// See Snort General Rule Options or Suricata Rule Options for more details.
	Keyword *string `json:"keyword,omitempty" tf:"keyword,omitempty"`

	// Set of strings for additional settings to use in stateful rule inspection.
	Settings []*string `json:"settings,omitempty" tf:"settings,omitempty"`
}

type RuleOptionParameters struct {

	// Keyword defined by open source detection systems like Snort or Suricata for stateful rule inspection.
	// See Snort General Rule Options or Suricata Rule Options for more details.
	// +kubebuilder:validation:Required
	Keyword *string `json:"keyword" tf:"keyword,omitempty"`

	// Set of strings for additional settings to use in stateful rule inspection.
	// +kubebuilder:validation:Optional
	Settings []*string `json:"settings,omitempty" tf:"settings,omitempty"`
}

type RuleVariablesObservation struct {

	// Set of configuration blocks that define IP address information. See IP Sets below for details.
	IPSets []IPSetsObservation `json:"ipSets,omitempty" tf:"ip_sets,omitempty"`

	// Set of configuration blocks that define port range information. See Port Sets below for details.
	PortSets []PortSetsObservation `json:"portSets,omitempty" tf:"port_sets,omitempty"`
}

type RuleVariablesParameters struct {

	// Set of configuration blocks that define IP address information. See IP Sets below for details.
	// +kubebuilder:validation:Optional
	IPSets []IPSetsParameters `json:"ipSets,omitempty" tf:"ip_sets,omitempty"`

	// Set of configuration blocks that define port range information. See Port Sets below for details.
	// +kubebuilder:validation:Optional
	PortSets []PortSetsParameters `json:"portSets,omitempty" tf:"port_sets,omitempty"`
}

type RulesSourceListObservation struct {

	// String value to specify whether domains in the target list are allowed or denied access. Valid values: ALLOWLIST, DENYLIST.
	GeneratedRulesType *string `json:"generatedRulesType,omitempty" tf:"generated_rules_type,omitempty"`

	// Set of types of domain specifications that are provided in the targets argument. Valid values: HTTP_HOST, TLS_SNI.
	TargetTypes []*string `json:"targetTypes,omitempty" tf:"target_types,omitempty"`

	// Set of domains that you want to inspect for in your traffic flows.
	Targets []*string `json:"targets,omitempty" tf:"targets,omitempty"`
}

type RulesSourceListParameters struct {

	// String value to specify whether domains in the target list are allowed or denied access. Valid values: ALLOWLIST, DENYLIST.
	// +kubebuilder:validation:Required
	GeneratedRulesType *string `json:"generatedRulesType" tf:"generated_rules_type,omitempty"`

	// Set of types of domain specifications that are provided in the targets argument. Valid values: HTTP_HOST, TLS_SNI.
	// +kubebuilder:validation:Required
	TargetTypes []*string `json:"targetTypes" tf:"target_types,omitempty"`

	// Set of domains that you want to inspect for in your traffic flows.
	// +kubebuilder:validation:Required
	Targets []*string `json:"targets" tf:"targets,omitempty"`
}

type RulesSourceObservation struct {

	// A configuration block containing stateful inspection criteria for a domain list rule group. See Rules Source List below for details.
	RulesSourceList []RulesSourceListObservation `json:"rulesSourceList,omitempty" tf:"rules_source_list,omitempty"`

	// The fully qualified name of a file in an S3 bucket that contains Suricata compatible intrusion preventions system (IPS) rules or the Suricata rules as a string. These rules contain stateful inspection criteria and the action to take for traffic that matches the criteria.
	RulesString *string `json:"rulesString,omitempty" tf:"rules_string,omitempty"`

	// Set of configuration blocks containing stateful inspection criteria for 5-tuple rules to be used together in a rule group. See Stateful Rule below for details.
	StatefulRule []StatefulRuleObservation `json:"statefulRule,omitempty" tf:"stateful_rule,omitempty"`

	// A configuration block containing stateless inspection criteria for a stateless rule group. See Stateless Rules and Custom Actions below for details.
	StatelessRulesAndCustomActions []StatelessRulesAndCustomActionsObservation `json:"statelessRulesAndCustomActions,omitempty" tf:"stateless_rules_and_custom_actions,omitempty"`
}

type RulesSourceParameters struct {

	// A configuration block containing stateful inspection criteria for a domain list rule group. See Rules Source List below for details.
	// +kubebuilder:validation:Optional
	RulesSourceList []RulesSourceListParameters `json:"rulesSourceList,omitempty" tf:"rules_source_list,omitempty"`

	// The fully qualified name of a file in an S3 bucket that contains Suricata compatible intrusion preventions system (IPS) rules or the Suricata rules as a string. These rules contain stateful inspection criteria and the action to take for traffic that matches the criteria.
	// +kubebuilder:validation:Optional
	RulesString *string `json:"rulesString,omitempty" tf:"rules_string,omitempty"`

	// Set of configuration blocks containing stateful inspection criteria for 5-tuple rules to be used together in a rule group. See Stateful Rule below for details.
	// +kubebuilder:validation:Optional
	StatefulRule []StatefulRuleParameters `json:"statefulRule,omitempty" tf:"stateful_rule,omitempty"`

	// A configuration block containing stateless inspection criteria for a stateless rule group. See Stateless Rules and Custom Actions below for details.
	// +kubebuilder:validation:Optional
	StatelessRulesAndCustomActions []StatelessRulesAndCustomActionsParameters `json:"statelessRulesAndCustomActions,omitempty" tf:"stateless_rules_and_custom_actions,omitempty"`
}

type SourceObservation struct {

	// An IP address or a block of IP addresses in CIDR notation. AWS Network Firewall supports all address ranges for IPv4.
	AddressDefinition *string `json:"addressDefinition,omitempty" tf:"address_definition,omitempty"`
}

type SourceParameters struct {

	// An IP address or a block of IP addresses in CIDR notation. AWS Network Firewall supports all address ranges for IPv4.
	// +kubebuilder:validation:Required
	AddressDefinition *string `json:"addressDefinition" tf:"address_definition,omitempty"`
}

type SourcePortObservation struct {

	// The lower limit of the port range. This must be less than or equal to the to_port.
	FromPort *float64 `json:"fromPort,omitempty" tf:"from_port,omitempty"`

	// The upper limit of the port range. This must be greater than or equal to the from_port.
	ToPort *float64 `json:"toPort,omitempty" tf:"to_port,omitempty"`
}

type SourcePortParameters struct {

	// The lower limit of the port range. This must be less than or equal to the to_port.
	// +kubebuilder:validation:Required
	FromPort *float64 `json:"fromPort" tf:"from_port,omitempty"`

	// The upper limit of the port range. This must be greater than or equal to the from_port.
	// +kubebuilder:validation:Optional
	ToPort *float64 `json:"toPort,omitempty" tf:"to_port,omitempty"`
}

type StatefulRuleObservation struct {

	// Action to take with packets in a traffic flow when the flow matches the stateful rule criteria. For all actions, AWS Network Firewall performs the specified action and discontinues stateful inspection of the traffic flow. Valid values: ALERT, DROP or PASS.
	Action *string `json:"action,omitempty" tf:"action,omitempty"`

	// A configuration block containing the stateful 5-tuple inspection criteria for the rule, used to inspect traffic flows. See Header below for details.
	Header []HeaderObservation `json:"header,omitempty" tf:"header,omitempty"`

	// Set of configuration blocks containing additional settings for a stateful rule. See Rule Option below for details.
	RuleOption []RuleOptionObservation `json:"ruleOption,omitempty" tf:"rule_option,omitempty"`
}

type StatefulRuleOptionsObservation struct {

	// Indicates how to manage the order of the rule evaluation for the rule group. Default value: DEFAULT_ACTION_ORDER. Valid values: DEFAULT_ACTION_ORDER, STRICT_ORDER.
	RuleOrder *string `json:"ruleOrder,omitempty" tf:"rule_order,omitempty"`
}

type StatefulRuleOptionsParameters struct {

	// Indicates how to manage the order of the rule evaluation for the rule group. Default value: DEFAULT_ACTION_ORDER. Valid values: DEFAULT_ACTION_ORDER, STRICT_ORDER.
	// +kubebuilder:validation:Required
	RuleOrder *string `json:"ruleOrder" tf:"rule_order,omitempty"`
}

type StatefulRuleParameters struct {

	// Action to take with packets in a traffic flow when the flow matches the stateful rule criteria. For all actions, AWS Network Firewall performs the specified action and discontinues stateful inspection of the traffic flow. Valid values: ALERT, DROP or PASS.
	// +kubebuilder:validation:Required
	Action *string `json:"action" tf:"action,omitempty"`

	// A configuration block containing the stateful 5-tuple inspection criteria for the rule, used to inspect traffic flows. See Header below for details.
	// +kubebuilder:validation:Required
	Header []HeaderParameters `json:"header" tf:"header,omitempty"`

	// Set of configuration blocks containing additional settings for a stateful rule. See Rule Option below for details.
	// +kubebuilder:validation:Required
	RuleOption []RuleOptionParameters `json:"ruleOption" tf:"rule_option,omitempty"`
}

type StatelessRuleObservation struct {

	// A setting that indicates the order in which to run this rule relative to all of the rules that are defined for a stateless rule group. AWS Network Firewall evaluates the rules in a rule group starting with the lowest priority setting.
	Priority *float64 `json:"priority,omitempty" tf:"priority,omitempty"`

	// A configuration block defining the stateless 5-tuple packet inspection criteria and the action to take on a packet that matches the criteria. See Rule Definition below for details.
	RuleDefinition []RuleDefinitionObservation `json:"ruleDefinition,omitempty" tf:"rule_definition,omitempty"`
}

type StatelessRuleParameters struct {

	// A setting that indicates the order in which to run this rule relative to all of the rules that are defined for a stateless rule group. AWS Network Firewall evaluates the rules in a rule group starting with the lowest priority setting.
	// +kubebuilder:validation:Required
	Priority *float64 `json:"priority" tf:"priority,omitempty"`

	// A configuration block defining the stateless 5-tuple packet inspection criteria and the action to take on a packet that matches the criteria. See Rule Definition below for details.
	// +kubebuilder:validation:Required
	RuleDefinition []RuleDefinitionParameters `json:"ruleDefinition" tf:"rule_definition,omitempty"`
}

type StatelessRulesAndCustomActionsObservation struct {

	// Set of configuration blocks containing custom action definitions that are available for use by the set of stateless rule. See Custom Action below for details.
	CustomAction []CustomActionObservation `json:"customAction,omitempty" tf:"custom_action,omitempty"`

	// Set of configuration blocks containing the stateless rules for use in the stateless rule group. See Stateless Rule below for details.
	StatelessRule []StatelessRuleObservation `json:"statelessRule,omitempty" tf:"stateless_rule,omitempty"`
}

type StatelessRulesAndCustomActionsParameters struct {

	// Set of configuration blocks containing custom action definitions that are available for use by the set of stateless rule. See Custom Action below for details.
	// +kubebuilder:validation:Optional
	CustomAction []CustomActionParameters `json:"customAction,omitempty" tf:"custom_action,omitempty"`

	// Set of configuration blocks containing the stateless rules for use in the stateless rule group. See Stateless Rule below for details.
	// +kubebuilder:validation:Required
	StatelessRule []StatelessRuleParameters `json:"statelessRule" tf:"stateless_rule,omitempty"`
}

type TCPFlagObservation struct {

	// Set of flags to look for in a packet. This setting can only specify values that are also specified in masks.
	// Valid values: FIN, SYN, RST, PSH, ACK, URG, ECE, CWR.
	Flags []*string `json:"flags,omitempty" tf:"flags,omitempty"`

	// Set of flags to consider in the inspection. To inspect all flags, leave this empty.
	// Valid values: FIN, SYN, RST, PSH, ACK, URG, ECE, CWR.
	Masks []*string `json:"masks,omitempty" tf:"masks,omitempty"`
}

type TCPFlagParameters struct {

	// Set of flags to look for in a packet. This setting can only specify values that are also specified in masks.
	// Valid values: FIN, SYN, RST, PSH, ACK, URG, ECE, CWR.
	// +kubebuilder:validation:Required
	Flags []*string `json:"flags" tf:"flags,omitempty"`

	// Set of flags to consider in the inspection. To inspect all flags, leave this empty.
	// Valid values: FIN, SYN, RST, PSH, ACK, URG, ECE, CWR.
	// +kubebuilder:validation:Optional
	Masks []*string `json:"masks,omitempty" tf:"masks,omitempty"`
}

// RuleGroupSpec defines the desired state of RuleGroup
type RuleGroupSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RuleGroupParameters `json:"forProvider"`
}

// RuleGroupStatus defines the observed state of RuleGroup.
type RuleGroupStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RuleGroupObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// RuleGroup is the Schema for the RuleGroups API. Provides an AWS Network Firewall Rule Group resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type RuleGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.capacity)",message="capacity is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.name)",message="name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.type)",message="type is a required parameter"
	Spec   RuleGroupSpec   `json:"spec"`
	Status RuleGroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RuleGroupList contains a list of RuleGroups
type RuleGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RuleGroup `json:"items"`
}

// Repository type metadata.
var (
	RuleGroup_Kind             = "RuleGroup"
	RuleGroup_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: RuleGroup_Kind}.String()
	RuleGroup_KindAPIVersion   = RuleGroup_Kind + "." + CRDGroupVersion.String()
	RuleGroup_GroupVersionKind = CRDGroupVersion.WithKind(RuleGroup_Kind)
)

func init() {
	SchemeBuilder.Register(&RuleGroup{}, &RuleGroupList{})
}
