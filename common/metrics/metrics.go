package metrics

import (
    "net/http"
    "strings"

    "github.com/prometheus/client_golang/prometheus"
    "github.com/prometheus/client_golang/prometheus/promhttp"
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

func NewService(conf Config) *Service {
    registry := prometheus.NewRegistry()
    register := prometheus.WrapRegistererWith(prometheus.Labels{LabelProcessName: conf.ProcessName, LabelHost: strings.Split(conf.ProcessInstance, ":")[0]}, registry)

    // set up global register
    globalRegister = register

    srv := Service{conf: conf, registry: register}

    srv.requestTotal = prometheus.NewCounterVec(
        prometheus.CounterOpts{
            Name: Namespace + "_http_request_total",
            Help: "http request total.",
        },
        []string{LabelHandler, LabelHTTPStatus, LabelOrigin, LabelAppCode},
    )
    register.MustRegister(srv.requestTotal)

    srv.requestDuration = prometheus.NewHistogramVec(
        prometheus.HistogramOpts{
            Name:    Namespace + "_http_request_duration_millisecond",
            Help:    "Histogram of latencies for HTTP requests.",
            Buckets: []float64{10, 30, 50, 70, 100, 200, 300, 400, 500, 1000, 2000, 5000},
        },
        []string{LabelHandler, LabelAppCode},
    )
    register.MustRegister(srv.requestDuration)
    register.MustRegister(prometheus.NewProcessCollector(prometheus.ProcessCollectorOpts{}))
    register.MustRegister(prometheus.NewGoCollector())

    srv.httpHandler = promhttp.InstrumentMetricHandler(
        registry, promhttp.HandlerFor(registry, promhttp.HandlerOpts{}),
    )
    return &srv
}
