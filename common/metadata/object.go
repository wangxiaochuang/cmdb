package metadata

import (
    "github.com/wxc/cmdb/common"
    "github.com/wxc/cmdb/common/mapstr"
)

const (
    ModelFieldID          = "id"
    ModelFieldObjCls      = "bk_classification_id"
    ModelFieldObjIcon     = "bk_obj_icon"
    ModelFieldObjectID    = "bk_obj_id"
    ModelFieldObjectName  = "bk_obj_name"
    ModelFieldIsHidden    = "bk_ishidden"
    ModelFieldIsPre       = "ispre"
    ModelFieldIsPaused    = "bk_ispaused"
    ModelFieldPosition    = "position"
    ModelFieldOwnerID     = "bk_supplier_account"
    ModelFieldDescription = "description"
    ModelFieldCreator     = "creator"
    ModelFieldModifier    = "modifier"
    ModelFieldCreateTime  = "create_time"
    ModelFieldLastTime    = "last_time"
)

type Object struct {
    ID         int64  `field:"id" json:"id" bson:"id"`
    ObjCls     string `field:"bk_classification_id" json:"bk_classification_id" bson:"bk_classification_id"`
    ObjIcon    string `field:"bk_obj_icon" json:"bk_obj_icon" bson:"bk_obj_icon"`
    ObjectID   string `field:"bk_obj_id" json:"bk_obj_id" bson:"bk_obj_id"`
    ObjectName string `field:"bk_obj_name" json:"bk_obj_name" bson:"bk_obj_name"`

    // IsHidden front-end don't display the object if IsHidden is true
    IsHidden bool `field:"bk_ishidden" json:"bk_ishidden" bson:"bk_ishidden"`

    IsPre       bool   `field:"ispre" json:"ispre" bson:"ispre"`
    IsPaused    bool   `field:"bk_ispaused" json:"bk_ispaused" bson:"bk_ispaused"`
    Position    string `field:"position" json:"position" bson:"position"`
    OwnerID     string `field:"bk_supplier_account" json:"bk_supplier_account" bson:"bk_supplier_account"`
    Description string `field:"description" json:"description" bson:"description"`
    Creator     string `field:"creator" json:"creator" bson:"creator"`
    Modifier    string `field:"modifier" json:"modifier" bson:"modifier"`
    CreateTime  *Time  `field:"create_time" json:"create_time" bson:"create_time"`
    LastTime    *Time  `field:"last_time" json:"last_time" bson:"last_time"`
}

func (o *Object) GetDefaultInstPropertyName() string {
    return common.DefaultInstName
}

// GetInstIDFieldName get instid filed
func (o *Object) GetInstIDFieldName() string {
    return GetInstIDFieldByObjID(o.ObjectID)

}

func GetInstIDFieldByObjID(objID string) string {
    switch objID {
    case common.BKInnerObjIDApp:
        return common.BKAppIDField
    case common.BKInnerObjIDSet:
        return common.BKSetIDField
    case common.BKInnerObjIDModule:
        return common.BKModuleIDField
    case common.BKInnerObjIDObject:
        return common.BKInstIDField
    case common.BKInnerObjIDHost:
        return common.BKHostIDField
    case common.BKInnerObjIDProc:
        return common.BKProcIDField
    case common.BKInnerObjIDPlat:
        return common.BKCloudIDField
    default:
        return common.BKInstIDField
    }

}

func GetInstNameFieldName(objID string) string {
    switch objID {
    case common.BKInnerObjIDApp:
        return common.BKAppNameField
    case common.BKInnerObjIDSet:
        return common.BKSetNameField
    case common.BKInnerObjIDModule:
        return common.BKModuleNameField
    case common.BKInnerObjIDHost:
        return common.BKHostInnerIPField
    case common.BKInnerObjIDProc:
        return common.BKProcNameField
    case common.BKInnerObjIDPlat:
        return common.BKCloudNameField
    default:
        return common.BKInstNameField
    }
}

// GetInstNameFieldName get the inst name
func (o *Object) GetInstNameFieldName() string {
    return GetInstNameFieldName(o.ObjectID)
}

