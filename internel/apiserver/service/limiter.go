package service

import (
    "sync"
    "time"

    "github.com/wxc/cmdb/common/metadata"
    "github.com/wxc/cmdb/common/zkclient"
)

type Limiter struct {
    zkCli        *zkclient.ZkClient
    rules        map[string]*metadata.LimiterRule
    lock         sync.RWMutex
    syncDuration time.Duration
}
