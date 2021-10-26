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

// LoadBalancerRulesGetter has a method to return a LoadBalancerRuleInterface.
// A group's client should implement this interface.
type LoadBalancerRulesGetter interface {
	LoadBalancerRules(namespace string) LoadBalancerRuleInterface
}

// LoadBalancerRuleInterface has methods to work with LoadBalancerRule resources.
type LoadBalancerRuleInterface interface {
	Create(ctx context.Context, loadBalancerRule *v1.LoadBalancerRule, opts metav1.CreateOptions) (*v1.LoadBalancerRule, error)
	Update(ctx context.Context, loadBalancerRule *v1.LoadBalancerRule, opts metav1.UpdateOptions) (*v1.LoadBalancerRule, error)
	UpdateStatus(ctx context.Context, loadBalancerRule *v1.LoadBalancerRule, opts metav1.UpdateOptions) (*v1.LoadBalancerRule, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.LoadBalancerRule, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.LoadBalancerRuleList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.LoadBalancerRule, err error)
	LoadBalancerRuleExpansion
}

// loadBalancerRules implements LoadBalancerRuleInterface
type loadBalancerRules struct {
	client rest.Interface
	ns     string
}

// newLoadBalancerRules returns a LoadBalancerRules
func newLoadBalancerRules(c *TmaxV1Client, namespace string) *loadBalancerRules {
	return &loadBalancerRules{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the loadBalancerRule, and returns the corresponding loadBalancerRule object, and an error if there is any.
func (c *loadBalancerRules) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.LoadBalancerRule, err error) {
	result = &v1.LoadBalancerRule{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("loadbalancerrules").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of LoadBalancerRules that match those selectors.
func (c *loadBalancerRules) List(ctx context.Context, opts metav1.ListOptions) (result *v1.LoadBalancerRuleList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.LoadBalancerRuleList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("loadbalancerrules").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested loadBalancerRules.
func (c *loadBalancerRules) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("loadbalancerrules").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a loadBalancerRule and creates it.  Returns the server's representation of the loadBalancerRule, and an error, if there is any.
func (c *loadBalancerRules) Create(ctx context.Context, loadBalancerRule *v1.LoadBalancerRule, opts metav1.CreateOptions) (result *v1.LoadBalancerRule, err error) {
	result = &v1.LoadBalancerRule{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("loadbalancerrules").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(loadBalancerRule).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a loadBalancerRule and updates it. Returns the server's representation of the loadBalancerRule, and an error, if there is any.
func (c *loadBalancerRules) Update(ctx context.Context, loadBalancerRule *v1.LoadBalancerRule, opts metav1.UpdateOptions) (result *v1.LoadBalancerRule, err error) {
	result = &v1.LoadBalancerRule{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("loadbalancerrules").
		Name(loadBalancerRule.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(loadBalancerRule).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *loadBalancerRules) UpdateStatus(ctx context.Context, loadBalancerRule *v1.LoadBalancerRule, opts metav1.UpdateOptions) (result *v1.LoadBalancerRule, err error) {
	result = &v1.LoadBalancerRule{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("loadbalancerrules").
		Name(loadBalancerRule.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(loadBalancerRule).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the loadBalancerRule and deletes it. Returns an error if one occurs.
func (c *loadBalancerRules) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("loadbalancerrules").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *loadBalancerRules) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("loadbalancerrules").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched loadBalancerRule.
func (c *loadBalancerRules) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.LoadBalancerRule, err error) {
	result = &v1.LoadBalancerRule{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("loadbalancerrules").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
