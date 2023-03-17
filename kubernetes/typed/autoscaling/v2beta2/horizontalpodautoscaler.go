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

package v2beta2

import (
	"context"
	json "encoding/json"
	"fmt"
	"time"

	v2beta2 "k8s.io/api/autoscaling/v2beta2"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	autoscalingv2beta2 "github.com/leigme/client-go-trace/applyconfigurations/autoscaling/v2beta2"
	scheme "github.com/leigme/client-go-trace/kubernetes/scheme"
	rest "github.com/leigme/client-go-trace/rest"
)

// HorizontalPodAutoscalersGetter has a method to return a HorizontalPodAutoscalerInterface.
// A group's client should implement this interface.
type HorizontalPodAutoscalersGetter interface {
	HorizontalPodAutoscalers(namespace string) HorizontalPodAutoscalerInterface
}

// HorizontalPodAutoscalerInterface has methods to work with HorizontalPodAutoscaler resources.
type HorizontalPodAutoscalerInterface interface {
	Create(ctx context.Context, horizontalPodAutoscaler *v2beta2.HorizontalPodAutoscaler, opts v1.CreateOptions) (*v2beta2.HorizontalPodAutoscaler, error)
	Update(ctx context.Context, horizontalPodAutoscaler *v2beta2.HorizontalPodAutoscaler, opts v1.UpdateOptions) (*v2beta2.HorizontalPodAutoscaler, error)
	UpdateStatus(ctx context.Context, horizontalPodAutoscaler *v2beta2.HorizontalPodAutoscaler, opts v1.UpdateOptions) (*v2beta2.HorizontalPodAutoscaler, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v2beta2.HorizontalPodAutoscaler, error)
	List(ctx context.Context, opts v1.ListOptions) (*v2beta2.HorizontalPodAutoscalerList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v2beta2.HorizontalPodAutoscaler, err error)
	Apply(ctx context.Context, horizontalPodAutoscaler *autoscalingv2beta2.HorizontalPodAutoscalerApplyConfiguration, opts v1.ApplyOptions) (result *v2beta2.HorizontalPodAutoscaler, err error)
	ApplyStatus(ctx context.Context, horizontalPodAutoscaler *autoscalingv2beta2.HorizontalPodAutoscalerApplyConfiguration, opts v1.ApplyOptions) (result *v2beta2.HorizontalPodAutoscaler, err error)
	HorizontalPodAutoscalerExpansion
}

// horizontalPodAutoscalers implements HorizontalPodAutoscalerInterface
type horizontalPodAutoscalers struct {
	client rest.Interface
	ns     string
}

