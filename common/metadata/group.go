package metadata

import (
    "github.com/wxc/cmdb/common/mapstr"
)

const (
    GroupFieldID              = "id"
    GroupFieldGroupID         = "bk_group_id"
    GroupFieldGroupName       = "bk_group_name"
    GroupFieldGroupIndex      = "bk_group_index"
    GroupFieldObjectID        = "bk_obj_id"
    GroupFieldSupplierAccount = "bk_supplier_account"
    GroupFieldIsDefault       = "bk_isdefault"
    GroupFieldIsPre           = "ispre"
)

type PropertyGroupObjectAtt struct {
    Condition struct {
        OwnerID    string `field:"bk_supplier_account" json:"bk_supplier_account"`
        ObjectID   string `field:"bk_obj_id" json:"bk_obj_id"`
        PropertyID string `field:"bk_property_id" json:"bk_property_id"`
    } `json:"condition"`
    Data struct {
        PropertyGroupID string `field:"bk_property_group" json:"bk_property_group"`
        PropertyIndex   int    `field:"bk_property_index" json:"bk_property_index"`
    } `json:"data"`
}

type Group struct {
    BizID      int64  `field:"bk_biz_id" json:"bk_biz_id" bson:"bk_biz_id"`
    ID         int64  `field:"id" json:"id" bson:"id"`
    GroupID    string `field:"bk_group_id" json:"bk_group_id" bson:"bk_group_id"`
    GroupName  string `field:"bk_group_name" json:"bk_group_name" bson:"bk_group_name"`
    GroupIndex int64  `field:"bk_group_index" json:"bk_group_index" bson:"bk_group_index"`
    ObjectID   string `field:"bk_obj_id" json:"bk_obj_id" bson:"bk_obj_id"`
    OwnerID    string `field:"bk_supplier_account" json:"bk_supplier_account" bson:"bk_supplier_account"`
    IsDefault  bool   `field:"bk_isdefault" json:"bk_isdefault" bson:"bk_isdefault"`
    IsPre      bool   `field:"ispre" json:"ispre" bson:"ispre"`
    IsCollapse bool   `field:"is_collapse" json:"is_collapse" bson:"is_collapse"`
}

// Parse load the data from mapstr group into group instance
func (cli *Group) Parse(data mapstr.MapStr) (*Group, error) {

    err := mapstr.SetValueToStructByTags(cli, data)
    if nil != err {
        return nil, err
    }

    return cli, err
}

func (cli *Group) ToMapStr() mapstr.MapStr {
    return mapstr.SetValueToMapStrByTags(cli)
}
