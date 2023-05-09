/*
Copyright 2022 Upbound Inc.
*/

package kms

import (
	"github.com/upbound/upjet/pkg/config"

	"github.com/spirosco/upbound-provider-aws/config/common"
)

// Configure adds configurations for kms group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("aws_kms_alias", func(r *config.Resource) {
		r.References["target_key_id"] = config.Reference{
			Type: "Key",
		}
		r.UseAsync = true
	})

	p.AddResourceConfigurator("aws_kms_ciphertext", func(r *config.Resource) {
		r.References["key_id"] = config.Reference{
			Type: "Key",
		}
		r.UseAsync = true
	})

	p.AddResourceConfigurator("aws_kms_grant", func(r *config.Resource) {
		r.References["key_id"] = config.Reference{
			Type:      "Key",
			Extractor: common.PathARNExtractor,
		}
	})

	p.AddResourceConfigurator("aws_kms_replica_key", func(r *config.Resource) {
		r.References["primary_key_arn"] = config.Reference{
			Type:      "Key",
			Extractor: common.PathARNExtractor,
		}
	})

	p.AddResourceConfigurator("aws_kms_replica_external_key", func(r *config.Resource) {
		r.References["primary_key_arn"] = config.Reference{
			Type:      "ExternalKey",
			Extractor: common.PathARNExtractor,
		}
	})
}
