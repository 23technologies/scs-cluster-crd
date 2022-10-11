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
	gardencorev1beta1 "github.com/gardener/gardener/pkg/apis/core/v1beta1"
	"github.com/sirupsen/logrus"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	scscommunityv1alpha1 "scs.community/v1alpha1/api/v1alpha1"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"
)

// ClusterReconciler reconciles a Cluster object
type ClusterReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=scs.community.my.domain,resources=clusters,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=scs.community.my.domain,resources=clusters/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=scs.community.my.domain,resources=clusters/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the Cluster object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.13.0/pkg/reconcile

/*
target-mk8s = gardener/cluster-api/gridscale/...
garden-controller():

	template shoot.yaml (name/k8s verison)
	kubectl apply shoot.yaml
*/
func (r *ClusterReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	_ = log.FromContext(ctx)
	logrus.Info("Cluster reconcile: ", req.String())

	var cluster scscommunityv1alpha1.Cluster

	err := r.Get(ctx, req.NamespacedName, &cluster)
	if err != nil {
		logrus.Warn(err)
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	k8sVersion := cluster.Spec.Kubernetes.Version
	clusterName := cluster.ObjectMeta.GetName()
	clusterNamespace := cluster.ObjectMeta.GetNamespace()
	logrus.Info("Cluster: ", clusterNamespace, "/", clusterName, " , k8s-version: ", k8sVersion)


	var purpose gardencorev1beta1.ShootPurpose = "testing"
	shoot := gardencorev1beta1.Shoot{
		TypeMeta:   metav1.TypeMeta{},
		ObjectMeta: metav1.ObjectMeta{
			Name: clusterName,
			Namespace: clusterNamespace,
		},
		Spec: gardencorev1beta1.ShootSpec{
			Addons:            &gardencorev1beta1.Addons{
				KubernetesDashboard: &gardencorev1beta1.KubernetesDashboard{},
				NginxIngress:        &gardencorev1beta1.NginxIngress{},
			},
			CloudProfileName:  "hcloud",
			//DNS:               nil,
			//Extensions:        nil,
			Hibernation:       nil,
			Kubernetes:        gardencorev1beta1.Kubernetes{
				Version:                     k8sVersion,
			},
			Networking:        gardencorev1beta1.Networking{},
			Maintenance:       nil,
			Monitoring:        nil,
			Provider:          gardencorev1beta1.Provider{},
			Purpose:           &purpose,
			Region:            "",
			SecretBindingName: "",
			SeedName:          nil,
			SeedSelector:      nil,
			Resources:         nil,
			Tolerations:       nil,
			ExposureClassName: nil,
			SystemComponents:  nil,
			ControlPlane:      nil,
		},
		Status:     gardencorev1beta1.ShootStatus{},
	}

	err = r.Create(ctx, &shoot)
	if err != nil{
		logrus.Warn(err)
	}


	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *ClusterReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&scscommunityv1alpha1.Cluster{}).
		Complete(r)
}
