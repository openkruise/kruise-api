/*
Copyright 2022 The Kruise Authors.

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

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	time "time"

	versioned "github.com/openkruise/kruise-api/client/clientset/versioned"
	internalinterfaces "github.com/openkruise/kruise-api/client/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/openkruise/kruise-api/client/listers/policy/v1alpha1"
	policyv1alpha1 "github.com/openkruise/kruise-api/policy/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// PodUnavailableBudgetInformer provides access to a shared informer and lister for
// PodUnavailableBudgets.
type PodUnavailableBudgetInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.PodUnavailableBudgetLister
}

type podUnavailableBudgetInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewPodUnavailableBudgetInformer constructs a new informer for PodUnavailableBudget type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewPodUnavailableBudgetInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredPodUnavailableBudgetInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredPodUnavailableBudgetInformer constructs a new informer for PodUnavailableBudget type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredPodUnavailableBudgetInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.PolicyV1alpha1().PodUnavailableBudgets(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.PolicyV1alpha1().PodUnavailableBudgets(namespace).Watch(context.TODO(), options)
			},
		},
		&policyv1alpha1.PodUnavailableBudget{},
		resyncPeriod,
		indexers,
	)
}

func (f *podUnavailableBudgetInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredPodUnavailableBudgetInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *podUnavailableBudgetInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&policyv1alpha1.PodUnavailableBudget{}, f.defaultInformer)
}

func (f *podUnavailableBudgetInformer) Lister() v1alpha1.PodUnavailableBudgetLister {
	return v1alpha1.NewPodUnavailableBudgetLister(f.Informer().GetIndexer())
}
