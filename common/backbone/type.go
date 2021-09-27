package backbone

import (
    "net/http"

    "github.com/wxc/cmdb/apimachinery"
    "github.com/wxc/cmdb/common/types"
)

type Config struct {
    RegisterPath string
    RegisterInfo types.ServerInfo
    CoreAPI      apimachinery.ClientSetInterface
}

type Server struct {
    ListenAddr   string
    ListenPort   uint
    Handler      http.Handler
    TLS          TLSConfig
    PProfEnabled bool
}

type TLSConfig struct {
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
