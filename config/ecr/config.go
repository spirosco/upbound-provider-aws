/*
Copyright 2021 Upbound Inc.
*/

package ecr

import (
	"github.com/upbound/upjet/pkg/config"

	"github.com/spirosco/upbound-provider-aws/config/common"
)

// Configure adds configurations for ecrs group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("aws_ecr_repository", func(r *config.Resource) {
		r.References = map[string]config.Reference{
			"encryption_configuration.kms_key": {
				Type:      "github.com/spirosco/upbound-provider-aws/apis/kms/v1beta1.Key",
				Extractor: common.PathARNExtractor,
			},
		}
		// Deletion takes a while.
		r.UseAsync = true
	})
}
