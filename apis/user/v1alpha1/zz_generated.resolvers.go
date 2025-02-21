// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by angryjet. DO NOT EDIT.

package v1alpha1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	errors "github.com/pkg/errors"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this UserToken.
func (mg *UserToken) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.LoginName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.LoginNameRef,
		Selector:     mg.Spec.ForProvider.LoginNameSelector,
		To: reference.To{
			List:    &UserList{},
			Managed: &User{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.LoginName")
	}
	mg.Spec.ForProvider.LoginName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.LoginNameRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.LoginName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.LoginNameRef,
		Selector:     mg.Spec.InitProvider.LoginNameSelector,
		To: reference.To{
			List:    &UserList{},
			Managed: &User{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.LoginName")
	}
	mg.Spec.InitProvider.LoginName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.LoginNameRef = rsp.ResolvedReference

	return nil
}
