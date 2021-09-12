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

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/codefresh-io/argo-rollouts/pkg/apis/rollouts/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ClusterAnalysisTemplateLister helps list ClusterAnalysisTemplates.
// All objects returned here must be treated as read-only.
type ClusterAnalysisTemplateLister interface {
	// List lists all ClusterAnalysisTemplates in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ClusterAnalysisTemplate, err error)
	// Get retrieves the ClusterAnalysisTemplate from the index for a given name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.ClusterAnalysisTemplate, error)
	ClusterAnalysisTemplateListerExpansion
}

// clusterAnalysisTemplateLister implements the ClusterAnalysisTemplateLister interface.
type clusterAnalysisTemplateLister struct {
	indexer cache.Indexer
}

// NewClusterAnalysisTemplateLister returns a new ClusterAnalysisTemplateLister.
func NewClusterAnalysisTemplateLister(indexer cache.Indexer) ClusterAnalysisTemplateLister {
	return &clusterAnalysisTemplateLister{indexer: indexer}
}

// List lists all ClusterAnalysisTemplates in the indexer.
func (s *clusterAnalysisTemplateLister) List(selector labels.Selector) (ret []*v1alpha1.ClusterAnalysisTemplate, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ClusterAnalysisTemplate))
	})
	return ret, err
}

// Get retrieves the ClusterAnalysisTemplate from the index for a given name.
func (s *clusterAnalysisTemplateLister) Get(name string) (*v1alpha1.ClusterAnalysisTemplate, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("clusteranalysistemplate"), name)
	}
	return obj.(*v1alpha1.ClusterAnalysisTemplate), nil
}
