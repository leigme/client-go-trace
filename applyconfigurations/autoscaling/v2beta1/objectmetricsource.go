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

package v2beta1

import (
	resource "k8s.io/apimachinery/pkg/api/resource"
	v1 "github.com/leigme/client-go-trace/applyconfigurations/meta/v1"
)

// ObjectMetricSourceApplyConfiguration represents an declarative configuration of the ObjectMetricSource type for use
// with apply.
type ObjectMetricSourceApplyConfiguration struct {
	Target       *CrossVersionObjectReferenceApplyConfiguration `json:"target,omitempty"`
	MetricName   *string                                        `json:"metricName,omitempty"`
	TargetValue  *resource.Quantity                             `json:"targetValue,omitempty"`
	Selector     *v1.LabelSelectorApplyConfiguration            `json:"selector,omitempty"`
	AverageValue *resource.Quantity                             `json:"averageValue,omitempty"`
}

// ObjectMetricSourceApplyConfiguration constructs an declarative configuration of the ObjectMetricSource type for use with
// apply.
func ObjectMetricSource() *ObjectMetricSourceApplyConfiguration {
	return &ObjectMetricSourceApplyConfiguration{}
}

// WithTarget sets the Target field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Target field is set to the value of the last call.
func (b *ObjectMetricSourceApplyConfiguration) WithTarget(value *CrossVersionObjectReferenceApplyConfiguration) *ObjectMetricSourceApplyConfiguration {
	b.Target = value
	return b
}

// WithMetricName sets the MetricName field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the MetricName field is set to the value of the last call.
func (b *ObjectMetricSourceApplyConfiguration) WithMetricName(value string) *ObjectMetricSourceApplyConfiguration {
	b.MetricName = &value
	return b
}

// WithTargetValue sets the TargetValue field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the TargetValue field is set to the value of the last call.
func (b *ObjectMetricSourceApplyConfiguration) WithTargetValue(value resource.Quantity) *ObjectMetricSourceApplyConfiguration {
	b.TargetValue = &value
	return b
}

// WithSelector sets the Selector field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Selector field is set to the value of the last call.
func (b *ObjectMetricSourceApplyConfiguration) WithSelector(value *v1.LabelSelectorApplyConfiguration) *ObjectMetricSourceApplyConfiguration {
	b.Selector = value
	return b
}

// WithAverageValue sets the AverageValue field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the AverageValue field is set to the value of the last call.
func (b *ObjectMetricSourceApplyConfiguration) WithAverageValue(value resource.Quantity) *ObjectMetricSourceApplyConfiguration {
	b.AverageValue = &value
	return b
}
