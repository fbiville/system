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
	v1alpha1 "github.com/projectriff/system/pkg/apis/stream/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ProcessorLister helps list Processors.
type ProcessorLister interface {
	// List lists all Processors in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.Processor, err error)
	// Processors returns an object that can list and get Processors.
	Processors(namespace string) ProcessorNamespaceLister
	ProcessorListerExpansion
}

// processorLister implements the ProcessorLister interface.
type processorLister struct {
	indexer cache.Indexer
}

// NewProcessorLister returns a new ProcessorLister.
func NewProcessorLister(indexer cache.Indexer) ProcessorLister {
	return &processorLister{indexer: indexer}
}

// List lists all Processors in the indexer.
func (s *processorLister) List(selector labels.Selector) (ret []*v1alpha1.Processor, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Processor))
	})
	return ret, err
}

// Processors returns an object that can list and get Processors.
func (s *processorLister) Processors(namespace string) ProcessorNamespaceLister {
	return processorNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ProcessorNamespaceLister helps list and get Processors.
type ProcessorNamespaceLister interface {
	// List lists all Processors in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.Processor, err error)
	// Get retrieves the Processor from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.Processor, error)
	ProcessorNamespaceListerExpansion
}

// processorNamespaceLister implements the ProcessorNamespaceLister
// interface.
type processorNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Processors in the indexer for a given namespace.
func (s processorNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.Processor, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Processor))
	})
	return ret, err
}

// Get retrieves the Processor from the indexer for a given namespace and name.
func (s processorNamespaceLister) Get(name string) (*v1alpha1.Processor, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("processor"), name)
	}
	return obj.(*v1alpha1.Processor), nil
}
