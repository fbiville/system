package v1alpha1

import "github.com/knative/pkg/apis"

func (ps *ProviderStatus) IsReady() bool {
	panic("implement me")
}

func (ps *ProviderStatus) GetReadyConditionType() apis.ConditionType {
	panic("implement me")
}

func (ps *ProviderStatus) GetObservedGeneration() int64 {
	panic("implement me")
}

