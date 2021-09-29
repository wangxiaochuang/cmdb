package common

import (
	"github.com/wxc/cmdb/common/types"
	"github.com/wxc/cmdb/common/version"
)

var identification string = "unknown"
var server *types.ServerInfo

func SetIdentification(id string) {
	if identification == "unknown" {
		version.ServiceName = id
		identification = id
	}
}

func GetIdentification() string {
	return identification
}

func SetServerInfo(srvInfo *types.ServerInfo) {
    server = srvInfo
}
