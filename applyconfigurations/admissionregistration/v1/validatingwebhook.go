/*
Copyright The Kubernetes Authors.

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

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

import (
	admissionregistrationv1 "k8s.io/api/admissionregistration/v1"
	metav1 "github.com/leigme/client-go-trace/applyconfigurations/meta/v1"
)

// ValidatingWebhookApplyConfiguration represents an declarative configuration of the ValidatingWebhook type for use
// with apply.
type ValidatingWebhookApplyConfiguration struct {
	Name                    *string                                    `json:"name,omitempty"`
	ClientConfig            *WebhookClientConfigApplyConfiguration     `json:"clientConfig,omitempty"`
	Rules                   []RuleWithOperationsApplyConfiguration     `json:"rules,omitempty"`
	FailurePolicy           *admissionregistrationv1.FailurePolicyType `json:"failurePolicy,omitempty"`
	MatchPolicy             *admissionregistrationv1.MatchPolicyType   `json:"matchPolicy,omitempty"`
	NamespaceSelector       *metav1.LabelSelectorApplyConfiguration    `json:"namespaceSelector,omitempty"`
	ObjectSelector          *metav1.LabelSelectorApplyConfiguration    `json:"objectSelector,omitempty"`
	SideEffects             *admissionregistrationv1.SideEffectClass   `json:"sideEffects,omitempty"`
	TimeoutSeconds          *int32                                     `json:"timeoutSeconds,omitempty"`
	AdmissionReviewVersions []string                                   `json:"admissionReviewVersions,omitempty"`
}

// ValidatingWebhookApplyConfiguration constructs an declarative configuration of the ValidatingWebhook type for use with
// apply.
func ValidatingWebhook() *ValidatingWebhookApplyConfiguration {
	return &ValidatingWebhookApplyConfiguration{}
}

// WithName sets the Name field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Name field is set to the value of the last call.
func (b *ValidatingWebhookApplyConfiguration) WithName(value string) *ValidatingWebhookApplyConfiguration {
	b.Name = &value
	return b
}

// WithClientConfig sets the ClientConfig field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ClientConfig field is set to the value of the last call.
func (b *ValidatingWebhookApplyConfiguration) WithClientConfig(value *WebhookClientConfigApplyConfiguration) *ValidatingWebhookApplyConfiguration {
	b.ClientConfig = value
	return b
}

// WithRules adds the given value to the Rules field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Rules field.
func (b *ValidatingWebhookApplyConfiguration) WithRules(values ...*RuleWithOperationsApplyConfiguration) *ValidatingWebhookApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithRules")
		}
		b.Rules = append(b.Rules, *values[i])
	}
	return b
}

// WithFailurePolicy sets the FailurePolicy field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the FailurePolicy field is set to the value of the last call.
func (b *ValidatingWebhookApplyConfiguration) WithFailurePolicy(value admissionregistrationv1.FailurePolicyType) *ValidatingWebhookApplyConfiguration {
	b.FailurePolicy = &value
	return b
}

// WithMatchPolicy sets the MatchPolicy field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the MatchPolicy field is set to the value of the last call.
func (b *ValidatingWebhookApplyConfiguration) WithMatchPolicy(value admissionregistrationv1.MatchPolicyType) *ValidatingWebhookApplyConfiguration {
	b.MatchPolicy = &value
	return b
}

// WithNamespaceSelector sets the NamespaceSelector field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the NamespaceSelector field is set to the value of the last call.
func (b *ValidatingWebhookApplyConfiguration) WithNamespaceSelector(value *metav1.LabelSelectorApplyConfiguration) *ValidatingWebhookApplyConfiguration {
	b.NamespaceSelector = value
	return b
}

// WithObjectSelector sets the ObjectSelector field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ObjectSelector field is set to the value of the last call.
func (b *ValidatingWebhookApplyConfiguration) WithObjectSelector(value *metav1.LabelSelectorApplyConfiguration) *ValidatingWebhookApplyConfiguration {
	b.ObjectSelector = value
	return b
}

// WithSideEffects sets the SideEffects field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the SideEffects field is set to the value of the last call.
func (b *ValidatingWebhookApplyConfiguration) WithSideEffects(value admissionregistrationv1.SideEffectClass) *ValidatingWebhookApplyConfiguration {
	b.SideEffects = &value
	return b
}

// WithTimeoutSeconds sets the TimeoutSeconds field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the TimeoutSeconds field is set to the value of the last call.
func (b *ValidatingWebhookApplyConfiguration) WithTimeoutSeconds(value int32) *ValidatingWebhookApplyConfiguration {
	b.TimeoutSeconds = &value
	return b
}

// WithAdmissionReviewVersions adds the given value to the AdmissionReviewVersions field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the AdmissionReviewVersions field.
func (b *ValidatingWebhookApplyConfiguration) WithAdmissionReviewVersions(values ...string) *ValidatingWebhookApplyConfiguration {
	for i := range values {
		b.AdmissionReviewVersions = append(b.AdmissionReviewVersions, values[i])
	}
	return b
}
