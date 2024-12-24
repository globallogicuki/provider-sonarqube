package usertoken

import "github.com/crossplane/upjet/pkg/config"

func Configure(p *config.Provider) {
	p.AddResourceConfigurator("sonarqube_user_token", func(r *config.Resource) {
		r.ShortGroup = "user"
		r.Kind = "UserToken"
		r.References["login_name"] = config.Reference{
			Type: "github.com/globallogicuki/provider-sonarqube/apis/user/v1alpha1.User",
		}
	})
}
