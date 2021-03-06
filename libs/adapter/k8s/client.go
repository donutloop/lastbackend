package k8s

import (
	"k8s.io/client-go/1.5/rest"
)

// LBExtend is used to interact with features provided by the Last.Backend group.

type LBClientInterface interface {
	GetRESTClient() *rest.RESTClient
}

type LBClient struct {
	*rest.RESTClient
}

func (c *LBClient) GetRESTClient() *rest.RESTClient {
	if c == nil {
		return nil
	}
	return c.RESTClient
}
