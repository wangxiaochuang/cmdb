package app

import (
    "context"
    "fmt"

    "github.com/wxc/cmdb/apimachinery/util"
    "github.com/wxc/cmdb/internel/apiserver/app/options"
    "github.com/wxc/cmdb/internel/apiserver/service"
    "github.com/wxc/cmdb/common/backbone"
    cc "github.com/wxc/cmdb/common/backbone/configcenter"
    "github.com/wxc/cmdb/common/types"
)

func Run(ctx context.Context, cancel context.CancelFunc, op *options.ServerOption) error {
    svrInfo, err := types.NewServerInfo(op.ServConf)
    if err != nil {
            return fmt.Errorf("wrap server info failed, err: %v", err)
    }

    _, err = util.NewClient(&util.TLSClientConfig{})
    if err != nil {
        return fmt.Errorf("new proxy client failed, err: %v", err)
    }

    _ = service.NewService()

    apiSvr := new(APIServer)
    input := &backbone.BackboneParameter{
        ConfigUpdate: apiSvr.onApiServerConfigUpdate,
        ConfigPath:   op.ServConf.ExConfig,
        Regdiscv:     op.ServConf.RegDiscover,
        SrvInfo:      svrInfo,
    }

    engine, err := backbone.NewBackbone(ctx, input)
    if err != nil {
        return fmt.Errorf("new backbone failed, err: %v", err)
    }
    fmt.Printf("engin: %v", engine)

    return nil
}

type APIServer struct {
    Core        *backbone.Engine
    Config      map[string]string
    configReady bool
}

func (h *APIServer) onApiServerConfigUpdate(previous, current cc.ProcessConfig) {
    h.configReady = true
}

const waitForSeconds = 180
