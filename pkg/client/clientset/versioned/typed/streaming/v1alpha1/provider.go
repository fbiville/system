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
	scheme "github.com/projectriff/system/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ProvidersGetter has a method to return a ProviderInterface.
// A group's client should implement this interface.
type ProvidersGetter interface {
	Providers(namespace string) ProviderInterface
}

// ProviderInterface has methods to work with Provider resources.
type ProviderInterface interface {
	Create(*v1alpha1.Provider) (*v1alpha1.Provider, error)
	Update(*v1alpha1.Provider) (*v1alpha1.Provider, error)
	UpdateStatus(*v1alpha1.Provider) (*v1alpha1.Provider, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.Provider, error)
	List(opts v1.ListOptions) (*v1alpha1.ProviderList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Provider, err error)
	ProviderExpansion
}

// providers implements ProviderInterface
type providers struct {
	client rest.Interface
	ns     string
}

// newProviders returns a Providers
func newProviders(c *StreamingV1alpha1Client, namespace string) *providers {
	return &providers{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the provider, and returns the corresponding provider object, and an error if there is any.
func (c *providers) Get(name string, options v1.GetOptions) (result *v1alpha1.Provider, err error) {
	result = &v1alpha1.Provider{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("providers").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Providers that match those selectors.
func (c *providers) List(opts v1.ListOptions) (result *v1alpha1.ProviderList, err error) {
	result = &v1alpha1.ProviderList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("providers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested providers.
func (c *providers) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("providers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a provider and creates it.  Returns the server's representation of the provider, and an error, if there is any.
func (c *providers) Create(provider *v1alpha1.Provider) (result *v1alpha1.Provider, err error) {
	result = &v1alpha1.Provider{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("providers").
		Body(provider).
		Do().
		Into(result)
	return
}

// Update takes the representation of a provider and updates it. Returns the server's representation of the provider, and an error, if there is any.
func (c *providers) Update(provider *v1alpha1.Provider) (result *v1alpha1.Provider, err error) {
	result = &v1alpha1.Provider{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("providers").
		Name(provider.Name).
		Body(provider).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *providers) UpdateStatus(provider *v1alpha1.Provider) (result *v1alpha1.Provider, err error) {
	result = &v1alpha1.Provider{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("providers").
		Name(provider.Name).
		SubResource("status").
		Body(provider).
		Do().
		Into(result)
	return
}

// Delete takes name of the provider and deletes it. Returns an error if one occurs.
func (c *providers) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("providers").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *providers) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("providers").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched provider.
func (c *providers) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Provider, err error) {
	result = &v1alpha1.Provider{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("providers").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
