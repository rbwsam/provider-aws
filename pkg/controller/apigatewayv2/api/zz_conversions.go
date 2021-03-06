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

// Code generated by ack-generate. DO NOT EDIT.

package api

import (
	"github.com/aws/aws-sdk-go/aws/awserr"
	svcsdk "github.com/aws/aws-sdk-go/service/apigatewayv2"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	svcapitypes "github.com/crossplane/provider-aws/apis/apigatewayv2/v1alpha1"
)

// NOTE(muvaf): We return pointers in case the function needs to start with an
// empty object, hence need to return a new pointer.
// TODO(muvaf): We can generate one-time boilerplate for these hooks but currently
// ACK doesn't support not generating if file exists.
// GenerateGetApisInput returns input for read
// operation.
func GenerateGetApisInput(cr *svcapitypes.API) *svcsdk.GetApisInput {
	res := preGenerateGetApisInput(cr, &svcsdk.GetApisInput{})

	return postGenerateGetApisInput(cr, res)
}

// GenerateAPI returns the current state in the form of *svcapitypes.API.
func GenerateAPI(resp *svcsdk.GetApisOutput) *svcapitypes.API {
	cr := &svcapitypes.API{}

	found := false
	for _, elem := range resp.Items {
		if elem.ApiEndpoint != nil {
			cr.Status.AtProvider.APIEndpoint = elem.ApiEndpoint
		}
		if elem.ApiGatewayManaged != nil {
			cr.Status.AtProvider.APIGatewayManaged = elem.ApiGatewayManaged
		}
		if elem.ApiId != nil {
			cr.Status.AtProvider.APIID = elem.ApiId
		}
		if elem.ApiKeySelectionExpression != nil {
			cr.Spec.ForProvider.APIKeySelectionExpression = elem.ApiKeySelectionExpression
		}
		if elem.CorsConfiguration != nil {
			f4 := &svcapitypes.Cors{}
			if elem.CorsConfiguration.AllowCredentials != nil {
				f4.AllowCredentials = elem.CorsConfiguration.AllowCredentials
			}
			if elem.CorsConfiguration.AllowHeaders != nil {
				f4f1 := []*string{}
				for _, f4f1iter := range elem.CorsConfiguration.AllowHeaders {
					var f4f1elem string
					f4f1elem = *f4f1iter
					f4f1 = append(f4f1, &f4f1elem)
				}
				f4.AllowHeaders = f4f1
			}
			if elem.CorsConfiguration.AllowMethods != nil {
				f4f2 := []*string{}
				for _, f4f2iter := range elem.CorsConfiguration.AllowMethods {
					var f4f2elem string
					f4f2elem = *f4f2iter
					f4f2 = append(f4f2, &f4f2elem)
				}
				f4.AllowMethods = f4f2
			}
			if elem.CorsConfiguration.AllowOrigins != nil {
				f4f3 := []*string{}
				for _, f4f3iter := range elem.CorsConfiguration.AllowOrigins {
					var f4f3elem string
					f4f3elem = *f4f3iter
					f4f3 = append(f4f3, &f4f3elem)
				}
				f4.AllowOrigins = f4f3
			}
			if elem.CorsConfiguration.ExposeHeaders != nil {
				f4f4 := []*string{}
				for _, f4f4iter := range elem.CorsConfiguration.ExposeHeaders {
					var f4f4elem string
					f4f4elem = *f4f4iter
					f4f4 = append(f4f4, &f4f4elem)
				}
				f4.ExposeHeaders = f4f4
			}
			if elem.CorsConfiguration.MaxAge != nil {
				f4.MaxAge = elem.CorsConfiguration.MaxAge
			}
			cr.Spec.ForProvider.CorsConfiguration = f4
		}
		if elem.CreatedDate != nil {
			cr.Status.AtProvider.CreatedDate = &metav1.Time{*elem.CreatedDate}
		}
		if elem.Description != nil {
			cr.Spec.ForProvider.Description = elem.Description
		}
		if elem.DisableExecuteApiEndpoint != nil {
			cr.Spec.ForProvider.DisableExecuteAPIEndpoint = elem.DisableExecuteApiEndpoint
		}
		if elem.DisableSchemaValidation != nil {
			cr.Spec.ForProvider.DisableSchemaValidation = elem.DisableSchemaValidation
		}
		if elem.ImportInfo != nil {
			f9 := []*string{}
			for _, f9iter := range elem.ImportInfo {
				var f9elem string
				f9elem = *f9iter
				f9 = append(f9, &f9elem)
			}
			cr.Status.AtProvider.ImportInfo = f9
		}
		if elem.Name != nil {
			cr.Status.AtProvider.Name = elem.Name
		}
		if elem.ProtocolType != nil {
			cr.Spec.ForProvider.ProtocolType = elem.ProtocolType
		}
		if elem.RouteSelectionExpression != nil {
			cr.Spec.ForProvider.RouteSelectionExpression = elem.RouteSelectionExpression
		}
		if elem.Tags != nil {
			f13 := map[string]*string{}
			for f13key, f13valiter := range elem.Tags {
				var f13val string
				f13val = *f13valiter
				f13[f13key] = &f13val
			}
			cr.Spec.ForProvider.Tags = f13
		}
		if elem.Version != nil {
			cr.Spec.ForProvider.Version = elem.Version
		}
		if elem.Warnings != nil {
			f15 := []*string{}
			for _, f15iter := range elem.Warnings {
				var f15elem string
				f15elem = *f15iter
				f15 = append(f15, &f15elem)
			}
			cr.Status.AtProvider.Warnings = f15
		}
		found = true
		break
	}
	if !found {
		return cr
	}

	return cr
}

