package azurebinding

import "github.com/crossplane/upjet/pkg/config"

// Configure configures a Sonarqube Azure Devops binding resource.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("sonarqube_azure_binding", func(r *config.Resource) {
		r.ShortGroup = "alm"
		r.Kind = "AzureBinding"
		r.References["alm_setting"] = config.Reference{
			TerraformName:     "sonarqube_alm_azure",
			RefFieldName:      "AlmSettingRef",
			SelectorFieldName: "AlmSettingSelector",
		}
		r.References["project"] = config.Reference{
			TerraformName:     "sonarqube_project",
			RefFieldName:      "ProjectRef",
			SelectorFieldName: "ProjectSelector",
		}
	})
}
