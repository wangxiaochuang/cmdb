package service

import (
    "github.com/wxc/cmdb/apimachinery"
    "github.com/wxc/cmdb/apimachinery/discovery"
    "github.com/wxc/cmdb/common/backbone"
    "github.com/emicklei/go-restful"
)

type Service interface {
    WebServices() []*restful.WebService
    SetConfig(engine *backbone.Engine, httpClient HTTPClient, discovery discovery.DiscoveryInterface,
            clientSet apimachinery.ClientSetInterface, cache redis.Client, limiter *Limiter)
}

func NewService() Service {
    return new(service)
}

type service struct {
    engine     *backbone.Engine
    client     HTTPClient
    discovery  discovery.DiscoveryInterface
    clientSet  apimachinery.ClientSetInterface
    authorizer ac.AuthorizeInterface
    cache      redis.Client
    limiter    *Limiter
}

func (s *service) SetConfig(engine *backbone.Engine, httpClient HTTPClient, discovery discovery.DiscoveryInterface,
        clientSet apimachinery.ClientSetInterface, cache redis.Client, limiter *Limiter) {
    s.engine = engine
    s.client = httpClient
    s.discovery = discovery
    s.clientSet = clientSet
    s.cache = cache
    s.limiter = limiter
    s.authorizer = iam.NewAuthorizer(clientSet)
}

func (s *service) WebServices() []*restful.WebService {
    allWebServices := make([]*restful.WebService, 0)
    return allWebServices
}
