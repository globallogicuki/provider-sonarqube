package setting

import "github.com/crossplane/upjet/pkg/config"

// Configure configures a SonarQube Settings resource.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("sonarqube_setting", func(r *config.Resource) {
		r.ShortGroup = "settings"
		r.Kind = "Setting"
	})
}
