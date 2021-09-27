package apimachinery

import (
    "github.com/wxc/cmdb/apimachinery/healthz"
)

type ClientSetInterface interface {
    Healthz() healthz.HealthzInterface
}
