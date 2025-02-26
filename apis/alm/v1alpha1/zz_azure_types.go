// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type AzureInitParameters struct {

	// (String) Unique key of the Azure Devops instance setting
	// Unique key of the Azure Devops instance setting
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	// (String, Sensitive) Azure Devops personal access token
	// Azure Devops personal access token
	PersonalAccessTokenSecretRef v1.SecretKeySelector `json:"personalAccessTokenSecretRef" tf:"-"`

	// (String) Azure API URL
	// Azure API URL
	URL *string `json:"url,omitempty" tf:"url,omitempty"`
}

type AzureObservation struct {

	// (String) The ID of this resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (String) Unique key of the Azure Devops instance setting
	// Unique key of the Azure Devops instance setting
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	// (String) Azure API URL
	// Azure API URL
	URL *string `json:"url,omitempty" tf:"url,omitempty"`
}

type AzureParameters struct {

	// (String) Unique key of the Azure Devops instance setting
	// Unique key of the Azure Devops instance setting
	// +kubebuilder:validation:Optional
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	// (String, Sensitive) Azure Devops personal access token
	// Azure Devops personal access token
	// +kubebuilder:validation:Optional
	PersonalAccessTokenSecretRef v1.SecretKeySelector `json:"personalAccessTokenSecretRef" tf:"-"`

	// (String) Azure API URL
	// Azure API URL
	// +kubebuilder:validation:Optional
	URL *string `json:"url,omitempty" tf:"url,omitempty"`
}

// AzureSpec defines the desired state of Azure
type AzureSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AzureParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider AzureInitParameters `json:"initProvider,omitempty"`
}

// AzureStatus defines the observed state of Azure.
type AzureStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AzureObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Azure is the Schema for the Azures API. Provides a Sonarqube Azure Devops Alm/Devops Platform Integration resource. This can be used to create and manage a Alm/Devops Platform Integration for Azure Devops.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,sonarqube}
type Azure struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.key) || (has(self.initProvider) && has(self.initProvider.key))",message="spec.forProvider.key is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.personalAccessTokenSecretRef)",message="spec.forProvider.personalAccessTokenSecretRef is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.url) || (has(self.initProvider) && has(self.initProvider.url))",message="spec.forProvider.url is a required parameter"
	Spec   AzureSpec   `json:"spec"`
	Status AzureStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AzureList contains a list of Azures
type AzureList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Azure `json:"items"`
}

// Repository type metadata.
var (
	Azure_Kind             = "Azure"
	Azure_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Azure_Kind}.String()
	Azure_KindAPIVersion   = Azure_Kind + "." + CRDGroupVersion.String()
	Azure_GroupVersionKind = CRDGroupVersion.WithKind(Azure_Kind)
)

func init() {
	SchemeBuilder.Register(&Azure{}, &AzureList{})
}
