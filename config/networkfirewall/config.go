/*
Copyright 2022 Upbound Inc.
*/

package networkfirewall

import (
	"github.com/upbound/upjet/pkg/config"

	"github.com/spirosco/upbound-provider-aws/config/common"
)

// Configure adds configurations for networkfirewall group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("aws_networkfirewall_firewall_policy", func(r *config.Resource) {
		r.References = config.References{
			"firewall_policy.stateless_rule_group_reference.resource_arn": {
				TerraformName: "aws_networkfirewall_rule_group",
				Extractor:     common.PathARNExtractor,
			},
			"firewall_policy.stateful_rule_group_reference.resource_arn": {
				TerraformName: "aws_networkfirewall_rule_group",
				Extractor:     common.PathARNExtractor,
			},
		}
	})
}
