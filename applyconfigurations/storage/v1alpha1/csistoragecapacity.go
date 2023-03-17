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

package v1alpha1

import (
	v1alpha1 "k8s.io/api/storage/v1alpha1"
	resource "k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	managedfields "k8s.io/apimachinery/pkg/util/managedfields"
	internal "github.com/leigme/client-go-trace/applyconfigurations/internal"
	v1 "github.com/leigme/client-go-trace/applyconfigurations/meta/v1"
)

// CSIStorageCapacityApplyConfiguration represents an declarative configuration of the CSIStorageCapacity type for use
// with apply.
type CSIStorageCapacityApplyConfiguration struct {
	v1.TypeMetaApplyConfiguration    `json:",inline"`
	*v1.ObjectMetaApplyConfiguration `json:"metadata,omitempty"`
	NodeTopology                     *v1.LabelSelectorApplyConfiguration `json:"nodeTopology,omitempty"`
	StorageClassName                 *string                             `json:"storageClassName,omitempty"`
	Capacity                         *resource.Quantity                  `json:"capacity,omitempty"`
	MaximumVolumeSize                *resource.Quantity                  `json:"maximumVolumeSize,omitempty"`
}

// CSIStorageCapacity constructs an declarative configuration of the CSIStorageCapacity type for use with
// apply.
func CSIStorageCapacity(name, namespace string) *CSIStorageCapacityApplyConfiguration {
	b := &CSIStorageCapacityApplyConfiguration{}
	b.WithName(name)
	b.WithNamespace(namespace)
	b.WithKind("CSIStorageCapacity")
	b.WithAPIVersion("storage.k8s.io/v1alpha1")
	return b
}

// ExtractCSIStorageCapacity extracts the applied configuration owned by fieldManager from
// cSIStorageCapacity. If no managedFields are found in cSIStorageCapacity for fieldManager, a
// CSIStorageCapacityApplyConfiguration is returned with only the Name, Namespace (if applicable),
// APIVersion and Kind populated. It is possible that no managed fields were found for because other
// field managers have taken ownership of all the fields previously owned by fieldManager, or because
// the fieldManager never owned fields any fields.
// cSIStorageCapacity must be a unmodified CSIStorageCapacity API object that was retrieved from the Kubernetes API.
// ExtractCSIStorageCapacity provides a way to perform a extract/modify-in-place/apply workflow.
// Note that an extracted apply configuration will contain fewer fields than what the fieldManager previously
// applied if another fieldManager has updated or force applied any of the previously applied fields.
// Experimental!
func ExtractCSIStorageCapacity(cSIStorageCapacity *v1alpha1.CSIStorageCapacity, fieldManager string) (*CSIStorageCapacityApplyConfiguration, error) {
	return extractCSIStorageCapacity(cSIStorageCapacity, fieldManager, "")
}

// ExtractCSIStorageCapacityStatus is the same as ExtractCSIStorageCapacity except
// that it extracts the status subresource applied configuration.
// Experimental!
func ExtractCSIStorageCapacityStatus(cSIStorageCapacity *v1alpha1.CSIStorageCapacity, fieldManager string) (*CSIStorageCapacityApplyConfiguration, error) {
	return extractCSIStorageCapacity(cSIStorageCapacity, fieldManager, "status")
}

func extractCSIStorageCapacity(cSIStorageCapacity *v1alpha1.CSIStorageCapacity, fieldManager string, subresource string) (*CSIStorageCapacityApplyConfiguration, error) {
	b := &CSIStorageCapacityApplyConfiguration{}
	err := managedfields.ExtractInto(cSIStorageCapacity, internal.Parser().Type("io.k8s.api.storage.v1alpha1.CSIStorageCapacity"), fieldManager, b, subresource)
	if err != nil {
		return nil, err
	}
	b.WithName(cSIStorageCapacity.Name)
	b.WithNamespace(cSIStorageCapacity.Namespace)

	b.WithKind("CSIStorageCapacity")
	b.WithAPIVersion("storage.k8s.io/v1alpha1")
	return b, nil
}

// WithKind sets the Kind field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Kind field is set to the value of the last call.
func (b *CSIStorageCapacityApplyConfiguration) WithKind(value string) *CSIStorageCapacityApplyConfiguration {
	b.Kind = &value
	return b
}

// WithAPIVersion sets the APIVersion field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the APIVersion field is set to the value of the last call.
func (b *CSIStorageCapacityApplyConfiguration) WithAPIVersion(value string) *CSIStorageCapacityApplyConfiguration {
	b.APIVersion = &value
	return b
}

// WithName sets the Name field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Name field is set to the value of the last call.
func (b *CSIStorageCapacityApplyConfiguration) WithName(value string) *CSIStorageCapacityApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.Name = &value
	return b
}

// WithGenerateName sets the GenerateName field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the GenerateName field is set to the value of the last call.
func (b *CSIStorageCapacityApplyConfiguration) WithGenerateName(value string) *CSIStorageCapacityApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.GenerateName = &value
	return b
}

// WithNamespace sets the Namespace field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Namespace field is set to the value of the last call.
func (b *CSIStorageCapacityApplyConfiguration) WithNamespace(value string) *CSIStorageCapacityApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.Namespace = &value
	return b
}

