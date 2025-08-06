package permissiontemplate

import "github.com/crossplane/upjet/pkg/config"

// Configure configures a SonarQube user resource.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("sonarqube_permission_template", func(r *config.Resource) {
		r.ShortGroup = "permissiontemplate"
		r.Kind = "PermissionTemplate"
	})
}
