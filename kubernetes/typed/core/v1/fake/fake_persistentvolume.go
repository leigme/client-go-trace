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

package fake

import (
	"context"
	json "encoding/json"
	"fmt"

	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	applyconfigurationscorev1 "github.com/leigme/client-go-trace/applyconfigurations/core/v1"
	testing "github.com/leigme/client-go-trace/testing"
)

// FakePersistentVolumes implements PersistentVolumeInterface
type FakePersistentVolumes struct {
	Fake *FakeCoreV1
}

var persistentvolumesResource = schema.GroupVersionResource{Group: "", Version: "v1", Resource: "persistentvolumes"}

var persistentvolumesKind = schema.GroupVersionKind{Group: "", Version: "v1", Kind: "PersistentVolume"}

// Get takes name of the persistentVolume, and returns the corresponding persistentVolume object, and an error if there is any.
func (c *FakePersistentVolumes) Get(ctx context.Context, name string, options v1.GetOptions) (result *corev1.PersistentVolume, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(persistentvolumesResource, name), &corev1.PersistentVolume{})
	if obj == nil {
		return nil, err
	}
	return obj.(*corev1.PersistentVolume), err
}

// List takes label and field selectors, and returns the list of PersistentVolumes that match those selectors.
func (c *FakePersistentVolumes) List(ctx context.Context, opts v1.ListOptions) (result *corev1.PersistentVolumeList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(persistentvolumesResource, persistentvolumesKind, opts), &corev1.PersistentVolumeList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &corev1.PersistentVolumeList{ListMeta: obj.(*corev1.PersistentVolumeList).ListMeta}
	for _, item := range obj.(*corev1.PersistentVolumeList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested persistentVolumes.
func (c *FakePersistentVolumes) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(persistentvolumesResource, opts))
}

// Create takes the representation of a persistentVolume and creates it.  Returns the server's representation of the persistentVolume, and an error, if there is any.
func (c *FakePersistentVolumes) Create(ctx context.Context, persistentVolume *corev1.PersistentVolume, opts v1.CreateOptions) (result *corev1.PersistentVolume, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(persistentvolumesResource, persistentVolume), &corev1.PersistentVolume{})
	if obj == nil {
		return nil, err
	}
	return obj.(*corev1.PersistentVolume), err
}

// Update takes the representation of a persistentVolume and updates it. Returns the server's representation of the persistentVolume, and an error, if there is any.
func (c *FakePersistentVolumes) Update(ctx context.Context, persistentVolume *corev1.PersistentVolume, opts v1.UpdateOptions) (result *corev1.PersistentVolume, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(persistentvolumesResource, persistentVolume), &corev1.PersistentVolume{})
	if obj == nil {
		return nil, err
	}
	return obj.(*corev1.PersistentVolume), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakePersistentVolumes) UpdateStatus(ctx context.Context, persistentVolume *corev1.PersistentVolume, opts v1.UpdateOptions) (*corev1.PersistentVolume, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(persistentvolumesResource, "status", persistentVolume), &corev1.PersistentVolume{})
	if obj == nil {
		return nil, err
	}
	return obj.(*corev1.PersistentVolume), err
}

// Delete takes name of the persistentVolume and deletes it. Returns an error if one occurs.
func (c *FakePersistentVolumes) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(persistentvolumesResource, name), &corev1.PersistentVolume{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakePersistentVolumes) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(persistentvolumesResource, listOpts)

	_, err := c.Fake.Invokes(action, &corev1.PersistentVolumeList{})
	return err
}

// Patch applies the patch and returns the patched persistentVolume.
func (c *FakePersistentVolumes) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *corev1.PersistentVolume, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(persistentvolumesResource, name, pt, data, subresources...), &corev1.PersistentVolume{})
	if obj == nil {
		return nil, err
	}
	return obj.(*corev1.PersistentVolume), err
}

// Apply takes the given apply declarative configuration, applies it and returns the applied persistentVolume.
func (c *FakePersistentVolumes) Apply(ctx context.Context, persistentVolume *applyconfigurationscorev1.PersistentVolumeApplyConfiguration, opts v1.ApplyOptions) (result *corev1.PersistentVolume, err error) {
	if persistentVolume == nil {
		return nil, fmt.Errorf("persistentVolume provided to Apply must not be nil")
	}
	data, err := json.Marshal(persistentVolume)
	if err != nil {
		return nil, err
	}
	name := persistentVolume.Name
	if name == nil {
		return nil, fmt.Errorf("persistentVolume.Name must be provided to Apply")
	}
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(persistentvolumesResource, *name, types.ApplyPatchType, data), &corev1.PersistentVolume{})
	if obj == nil {
		return nil, err
	}
	return obj.(*corev1.PersistentVolume), err
}

// ApplyStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
func (c *FakePersistentVolumes) ApplyStatus(ctx context.Context, persistentVolume *applyconfigurationscorev1.PersistentVolumeApplyConfiguration, opts v1.ApplyOptions) (result *corev1.PersistentVolume, err error) {
	if persistentVolume == nil {
		return nil, fmt.Errorf("persistentVolume provided to Apply must not be nil")
	}
	data, err := json.Marshal(persistentVolume)
	if err != nil {
		return nil, err
	}
	name := persistentVolume.Name
	if name == nil {
		return nil, fmt.Errorf("persistentVolume.Name must be provided to Apply")
	}
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(persistentvolumesResource, *name, types.ApplyPatchType, data, "status"), &corev1.PersistentVolume{})
	if obj == nil {
		return nil, err
	}
	return obj.(*corev1.PersistentVolume), err
}
