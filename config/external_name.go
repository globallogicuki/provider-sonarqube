/*
Copyright 2022 Upbound Inc.
*/

package config

import "github.com/crossplane/upjet/pkg/config"

// ExternalNameConfigs contains all external name configurations for this
// provider.
var ExternalNameConfigs = map[string]config.ExternalName{
	"sonarqube_project":             config.IdentifierFromProvider,
	"sonarqube_user":                config.IdentifierFromProvider,
	"sonarqube_user_token":          config.IdentifierFromProvider,
	"sonarqube_rule":                config.IdentifierFromProvider,
	"sonarqube_setting":             config.IdentifierFromProvider,
	"sonarqube_alm_azure":           config.IdentifierFromProvider,
	"sonarqube_alm_github":          config.IdentifierFromProvider,
	"sonarqube_alm_gitlab":          config.IdentifierFromProvider,
	"sonarqube_permission_template": config.IdentifierFromProvider,
	"sonarqube_permissions":         config.IdentifierFromProvider,
	"sonarqube_group":               config.IdentifierFromProvider,
}

// ExternalNameConfigurations applies all external name configs listed in the
// table ExternalNameConfigs and sets the version of those resources to v1beta1
// assuming they will be tested.
func ExternalNameConfigurations() config.ResourceOption {
	return func(r *config.Resource) {
		if e, ok := ExternalNameConfigs[r.Name]; ok {
			r.ExternalName = e
		}
	}
}

// ExternalNameConfigured returns the list of all resources whose external name
// is configured manually.
func ExternalNameConfigured() []string {
	l := make([]string, len(ExternalNameConfigs))
	i := 0
	for name := range ExternalNameConfigs {
		// $ is added to match the exact string since the format is regex.
		l[i] = name + "$"
		i++
	}
	return l
}