// GenerateCreateApiInput returns a create input.
func GenerateCreateApiInput(cr *svcapitypes.API) *svcsdk.CreateApiInput {
	res := preGenerateCreateApiInput(cr, &svcsdk.CreateApiInput{})

	if cr.Spec.ForProvider.APIKeySelectionExpression != nil {
		res.SetApiKeySelectionExpression(*cr.Spec.ForProvider.APIKeySelectionExpression)
	}
	if cr.Spec.ForProvider.CorsConfiguration != nil {
		f1 := &svcsdk.Cors{}
		if cr.Spec.ForProvider.CorsConfiguration.AllowCredentials != nil {
			f1.SetAllowCredentials(*cr.Spec.ForProvider.CorsConfiguration.AllowCredentials)
		}
		if cr.Spec.ForProvider.CorsConfiguration.AllowHeaders != nil {
			f1f1 := []*string{}
			for _, f1f1iter := range cr.Spec.ForProvider.CorsConfiguration.AllowHeaders {
				var f1f1elem string
				f1f1elem = *f1f1iter
				f1f1 = append(f1f1, &f1f1elem)
			}
			f1.SetAllowHeaders(f1f1)
		}
		if cr.Spec.ForProvider.CorsConfiguration.AllowMethods != nil {
			f1f2 := []*string{}
			for _, f1f2iter := range cr.Spec.ForProvider.CorsConfiguration.AllowMethods {
				var f1f2elem string
				f1f2elem = *f1f2iter
				f1f2 = append(f1f2, &f1f2elem)
			}
			f1.SetAllowMethods(f1f2)
		}
		if cr.Spec.ForProvider.CorsConfiguration.AllowOrigins != nil {
			f1f3 := []*string{}
			for _, f1f3iter := range cr.Spec.ForProvider.CorsConfiguration.AllowOrigins {
				var f1f3elem string
				f1f3elem = *f1f3iter
				f1f3 = append(f1f3, &f1f3elem)
			}
			f1.SetAllowOrigins(f1f3)
		}
		if cr.Spec.ForProvider.CorsConfiguration.ExposeHeaders != nil {
			f1f4 := []*string{}
			for _, f1f4iter := range cr.Spec.ForProvider.CorsConfiguration.ExposeHeaders {
				var f1f4elem string
				f1f4elem = *f1f4iter
				f1f4 = append(f1f4, &f1f4elem)
			}
			f1.SetExposeHeaders(f1f4)
		}
		if cr.Spec.ForProvider.CorsConfiguration.MaxAge != nil {
			f1.SetMaxAge(*cr.Spec.ForProvider.CorsConfiguration.MaxAge)
		}
		res.SetCorsConfiguration(f1)
	}
	if cr.Spec.ForProvider.CredentialsARN != nil {
		res.SetCredentialsArn(*cr.Spec.ForProvider.CredentialsARN)
	}
	if cr.Spec.ForProvider.Description != nil {
		res.SetDescription(*cr.Spec.ForProvider.Description)
	}
	if cr.Spec.ForProvider.DisableExecuteAPIEndpoint != nil {
		res.SetDisableExecuteApiEndpoint(*cr.Spec.ForProvider.DisableExecuteAPIEndpoint)
	}
	if cr.Spec.ForProvider.DisableSchemaValidation != nil {
		res.SetDisableSchemaValidation(*cr.Spec.ForProvider.DisableSchemaValidation)
	}
	if cr.Spec.ForProvider.ProtocolType != nil {
		res.SetProtocolType(*cr.Spec.ForProvider.ProtocolType)
	}
	if cr.Spec.ForProvider.RouteKey != nil {
		res.SetRouteKey(*cr.Spec.ForProvider.RouteKey)
	}
	if cr.Spec.ForProvider.RouteSelectionExpression != nil {
		res.SetRouteSelectionExpression(*cr.Spec.ForProvider.RouteSelectionExpression)
	}
	if cr.Spec.ForProvider.Tags != nil {
		f9 := map[string]*string{}
		for f9key, f9valiter := range cr.Spec.ForProvider.Tags {
			var f9val string
			f9val = *f9valiter
			f9[f9key] = &f9val
		}
		res.SetTags(f9)
	}
	if cr.Spec.ForProvider.Target != nil {
		res.SetTarget(*cr.Spec.ForProvider.Target)
	}
	if cr.Spec.ForProvider.Version != nil {
		res.SetVersion(*cr.Spec.ForProvider.Version)
	}

	return postGenerateCreateApiInput(cr, res)
}

// GenerateDeleteApiInput returns a deletion input.
func GenerateDeleteApiInput(cr *svcapitypes.API) *svcsdk.DeleteApiInput {
	res := preGenerateDeleteApiInput(cr, &svcsdk.DeleteApiInput{})

	if cr.Status.AtProvider.APIID != nil {
		res.SetApiId(*cr.Status.AtProvider.APIID)
	}

	return postGenerateDeleteApiInput(cr, res)
}

// IsNotFound returns whether the given error is of type NotFound or not.
func IsNotFound(err error) bool {
	awsErr, ok := err.(awserr.Error)
	return ok && awsErr.Code() == "NotFoundException"
}
