package group

import "github.com/crossplane/upjet/pkg/config"

// Configure configures a SonarQube group resource.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("sonarqube_group", func(r *config.Resource) {
		r.ShortGroup = "group"
		r.Kind = "Group"
	})
}
