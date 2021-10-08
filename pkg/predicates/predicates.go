package predicates

import (
	"reflect"

	"sigs.k8s.io/controller-runtime/pkg/event"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
)

// This code is based on an OpenShift article:
// https://www.openshift.com/blog/kubernetes-operators-best-practices

type ResourceGenerationOrFinalizerChangedPredicate struct {
	predicate.Funcs
}

// Update implements default UpdateEvent filter for validating resource version change
func (ResourceGenerationOrFinalizerChangedPredicate) UpdateFunc(e event.UpdateEvent) bool {
	if e.ObjectNew.GetGeneration() == e.ObjectNew.GetGeneration() && reflect.DeepEqual(e.ObjectNew.GetFinalizers(), e.ObjectNew.GetFinalizers()) {
		return false
	}
	return true
}

func HelmSecretFilterPredicate() predicate.Predicate {
	return predicate.Funcs{
		CreateFunc: func(e event.CreateEvent) bool {
			return true
		},
		DeleteFunc: func(e event.DeleteEvent) bool {
			return true
		},
		GenericFunc: func(event.GenericEvent) bool {
			return false
		},
		UpdateFunc: func(e event.UpdateEvent) bool {
			return true
		},
	}
}
