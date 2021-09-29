package registerdiscover

import (
    "github.com/wxc/cmdb/common/backbone/service_mange/zk"
)

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
