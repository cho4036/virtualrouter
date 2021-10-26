/*
Copyright The Kubernetes Authors.

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

package v1

import (
	"context"
	time "time"

	networkcontrollerv1 "github.com/tmax-cloud/virtualrouter/pkg/apis/networkcontroller/v1"
	versioned "github.com/tmax-cloud/virtualrouter/pkg/client/clientset/versioned"
	internalinterfaces "github.com/tmax-cloud/virtualrouter/pkg/client/informers/externalversions/internalinterfaces"
	v1 "github.com/tmax-cloud/virtualrouter/pkg/client/listers/networkcontroller/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// LoadBalancerRuleInformer provides access to a shared informer and lister for
// LoadBalancerRules.
type LoadBalancerRuleInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.LoadBalancerRuleLister
}

type loadBalancerRuleInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewLoadBalancerRuleInformer constructs a new informer for LoadBalancerRule type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewLoadBalancerRuleInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredLoadBalancerRuleInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredLoadBalancerRuleInformer constructs a new informer for LoadBalancerRule type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredLoadBalancerRuleInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.TmaxV1().LoadBalancerRules(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.TmaxV1().LoadBalancerRules(namespace).Watch(context.TODO(), options)
			},
		},
		&networkcontrollerv1.LoadBalancerRule{},
		resyncPeriod,
		indexers,
	)
}

func (f *loadBalancerRuleInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredLoadBalancerRuleInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *loadBalancerRuleInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&networkcontrollerv1.LoadBalancerRule{}, f.defaultInformer)
}

func (f *loadBalancerRuleInformer) Lister() v1.LoadBalancerRuleLister {
	return v1.NewLoadBalancerRuleLister(f.Informer().GetIndexer())
}
