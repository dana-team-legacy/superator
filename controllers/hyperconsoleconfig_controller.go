/*
Copyright 2022.

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
	"github.com/go-logr/logr"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/source"
	"sigs.k8s.io/kubebuilder-declarative-pattern/pkg/patterns/addon"
	"sigs.k8s.io/kubebuilder-declarative-pattern/pkg/patterns/addon/pkg/status"
	"sigs.k8s.io/kubebuilder-declarative-pattern/pkg/patterns/declarative"

	configv1alpha1 "metoperator/api/v1alpha1"
)

var _ reconcile.Reconciler = &HyperConsoleConfigReconciler{}

// HyperConsoleConfigReconciler reconciles a HyperConsoleConfig object
type HyperConsoleConfigReconciler struct {
	client.Client
	Log    logr.Logger
	Scheme *runtime.Scheme

	declarative.Reconciler
}

//+kubebuilder:rbac:groups=config.dana.io,resources=hyperconsoleconfigs,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=config.dana.io,resources=hyperconsoleconfigs/status,verbs=get;update;patch

// SetupWithManager sets up the controller with the Manager.
func (r *HyperConsoleConfigReconciler) SetupWithManager(mgr ctrl.Manager) error {
	addon.Init()

	labels := map[string]string{
		"k8s-app": "hyperconsoleconfig",
	}

	watchLabels := declarative.SourceLabel(mgr.GetScheme())

	if err := r.Reconciler.Init(mgr, &configv1alpha1.HyperConsoleConfig{},
		declarative.WithObjectTransform(declarative.AddLabels(labels)),
		declarative.WithOwner(declarative.SourceAsOwner),
		declarative.WithLabels(watchLabels),
		declarative.WithStatus(status.NewBasic(mgr.GetClient())),
		declarative.WithApplyKustomize(),
		declarative.WithApplyPrune(),
		declarative.WithReconcileMetrics(0, nil),
		declarative.WithObjectTransform(addon.ApplyPatches),
	); err != nil {
		return err
	}

	c, err := controller.New("hyperconsoleconfig-controller", mgr, controller.Options{Reconciler: r})
	if err != nil {
		return err
	}

	// Watch for changes to HyperConsoleConfig
	err = c.Watch(&source.Kind{Type: &configv1alpha1.HyperConsoleConfig{}}, &handler.EnqueueRequestForObject{})
	if err != nil {
		return err
	}

	// Watch for changes to deployed objects
	_, err = declarative.WatchAll(mgr.GetConfig(), c, r, watchLabels)
	if err != nil {
		return err
	}

	return nil
}
