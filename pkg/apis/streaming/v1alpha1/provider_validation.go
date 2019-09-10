package v1alpha1

import (
	"context"
	"github.com/knative/pkg/apis"
	systemapis "github.com/projectriff/system/pkg/apis"
	"k8s.io/apimachinery/pkg/api/equality"
)

func (p *Provider) Validate(ctx context.Context) *apis.FieldError {
	errs := &apis.FieldError{}
	errs = errs.Also(systemapis.ValidateObjectMetadata(p.GetObjectMeta()).ViaField("metadata"))
	errs = errs.Also(p.Spec.Validate(ctx).ViaField("spec"))
	return errs
}

func (ps *ProviderSpec) Validate(ctx context.Context) *apis.FieldError {
	if equality.Semantic.DeepEqual(ps, &ProviderSpec{}) {
		return apis.ErrMissingField(apis.CurrentField)
	}

	errs := &apis.FieldError{}

	if ps.BrokerType == "" {
		errs = errs.Also(apis.ErrMissingField("brokerType"))
	}

	return errs
}
