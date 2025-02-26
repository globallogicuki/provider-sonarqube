// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	project "github.com/globallogicuki/provider-sonarqube/internal/controller/project/project"
	providerconfig "github.com/globallogicuki/provider-sonarqube/internal/controller/providerconfig"
	rule "github.com/globallogicuki/provider-sonarqube/internal/controller/rule/rule"
	setting "github.com/globallogicuki/provider-sonarqube/internal/controller/settings/setting"
	user "github.com/globallogicuki/provider-sonarqube/internal/controller/user/user"
	usertoken "github.com/globallogicuki/provider-sonarqube/internal/controller/user/usertoken"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		project.Setup,
		providerconfig.Setup,
		rule.Setup,
		setting.Setup,
		user.Setup,
		usertoken.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
