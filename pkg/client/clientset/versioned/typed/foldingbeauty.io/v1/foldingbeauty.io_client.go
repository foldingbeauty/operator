// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/foldingbeauty/operator/pkg/apis/foldingbeauty.io/v1"
	"github.com/foldingbeauty/operator/pkg/client/clientset/versioned/scheme"
	serializer "k8s.io/apimachinery/pkg/runtime/serializer"
	rest "k8s.io/client-go/rest"
)

type FoldingbeautyV1Interface interface {
	RESTClient() rest.Interface
	MinersGetter
}

// FoldingbeautyV1Client is used to interact with features provided by the foldingbeauty.io group.
type FoldingbeautyV1Client struct {
	restClient rest.Interface
}

func (c *FoldingbeautyV1Client) Miners(namespace string) MinerInterface {
	return newMiners(c, namespace)
}

// NewForConfig creates a new FoldingbeautyV1Client for the given config.
func NewForConfig(c *rest.Config) (*FoldingbeautyV1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}
	return &FoldingbeautyV1Client{client}, nil
}

// NewForConfigOrDie creates a new FoldingbeautyV1Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *FoldingbeautyV1Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new FoldingbeautyV1Client for the given RESTClient.
func New(c rest.Interface) *FoldingbeautyV1Client {
	return &FoldingbeautyV1Client{c}
}

func setConfigDefaults(config *rest.Config) error {
	gv := v1.SchemeGroupVersion
	config.GroupVersion = &gv
	config.APIPath = "/apis"
	config.NegotiatedSerializer = serializer.DirectCodecFactory{CodecFactory: scheme.Codecs}

	if config.UserAgent == "" {
		config.UserAgent = rest.DefaultKubernetesUserAgent()
	}

	return nil
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FoldingbeautyV1Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}
