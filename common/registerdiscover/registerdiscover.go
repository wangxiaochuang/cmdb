package registerdiscover

import (
    "github.com/wxc/cmdb/common/backbone/service_mange/zk"
)

type DiscoverEvent struct {
    Err    error
    Key    string
    Server []string
    Nodes  []string
}

type RegDiscover struct {
    rdServer RegDiscvServer
}

func NewRegDiscoverEx(client *zk.ZkClient) *RegDiscover {
    regDiscv := &RegDiscover{
        rdServer: nil,
    }

    regDiscv.rdServer = RegDiscvServer(NewZkRegDiscv(client))

    return regDiscv
}

func (rd *RegDiscover) RegisterAndWatchService(key string, data []byte) error {
    return rd.rdServer.RegisterAndWatch(key, data)
}

func (rd *RegDiscover) GetServNodes(key string) ([]string, error) {
    return rd.rdServer.GetServNodes(key)
}

func (rd *RegDiscover) DiscoverService(key string) (<-chan *DiscoverEvent, error) {
    return rd.rdServer.Discover(key)
}

func (rd *RegDiscover) Ping() error {
    return rd.rdServer.Ping()
}

// Cancel to stop server register and discover
func (rd *RegDiscover) Cancel() {
    rd.rdServer.Cancel()
}

// ClearRegisterPath to delete server register path from zk
func (rd *RegDiscover) ClearRegisterPath() error {
    return rd.rdServer.ClearRegisterPath()
}
