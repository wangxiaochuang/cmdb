package config

import (
    "fmt"
    "strconv"
    "strings"
)

type CCAPIConfig struct {
    AddrPort    string
    RegDiscover string
    RegisterIP  string
    ExConfig    string
    Qps         int64
    Burst       int64
}

func NewCCAPIConfig() *CCAPIConfig {
    return &CCAPIConfig{
        AddrPort:    "127.0.0.1:8081",
        RegDiscover: "",
        RegisterIP:  "",
        Qps:         1000,
        Burst:       2000,
    }
}

func (conf *CCAPIConfig) GetAddress() (string, error) {
    addrPort := strings.TrimSpace(conf.AddrPort)
    if err := checkAddrPort(addrPort); err != nil {
        return "", err
    }
    if isIPV6(addrPort) {
        return getIPV6Adrr(addrPort)
    }
    return getIPV4Adrr(addrPort)
}

// GetPort get the port
func (conf *CCAPIConfig) GetPort() (uint, error) {
    addrPort := strings.TrimSpace(conf.AddrPort)
    if err := checkAddrPort(addrPort); err != nil {
        return uint(0), err
    }
    if isIPV6(addrPort) {
        return getIPV6Port(addrPort)
    }
    return getIPV4Port(addrPort)
}

func checkAddrPort(addrPort string) error {
    if strings.Count(addrPort, ":") == 0 {
            return fmt.Errorf("the value of flag[AddrPort: %s] is wrong", addrPort)
    }
    return nil
}

func isIPV6(addrPort string) bool {
    return strings.Count(addrPort, ":") > 1
}

func getIPV6Adrr(addrPort string) (string, error) {
    idx := strings.LastIndex(addrPort, ":")
    return addrPort[:idx], nil
}

func getIPV4Adrr(addrPort string) (string, error) {
    idx := strings.LastIndex(addrPort, ":")
    return addrPort[:idx], nil
}

func getIPV6Port(addrPort string) (uint, error) {
    return getPortFunc(addrPort)
}

func getIPV4Port(addrPort string) (uint, error) {
        return getPortFunc(addrPort)
}

func getPortFunc(addrPort string) (uint, error) {
    idx := strings.LastIndex(addrPort, ":")
    if len(addrPort[idx:]) < 2 {
        return 0, fmt.Errorf("the value of flag[AddrPort: %s] is wrong", addrPort)
    }
    port, err := strconv.ParseUint(addrPort[idx+1:], 10, 0)
    if err != nil {
        return uint(0), err
    }
    return uint(port), nil
}
