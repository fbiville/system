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
	internalinterfaces "github.com/projectriff/system/pkg/client/informers/externalversions/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// Processors returns a ProcessorInformer.
	Processors() ProcessorInformer
	// Providers returns a ProviderInformer.
	Providers() ProviderInformer
	// Streams returns a StreamInformer.
	Streams() StreamInformer
}

type version struct {
	factory          internalinterfaces.SharedInformerFactory
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &version{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}

// Processors returns a ProcessorInformer.
func (v *version) Processors() ProcessorInformer {
	return &processorInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// Providers returns a ProviderInformer.
func (v *version) Providers() ProviderInformer {
	return &providerInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// Streams returns a StreamInformer.
func (v *version) Streams() StreamInformer {
	return &streamInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}
