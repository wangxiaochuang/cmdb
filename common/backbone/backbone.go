package backbone

import (
    "sync"

    "github.com/wxc/cmdb/apimachinery"
    "github.com/wxc/cmdb/apimachinery/discovery"
    "github.com/wxc/cmdb/apimachinery/util"
    "github.com/wxc/cmdb/common/backbone/service_mange/zk"
    "github.com/wxc/cmdb/common/errors"
    "github.com/wxc/cmdb/common/language"
    "github.com/wxc/cmdb/common/metrics"
    "github.com/wxc/cmdb/common/types"
)

const maxRetry = 200

type Engine struct {
    CoreAPI             apimachinery.ClientSetInterface
    apiMachineryConfig  *util.APIMachineryConfig

    client              *zk.ZkClient
    ServiceManageInterface discovery.ServiceManageInterface
    SvcDisc             ServiceRegisterInterface
    discovery           discovery.DiscoveryInterface
    metric              *metrics.Service

    sync.Mutex

    RegisterPath string
    server       Server
    srvInfo      *types.ServerInfo

    Language language.CCLanguageIf
    CCErr    errors.CCErrorIf
    CCCtx    CCContextInterface
}

func (e *Engine) Discovery() discovery.DiscoveryInterface {
    return e.discovery
}

func (e *Engine) ApiMachineryConfig() *util.APIMachineryConfig {
    return e.apiMachineryConfig
}

func (e *Engine) ServiceManageClient() *zk.ZkClient {
    return e.client
}

func (e *Engine) Metric() *metrics.Service {
    return e.metric
}


