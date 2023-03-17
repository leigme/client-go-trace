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

// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	"context"
	json "encoding/json"
	"fmt"
	"time"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	corev1 "github.com/leigme/client-go-trace/applyconfigurations/core/v1"
	scheme "github.com/leigme/client-go-trace/kubernetes/scheme"
	rest "github.com/leigme/client-go-trace/rest"
)

// PersistentVolumeClaimsGetter has a method to return a PersistentVolumeClaimInterface.
// A group's client should implement this interface.
type PersistentVolumeClaimsGetter interface {
	PersistentVolumeClaims(namespace string) PersistentVolumeClaimInterface
}

// PersistentVolumeClaimInterface has methods to work with PersistentVolumeClaim resources.
type PersistentVolumeClaimInterface interface {
	Create(ctx context.Context, persistentVolumeClaim *v1.PersistentVolumeClaim, opts metav1.CreateOptions) (*v1.PersistentVolumeClaim, error)
	Update(ctx context.Context, persistentVolumeClaim *v1.PersistentVolumeClaim, opts metav1.UpdateOptions) (*v1.PersistentVolumeClaim, error)
	UpdateStatus(ctx context.Context, persistentVolumeClaim *v1.PersistentVolumeClaim, opts metav1.UpdateOptions) (*v1.PersistentVolumeClaim, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.PersistentVolumeClaim, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.PersistentVolumeClaimList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.PersistentVolumeClaim, err error)
	Apply(ctx context.Context, persistentVolumeClaim *corev1.PersistentVolumeClaimApplyConfiguration, opts metav1.ApplyOptions) (result *v1.PersistentVolumeClaim, err error)
	ApplyStatus(ctx context.Context, persistentVolumeClaim *corev1.PersistentVolumeClaimApplyConfiguration, opts metav1.ApplyOptions) (result *v1.PersistentVolumeClaim, err error)
	PersistentVolumeClaimExpansion
}

// persistentVolumeClaims implements PersistentVolumeClaimInterface
type persistentVolumeClaims struct {
	client rest.Interface
	ns     string
}

// newPersistentVolumeClaims returns a PersistentVolumeClaims
func newPersistentVolumeClaims(c *CoreV1Client, namespace string) *persistentVolumeClaims {
	return &persistentVolumeClaims{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the persistentVolumeClaim, and returns the corresponding persistentVolumeClaim object, and an error if there is any.
func (c *persistentVolumeClaims) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.PersistentVolumeClaim, err error) {
	result = &v1.PersistentVolumeClaim{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("persistentvolumeclaims").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of PersistentVolumeClaims that match those selectors.
func (c *persistentVolumeClaims) List(ctx context.Context, opts metav1.ListOptions) (result *v1.PersistentVolumeClaimList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.PersistentVolumeClaimList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("persistentvolumeclaims").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested persistentVolumeClaims.
func (c *persistentVolumeClaims) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("persistentvolumeclaims").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a persistentVolumeClaim and creates it.  Returns the server's representation of the persistentVolumeClaim, and an error, if there is any.
func (c *persistentVolumeClaims) Create(ctx context.Context, persistentVolumeClaim *v1.PersistentVolumeClaim, opts metav1.CreateOptions) (result *v1.PersistentVolumeClaim, err error) {
	result = &v1.PersistentVolumeClaim{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("persistentvolumeclaims").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(persistentVolumeClaim).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a persistentVolumeClaim and updates it. Returns the server's representation of the persistentVolumeClaim, and an error, if there is any.
func (c *persistentVolumeClaims) Update(ctx context.Context, persistentVolumeClaim *v1.PersistentVolumeClaim, opts metav1.UpdateOptions) (result *v1.PersistentVolumeClaim, err error) {
	result = &v1.PersistentVolumeClaim{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("persistentvolumeclaims").
		Name(persistentVolumeClaim.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(persistentVolumeClaim).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *persistentVolumeClaims) UpdateStatus(ctx context.Context, persistentVolumeClaim *v1.PersistentVolumeClaim, opts metav1.UpdateOptions) (result *v1.PersistentVolumeClaim, err error) {
	result = &v1.PersistentVolumeClaim{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("persistentvolumeclaims").
		Name(persistentVolumeClaim.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(persistentVolumeClaim).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the persistentVolumeClaim and deletes it. Returns an error if one occurs.
func (c *persistentVolumeClaims) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("persistentvolumeclaims").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *persistentVolumeClaims) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("persistentvolumeclaims").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched persistentVolumeClaim.
func (c *persistentVolumeClaims) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.PersistentVolumeClaim, err error) {
	result = &v1.PersistentVolumeClaim{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("persistentvolumeclaims").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}

// Apply takes the given apply declarative configuration, applies it and returns the applied persistentVolumeClaim.
func (c *persistentVolumeClaims) Apply(ctx context.Context, persistentVolumeClaim *corev1.PersistentVolumeClaimApplyConfiguration, opts metav1.ApplyOptions) (result *v1.PersistentVolumeClaim, err error) {
	if persistentVolumeClaim == nil {
		return nil, fmt.Errorf("persistentVolumeClaim provided to Apply must not be nil")
	}
	patchOpts := opts.ToPatchOptions()
	data, err := json.Marshal(persistentVolumeClaim)
	if err != nil {
		return nil, err
	}
	name := persistentVolumeClaim.Name
	if name == nil {
		return nil, fmt.Errorf("persistentVolumeClaim.Name must be provided to Apply")
	}
	result = &v1.PersistentVolumeClaim{}
	err = c.client.Patch(types.ApplyPatchType).
		Namespace(c.ns).
		Resource("persistentvolumeclaims").
		Name(*name).
		VersionedParams(&patchOpts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}

// ApplyStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
func (c *persistentVolumeClaims) ApplyStatus(ctx context.Context, persistentVolumeClaim *corev1.PersistentVolumeClaimApplyConfiguration, opts metav1.ApplyOptions) (result *v1.PersistentVolumeClaim, err error) {
	if persistentVolumeClaim == nil {
		return nil, fmt.Errorf("persistentVolumeClaim provided to Apply must not be nil")
	}
	patchOpts := opts.ToPatchOptions()
	data, err := json.Marshal(persistentVolumeClaim)
	if err != nil {
		return nil, err
	}

	name := persistentVolumeClaim.Name
	if name == nil {
		return nil, fmt.Errorf("persistentVolumeClaim.Name must be provided to Apply")
	}

	result = &v1.PersistentVolumeClaim{}
	err = c.client.Patch(types.ApplyPatchType).
		Namespace(c.ns).
		Resource("persistentvolumeclaims").
		Name(*name).
		SubResource("status").
		VersionedParams(&patchOpts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
