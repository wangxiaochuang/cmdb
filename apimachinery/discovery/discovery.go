package discovery

import (
    "fmt"

    "github.com/wxc/cmdb/common"
    "github.com/wxc/cmdb/common/backbone/service_mange/zk"
    "github.com/wxc/cmdb/common/blog"
    "github.com/wxc/cmdb/common/registerdiscover"
    "github.com/wxc/cmdb/common/types"
)

type ServiceManageInterface interface {
    // 判断当前进程是否为master 进程， 服务注册节点的第一个节点
    IsMaster() bool
}

type DiscoveryInterface interface {
    ApiServer() Interface
    MigrateServer() Interface
    EventServer() Interface
    HostServer() Interface
    ProcServer() Interface
    TopoServer() Interface
    DataCollect() Interface
    GseProcServer() Interface
    CoreService() Interface
    OperationServer() Interface
    TaskServer() Interface
    CloudServer() Interface
    AuthServer() Interface
    Server(name string) Interface
    CacheService() Interface
    ServiceManageInterface
}

type Interface interface {
        // 获取注册在zk上的所有服务节点
        GetServers() ([]string, error)
        // 最新的服务节点信息存放在该channel里，可被用来>消费，以监听服务节点的变化
        GetServersChan() chan []string
}

func NewServiceDiscovery(client *zk.ZkClient) (DiscoveryInterface, error) {
    disc := registerdiscover.NewRegDiscoverEx(client)

    d := &discover{
        servers: make(map[string]*server),
    }

    curServiceName := common.GetIdentification()
    services := types.GetDiscoveryService()
    services[curServiceName] = struct{}{}
    for component := range services {
        if component == types.CC_MODULE_WEBSERVER && curServiceName != types.CC_MODULE_WEBSERVER {
            continue
        }
        // /cc/services/endpoints/*
        path := fmt.Sprintf("%s/%s", types.CC_SERV_BASEPATH, component)
        // 后台自动获取监控配置更新
        svr, err := newServerDiscover(disc, path, component)
        if err != nil {
            return nil, fmt.Errorf("discover %s failed, err: %v", component, err)
        }

        d.servers[component] = svr
    }

    return d, nil
}

type discover struct {
    servers map[string]*server
}

func (d *discover) ApiServer() Interface {
    return d.servers[types.CC_MODULE_APISERVER]
}

func (d *discover) MigrateServer() Interface {
    return d.servers[types.CC_MODULE_MIGRATE]
}

func (d *discover) EventServer() Interface {
    return d.servers[types.CC_MODULE_EVENTSERVER]
}

func (d *discover) HostServer() Interface {
    return d.servers[types.CC_MODULE_HOST]
}

func (d *discover) ProcServer() Interface {
    return d.servers[types.CC_MODULE_PROC]
}

func (d *discover) TopoServer() Interface {
    return d.servers[types.CC_MODULE_TOPO]
}

func (d *discover) DataCollect() Interface {
    return d.servers[types.CC_MODULE_DATACOLLECTION]
}

func (d *discover) GseProcServer() Interface {
    return d.servers[types.GSE_MODULE_PROCSERVER]
}

func (d *discover) CoreService() Interface {
    return d.servers[types.CC_MODULE_CORESERVICE]
}

func (d *discover) OperationServer() Interface {
    return d.servers[types.CC_MODULE_OPERATION]
}

func (d *discover) TaskServer() Interface {
    return d.servers[types.CC_MODULE_TASK]
}

func (d *discover) CloudServer() Interface {
    return d.servers[types.CC_MODULE_CLOUD]
}

func (d *discover) AuthServer() Interface {
    return d.servers[types.CC_MODULE_AUTH]
}

func (d *discover) CacheService() Interface {
    return d.servers[types.CC_MODULE_CACHESERVICE]
}

func (d *discover) IsMaster() bool {
    return d.servers[common.GetIdentification()].IsMaster(common.GetServerInfo().UUID)
}

func (d *discover) Server(name string) Interface {
    if svr, ok := d.servers[name]; ok {
        return svr
    }
    blog.V(5).Infof("not found server. name: %s", name)

    return emptyServerInst
}
