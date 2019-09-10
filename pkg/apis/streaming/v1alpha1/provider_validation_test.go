/*
Copyright 2018 The Knative Authors

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

package v1alpha1

import (
	"context"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/knative/pkg/apis"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestValidateProvider(t *testing.T) {
	for _, c := range []struct {
		name string
		s    *Provider
		want *apis.FieldError
	}{{
		name: "valid",
		s: &Provider{
			ObjectMeta: metav1.ObjectMeta{
				Name: "test-provider",
			},
			Spec: ProviderSpec{
				BrokerType: "kafka",
			},
		},
		want: nil,
	}, {
		name: "validates metadata",
		s: &Provider{
			ObjectMeta: metav1.ObjectMeta{},
			Spec: ProviderSpec{
				BrokerType: "kafka",
			},
		},
		want: apis.ErrMissingOneOf("metadata.generateName", "metadata.name"),
	}, {
		name: "validates spec",
		s: &Provider{
			ObjectMeta: metav1.ObjectMeta{
				Name: "test-provider",
			},
			Spec: ProviderSpec{},
		},
		want: apis.ErrMissingField("spec"),
	}} {
		name := c.name
		t.Run(name, func(t *testing.T) {
			got := c.s.Validate(context.Background())
			if diff := cmp.Diff(c.want.Error(), got.Error()); diff != "" {
				t.Errorf("validateProvider(%s) (-want, +got) = %v", name, diff)
			}
		})
	}
}

func TestValidateProviderSpec(t *testing.T) {
	for _, c := range []struct {
		name string
		s    *ProviderSpec
		want *apis.FieldError
	}{
		{
			name: "valid",
			s: &ProviderSpec{
				BrokerType: "kafka",
			},
			want: nil,
		},
		{
			name: "empty",
			s:    &ProviderSpec{},
			want: apis.ErrMissingField(apis.CurrentField),
		},
		{
			name: "requires broker type",
			s: &ProviderSpec{
				BrokerType: "",
				Config:     map[string]string{"foo": "bar"},
			},
			want: apis.ErrMissingField("brokerType"),
		},
	} {
		name := c.name
		t.Run(name, func(t *testing.T) {
			got := c.s.Validate(context.Background())
			if diff := cmp.Diff(c.want.Error(), got.Error()); diff != "" {
				t.Errorf("validateProviderSpec(%s) (-want, +got) = %v", name, diff)
			}
		})
	}
}
