package main

import (
    "runtime"

    "github.com/wxc/cmdb/common"
    "github.com/wxc/cmdb/common/blog"
    "github.com/wxc/cmdb/common/types"
)

func main() {
    common.SetIdentification(types.CC_MODULE_APISERVER)
    runtime.GOMAXPROCS(runtime.NumCPU())

    blog.InitLogs()
    defer blog.CloseLogs()
}
