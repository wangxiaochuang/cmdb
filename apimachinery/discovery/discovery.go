package discovery

type ServiceManageInterface interface {
    // 判断当前进程是否为master 进程， 服务注册节点的第一个节点
    IsMaster() bool
}

type DiscoveryInterface interface {
    ApiServer() Interface
}

type Interface interface {
        // 获取注册在zk上的所有服务节点
        GetServers() ([]string, error)
        // 最新的服务节点信息存放在该channel里，可被用来>消费，以监听服务节点的变化
        GetServersChan() chan []string
}
