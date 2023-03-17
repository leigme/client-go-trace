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

// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	v1 "k8s.io/api/admissionregistration/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"github.com/leigme/client-go-trace/tools/cache"
)

// MutatingWebhookConfigurationLister helps list MutatingWebhookConfigurations.
// All objects returned here must be treated as read-only.
type MutatingWebhookConfigurationLister interface {
	// List lists all MutatingWebhookConfigurations in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.MutatingWebhookConfiguration, err error)
	// Get retrieves the MutatingWebhookConfiguration from the index for a given name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1.MutatingWebhookConfiguration, error)
	MutatingWebhookConfigurationListerExpansion
}

// mutatingWebhookConfigurationLister implements the MutatingWebhookConfigurationLister interface.
type mutatingWebhookConfigurationLister struct {
	indexer cache.Indexer
}

// NewMutatingWebhookConfigurationLister returns a new MutatingWebhookConfigurationLister.
func NewMutatingWebhookConfigurationLister(indexer cache.Indexer) MutatingWebhookConfigurationLister {
	return &mutatingWebhookConfigurationLister{indexer: indexer}
}

// List lists all MutatingWebhookConfigurations in the indexer.
func (s *mutatingWebhookConfigurationLister) List(selector labels.Selector) (ret []*v1.MutatingWebhookConfiguration, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.MutatingWebhookConfiguration))
	})
	return ret, err
}

// Get retrieves the MutatingWebhookConfiguration from the index for a given name.
func (s *mutatingWebhookConfigurationLister) Get(name string) (*v1.MutatingWebhookConfiguration, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("mutatingwebhookconfiguration"), name)
	}
	return obj.(*v1.MutatingWebhookConfiguration), nil
}
