/*
Copyright 2024 Nokia.

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

package configset

import (
	"context"

	"github.com/henderiw/logger/log"
	invv1alpha1 "github.com/sdcio/config-server/apis/inv/v1alpha1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/util/workqueue"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/event"
)

type targetEventHandler struct {
	client client.Client
}

// Create enqueues a request for all ip allocation within the ipam
func (r *targetEventHandler) Create(ctx context.Context, evt event.CreateEvent, q workqueue.RateLimitingInterface) {
	r.add(ctx, evt.Object, q)
}

// Create enqueues a request for all ip allocation within the ipam
func (r *targetEventHandler) Update(ctx context.Context, evt event.UpdateEvent, q workqueue.RateLimitingInterface) {
	r.add(ctx, evt.ObjectOld, q)
	r.add(ctx, evt.ObjectNew, q)
}

// Create enqueues a request for all ip allocation within the ipam
func (r *targetEventHandler) Delete(ctx context.Context, evt event.DeleteEvent, q workqueue.RateLimitingInterface) {
	r.add(ctx, evt.Object, q)
}

// Create enqueues a request for all ip allocation within the ipam
func (r *targetEventHandler) Generic(ctx context.Context, evt event.GenericEvent, q workqueue.RateLimitingInterface) {
	r.add(ctx, evt.Object, q)
}

func (r *targetEventHandler) add(ctx context.Context, obj runtime.Object, queue adder) {
	cr, ok := obj.(*invv1alpha1.Target)
	if !ok {
		return
	}
	log := log.FromContext(ctx)

	log.Info("event", "gvk", invv1alpha1.TargetGroupVersionKind.String(), "name", cr.GetName())

	// list all the configs of the particular target that got changed
	/*
		opts := []client.ListOption{
			client.InNamespace(cr.Namespace),
			client.MatchingLabels{
				configv1alpha1.TargetNameKey:      cr.GetName(),
				configv1alpha1.TargetNamespaceKey: cr.GetNamespace(),
			},
		}
		configs := &configv1alpha1.ConfigSetList{}
		if err := r.client.List(ctx, configs, opts...); err != nil {
			log.Error("cannot list targets", "error", err)
			return
		}
		for _, config := range configs.Items {
			key := types.NamespacedName{
				Namespace: config.Namespace,
				Name:      config.Name}
			log.Info("event requeue config", "key", key.String())
			queue.Add(reconcile.Request{NamespacedName: key})
		}
	*/
}
