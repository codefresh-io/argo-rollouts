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

package v1alpha1

import (
	"context"
	time "time"

	rolloutsv1alpha1 "github.com/codefresh-io/argo-rollouts/pkg/apis/rollouts/v1alpha1"
	versioned "github.com/codefresh-io/argo-rollouts/pkg/client/clientset/versioned"
	internalinterfaces "github.com/codefresh-io/argo-rollouts/pkg/client/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/codefresh-io/argo-rollouts/pkg/client/listers/rollouts/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// AnalysisTemplateInformer provides access to a shared informer and lister for
// AnalysisTemplates.
type AnalysisTemplateInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.AnalysisTemplateLister
}

type analysisTemplateInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewAnalysisTemplateInformer constructs a new informer for AnalysisTemplate type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewAnalysisTemplateInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredAnalysisTemplateInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredAnalysisTemplateInformer constructs a new informer for AnalysisTemplate type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredAnalysisTemplateInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ArgoprojV1alpha1().AnalysisTemplates(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ArgoprojV1alpha1().AnalysisTemplates(namespace).Watch(context.TODO(), options)
			},
		},
		&rolloutsv1alpha1.AnalysisTemplate{},
		resyncPeriod,
		indexers,
	)
}

func (f *analysisTemplateInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredAnalysisTemplateInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *analysisTemplateInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&rolloutsv1alpha1.AnalysisTemplate{}, f.defaultInformer)
}

func (f *analysisTemplateInformer) Lister() v1alpha1.AnalysisTemplateLister {
	return v1alpha1.NewAnalysisTemplateLister(f.Informer().GetIndexer())
}
