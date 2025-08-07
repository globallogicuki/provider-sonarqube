/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	ujconfig "github.com/crossplane/upjet/pkg/config"
	"github.com/globallogicuki/provider-sonarqube/config/almazure"
	"github.com/globallogicuki/provider-sonarqube/config/almgithub"
	"github.com/globallogicuki/provider-sonarqube/config/almgitlab"
	"github.com/globallogicuki/provider-sonarqube/config/group"
	"github.com/globallogicuki/provider-sonarqube/config/permissions"
	"github.com/globallogicuki/provider-sonarqube/config/permissiontemplate"
	"github.com/globallogicuki/provider-sonarqube/config/project"
	"github.com/globallogicuki/provider-sonarqube/config/rule"
	"github.com/globallogicuki/provider-sonarqube/config/setting"
	"github.com/globallogicuki/provider-sonarqube/config/user"
	"github.com/globallogicuki/provider-sonarqube/config/usertoken"
)

const (
	resourcePrefix = "sonarqube"
	modulePath     = "github.com/globallogicuki/provider-sonarqube"
)

//go:embed schema.json
var providerSchema string

//go:embed provider-metadata.yaml
var providerMetadata string

// GetProvider returns provider configuration
func GetProvider() *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithRootGroup("sonarqube.crossplane.io"),
		ujconfig.WithIncludeList(ExternalNameConfigured()),
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
		))

	for _, configure := range []func(provider *ujconfig.Provider){
		// add custom config functions
		project.Configure,
		user.Configure,
		usertoken.Configure,
		rule.Configure,
		setting.Configure,
		almazure.Configure,
		almgithub.Configure,
		almgitlab.Configure,
		permissiontemplate.Configure,
		permissions.Configure,
		group.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
