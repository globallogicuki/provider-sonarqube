package almgitlab

import "github.com/crossplane/upjet/pkg/config"

// Configure configures a SonarQube GitLab ALM/DevOps Platform Integration resource.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("sonarqube_alm_gitlab", func(r *config.Resource) {
		r.ShortGroup = "alm"
		r.Kind = "GitLab"
	})
}
