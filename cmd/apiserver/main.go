package main

import (
    "context"
    "fmt"
    "os"
    "runtime"

    "github.com/wxc/cmdb/internel/apiserver/app"
    "github.com/wxc/cmdb/internel/apiserver/app/options"
    "github.com/wxc/cmdb/common"
    "github.com/wxc/cmdb/common/blog"
    "github.com/wxc/cmdb/common/types"
    "github.com/wxc/cmdb/common/util"

    "github.com/spf13/pflag"
)

func main() {
	common.SetIdentification(types.CC_MODULE_APISERVER)
	runtime.GOMAXPROCS(runtime.NumCPU())

	blog.InitLogs()
	defer blog.CloseLogs()

        op := options.NewServerOption()
        op.AddFlags(pflag.CommandLine)

        util.InitFlags()

        ctx, cancel := context.WithCancel(context.Background())
        if err := app.Run(ctx, cancel, op); err != nil {
            fmt.Fprintf(os.Stderr, "%v\n", err)
            blog.CloseLogs()
            os.Exit(1)
        }
}
