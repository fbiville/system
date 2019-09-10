/*
 * Copyright 2019 The original author or authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
package v1alpha1

import (
	v1alpha1 "github.com/projectriff/system/pkg/apis/streaming/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ProviderLister helps list Providers.
type ProviderLister interface {
	// List lists all Providers in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.Provider, err error)
	// Providers returns an object that can list and get Providers.
	Providers(namespace string) ProviderNamespaceLister
	ProviderListerExpansion
}

// providerLister implements the ProviderLister interface.
type providerLister struct {
	indexer cache.Indexer
}

// NewProviderLister returns a new ProviderLister.
func NewProviderLister(indexer cache.Indexer) ProviderLister {
	return &providerLister{indexer: indexer}
}

// List lists all Providers in the indexer.
func (s *providerLister) List(selector labels.Selector) (ret []*v1alpha1.Provider, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Provider))
	})
	return ret, err
}

// Providers returns an object that can list and get Providers.
func (s *providerLister) Providers(namespace string) ProviderNamespaceLister {
	return providerNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ProviderNamespaceLister helps list and get Providers.
type ProviderNamespaceLister interface {
	// List lists all Providers in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.Provider, err error)
	// Get retrieves the Provider from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.Provider, error)
	ProviderNamespaceListerExpansion
}

// providerNamespaceLister implements the ProviderNamespaceLister
// interface.
type providerNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Providers in the indexer for a given namespace.
func (s providerNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.Provider, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Provider))
	})
	return ret, err
}

// Get retrieves the Provider from the indexer for a given namespace and name.
func (s providerNamespaceLister) Get(name string) (*v1alpha1.Provider, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("provider"), name)
	}
	return obj.(*v1alpha1.Provider), nil
}
