package rule

import "github.com/crossplane/upjet/pkg/config"

// Configure configures a SonarQube Rules resource.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("sonarqube_rule", func(r *config.Resource) {
		r.ShortGroup = "rule"
		r.Kind = "Rule"
	})
}
