package almgithub

import "github.com/crossplane/upjet/pkg/config"

// Configure configures a SonarQube GitHub ALM/DevOps Platform Integration resource.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("sonarqube_alm_github", func(r *config.Resource) {
		r.ShortGroup = "alm"
		r.Kind = "GitHub"
	})
}
