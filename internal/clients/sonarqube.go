/*
Copyright 2021 Upbound Inc.
*/

package clients

import (
	"context"
	"encoding/json"

	"github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/crossplane/upjet/pkg/terraform"

	"github.com/globallogicuki/provider-sonarqube/apis/v1beta1"
)

const (
	// error messages
	errNoProviderConfig     = "no providerConfigRef provided"
	errGetProviderConfig    = "cannot get referenced ProviderConfig"
	errTrackUsage           = "cannot track ProviderConfig usage"
	errExtractCredentials   = "cannot extract credentials"
	errUnmarshalCredentials = "cannot unmarshal SonarQube credentials as JSON"

	// provider config variables
	user                  = "user"
	pass                  = "pass"
	token                 = "token"
	host                  = "host"
	installedVersion      = "installed_version"
	tlsInsecureSkipVerify = "tls_insecure_skip_verify"
	anonymizeUserOnDelete = "anonymize_user_on_delete"
)

type sonarqubeConfig struct {
	User                  *string `json:"user,omitempty"`
	Pass                  *string `json:"pass,omitempty"`
	Token                 *string `json:"token,omitempty"`
	Host                  *string `json:"host"`
	InstalledVersion      *string `json:"installed_version,omitempty"`
	TLSInsecureSkipVerify *bool   `json:"tls_insecure_skip_verify,omitempty"`
	AnonymizeUserOnDelete *bool   `json:"anonymize_user_on_delete,omitempty"`
}

func terraformProviderConfigurationBuilder(creds sonarqubeConfig) terraform.ProviderConfiguration {
	cnf := terraform.ProviderConfiguration{}

	if creds.User != nil {
		cnf[user] = *creds.User
	}

	if creds.Pass != nil {
		cnf[pass] = *creds.Pass
	}

	if creds.Token != nil {
		cnf[token] = *creds.Token
	}

	if creds.Host != nil {
		cnf[host] = *creds.Host
	}

	if creds.InstalledVersion != nil {
		cnf[installedVersion] = *creds.InstalledVersion
	}

	if creds.TLSInsecureSkipVerify != nil {
		cnf[tlsInsecureSkipVerify] = *creds.TLSInsecureSkipVerify
	}

	if creds.AnonymizeUserOnDelete != nil {
		cnf[anonymizeUserOnDelete] = *creds.AnonymizeUserOnDelete
	}

	return cnf
}

// TerraformSetupBuilder builds Terraform a terraform.SetupFn function which
// returns Terraform provider setup configuration
func TerraformSetupBuilder(version, providerSource, providerVersion string) terraform.SetupFn {
	return func(ctx context.Context, client client.Client, mg resource.Managed) (terraform.Setup, error) {
		ps := terraform.Setup{
			Version: version,
			Requirement: terraform.ProviderRequirement{
				Source:  providerSource,
				Version: providerVersion,
			},
		}

		configRef := mg.GetProviderConfigReference()
		if configRef == nil {
			return ps, errors.New(errNoProviderConfig)
		}
		pc := &v1beta1.ProviderConfig{}
		if err := client.Get(ctx, types.NamespacedName{Name: configRef.Name}, pc); err != nil {
			return ps, errors.Wrap(err, errGetProviderConfig)
		}

		t := resource.NewProviderConfigUsageTracker(client, &v1beta1.ProviderConfigUsage{})
		if err := t.Track(ctx, mg); err != nil {
			return ps, errors.Wrap(err, errTrackUsage)
		}

		data, err := resource.CommonCredentialExtractor(
			ctx,
			pc.Spec.Credentials.Source,
			client,
			pc.Spec.Credentials.CommonCredentialSelectors,
		)
		if err != nil {
			return ps, errors.Wrap(err, errExtractCredentials)
		}
		creds := sonarqubeConfig{}
		if err := json.Unmarshal(data, &creds); err != nil {
			return ps, errors.Wrap(err, errUnmarshalCredentials)
		}

		ps.Configuration = terraformProviderConfigurationBuilder(creds)

		return ps, nil
	}
}
