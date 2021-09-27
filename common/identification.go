package common

import (
    "github.com/wxc/cmdb/common/version"
)

var identification string = "unknown"

func SetIdentification(id string) {
    if identification == "unknown" {
        version.ServiceName = id
        identification = id
    }
}

func GetIdentification() string {
    return identification
}
