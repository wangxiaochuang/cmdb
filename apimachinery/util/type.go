package util

import (
    "fmt"
    "time"
    "github.com/wxc/cmdb/apimachinery/discovery"
    "github.com/wxc/cmdb/apimachinery/flowctrl"
    cc "github.com/wxc/cmdb/common/backbone/configcenter"
    "github.com/prometheus/client_golang/prometheus"
)

type APIMachineryConfig struct {
    // request's qps value
    QPS int64
    // request's burst value
    Burst     int64
    TLSConfig *TLSClientConfig
}

type Capability struct {
    Client     HttpClient
    Discover   discovery.Interface
    Throttle   flowctrl.RateLimiter
    Mock       MockInfo
    MetricOpts MetricOption
    // the max tolerance api request latency time, if exceeded this time, then
    // this request will be logged and warned.
    ToleranceLatencyTime time.Duration
}

type MetricOption struct {
    // prometheus metric register
    Register prometheus.Registerer
    // if not set, use default buckets value
    DurationBuckets []float64
}

type MockInfo struct {
    Mocked      bool
    SetMockData bool
    MockData    interface{}
}

type TLSClientConfig struct {
    // Server should be accessed without verifying the TLS certificate. For testing only.
    InsecureSkipVerify bool
    // Server requires TLS client certificate authentication
    CertFile string
    // Server requires TLS client certificate authentication
    KeyFile string
    // Trusted root certificates for server
    CAFile string
    // the password to decrypt the certificate
    Password string
}

func NewTLSClientConfigFromConfig(prefix string, config map[string]string) (TLSClientConfig, error) {
        tlsConfig := TLSClientConfig{}

        skipVerifyKey := fmt.Sprintf("%s.insecureSkipVerify", prefix)
        if val, err := cc.String(skipVerifyKey); err == nil {
                skipVerifyVal := val
                if skipVerifyVal == "true" {
                        tlsConfig.InsecureSkipVerify = true
                }
        }

        certFileKey := fmt.Sprintf("%s.certFile", prefix)
        if val, err := cc.String(certFileKey); err == nil {
                tlsConfig.CertFile = val
        }

        keyFileKey := fmt.Sprintf("%s.keyFile", prefix)
        if val, err := cc.String(keyFileKey); err == nil {
                tlsConfig.KeyFile = val
        }

        caFileKey := fmt.Sprintf("%s.caFile", prefix)
        if val, err := cc.String(caFileKey); err == nil {
                tlsConfig.CAFile = val
        }

        passwordKey := fmt.Sprintf("%s.password", prefix)
        if val, err := cc.String(passwordKey); err == nil {
                tlsConfig.Password = val
        }

        return tlsConfig, nil
}

