package backbone

import (
    "encoding/json"
    "github.com/wxc/cmdb/common/backbone/service_mange/zk"
    "github.com/wxc/cmdb/common/registerdiscover"
    "github.com/wxc/cmdb/common/types"
    "github.com/wxc/cmdb/framework/core/errors"
)

type ServiceRegisterInterface interface {
    // Ping to ping server
    Ping() error
    // register local server info, it can only be called for once.
    Register(path string, c types.ServerInfo) error
    // Cancel to stop server register and discover
    Cancel()
    // ClearRegisterPath to delete server register path from zk
    ClearRegisterPath() error
}

func NewServiceRegister(client *zk.ZkClient) (ServiceRegisterInterface, error) {
    s := new(serviceRegister)
    s.client = registerdiscover.NewRegDiscoverEx(client)
    return s, nil
}

type serviceRegister struct {
    client *registerdiscover.RegDiscover
}

func (s *serviceRegister) Register(path string, c types.ServerInfo) error {
    if c.RegisterIP == "0.0.0.0" {
        return errors.New("register ip can not be 0.0.0.0")
    }

    js, err := json.Marshal(c)
    if err != nil {
        return err
    }

    return s.client.RegisterAndWatchService(path, js)
}

func (s *serviceRegister) Ping() error {
    return s.client.Ping()
}

// Cancel to stop server register and discover
func (s *serviceRegister) Cancel() {
    s.client.Cancel()
}

// ClearRegisterPath to delete server register path from zk
func (s *serviceRegister) ClearRegisterPath() error {
    return s.client.ClearRegisterPath()
}