func (o *Object) GetObjectType() string {
    switch o.ObjectID {
    case common.BKInnerObjIDApp:
        return o.ObjectID
    case common.BKInnerObjIDSet:
        return o.ObjectID
    case common.BKInnerObjIDModule:
        return o.ObjectID
    case common.BKInnerObjIDHost:
        return o.ObjectID
    case common.BKInnerObjIDProc:
        return o.ObjectID
    case common.BKInnerObjIDPlat:
        return o.ObjectID
    default:
        return common.BKInnerObjIDObject
    }
}

// GetObjectID get the object type
func (o *Object) GetObjectID() string {
    return o.ObjectID
}

// IsCommon is common object
func (o *Object) IsCommon() bool {
    return IsCommon(o.ObjectID)
}

func IsCommon(objID string) bool {
    switch objID {
    case common.BKInnerObjIDApp:
        return false
    case common.BKInnerObjIDSet:
        return false
    case common.BKInnerObjIDModule:
        return false
    case common.BKInnerObjIDHost:
        return false
    case common.BKInnerObjIDProc:
        return false
    case common.BKInnerObjIDPlat:
        return false
    default:
        return true
    }
}

// Parse load the data from mapstr object into object instance
func (o *Object) Parse(data mapstr.MapStr) (*Object, error) {

    err := mapstr.SetValueToStructByTags(o, data)
    if nil != err {
        return nil, err
    }

    return o, err
}

func (o *Object) ToMapStr() mapstr.MapStr {
    return mapstr.SetValueToMapStrByTags(o)
}

// MainLineObject main line object definition
type MainLineObject struct {
    Object        `json:",inline"`
    AssociationID string `json:"bk_asst_obj_id"`
}

type ObjectClsDes struct {
    ID      int    `json:"id" bson:"id"`
    ClsID   string `json:"bk_classification_id" bson:"bk_classification_id"`
    ClsName string `json:"bk_classification_name" bson:"bk_classification_name"`
    ClsType string `json:"bk_classification_type" bson:"bk_classification_type" `
    ClsIcon string `json:"bk_classification_icon" bson:"bk_classification_icon"`
}

type InnerModule struct {
    ModuleID         int64  `field:"bk_module_id" json:"bk_module_id" bson:"bk_module_id" mapstructure:"bk_module_id"`
    ModuleName       string `field:"bk_module_name" bson:"bk_module_name" json:"bk_module_name" mapstructure:"bk_module_name"`
    Default          int64  `field:"default" bson:"default" json:"default" mapstructure:"default"`
    HostApplyEnabled bool   `field:"host_apply_enabled" bson:"host_apply_enabled" json:"host_apply_enabled" mapstructure:"host_apply_enabled"`
}

type InnterAppTopo struct {
    SetID   int64         `json:"bk_set_id" field:"bk_set_id"`
    SetName string        `json:"bk_set_name" field:"bk_set_name"`
    Module  []InnerModule `json:"module" field:"module"`
}

// TopoItem define topo item
type TopoItem struct {
    ClassificationID string `json:"bk_classification_id"`
    Position         string `json:"position"`
    ObjID            string `json:"bk_obj_id"`
    OwnerID          string `json:"bk_supplier_account"`
    ObjName          string `json:"bk_obj_name"`
}

// ObjectTopo define the common object topo
type ObjectTopo struct {
    LabelType string   `json:"label_type"`
    LabelName string   `json:"label_name"`
    Label     string   `json:"label"`
    From      TopoItem `json:"from"`
    To        TopoItem `json:"to"`
    Arrows    string   `json:"arrows"`
}

//ObjectCountParams define parameter of search objects count
type ObjectCountParams struct {
    Condition ObjectIDArray `json:"condition"`
}

type ObjectIDArray struct {
    ObjectIDs []string `json:"obj_ids"`
}

//ObjectCountResult result by searching object count
type ObjectCountResult struct {
    BaseResp `json:",inline"`
    Data     []ObjectCountDetails `json:"data"`
}

//ObjectCountDetails one object count or error message of searching
type ObjectCountDetails struct {
    ObjectID  string `json:"bk_obj_id"`
    InstCount uint64 `json:"inst_count"`
    Error     string `json:"error"`
}
