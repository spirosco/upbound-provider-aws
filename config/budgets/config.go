package budgets

import (
	"github.com/upbound/upjet/pkg/config"
)

// Configure adds configurations for budgets group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("aws_budgets_budget_action", func(r *config.Resource) {
		r.References["definition.iam_action_definition.aws_iam_role.example.name"] = config.Reference{
			Type: "github.com/spirosco/upbound-provider-aws/apis/iam/v1beta1.Role",
		}
	})
}
