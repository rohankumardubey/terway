/*
Copyright 2021 Terway Authors.

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

package v1beta1

import (
	"context"
	"time"

	v1beta1 "github.com/AliyunContainerService/terway/pkg/apis/network.alibabacloud.com/v1beta1"
	scheme "github.com/AliyunContainerService/terway/pkg/generated/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// PodENIsGetter has a method to return a PodENIInterface.
// A group's client should implement this interface.
type PodENIsGetter interface {
	PodENIs(namespace string) PodENIInterface
}

// PodENIInterface has methods to work with PodENI resources.
type PodENIInterface interface {
	Create(ctx context.Context, podENI *v1beta1.PodENI, opts v1.CreateOptions) (*v1beta1.PodENI, error)
	Update(ctx context.Context, podENI *v1beta1.PodENI, opts v1.UpdateOptions) (*v1beta1.PodENI, error)
	UpdateStatus(ctx context.Context, podENI *v1beta1.PodENI, opts v1.UpdateOptions) (*v1beta1.PodENI, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1beta1.PodENI, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1beta1.PodENIList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.PodENI, err error)
	PodENIExpansion
}

// podENIs implements PodENIInterface
type podENIs struct {
	client rest.Interface
	ns     string
}

// newPodENIs returns a PodENIs
func newPodENIs(c *NetworkV1beta1Client, namespace string) *podENIs {
	return &podENIs{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the podENI, and returns the corresponding podENI object, and an error if there is any.
func (c *podENIs) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.PodENI, err error) {
	result = &v1beta1.PodENI{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("podenis").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of PodENIs that match those selectors.
func (c *podENIs) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.PodENIList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1beta1.PodENIList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("podenis").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested podENIs.
func (c *podENIs) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("podenis").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a podENI and creates it.  Returns the server's representation of the podENI, and an error, if there is any.
func (c *podENIs) Create(ctx context.Context, podENI *v1beta1.PodENI, opts v1.CreateOptions) (result *v1beta1.PodENI, err error) {
	result = &v1beta1.PodENI{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("podenis").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(podENI).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a podENI and updates it. Returns the server's representation of the podENI, and an error, if there is any.
func (c *podENIs) Update(ctx context.Context, podENI *v1beta1.PodENI, opts v1.UpdateOptions) (result *v1beta1.PodENI, err error) {
	result = &v1beta1.PodENI{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("podenis").
		Name(podENI.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(podENI).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *podENIs) UpdateStatus(ctx context.Context, podENI *v1beta1.PodENI, opts v1.UpdateOptions) (result *v1beta1.PodENI, err error) {
	result = &v1beta1.PodENI{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("podenis").
		Name(podENI.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(podENI).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the podENI and deletes it. Returns an error if one occurs.
func (c *podENIs) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("podenis").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *podENIs) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("podenis").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched podENI.
func (c *podENIs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.PodENI, err error) {
	result = &v1beta1.PodENI{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("podenis").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
