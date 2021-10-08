package app

import (
    "context"
    "fmt"

    "github.com/wxc/cmdb/apimachinery/util"
    "github.com/wxc/cmdb/internel/apiserver/app/options"
    "github.com/wxc/cmdb/internel/apiserver/service"
    "github.com/wxc/cmdb/common/backbone"
    cc "github.com/wxc/cmdb/common/backbone/configcenter"
    "github.com/wxc/cmdb/common/blog"
    "github.com/wxc/cmdb/common/types"
    "github.com/wxc/cmdb/storage/dal/redis"

    "github.com/emicklei/go-restful"
)

func Run(ctx context.Context, cancel context.CancelFunc, op *options.ServerOption) error {
    svrInfo, err := types.NewServerInfo(op.ServConf)
    if err != nil {
            return fmt.Errorf("wrap server info failed, err: %v", err)
    }
    //{
    //    IP: "127.0.0.1",
    //    Port: 80,
    //    RegisterIP: "127.0.0.1",
    //    HostName: "wxc's ...",
    //    Scheme: "http",
    //    Version: "...",
    //    Pid: xxx,
    //    UUID: "xxx"
    //}

    // http.Client
    client, err := util.NewClient(&util.TLSClientConfig{})
    if err != nil {
        return fmt.Errorf("new proxy client failed, err: %v", err)
    }

    svc := service.NewService()

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

    redisConf, err := engine.WithRedis()
    if err != nil {
        return err
    }
    cache, err := redis.NewFromConfig(redisConf)
    if err != nil {
        return fmt.Errorf("connect redis server failed, err: %s", err.Error())
    }

    limiter := service.NewLimiter(engine.ServiceManageClient().Client())
    err = limiter.SyncLimiterRules()
    if err != nil {
        blog.Infof("SyncLimiterRules failed, err: %v", err)
        return err
    }

    svc.SetConfig(engine, client, engine.Discovery(), engine.CoreAPI, cache, limiter)

    ctnr := restful.NewContainer()
    ctnr.Router(restful.CurlyRouter{})
    for _, item := range svc.WebServices() {
        ctnr.Add(item)
    }
    apiSvr.Core = engine

    err = backbone.StartServer(ctx, cancel, engine, ctnr, false)
    if err != nil {
        return err
    }

    select {
    case <-ctx.Done():
    }
    return nil

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
