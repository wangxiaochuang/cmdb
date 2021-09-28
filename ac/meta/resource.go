package meta

type ResourceType string

func (r ResourceType) String() string {
    return string(r)
}

const (
    Business                 ResourceType = "business"
    Model                    ResourceType = "model"
    ModelModule              ResourceType = "modelModule"
    ModelSet                 ResourceType = "modelSet"
    MainlineModel            ResourceType = "mainlineObject"
    MainlineModelTopology    ResourceType = "mainlineObjectTopology"
    MainlineInstanceTopology ResourceType = "mainlineInstanceTopology"
    MainlineInstance         ResourceType = "mainlineInstance"
    AssociationType          ResourceType = "associationType"
    ModelAssociation         ResourceType = "modelAssociation"
    ModelInstanceAssociation ResourceType = "modelInstanceAssociation"
    ModelInstance            ResourceType = "modelInstance"
    ModelInstanceTopology    ResourceType = "modelInstanceTopology"
    ModelTopology            ResourceType = "modelTopology"
    ModelClassification      ResourceType = "modelClassification"
    ModelAttributeGroup      ResourceType = "modelAttributeGroup"
    ModelAttribute           ResourceType = "modelAttribute"
    ModelUnique              ResourceType = "modelUnique"
    HostFavorite             ResourceType = "hostFavorite"
    Process                  ResourceType = "process"
    ProcessServiceCategory   ResourceType = "processServiceCategory"
    ProcessServiceTemplate   ResourceType = "processServiceTemplate"
    ProcessTemplate          ResourceType = "processTemplate"
    ProcessServiceInstance   ResourceType = "processServiceInstance"
    BizTopology              ResourceType = "bizTopology"
    HostInstance             ResourceType = "hostInstance"
    NetDataCollector         ResourceType = "netDataCollector"
    DynamicGrouping          ResourceType = "dynamicGrouping" // 动态>分组
    EventWatch               ResourceType = "eventWatch"
    CloudAreaInstance        ResourceType = "plat"
    AuditLog                 ResourceType = "auditlog"   // 操作审计
    UserCustom               ResourceType = "usercustom" // 用户自定义
    SystemBase               ResourceType = "systemBase"
    InstallBK                ResourceType = "installBK"
    SystemConfig             ResourceType = "systemConfig"
    SetTemplate              ResourceType = "setTemplate"
    OperationStatistic       ResourceType = "operationStatistic" // 运
营统计
    HostApply                ResourceType = "hostApply"
    ResourcePoolDirectory    ResourceType = "resourcePoolDirectory"
    CloudAccount             ResourceType = "cloudAccount"
    CloudResourceTask        ResourceType = "cloudResourceTask"
    ConfigAdmin              ResourceType = "configAdmin"
)

const (
    NetCollector = "netCollector"
    NetDevice    = "netDevice"
    NetProperty  = "netProperty"
    NetReport    = "netReport"
)

type ResourceDescribe struct {
    Type    ResourceType
    Actions []Action
}
