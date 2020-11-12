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

package deployment

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
// GenerateGetDeploymentsInput returns input for read
// operation.
func GenerateGetDeploymentsInput(cr *svcapitypes.Deployment) *svcsdk.GetDeploymentsInput {
	res := preGenerateGetDeploymentsInput(cr, &svcsdk.GetDeploymentsInput{})

	return postGenerateGetDeploymentsInput(cr, res)
}

// GenerateDeployment returns the current state in the form of *svcapitypes.Deployment.
func GenerateDeployment(resp *svcsdk.GetDeploymentsOutput) *svcapitypes.Deployment {
	cr := &svcapitypes.Deployment{}

	found := false
	for _, elem := range resp.Items {
		if elem.AutoDeployed != nil {
			cr.Status.AtProvider.AutoDeployed = elem.AutoDeployed
		}
		if elem.CreatedDate != nil {
			cr.Status.AtProvider.CreatedDate = &metav1.Time{*elem.CreatedDate}
		}
		if elem.DeploymentId != nil {
			cr.Status.AtProvider.DeploymentID = elem.DeploymentId
		}
		if elem.DeploymentStatus != nil {
			cr.Status.AtProvider.DeploymentStatus = elem.DeploymentStatus
		}
		if elem.DeploymentStatusMessage != nil {
			cr.Status.AtProvider.DeploymentStatusMessage = elem.DeploymentStatusMessage
		}
		if elem.Description != nil {
			cr.Spec.ForProvider.Description = elem.Description
		}
		found = true
		break
	}
	if !found {
		return cr
	}

	return cr
}

// GenerateCreateDeploymentInput returns a create input.
func GenerateCreateDeploymentInput(cr *svcapitypes.Deployment) *svcsdk.CreateDeploymentInput {
	res := preGenerateCreateDeploymentInput(cr, &svcsdk.CreateDeploymentInput{})

	if cr.Spec.ForProvider.Description != nil {
		res.SetDescription(*cr.Spec.ForProvider.Description)
	}
	if cr.Spec.ForProvider.StageName != nil {
		res.SetStageName(*cr.Spec.ForProvider.StageName)
	}

	return postGenerateCreateDeploymentInput(cr, res)
}

// GenerateDeleteDeploymentInput returns a deletion input.
func GenerateDeleteDeploymentInput(cr *svcapitypes.Deployment) *svcsdk.DeleteDeploymentInput {
	res := preGenerateDeleteDeploymentInput(cr, &svcsdk.DeleteDeploymentInput{})

	if cr.Status.AtProvider.DeploymentID != nil {
		res.SetDeploymentId(*cr.Status.AtProvider.DeploymentID)
	}

	return postGenerateDeleteDeploymentInput(cr, res)
}

// IsNotFound returns whether the given error is of type NotFound or not.
func IsNotFound(err error) bool {
	awsErr, ok := err.(awserr.Error)
	return ok && awsErr.Code() == "NotFoundException"
}
