package licensemanager

import (
	"github.com/upbound/upjet/pkg/config"

	"github.com/spirosco/upbound-provider-aws/config/common"
)

// Configure adds configurations for licensemanager group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("aws_licensemanager_association", func(r *config.Resource) {
		r.References["license_configuration_arn"] = config.Reference{
			Type:      "LicenseConfiguration",
			Extractor: common.PathARNExtractor,
		}
	})
}