// newHorizontalPodAutoscalers returns a HorizontalPodAutoscalers
func newHorizontalPodAutoscalers(c *AutoscalingV2beta2Client, namespace string) *horizontalPodAutoscalers {
	return &horizontalPodAutoscalers{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the horizontalPodAutoscaler, and returns the corresponding horizontalPodAutoscaler object, and an error if there is any.
func (c *horizontalPodAutoscalers) Get(ctx context.Context, name string, options v1.GetOptions) (result *v2beta2.HorizontalPodAutoscaler, err error) {
	result = &v2beta2.HorizontalPodAutoscaler{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("horizontalpodautoscalers").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of HorizontalPodAutoscalers that match those selectors.
func (c *horizontalPodAutoscalers) List(ctx context.Context, opts v1.ListOptions) (result *v2beta2.HorizontalPodAutoscalerList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v2beta2.HorizontalPodAutoscalerList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("horizontalpodautoscalers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested horizontalPodAutoscalers.
func (c *horizontalPodAutoscalers) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("horizontalpodautoscalers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a horizontalPodAutoscaler and creates it.  Returns the server's representation of the horizontalPodAutoscaler, and an error, if there is any.
func (c *horizontalPodAutoscalers) Create(ctx context.Context, horizontalPodAutoscaler *v2beta2.HorizontalPodAutoscaler, opts v1.CreateOptions) (result *v2beta2.HorizontalPodAutoscaler, err error) {
	result = &v2beta2.HorizontalPodAutoscaler{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("horizontalpodautoscalers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(horizontalPodAutoscaler).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a horizontalPodAutoscaler and updates it. Returns the server's representation of the horizontalPodAutoscaler, and an error, if there is any.
func (c *horizontalPodAutoscalers) Update(ctx context.Context, horizontalPodAutoscaler *v2beta2.HorizontalPodAutoscaler, opts v1.UpdateOptions) (result *v2beta2.HorizontalPodAutoscaler, err error) {
	result = &v2beta2.HorizontalPodAutoscaler{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("horizontalpodautoscalers").
		Name(horizontalPodAutoscaler.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(horizontalPodAutoscaler).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *horizontalPodAutoscalers) UpdateStatus(ctx context.Context, horizontalPodAutoscaler *v2beta2.HorizontalPodAutoscaler, opts v1.UpdateOptions) (result *v2beta2.HorizontalPodAutoscaler, err error) {
	result = &v2beta2.HorizontalPodAutoscaler{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("horizontalpodautoscalers").
		Name(horizontalPodAutoscaler.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(horizontalPodAutoscaler).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the horizontalPodAutoscaler and deletes it. Returns an error if one occurs.
func (c *horizontalPodAutoscalers) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("horizontalpodautoscalers").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *horizontalPodAutoscalers) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("horizontalpodautoscalers").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched horizontalPodAutoscaler.
func (c *horizontalPodAutoscalers) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v2beta2.HorizontalPodAutoscaler, err error) {
	result = &v2beta2.HorizontalPodAutoscaler{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("horizontalpodautoscalers").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}

// Apply takes the given apply declarative configuration, applies it and returns the applied horizontalPodAutoscaler.
func (c *horizontalPodAutoscalers) Apply(ctx context.Context, horizontalPodAutoscaler *autoscalingv2beta2.HorizontalPodAutoscalerApplyConfiguration, opts v1.ApplyOptions) (result *v2beta2.HorizontalPodAutoscaler, err error) {
	if horizontalPodAutoscaler == nil {
		return nil, fmt.Errorf("horizontalPodAutoscaler provided to Apply must not be nil")
	}
	patchOpts := opts.ToPatchOptions()
	data, err := json.Marshal(horizontalPodAutoscaler)
	if err != nil {
		return nil, err
	}
	name := horizontalPodAutoscaler.Name
	if name == nil {
		return nil, fmt.Errorf("horizontalPodAutoscaler.Name must be provided to Apply")
	}
	result = &v2beta2.HorizontalPodAutoscaler{}
	err = c.client.Patch(types.ApplyPatchType).
		Namespace(c.ns).
		Resource("horizontalpodautoscalers").
		Name(*name).
		VersionedParams(&patchOpts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}

// ApplyStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
func (c *horizontalPodAutoscalers) ApplyStatus(ctx context.Context, horizontalPodAutoscaler *autoscalingv2beta2.HorizontalPodAutoscalerApplyConfiguration, opts v1.ApplyOptions) (result *v2beta2.HorizontalPodAutoscaler, err error) {
	if horizontalPodAutoscaler == nil {
		return nil, fmt.Errorf("horizontalPodAutoscaler provided to Apply must not be nil")
	}
	patchOpts := opts.ToPatchOptions()
	data, err := json.Marshal(horizontalPodAutoscaler)
	if err != nil {
		return nil, err
	}

	name := horizontalPodAutoscaler.Name
	if name == nil {
		return nil, fmt.Errorf("horizontalPodAutoscaler.Name must be provided to Apply")
	}

	result = &v2beta2.HorizontalPodAutoscaler{}
	err = c.client.Patch(types.ApplyPatchType).
		Namespace(c.ns).
		Resource("horizontalpodautoscalers").
		Name(*name).
		SubResource("status").
		VersionedParams(&patchOpts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
