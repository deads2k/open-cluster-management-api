// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1 "github.com/open-cluster-management/api/client/work/clientset/versioned/typed/work/v1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeWorkV1 struct {
	*testing.Fake
}

func (c *FakeWorkV1) ManifestWorks(namespace string) v1.ManifestWorkInterface {
	return &FakeManifestWorks{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeWorkV1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
