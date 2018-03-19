// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/foldingbeauty/operator/pkg/apis/foldingbeauty.io/v1"
	scheme "github.com/foldingbeauty/operator/pkg/client/clientset/versioned/scheme"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// MinersGetter has a method to return a MinerInterface.
// A group's client should implement this interface.
type MinersGetter interface {
	Miners(namespace string) MinerInterface
}

// MinerInterface has methods to work with Miner resources.
type MinerInterface interface {
	Create(*v1.Miner) (*v1.Miner, error)
	Update(*v1.Miner) (*v1.Miner, error)
	Delete(name string, options *meta_v1.DeleteOptions) error
	DeleteCollection(options *meta_v1.DeleteOptions, listOptions meta_v1.ListOptions) error
	Get(name string, options meta_v1.GetOptions) (*v1.Miner, error)
	List(opts meta_v1.ListOptions) (*v1.MinerList, error)
	Watch(opts meta_v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.Miner, err error)
	MinerExpansion
}

// miners implements MinerInterface
type miners struct {
	client rest.Interface
	ns     string
}

// newMiners returns a Miners
func newMiners(c *FoldingbeautyV1Client, namespace string) *miners {
	return &miners{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the miner, and returns the corresponding miner object, and an error if there is any.
func (c *miners) Get(name string, options meta_v1.GetOptions) (result *v1.Miner, err error) {
	result = &v1.Miner{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("miners").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Miners that match those selectors.
func (c *miners) List(opts meta_v1.ListOptions) (result *v1.MinerList, err error) {
	result = &v1.MinerList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("miners").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested miners.
func (c *miners) Watch(opts meta_v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("miners").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a miner and creates it.  Returns the server's representation of the miner, and an error, if there is any.
func (c *miners) Create(miner *v1.Miner) (result *v1.Miner, err error) {
	result = &v1.Miner{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("miners").
		Body(miner).
		Do().
		Into(result)
	return
}

// Update takes the representation of a miner and updates it. Returns the server's representation of the miner, and an error, if there is any.
func (c *miners) Update(miner *v1.Miner) (result *v1.Miner, err error) {
	result = &v1.Miner{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("miners").
		Name(miner.Name).
		Body(miner).
		Do().
		Into(result)
	return
}

// Delete takes name of the miner and deletes it. Returns an error if one occurs.
func (c *miners) Delete(name string, options *meta_v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("miners").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *miners) DeleteCollection(options *meta_v1.DeleteOptions, listOptions meta_v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("miners").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched miner.
func (c *miners) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.Miner, err error) {
	result = &v1.Miner{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("miners").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
