package almazure

import "github.com/crossplane/upjet/pkg/config"

// Configure configures a SonarQube Azure DevOps ALM/DevOps Platform Integration resource.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("sonarqube_alm_azure", func(r *config.Resource) {
		r.ShortGroup = "alm"
		r.Kind = "Azure"
	})
}
