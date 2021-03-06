package fake

import (
	network_v1 "github.com/openshift/origin/pkg/sdn/apis/network/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeClusterNetworks implements ClusterNetworkInterface
type FakeClusterNetworks struct {
	Fake *FakeNetworkV1
}

var clusternetworksResource = schema.GroupVersionResource{Group: "network.openshift.io", Version: "v1", Resource: "clusternetworks"}

var clusternetworksKind = schema.GroupVersionKind{Group: "network.openshift.io", Version: "v1", Kind: "ClusterNetwork"}

// Get takes name of the clusterNetwork, and returns the corresponding clusterNetwork object, and an error if there is any.
func (c *FakeClusterNetworks) Get(name string, options v1.GetOptions) (result *network_v1.ClusterNetwork, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(clusternetworksResource, name), &network_v1.ClusterNetwork{})
	if obj == nil {
		return nil, err
	}
	return obj.(*network_v1.ClusterNetwork), err
}

// List takes label and field selectors, and returns the list of ClusterNetworks that match those selectors.
func (c *FakeClusterNetworks) List(opts v1.ListOptions) (result *network_v1.ClusterNetworkList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(clusternetworksResource, clusternetworksKind, opts), &network_v1.ClusterNetworkList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &network_v1.ClusterNetworkList{}
	for _, item := range obj.(*network_v1.ClusterNetworkList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested clusterNetworks.
func (c *FakeClusterNetworks) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(clusternetworksResource, opts))
}

// Create takes the representation of a clusterNetwork and creates it.  Returns the server's representation of the clusterNetwork, and an error, if there is any.
func (c *FakeClusterNetworks) Create(clusterNetwork *network_v1.ClusterNetwork) (result *network_v1.ClusterNetwork, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(clusternetworksResource, clusterNetwork), &network_v1.ClusterNetwork{})
	if obj == nil {
		return nil, err
	}
	return obj.(*network_v1.ClusterNetwork), err
}

// Update takes the representation of a clusterNetwork and updates it. Returns the server's representation of the clusterNetwork, and an error, if there is any.
func (c *FakeClusterNetworks) Update(clusterNetwork *network_v1.ClusterNetwork) (result *network_v1.ClusterNetwork, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(clusternetworksResource, clusterNetwork), &network_v1.ClusterNetwork{})
	if obj == nil {
		return nil, err
	}
	return obj.(*network_v1.ClusterNetwork), err
}

// Delete takes name of the clusterNetwork and deletes it. Returns an error if one occurs.
func (c *FakeClusterNetworks) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(clusternetworksResource, name), &network_v1.ClusterNetwork{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeClusterNetworks) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(clusternetworksResource, listOptions)

	_, err := c.Fake.Invokes(action, &network_v1.ClusterNetworkList{})
	return err
}

// Patch applies the patch and returns the patched clusterNetwork.
func (c *FakeClusterNetworks) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *network_v1.ClusterNetwork, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(clusternetworksResource, name, data, subresources...), &network_v1.ClusterNetwork{})
	if obj == nil {
		return nil, err
	}
	return obj.(*network_v1.ClusterNetwork), err
}
