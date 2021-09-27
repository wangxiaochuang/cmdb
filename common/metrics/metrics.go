package metrics

import (
    "net/http"
    "github.com/prometheus/client_golang/prometheus"
)

var globalRegister prometheus.Registerer

func init() {
    globalRegister = prometheus.DefaultRegisterer
}

func Register() prometheus.Registerer {
    return globalRegister
}

const Namespace = "cmdb"

const (
    LabelHandler     = "handler"
    LabelHTTPStatus  = "status_code"
    LabelOrigin      = "origin"
    LabelProcessName = "process_name"
    LabelAppCode     = "app_code"
    LabelHost        = "host"
)

// labels
const (
    KeySelectedRoutePath string = "SelectedRoutePath"
)

// Config metrics config
type Config struct {
    ProcessName     string
    ProcessInstance string
}

type Service struct {
    conf Config

    httpHandler http.Handler

    registry        prometheus.Registerer
    requestTotal    *prometheus.CounterVec
    requestDuration *prometheus.HistogramVec
}
