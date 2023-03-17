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
	apicertificatesv1 "k8s.io/api/certificates/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	managedfields "k8s.io/apimachinery/pkg/util/managedfields"
	internal "github.com/leigme/client-go-trace/applyconfigurations/internal"
	v1 "github.com/leigme/client-go-trace/applyconfigurations/meta/v1"
)

// CertificateSigningRequestApplyConfiguration represents an declarative configuration of the CertificateSigningRequest type for use
// with apply.
type CertificateSigningRequestApplyConfiguration struct {
	v1.TypeMetaApplyConfiguration    `json:",inline"`
	*v1.ObjectMetaApplyConfiguration `json:"metadata,omitempty"`
	Spec                             *CertificateSigningRequestSpecApplyConfiguration   `json:"spec,omitempty"`
	Status                           *CertificateSigningRequestStatusApplyConfiguration `json:"status,omitempty"`
}

// CertificateSigningRequest constructs an declarative configuration of the CertificateSigningRequest type for use with
// apply.
func CertificateSigningRequest(name string) *CertificateSigningRequestApplyConfiguration {
	b := &CertificateSigningRequestApplyConfiguration{}
	b.WithName(name)
	b.WithKind("CertificateSigningRequest")
	b.WithAPIVersion("certificates.k8s.io/v1")
	return b
}

// ExtractCertificateSigningRequest extracts the applied configuration owned by fieldManager from
// certificateSigningRequest. If no managedFields are found in certificateSigningRequest for fieldManager, a
// CertificateSigningRequestApplyConfiguration is returned with only the Name, Namespace (if applicable),
// APIVersion and Kind populated. It is possible that no managed fields were found for because other
// field managers have taken ownership of all the fields previously owned by fieldManager, or because
// the fieldManager never owned fields any fields.
// certificateSigningRequest must be a unmodified CertificateSigningRequest API object that was retrieved from the Kubernetes API.
// ExtractCertificateSigningRequest provides a way to perform a extract/modify-in-place/apply workflow.
// Note that an extracted apply configuration will contain fewer fields than what the fieldManager previously
// applied if another fieldManager has updated or force applied any of the previously applied fields.
// Experimental!
func ExtractCertificateSigningRequest(certificateSigningRequest *apicertificatesv1.CertificateSigningRequest, fieldManager string) (*CertificateSigningRequestApplyConfiguration, error) {
	return extractCertificateSigningRequest(certificateSigningRequest, fieldManager, "")
}

// ExtractCertificateSigningRequestStatus is the same as ExtractCertificateSigningRequest except
// that it extracts the status subresource applied configuration.
// Experimental!
func ExtractCertificateSigningRequestStatus(certificateSigningRequest *apicertificatesv1.CertificateSigningRequest, fieldManager string) (*CertificateSigningRequestApplyConfiguration, error) {
	return extractCertificateSigningRequest(certificateSigningRequest, fieldManager, "status")
}

func extractCertificateSigningRequest(certificateSigningRequest *apicertificatesv1.CertificateSigningRequest, fieldManager string, subresource string) (*CertificateSigningRequestApplyConfiguration, error) {
	b := &CertificateSigningRequestApplyConfiguration{}
	err := managedfields.ExtractInto(certificateSigningRequest, internal.Parser().Type("io.k8s.api.certificates.v1.CertificateSigningRequest"), fieldManager, b, subresource)
	if err != nil {
		return nil, err
	}
	b.WithName(certificateSigningRequest.Name)

	b.WithKind("CertificateSigningRequest")
	b.WithAPIVersion("certificates.k8s.io/v1")
	return b, nil
}

// WithKind sets the Kind field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Kind field is set to the value of the last call.
func (b *CertificateSigningRequestApplyConfiguration) WithKind(value string) *CertificateSigningRequestApplyConfiguration {
	b.Kind = &value
	return b
}

// WithAPIVersion sets the APIVersion field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the APIVersion field is set to the value of the last call.
func (b *CertificateSigningRequestApplyConfiguration) WithAPIVersion(value string) *CertificateSigningRequestApplyConfiguration {
	b.APIVersion = &value
	return b
}

// WithName sets the Name field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Name field is set to the value of the last call.
func (b *CertificateSigningRequestApplyConfiguration) WithName(value string) *CertificateSigningRequestApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.Name = &value
	return b
}

