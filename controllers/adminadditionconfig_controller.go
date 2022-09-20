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
	mymc "github.com/Dana-Team/ManifestController/pkg/ManifestController"
	"github.com/go-logr/logr"
	"k8s.io/apimachinery/pkg/runtime"
	configv1alpha1 "metoperator/api/v1alpha1"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/source"
	"sigs.k8s.io/kubebuilder-declarative-pattern/pkg/patterns/addon"
	"sigs.k8s.io/kubebuilder-declarative-pattern/pkg/patterns/addon/pkg/loaders"
	"sigs.k8s.io/kubebuilder-declarative-pattern/pkg/patterns/addon/pkg/status"
	"sigs.k8s.io/kubebuilder-declarative-pattern/pkg/patterns/declarative"
)

var _ reconcile.Reconciler = &AdminAdditionConfigReconciler{}

// AdminAdditionConfigReconciler reconciles a AdminAdditionConfig object
type AdminAdditionConfigReconciler struct {
	client.Client
	Log    logr.Logger
	Scheme *runtime.Scheme

	declarative.Reconciler
}

//+kubebuilder:rbac:groups=config.dana.io,resources=adminadditionconfigs,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=config.dana.io,resources=adminadditionconfigs/status,verbs=get;update;patch

// SetupWithManager sets up the controller with the Manager.
func (r *AdminAdditionConfigReconciler) SetupWithManager(mgr ctrl.Manager) error {
	addon.Init()

	labels := map[string]string{
		"k8s-app": "adminadditionconfig",
	}

	watchLabels := declarative.SourceLabel(mgr.GetScheme())
	myManifestController, _ := mymc.NewManifestLoader(loaders.FlagChannel)
	if err := r.Reconciler.Init(mgr, &configv1alpha1.AdminAdditionConfig{},
		declarative.WithObjectTransform(declarative.AddLabels(labels)),
		declarative.WithOwner(declarative.SourceAsOwner),
		declarative.WithLabels(watchLabels),
		declarative.WithStatus(status.NewBasic(mgr.GetClient())),
		declarative.WithApplyKustomize(),
		declarative.WithApplyPrune(),
		declarative.WithManifestController(myManifestController),
		declarative.WithReconcileMetrics(0, nil),
		declarative.WithObjectTransform(addon.ApplyPatches),
	); err != nil {
		return err
	}
	c, err := controller.New("adminadditionconfig-controller", mgr, controller.Options{Reconciler: r})
	if err != nil {
		return err
	}

	// Watch for changes to AdminAdditionConfig
	err = c.Watch(&source.Kind{Type: &configv1alpha1.AdminAdditionConfig{}}, &handler.EnqueueRequestForObject{})
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
