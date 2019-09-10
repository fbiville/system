package v1alpha1

import (
	knapis "github.com/knative/pkg/apis"
	duckv1beta1 "github.com/knative/pkg/apis/duck/v1beta1"
	"github.com/knative/pkg/kmeta"
	"github.com/projectriff/system/pkg/apis"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type Provider struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ProviderSpec   `json:"spec"`
	Status ProviderStatus `json:"status"`
}

var (
	_ knapis.Validatable = (*Provider)(nil)
	_ knapis.Defaultable = (*Provider)(nil)
	_ kmeta.OwnerRefable = (*Provider)(nil)
	_ apis.Object        = (*Provider)(nil)
)

type ProviderSpec struct {
	BrokerType string            `json:"brokerType"`
	Config     map[string]string `json:"config"`
}

type ProviderStatus struct {
	duckv1beta1.Status `json:",inline"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type ProviderList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []Provider `json:"items"`
}

func (*Provider) GetGroupVersionKind() schema.GroupVersionKind {
	return SchemeGroupVersion.WithKind("Provider")
}

func (p *Provider) GetStatus() apis.Status {
	return &p.Status
}
