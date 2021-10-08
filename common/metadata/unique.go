package metadata

import (
    "fmt"
    "sort"
    "strings"

    "github.com/wxc/cmdb/common/mapstr"
)

type ObjectUnique struct {
    ID       uint64      `json:"id" bson:"id"`
    ObjID    string      `json:"bk_obj_id" bson:"bk_obj_id"`
    Keys     []UniqueKey `json:"keys" bson:"keys"`
    Ispre    bool        `json:"ispre" bson:"ispre"`
    OwnerID  string      `json:"bk_supplier_account" bson:"bk_supplier_account"`
    LastTime Time        `json:"last_time" bson:"last_time"`
}

// Parse load the data from mapstr attribute into ObjectUnique instance
func (cli *ObjectUnique) Parse(data mapstr.MapStr) (*ObjectUnique, error) {

    err := mapstr.SetValueToStructByTags(cli, data)
    if nil != err {
        return nil, err
    }

    return cli, err
}

func (u ObjectUnique) KeysHash() string {
    keys := []string{}
    for _, key := range u.Keys {
        keys = append(keys, fmt.Sprintf("%s:%d", key.Kind, key.ID))
    }
    sort.Strings(keys)
    return strings.Join(keys, "#")
}

type UniqueKey struct {
    Kind string `json:"key_kind" bson:"key_kind"`
    ID   uint64 `json:"key_id" bson:"key_id"`
}

const (
    UniqueKeyKindProperty    = "property"
    UniqueKeyKindAssociation = "association"
)

type CreateUniqueRequest struct {
    ObjID string      `json:"bk_obj_id" bson:"bk_obj_id"`
    Keys  []UniqueKey `json:"keys" bson:"keys"`
}

type CreateUniqueResult struct {
    BaseResp
    Data RspID `json:"data"`
}

type UpdateUniqueRequest struct {
    Keys     []UniqueKey `json:"keys" bson:"keys"`
    LastTime Time        `json:"last_time" bson:"last_time"`
}

type UpdateUniqueResult struct {
    BaseResp
}

type DeleteUniqueRequest struct {
    ID    uint64 `json:"id"`
    ObjID string `json:"bk_obj_id"`
}

type DeleteUniqueResult struct {
    BaseResp
}

type SearchUniqueRequest struct {
    ObjID string `json:"bk_obj_id"`
}

type SearchUniqueResult struct {
    BaseResp
    Data []ObjectUnique `json:"data"`
}

type QueryUniqueResult struct {
    Count uint64         `json:"count"`
    Info  []ObjectUnique `json:"info"`
}
