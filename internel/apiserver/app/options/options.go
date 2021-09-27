package options

import (
    "github.com/wxc/cmdb/common/auth"
    "github.com/wxc/cmdb/common/core/cc/config"

    "github.com/spf13/pflag"
)

type ServerOption struct {
    ServConf *config.CCAPIConfig
}

func NewServerOption() *ServerOption {
    s := ServerOption{
        ServConf: config.NewCCAPIConfig(),
    }
    return &s
}

func (s *ServerOption) AddFlags(fs *pflag.FlagSet) {
    fs.StringVar(&s.ServConf.AddrPort, "addrport", "127.0.0.1:50001", "The ip address and port for the serve on")
    fs.StringVar(&s.ServConf.RegDiscover, "regdiscv", "", "hosts of register and discover server. e.g: 127.0.0.1:2181")
    fs.StringVar(&s.ServConf.ExConfig, "config", "", "The config path. e.g ")
    fs.StringVar(&s.ServConf.RegisterIP, "register-ip", "", "the ip address registered on zookeeper, it can be domain")
    fs.Var(auth.EnableAuthFlag, "enable-auth", "The auth center enable status, true for enabled, false for disabled")
}
