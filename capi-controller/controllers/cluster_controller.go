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
	"context"
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	capiv1alpha3 "sigs.k8s.io/cluster-api/api/v1alpha3"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/event"
	"sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/predicate"

	k8sv1alpha1 "github.com/23technologies/scs-cluster-crd/capi-controller/api/v1alpha1"
)

// ClusterReconciler reconciles a Cluster object
type ClusterReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=k8s.scs.community,resources=clusters,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=k8s.scs.community,resources=clusters/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=k8s.scs.community,resources=clusters/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the Cluster object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.13.0/pkg/reconcile
func (r *ClusterReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	_ = log.FromContext(ctx)

	var myCluster k8sv1alpha1.Cluster
	var myCapiCluster capiv1alpha3.Cluster
	err := r.Get(ctx, req.NamespacedName, &myCluster)
	// Successfully retrieved cluster-object
	ctrl.Log.Info("Received new Cluster-Event: " + myCluster.Name + " k8s-version: " + myCluster.Spec.Kubernetes.Version)
	if err != nil {
		ctrl.Log.Error(err, "Problem")
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}
	labels := make(map[string]string)
	labels["k8s.scs.community/version"] = myCluster.Spec.Kubernetes.Version
	myCapiCluster = capiv1alpha3.Cluster{
		ObjectMeta: metav1.ObjectMeta{
			Name:      myCluster.Name,
			Namespace: myCluster.Namespace,
			Labels:    labels,
		},
		Spec: capiv1alpha3.ClusterSpec{},
	}

	err = r.Create(ctx, &myCapiCluster)
	if err != nil {
		ctrl.Log.Error(err, "Problem while creating capi-cluster")
		return ctrl.Result{}, err
	}

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *ClusterReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&k8sv1alpha1.Cluster{}).
		WithEventFilter(predicate.Funcs{
			UpdateFunc: func(e event.UpdateEvent) bool {

				oldGeneration := e.ObjectOld.GetGeneration()
				newGeneration := e.ObjectNew.GetGeneration()
				// Generation is only updated on spec changes (also on deletion),
				// not metadata or status
				// Filter out events where the generation hasn't changed to
				// avoid being triggered by status updates
				if oldGeneration == newGeneration {
					fmt.Println("metadata or status change")
					return false
				}
				return true
			},
			DeleteFunc: func(e event.DeleteEvent) bool {
				// The reconciler adds a finalizer, so we perform clean-up
				// when the delete timestamp is added
				// Suppress Delete events to avoid filtering them out in the Reconcile function
				fmt.Println("delete event")
				return false
			},
		}).
		Complete(r)
}
