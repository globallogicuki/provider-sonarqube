package permissions

import "github.com/crossplane/upjet/pkg/config"

// Configure configures a SonarQube user resource.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("sonarqube_permissions", func(r *config.Resource) {
		r.ShortGroup = "permissions"
		r.Kind = "Permissions"
		r.References["template_id"] = config.Reference{
			TerraformName:     "sonarqube_permission_template",
			RefFieldName:      "TemplateIDRef",
			SelectorFieldName: "TemplateIDSelector",
		}
		r.References["group_name"] = config.Reference{
			TerraformName:     "sonarqube_group",
			RefFieldName:      "GroupNameRef",
			SelectorFieldName: "GroupNameSelector",
		}
	})
}
