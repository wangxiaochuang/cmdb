package zk

import (
    "context"
    "fmt"
    "strings"
    "time"

    "github.com/wxc/cmdb/common/zkclient"
)

// ZkClient do service register and discover by zookeeper
type ZkClient struct {
    zkCli          *zkclient.ZkClient
    cancel         context.CancelFunc
    rootCxt        context.Context
    sessionTimeOut time.Duration
}

func NewZkClient(zkAddress string, timeOut time.Duration) *ZkClient {
    zkAddresses := strings.Split(zkAddress, ",")
    return &ZkClient{
        zkCli:          zkclient.NewZkClient(zkAddresses),
        sessionTimeOut: timeOut,
    }
}

func (zk *ZkClient) Ping() error {
    return zk.zkCli.Ping()
}

// Start used to run register and discover server
func (zk *ZkClient) Start() error {
    // connect zookeeper
    if err := zk.zkCli.ConnectEx(zk.sessionTimeOut); err != nil {
        return fmt.Errorf("fail to connect zookeeper, err: %+v", err)
    }

    // create root context
    zk.rootCxt, zk.cancel = context.WithCancel(context.Background())

    return nil
}

func (zk *ZkClient) Stop() error {
    // close the connection of zookeeper
    zk.zkCli.Close()

    zk.cancel()

    return nil
}

// Client return zk client
func (zk *ZkClient) Client() *zkclient.ZkClient {
    return zk.zkCli
}

// SessionTimeOut client session time out
func (zk *ZkClient) SessionTimeOut() time.Duration {
    return zk.sessionTimeOut
}

// WithCancel context with cancel
func (zk *ZkClient) WithCancel() (context.Context, context.CancelFunc) {
    return context.WithCancel(zk.rootCxt)
}
