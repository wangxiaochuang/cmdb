package apimachinery

import (
    "github.com/wxc/cmdb/apimachinery/healthz"
    "github.com/wxc/cmdb/apimachinery/authserver"
)

type ClientSetInterface interface {
    AuthServer() authserver.AuthServerClientInterface
    Healthz() healthz.HealthzInterface
}
