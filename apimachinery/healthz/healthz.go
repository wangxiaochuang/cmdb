package healthz

import (
    "fmt"
    "github.com/wxc/cmdb/apimachinery/discovery"
    "github.com/wxc/cmdb/apimachinery/util"
)

type HealthzInterface interface {
    HealthCheck(moduleName string) (healthy bool, err error)
}

func NewHealthzClient(capability *util.Capability, disc discovery.DiscoveryInterface) HealthzInterface {
    return &health{
        capability: capability,
        disc:       disc,
    }
}

type health struct {
    capability *util.Capability
    disc       discovery.DiscoveryInterface
}

func (h *health) HealthCheck(moduleName string) (healthy bool, err error) {
    switch moduleName {
        default:
            panic("in HealthCheck")
            return false, fmt.Errorf("unsupported health module: %s", moduleName)
    }
}
