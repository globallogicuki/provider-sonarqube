package project

import "github.com/crossplane/upjet/pkg/config"

// Configure configures a SonarQube project resource.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("sonarqube_project", func(r *config.Resource) {
		r.ShortGroup = "project"
	})
}
