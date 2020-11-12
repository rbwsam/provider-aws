/*
Copyright 2020 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1alpha1

import (
	"context"

	"github.com/crossplane/crossplane-runtime/pkg/reference"
	"github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/pkg/errors"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// APIID extracts the resolved API's ID.
func APIID() reference.ExtractValueFn {
	return func(mg resource.Managed) string {
		cr, ok := mg.(*API)
		if !ok {
			return ""
		}
		return reference.FromPtrValue(cr.Status.AtProvider.APIID)
	}
}

// ResolveReferences of this Cluster
func (mg *Stage) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	// Resolve spec.forProvider.apiId
	rsp, err := r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.APIID),
		Reference:    mg.Spec.ForProvider.APIIDRef,
		Selector:     mg.Spec.ForProvider.APIIDSelector,
		To:           reference.To{Managed: &API{}, List: &APIList{}},
		Extract:      APIID(),
	})
	if err != nil {
		return errors.Wrap(err, "spec.forProvider.apiId")
	}
	mg.Spec.ForProvider.APIID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.APIIDRef = rsp.ResolvedReference
	return nil
}

// ResolveReferences of this Cluster
func (mg *Route) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	// Resolve spec.forProvider.apiId
	rsp, err := r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.APIID),
		Reference:    mg.Spec.ForProvider.APIIDRef,
		Selector:     mg.Spec.ForProvider.APIIDSelector,
		To:           reference.To{Managed: &API{}, List: &APIList{}},
		Extract:      APIID(),
	})
	if err != nil {
		return errors.Wrap(err, "spec.forProvider.apiId")
	}
	mg.Spec.ForProvider.APIID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.APIIDRef = rsp.ResolvedReference
	return nil
}