// WithGenerateName sets the GenerateName field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the GenerateName field is set to the value of the last call.
func (b *CertificateSigningRequestApplyConfiguration) WithGenerateName(value string) *CertificateSigningRequestApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.GenerateName = &value
	return b
}

// WithNamespace sets the Namespace field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Namespace field is set to the value of the last call.
func (b *CertificateSigningRequestApplyConfiguration) WithNamespace(value string) *CertificateSigningRequestApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.Namespace = &value
	return b
}

// WithSelfLink sets the SelfLink field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the SelfLink field is set to the value of the last call.
func (b *CertificateSigningRequestApplyConfiguration) WithSelfLink(value string) *CertificateSigningRequestApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.SelfLink = &value
	return b
}

// WithUID sets the UID field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the UID field is set to the value of the last call.
func (b *CertificateSigningRequestApplyConfiguration) WithUID(value types.UID) *CertificateSigningRequestApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.UID = &value
	return b
}

// WithResourceVersion sets the ResourceVersion field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ResourceVersion field is set to the value of the last call.
func (b *CertificateSigningRequestApplyConfiguration) WithResourceVersion(value string) *CertificateSigningRequestApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.ResourceVersion = &value
	return b
}

// WithGeneration sets the Generation field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Generation field is set to the value of the last call.
func (b *CertificateSigningRequestApplyConfiguration) WithGeneration(value int64) *CertificateSigningRequestApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.Generation = &value
	return b
}

// WithCreationTimestamp sets the CreationTimestamp field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the CreationTimestamp field is set to the value of the last call.
func (b *CertificateSigningRequestApplyConfiguration) WithCreationTimestamp(value metav1.Time) *CertificateSigningRequestApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.CreationTimestamp = &value
	return b
}

// WithDeletionTimestamp sets the DeletionTimestamp field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the DeletionTimestamp field is set to the value of the last call.
func (b *CertificateSigningRequestApplyConfiguration) WithDeletionTimestamp(value metav1.Time) *CertificateSigningRequestApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.DeletionTimestamp = &value
	return b
}

// WithDeletionGracePeriodSeconds sets the DeletionGracePeriodSeconds field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the DeletionGracePeriodSeconds field is set to the value of the last call.
func (b *CertificateSigningRequestApplyConfiguration) WithDeletionGracePeriodSeconds(value int64) *CertificateSigningRequestApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.DeletionGracePeriodSeconds = &value
	return b
}

// WithLabels puts the entries into the Labels field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, the entries provided by each call will be put on the Labels field,
// overwriting an existing map entries in Labels field with the same key.
func (b *CertificateSigningRequestApplyConfiguration) WithLabels(entries map[string]string) *CertificateSigningRequestApplyConfiguration {
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
func (b *CertificateSigningRequestApplyConfiguration) WithAnnotations(entries map[string]string) *CertificateSigningRequestApplyConfiguration {
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
func (b *CertificateSigningRequestApplyConfiguration) WithOwnerReferences(values ...*v1.OwnerReferenceApplyConfiguration) *CertificateSigningRequestApplyConfiguration {
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
func (b *CertificateSigningRequestApplyConfiguration) WithFinalizers(values ...string) *CertificateSigningRequestApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	for i := range values {
		b.Finalizers = append(b.Finalizers, values[i])
	}
	return b
}

// WithClusterName sets the ClusterName field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ClusterName field is set to the value of the last call.
func (b *CertificateSigningRequestApplyConfiguration) WithClusterName(value string) *CertificateSigningRequestApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.ClusterName = &value
	return b
}

func (b *CertificateSigningRequestApplyConfiguration) ensureObjectMetaApplyConfigurationExists() {
	if b.ObjectMetaApplyConfiguration == nil {
		b.ObjectMetaApplyConfiguration = &v1.ObjectMetaApplyConfiguration{}
	}
}

// WithSpec sets the Spec field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Spec field is set to the value of the last call.
func (b *CertificateSigningRequestApplyConfiguration) WithSpec(value *CertificateSigningRequestSpecApplyConfiguration) *CertificateSigningRequestApplyConfiguration {
	b.Spec = value
	return b
}

// WithStatus sets the Status field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Status field is set to the value of the last call.
func (b *CertificateSigningRequestApplyConfiguration) WithStatus(value *CertificateSigningRequestStatusApplyConfiguration) *CertificateSigningRequestApplyConfiguration {
	b.Status = value
	return b
}
