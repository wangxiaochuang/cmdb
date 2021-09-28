package backbone

import (
    "context"
    "fmt"
    "sync"

    "github.com/wxc/cmdb/apimachinery"
    "github.com/wxc/cmdb/apimachinery/discovery"
    "github.com/wxc/cmdb/apimachinery/util"
    cc "github.com/wxc/cmdb/common/backbone/configcenter"
    "github.com/wxc/cmdb/common/backbone/service_mange/zk"
    "github.com/wxc/cmdb/common/errors"
    "github.com/wxc/cmdb/common/language"
    "github.com/wxc/cmdb/common/metrics"
    "github.com/wxc/cmdb/common/metrics"
    "github.com/wxc/cmdb/common/types"

    "github.com/rs/xid"
)

const maxRetry = 200

type BackboneParameter struct {
    // ConfigUpdate handle process config change
    ConfigUpdate cc.ProcHandlerFunc
    ExtraUpdate  cc.ProcHandlerFunc

    // service component addr
    Regdiscv string
    // config path
    ConfigPath string
    // http server parameter
    SrvInfo *types.ServerInfo
}

func validateParameter(input *BackboneParameter) error {
    if input.Regdiscv == "" {
        return fmt.Errorf("regdiscv can not be emtpy")
    }
    if input.SrvInfo.IP == "" {
        return fmt.Errorf("addrport ip can not be emtpy")
    }
    if input.SrvInfo.Port <= 0 || input.SrvInfo.Port > 65535 {
        return fmt.Errorf("addrport port must be 1-65535")
    }
    if input.ConfigUpdate == nil && input.ExtraUpdate == nil {
        return fmt.Errorf("service config change funcation can not beemtpy")
    }
    // to prevent other components which doesn't set it from failing
    if input.SrvInfo.RegisterIP == "" {
        input.SrvInfo.RegisterIP = input.SrvInfo.IP
    }
    if input.SrvInfo.UUID == "" {
        input.SrvInfo.UUID = xid.New().String()
    }
    return nil
}

func NewBackbone(ctx context.Context, input *BackboneParameter) (*Engine, error) {
    if err := validateParameter(input); err != nil {
        return nil, err
    }

    metricService := metrics.NewService(metrics.Config{ProcessName: common.GetIdentification(), ProcessInstance: input.SrvInfo.Instance()})
    return nil, nil
}

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


