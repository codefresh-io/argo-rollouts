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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1alpha1 "github.com/codefresh-io/argo-rollouts/pkg/apis/rollouts/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeAnalysisTemplates implements AnalysisTemplateInterface
type FakeAnalysisTemplates struct {
	Fake *FakeArgoprojV1alpha1
	ns   string
}

var analysistemplatesResource = schema.GroupVersionResource{Group: "argoproj.io", Version: "v1alpha1", Resource: "analysistemplates"}

var analysistemplatesKind = schema.GroupVersionKind{Group: "argoproj.io", Version: "v1alpha1", Kind: "AnalysisTemplate"}

// Get takes name of the analysisTemplate, and returns the corresponding analysisTemplate object, and an error if there is any.
func (c *FakeAnalysisTemplates) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.AnalysisTemplate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(analysistemplatesResource, c.ns, name), &v1alpha1.AnalysisTemplate{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AnalysisTemplate), err
}

// List takes label and field selectors, and returns the list of AnalysisTemplates that match those selectors.
func (c *FakeAnalysisTemplates) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.AnalysisTemplateList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(analysistemplatesResource, analysistemplatesKind, c.ns, opts), &v1alpha1.AnalysisTemplateList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AnalysisTemplateList{ListMeta: obj.(*v1alpha1.AnalysisTemplateList).ListMeta}
	for _, item := range obj.(*v1alpha1.AnalysisTemplateList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested analysisTemplates.
func (c *FakeAnalysisTemplates) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(analysistemplatesResource, c.ns, opts))

}

// Create takes the representation of a analysisTemplate and creates it.  Returns the server's representation of the analysisTemplate, and an error, if there is any.
func (c *FakeAnalysisTemplates) Create(ctx context.Context, analysisTemplate *v1alpha1.AnalysisTemplate, opts v1.CreateOptions) (result *v1alpha1.AnalysisTemplate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(analysistemplatesResource, c.ns, analysisTemplate), &v1alpha1.AnalysisTemplate{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AnalysisTemplate), err
}

// Update takes the representation of a analysisTemplate and updates it. Returns the server's representation of the analysisTemplate, and an error, if there is any.
func (c *FakeAnalysisTemplates) Update(ctx context.Context, analysisTemplate *v1alpha1.AnalysisTemplate, opts v1.UpdateOptions) (result *v1alpha1.AnalysisTemplate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(analysistemplatesResource, c.ns, analysisTemplate), &v1alpha1.AnalysisTemplate{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AnalysisTemplate), err
}

// Delete takes name of the analysisTemplate and deletes it. Returns an error if one occurs.
func (c *FakeAnalysisTemplates) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(analysistemplatesResource, c.ns, name), &v1alpha1.AnalysisTemplate{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAnalysisTemplates) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(analysistemplatesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.AnalysisTemplateList{})
	return err
}

// Patch applies the patch and returns the patched analysisTemplate.
func (c *FakeAnalysisTemplates) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.AnalysisTemplate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(analysistemplatesResource, c.ns, name, pt, data, subresources...), &v1alpha1.AnalysisTemplate{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AnalysisTemplate), err
}