// WithSelfLink sets the SelfLink field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the SelfLink field is set to the value of the last call.
func (b *CSIStorageCapacityApplyConfiguration) WithSelfLink(value string) *CSIStorageCapacityApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.SelfLink = &value
	return b
}

// WithUID sets the UID field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the UID field is set to the value of the last call.
func (b *CSIStorageCapacityApplyConfiguration) WithUID(value types.UID) *CSIStorageCapacityApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.UID = &value
	return b
}

// WithResourceVersion sets the ResourceVersion field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ResourceVersion field is set to the value of the last call.
func (b *CSIStorageCapacityApplyConfiguration) WithResourceVersion(value string) *CSIStorageCapacityApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.ResourceVersion = &value
	return b
}

// WithGeneration sets the Generation field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Generation field is set to the value of the last call.
func (b *CSIStorageCapacityApplyConfiguration) WithGeneration(value int64) *CSIStorageCapacityApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.Generation = &value
	return b
}

// WithCreationTimestamp sets the CreationTimestamp field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the CreationTimestamp field is set to the value of the last call.
func (b *CSIStorageCapacityApplyConfiguration) WithCreationTimestamp(value metav1.Time) *CSIStorageCapacityApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.CreationTimestamp = &value
	return b
}

// WithDeletionTimestamp sets the DeletionTimestamp field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the DeletionTimestamp field is set to the value of the last call.
func (b *CSIStorageCapacityApplyConfiguration) WithDeletionTimestamp(value metav1.Time) *CSIStorageCapacityApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.DeletionTimestamp = &value
	return b
}

// WithDeletionGracePeriodSeconds sets the DeletionGracePeriodSeconds field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the DeletionGracePeriodSeconds field is set to the value of the last call.
func (b *CSIStorageCapacityApplyConfiguration) WithDeletionGracePeriodSeconds(value int64) *CSIStorageCapacityApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.DeletionGracePeriodSeconds = &value
	return b
}

// WithLabels puts the entries into the Labels field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, the entries provided by each call will be put on the Labels field,
// overwriting an existing map entries in Labels field with the same key.
func (b *CSIStorageCapacityApplyConfiguration) WithLabels(entries map[string]string) *CSIStorageCapacityApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	if b.Labels == nil && len(entries) > 0 {
		b.Labels = make(map[string]string, len(entries))
	}
	for k, v := range entries {
		b.Labels[k] = v
	}
	return b
}

// WithAnnotations puts the entries into the Annotations field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, the entries provided by each call will be put on the Annotations field,
// overwriting an existing map entries in Annotations field with the same key.
func (b *CSIStorageCapacityApplyConfiguration) WithAnnotations(entries map[string]string) *CSIStorageCapacityApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	if b.Annotations == nil && len(entries) > 0 {
		b.Annotations = make(map[string]string, len(entries))
	}
	for k, v := range entries {
		b.Annotations[k] = v
	}
	return b
}

// WithOwnerReferences adds the given value to the OwnerReferences field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the OwnerReferences field.
func (b *CSIStorageCapacityApplyConfiguration) WithOwnerReferences(values ...*v1.OwnerReferenceApplyConfiguration) *CSIStorageCapacityApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithOwnerReferences")
		}
		b.OwnerReferences = append(b.OwnerReferences, *values[i])
	}
	return b
}

// WithFinalizers adds the given value to the Finalizers field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Finalizers field.
func (b *CSIStorageCapacityApplyConfiguration) WithFinalizers(values ...string) *CSIStorageCapacityApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	for i := range values {
		b.Finalizers = append(b.Finalizers, values[i])
	}
	return b
}

// WithClusterName sets the ClusterName field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ClusterName field is set to the value of the last call.
func (b *CSIStorageCapacityApplyConfiguration) WithClusterName(value string) *CSIStorageCapacityApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.ClusterName = &value
	return b
}

func (b *CSIStorageCapacityApplyConfiguration) ensureObjectMetaApplyConfigurationExists() {
	if b.ObjectMetaApplyConfiguration == nil {
		b.ObjectMetaApplyConfiguration = &v1.ObjectMetaApplyConfiguration{}
	}
}

// WithNodeTopology sets the NodeTopology field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the NodeTopology field is set to the value of the last call.
func (b *CSIStorageCapacityApplyConfiguration) WithNodeTopology(value *v1.LabelSelectorApplyConfiguration) *CSIStorageCapacityApplyConfiguration {
	b.NodeTopology = value
	return b
}

// WithStorageClassName sets the StorageClassName field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the StorageClassName field is set to the value of the last call.
func (b *CSIStorageCapacityApplyConfiguration) WithStorageClassName(value string) *CSIStorageCapacityApplyConfiguration {
	b.StorageClassName = &value
	return b
}

// WithCapacity sets the Capacity field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Capacity field is set to the value of the last call.
func (b *CSIStorageCapacityApplyConfiguration) WithCapacity(value resource.Quantity) *CSIStorageCapacityApplyConfiguration {
	b.Capacity = &value
	return b
}

// WithMaximumVolumeSize sets the MaximumVolumeSize field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the MaximumVolumeSize field is set to the value of the last call.
func (b *CSIStorageCapacityApplyConfiguration) WithMaximumVolumeSize(value resource.Quantity) *CSIStorageCapacityApplyConfiguration {
	b.MaximumVolumeSize = &value
	return b
}
