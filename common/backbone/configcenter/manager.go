package configcenter

import (
    "github.com/wxc/cmdb/common"
    crd "github.com/wxc/cmdb/common/confregdiscover"
)

type ConfigCenter struct {
    Type       string
    ConfigCenterDetail crd.ConfRegDiscvIf
}

var (
    configCenterGroup []*ConfigCenter
    configCenterType = common.BKDefaultConfigCenter
)

func SetConfigCenterType(serverType string){
    configCenterType = serverType
}

func AddConfigCenter(configCenter *ConfigCenter) {
    configCenterGroup = append(configCenterGroup, configCenter)
}

func CurrentConfigCenter() crd.ConfRegDiscvIf {
    var defaultConfigCenter *ConfigCenter
    for _, center := range configCenterGroup {
        if center.Type == configCenterType {
            return center.ConfigCenterDetail
        }
        if common.BKDefaultConfigCenter == center.Type {
            defaultConfigCenter = center
        }
    }
    if nil != defaultConfigCenter {
        return defaultConfigCenter.ConfigCenterDetail
    }
    return nil
}
