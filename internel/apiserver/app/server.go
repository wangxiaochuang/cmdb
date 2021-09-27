package app

import (
    "context"
    "fmt"

    "github.com/wxc/cmdb/apimachinery/util"
    "github.com/wxc/cmdb/internel/apiserver/app/options"
    "github.com/wxc/cmdb/internel/apiserver/service"
    "github.com/wxc/cmdb/common/types"
)

func Run(ctx context.Context, cancel context.CancelFunc, op *options.ServerOption) error {
    svrInfo, err := types.NewServerInfo(op.ServConf)
    if err != nil {
            return fmt.Errorf("wrap server info failed, err: %v", err)
    }

    client, err := util.NewClient(&util.TLSClientConfig{})
    if err != nil {
        return fmt.Errorf("new proxy client failed, err: %v", err)
    }

    svc := service.NewService()
}
