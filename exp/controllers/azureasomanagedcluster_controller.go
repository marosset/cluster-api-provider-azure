/*
Copyright 2024 The Kubernetes Authors.

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

	infrav1exp "sigs.k8s.io/cluster-api-provider-azure/exp/api/v1alpha1"
	"sigs.k8s.io/cluster-api-provider-azure/util/tele"
	"sigs.k8s.io/cluster-api/util/predicates"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// AzureASOManagedClusterReconciler reconciles a AzureASOManagedCluster object.
type AzureASOManagedClusterReconciler struct {
	client.Client
	WatchFilterValue string
}

// SetupWithManager sets up the controller with the Manager.
func (r *AzureASOManagedClusterReconciler) SetupWithManager(ctx context.Context, mgr ctrl.Manager) error {
	_, log, done := tele.StartSpanWithLogger(ctx,
		"controllers.AzureASOManagedClusterReconciler.SetupWithManager",
		tele.KVP("controller", infrav1exp.AzureASOManagedClusterKind),
	)
	defer done()

	_, err := ctrl.NewControllerManagedBy(mgr).
		For(&infrav1exp.AzureASOManagedCluster{}).
		WithEventFilter(predicates.ResourceHasFilterLabel(log, r.WatchFilterValue)).
		Build(r)
	if err != nil {
		return err
	}

	return nil
}

//+kubebuilder:rbac:groups=infrastructure.cluster.x-k8s.io,resources=azureasomanagedclusters,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=infrastructure.cluster.x-k8s.io,resources=azureasomanagedclusters/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=infrastructure.cluster.x-k8s.io,resources=azureasomanagedclusters/finalizers,verbs=update

// Reconcile reconciles an AzureASOManagedCluster.
func (r *AzureASOManagedClusterReconciler) Reconcile(ctx context.Context, req ctrl.Request) (result ctrl.Result, resultErr error) {
	return ctrl.Result{}, nil
}