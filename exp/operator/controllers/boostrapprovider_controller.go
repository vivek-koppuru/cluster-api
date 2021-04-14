/*
Copyright 2021 The Kubernetes Authors.

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

package controllers

import (
	"context"

	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"

	operatorv1alpha4 "sigs.k8s.io/cluster-api/exp/operator/api/v1alpha4"
)

// BootstrapProviderReconciler reconciles a BootstrapProvider object.
type BootstrapProviderReconciler struct {
	Client client.Client
}

// +kubebuilder:rbac:groups=operator.cluster.x-k8s.io,resources=bootstrapproviders,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=operator.cluster.x-k8s.io,resources=bootstrapproviders/status,verbs=get;update;patch

func (r *BootstrapProviderReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	log := ctrl.LoggerFrom(ctx)
	log.V(4).Info("Reconcile BootstrapProvider")

	return ctrl.Result{}, nil
}

func (r *BootstrapProviderReconciler) SetupWithManager(mgr ctrl.Manager, options controller.Options) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&operatorv1alpha4.BootstrapProvider{}).
		WithOptions(options).
		Complete(r)
}
