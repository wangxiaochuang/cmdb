package apimachinery

import (
    "github.com/wxc/cmdb/apimachinery/discovery"
    "github.com/wxc/cmdb/apimachinery/flowctrl"
    "github.com/wxc/cmdb/apimachinery/healthz"
    "github.com/wxc/cmdb/apimachinery/authserver"
    "github.com/wxc/cmdb/apimachinery/util"
)

type ClientSetInterface interface {
    AuthServer() authserver.AuthServerClientInterface
    Healthz() healthz.HealthzInterface
}

func NewApiMachinery(c *util.APIMachineryConfig, discover discovery.DiscoveryInterface) (ClientSetInterface, error) {
    // http.client
    client, err := util.NewClient(c.TLSConfig)
    if err != nil {
        return nil, err
    }

    flowcontrol := flowctrl.NewRateLimiter(c.QPS, c.Burst)
    return NewClientSet(client, discover, flowcontrol), nil
}

func NewClientSet(client util.HttpClient, discover discovery.DiscoveryInterface, throttle flowctrl.RateLimiter) ClientSetInterface {
    return &ClientSet{
        version:  "v3",
        client:   client,
        discover: discover,
        throttle: throttle,
    }
}

type ClientSet struct {
    version string
    client util.HttpClient
    discover discovery.DiscoveryInterface
    throttle flowctrl.RateLimiter
    Mock util.MockInfo
}

func (cs *ClientSet) Healthz() healthz.HealthzInterface {
    c := &util.Capability{
        Client:   cs.client,
        Throttle: cs.throttle,
    }
    return healthz.NewHealthzClient(c, cs.discover)
}

func (cs *ClientSet) AuthServer() authserver.AuthServerClientInterface {
    c := &util.Capability{
        Client:   cs.client,
        Discover: cs.discover.AuthServer(),
        Throttle: cs.throttle,
        Mock:     cs.Mock,
    }
    return authserver.NewAuthServerClientInterface(c, cs.version)
}
