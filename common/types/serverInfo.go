package types

import (
    "fmt"
    "os"

    "github.com/wxc/cmdb/common/core/cc/config"
    "github.com/wxc/cmdb/common/version"

    "github.com/rs/xid"
)

const (
	CC_MODULE_APISERVER = "apiserver"
)

type ServerInfo struct {
    IP         string `json:"ip"`
    Port       uint   `json:"port"`
    RegisterIP string `json:"registerip"`
    HostName   string `json:"hostname"`
    Scheme     string `json:"scheme"`
    Version    string `json:"version"`
    Pid        int    `json:"pid"`
    // UUID is used to distinguish which service is master in zookeeper
    UUID string `json:"uuid"`
}

func NewServerInfo(conf *config.CCAPIConfig) (*ServerInfo, error) {
    ip, err := conf.GetAddress()
    if err != nil {
        return nil, err
    }

    port, err := conf.GetPort()
    if err != nil {
        return nil, err
    }

    registerIP := conf.RegisterIP
    if registerIP == "" {
        registerIP = ip
    }

    hostname, err := os.Hostname()
    if err != nil {
        return nil, err
    }

    info := &ServerInfo{
        IP:         ip,
        Port:       port,
        RegisterIP: registerIP,
        HostName:   hostname,
        Scheme:     "http",
        Version:    version.GetVersion(),
        Pid:        os.Getpid(),
        UUID:       xid.New().String(),
    }
    return info, nil
}

// APIServerServInfo apiserver informaiton
type APIServerServInfo struct {
    ServerInfo
}

// WebServerInfo web server information
type WebServerInfo struct {
    ServerInfo
}

// AuditControllerServInfo audit-controller server information
type AuditControllerServInfo struct {
    ServerInfo
}

// HostControllerServInfo host-controller server information
type HostControllerServInfo struct {
    ServerInfo
}

type MigrateControllerServInfo struct {
    ServerInfo
}

// ObjectControllerServInfo object-controller server information
type ObjectControllerServInfo struct {
    ServerInfo
}

// ProcControllerServInfo proc-controller server information
type ProcControllerServInfo struct {
    ServerInfo
}

// DataCollectionServInfo data-conllection server information
type DataCollectionServInfo struct {
    ServerInfo
}

// HostServerInfo host server information
type HostServerInfo struct {
    ServerInfo
}

// MigrateServInfo migrate server information
type MigrateServInfo struct {
    ServerInfo
}

// ProcServInfo proc server information
type ProcServInfo struct {
    ServerInfo
}

// TopoServInfo topo server information
type TopoServInfo struct {
    ServerInfo
}

// EventServInfo topo server information
type EventServInfo struct {
    ServerInfo
}

// Address convert struct to host address
func (s *ServerInfo) RegisterAddress() string {
    if s == nil {
        return ""
    }
    return fmt.Sprintf("%s://%s:%d", s.Scheme, s.RegisterIP, s.Port)
}

func (s *ServerInfo) Instance() string {
    if s == nil {
        return ""
    }
    return fmt.Sprintf("%s:%d", s.IP, s.Port)
}
