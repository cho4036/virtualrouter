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

package v1

import (
	"context"
	"time"

	v1 "github.com/tmax-cloud/virtualrouter/pkg/apis/networkcontroller/v1"
	scheme "github.com/tmax-cloud/virtualrouter/pkg/client/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// FireWallRulesGetter has a method to return a FireWallRuleInterface.
// A group's client should implement this interface.
type FireWallRulesGetter interface {
	FireWallRules(namespace string) FireWallRuleInterface
}

// FireWallRuleInterface has methods to work with FireWallRule resources.
type FireWallRuleInterface interface {
	Create(ctx context.Context, fireWallRule *v1.FireWallRule, opts metav1.CreateOptions) (*v1.FireWallRule, error)
	Update(ctx context.Context, fireWallRule *v1.FireWallRule, opts metav1.UpdateOptions) (*v1.FireWallRule, error)
	UpdateStatus(ctx context.Context, fireWallRule *v1.FireWallRule, opts metav1.UpdateOptions) (*v1.FireWallRule, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.FireWallRule, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.FireWallRuleList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.FireWallRule, err error)
	FireWallRuleExpansion
}

// fireWallRules implements FireWallRuleInterface
type fireWallRules struct {
	client rest.Interface
	ns     string
}

// newFireWallRules returns a FireWallRules
func newFireWallRules(c *TmaxV1Client, namespace string) *fireWallRules {
	return &fireWallRules{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the fireWallRule, and returns the corresponding fireWallRule object, and an error if there is any.
func (c *fireWallRules) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.FireWallRule, err error) {
	result = &v1.FireWallRule{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("firewallrules").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of FireWallRules that match those selectors.
func (c *fireWallRules) List(ctx context.Context, opts metav1.ListOptions) (result *v1.FireWallRuleList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.FireWallRuleList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("firewallrules").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested fireWallRules.
func (c *fireWallRules) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("firewallrules").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a fireWallRule and creates it.  Returns the server's representation of the fireWallRule, and an error, if there is any.
func (c *fireWallRules) Create(ctx context.Context, fireWallRule *v1.FireWallRule, opts metav1.CreateOptions) (result *v1.FireWallRule, err error) {
	result = &v1.FireWallRule{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("firewallrules").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(fireWallRule).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a fireWallRule and updates it. Returns the server's representation of the fireWallRule, and an error, if there is any.
func (c *fireWallRules) Update(ctx context.Context, fireWallRule *v1.FireWallRule, opts metav1.UpdateOptions) (result *v1.FireWallRule, err error) {
	result = &v1.FireWallRule{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("firewallrules").
		Name(fireWallRule.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(fireWallRule).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *fireWallRules) UpdateStatus(ctx context.Context, fireWallRule *v1.FireWallRule, opts metav1.UpdateOptions) (result *v1.FireWallRule, err error) {
	result = &v1.FireWallRule{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("firewallrules").
		Name(fireWallRule.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(fireWallRule).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the fireWallRule and deletes it. Returns an error if one occurs.
func (c *fireWallRules) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("firewallrules").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *fireWallRules) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("firewallrules").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched fireWallRule.
func (c *fireWallRules) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.FireWallRule, err error) {
	result = &v1.FireWallRule{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("firewallrules").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
