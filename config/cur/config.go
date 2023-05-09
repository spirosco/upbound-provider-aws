/*
Copyright 2022 Upbound Inc.
*/

package cur

import (
	"github.com/upbound/upjet/pkg/config"
)

// Configure adds configurations for cur group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("aws_cur_report_definition", func(r *config.Resource) {
		r.References["s3_bucket"] = config.Reference{
			Type: "github.com/spirosco/upbound-provider-aws/apis/s3/v1beta1.Bucket",
		}
	})
}
