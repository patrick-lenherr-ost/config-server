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
// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1alpha1 "github.com/sdcio/config-server/apis/inv/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeDiscoveryRules implements DiscoveryRuleInterface
type FakeDiscoveryRules struct {
	Fake *FakeInvV1alpha1
	ns   string
}

var discoveryrulesResource = v1alpha1.SchemeGroupVersion.WithResource("discoveryrules")

var discoveryrulesKind = v1alpha1.SchemeGroupVersion.WithKind("DiscoveryRule")

// Get takes name of the discoveryRule, and returns the corresponding discoveryRule object, and an error if there is any.
func (c *FakeDiscoveryRules) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.DiscoveryRule, err error) {
	emptyResult := &v1alpha1.DiscoveryRule{}
	obj, err := c.Fake.
		Invokes(testing.NewGetActionWithOptions(discoveryrulesResource, c.ns, name, options), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.DiscoveryRule), err
}

// List takes label and field selectors, and returns the list of DiscoveryRules that match those selectors.
func (c *FakeDiscoveryRules) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.DiscoveryRuleList, err error) {
	emptyResult := &v1alpha1.DiscoveryRuleList{}
	obj, err := c.Fake.
		Invokes(testing.NewListActionWithOptions(discoveryrulesResource, discoveryrulesKind, c.ns, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.DiscoveryRuleList{ListMeta: obj.(*v1alpha1.DiscoveryRuleList).ListMeta}
	for _, item := range obj.(*v1alpha1.DiscoveryRuleList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested discoveryRules.
func (c *FakeDiscoveryRules) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchActionWithOptions(discoveryrulesResource, c.ns, opts))

}

// Create takes the representation of a discoveryRule and creates it.  Returns the server's representation of the discoveryRule, and an error, if there is any.
func (c *FakeDiscoveryRules) Create(ctx context.Context, discoveryRule *v1alpha1.DiscoveryRule, opts v1.CreateOptions) (result *v1alpha1.DiscoveryRule, err error) {
	emptyResult := &v1alpha1.DiscoveryRule{}
	obj, err := c.Fake.
		Invokes(testing.NewCreateActionWithOptions(discoveryrulesResource, c.ns, discoveryRule, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.DiscoveryRule), err
}

// Update takes the representation of a discoveryRule and updates it. Returns the server's representation of the discoveryRule, and an error, if there is any.
func (c *FakeDiscoveryRules) Update(ctx context.Context, discoveryRule *v1alpha1.DiscoveryRule, opts v1.UpdateOptions) (result *v1alpha1.DiscoveryRule, err error) {
	emptyResult := &v1alpha1.DiscoveryRule{}
	obj, err := c.Fake.
		Invokes(testing.NewUpdateActionWithOptions(discoveryrulesResource, c.ns, discoveryRule, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.DiscoveryRule), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeDiscoveryRules) UpdateStatus(ctx context.Context, discoveryRule *v1alpha1.DiscoveryRule, opts v1.UpdateOptions) (result *v1alpha1.DiscoveryRule, err error) {
	emptyResult := &v1alpha1.DiscoveryRule{}
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceActionWithOptions(discoveryrulesResource, "status", c.ns, discoveryRule, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.DiscoveryRule), err
}

// Delete takes name of the discoveryRule and deletes it. Returns an error if one occurs.
func (c *FakeDiscoveryRules) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(discoveryrulesResource, c.ns, name, opts), &v1alpha1.DiscoveryRule{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeDiscoveryRules) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionActionWithOptions(discoveryrulesResource, c.ns, opts, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.DiscoveryRuleList{})
	return err
}

// Patch applies the patch and returns the patched discoveryRule.
func (c *FakeDiscoveryRules) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.DiscoveryRule, err error) {
	emptyResult := &v1alpha1.DiscoveryRule{}
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceActionWithOptions(discoveryrulesResource, c.ns, name, pt, data, opts, subresources...), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.DiscoveryRule), err
}
