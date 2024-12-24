package user

import "github.com/crossplane/upjet/pkg/config"

func Configure(p *config.Provider) {
	p.AddResourceConfigurator("sonarqube_user", func(r *config.Resource) {
		r.ShortGroup = "user"
		r.Kind = "User"
	})
}
