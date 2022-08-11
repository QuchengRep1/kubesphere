/*
Copyright 2020 The KubeSphere Authors.

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

package v2beta1

import (
	"context"
	"time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
	v2beta1 "kubesphere.io/api/alerting/v2beta1"
	scheme "kubesphere.io/kubesphere/pkg/client/clientset/versioned/scheme"
)

// GlobalRuleGroupsGetter has a method to return a GlobalRuleGroupInterface.
// A group's client should implement this interface.
type GlobalRuleGroupsGetter interface {
	GlobalRuleGroups(namespace string) GlobalRuleGroupInterface
}

// GlobalRuleGroupInterface has methods to work with GlobalRuleGroup resources.
type GlobalRuleGroupInterface interface {
	Create(ctx context.Context, globalRuleGroup *v2beta1.GlobalRuleGroup, opts v1.CreateOptions) (*v2beta1.GlobalRuleGroup, error)
	Update(ctx context.Context, globalRuleGroup *v2beta1.GlobalRuleGroup, opts v1.UpdateOptions) (*v2beta1.GlobalRuleGroup, error)
	UpdateStatus(ctx context.Context, globalRuleGroup *v2beta1.GlobalRuleGroup, opts v1.UpdateOptions) (*v2beta1.GlobalRuleGroup, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v2beta1.GlobalRuleGroup, error)
	List(ctx context.Context, opts v1.ListOptions) (*v2beta1.GlobalRuleGroupList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v2beta1.GlobalRuleGroup, err error)
	GlobalRuleGroupExpansion
}

// globalRuleGroups implements GlobalRuleGroupInterface
type globalRuleGroups struct {
	client rest.Interface
	ns     string
}

// newGlobalRuleGroups returns a GlobalRuleGroups
func newGlobalRuleGroups(c *AlertingV2beta1Client, namespace string) *globalRuleGroups {
	return &globalRuleGroups{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the globalRuleGroup, and returns the corresponding globalRuleGroup object, and an error if there is any.
func (c *globalRuleGroups) Get(ctx context.Context, name string, options v1.GetOptions) (result *v2beta1.GlobalRuleGroup, err error) {
	result = &v2beta1.GlobalRuleGroup{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("globalrulegroups").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of GlobalRuleGroups that match those selectors.
func (c *globalRuleGroups) List(ctx context.Context, opts v1.ListOptions) (result *v2beta1.GlobalRuleGroupList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v2beta1.GlobalRuleGroupList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("globalrulegroups").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested globalRuleGroups.
func (c *globalRuleGroups) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("globalrulegroups").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a globalRuleGroup and creates it.  Returns the server's representation of the globalRuleGroup, and an error, if there is any.
func (c *globalRuleGroups) Create(ctx context.Context, globalRuleGroup *v2beta1.GlobalRuleGroup, opts v1.CreateOptions) (result *v2beta1.GlobalRuleGroup, err error) {
	result = &v2beta1.GlobalRuleGroup{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("globalrulegroups").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(globalRuleGroup).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a globalRuleGroup and updates it. Returns the server's representation of the globalRuleGroup, and an error, if there is any.
func (c *globalRuleGroups) Update(ctx context.Context, globalRuleGroup *v2beta1.GlobalRuleGroup, opts v1.UpdateOptions) (result *v2beta1.GlobalRuleGroup, err error) {
	result = &v2beta1.GlobalRuleGroup{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("globalrulegroups").
		Name(globalRuleGroup.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(globalRuleGroup).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *globalRuleGroups) UpdateStatus(ctx context.Context, globalRuleGroup *v2beta1.GlobalRuleGroup, opts v1.UpdateOptions) (result *v2beta1.GlobalRuleGroup, err error) {
	result = &v2beta1.GlobalRuleGroup{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("globalrulegroups").
		Name(globalRuleGroup.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(globalRuleGroup).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the globalRuleGroup and deletes it. Returns an error if one occurs.
func (c *globalRuleGroups) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("globalrulegroups").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *globalRuleGroups) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("globalrulegroups").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched globalRuleGroup.
func (c *globalRuleGroups) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v2beta1.GlobalRuleGroup, err error) {
	result = &v2beta1.GlobalRuleGroup{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("globalrulegroups").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
