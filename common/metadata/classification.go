package metadata

import (
    "github.com/wxc/cmdb/common/mapstr"
)

const (
    ClassificationFieldID        = "id"
    ClassFieldClassificationID   = "bk_classification_id"
    ClassFieldClassificationName = "bk_classification_name"
)

// Classification the classification metadata definition
type Classification struct {
    ID                 int64  `field:"id" json:"id" bson:"id"`
    ClassificationID   string `field:"bk_classification_id"  json:"bk_classification_id" bson:"bk_classification_id"`
    ClassificationName string `field:"bk_classification_name" json:"bk_classification_name" bson:"bk_classification_name"`
    ClassificationType string `field:"bk_classification_type" json:"bk_classification_type" bson:"bk_classification_type"`
    ClassificationIcon string `field:"bk_classification_icon" json:"bk_classification_icon" bson:"bk_classification_icon"`
    OwnerID            string `field:"bk_supplier_account" json:"bk_supplier_account" bson:"bk_supplier_account"  `
}

func (cli *Classification) Parse(data mapstr.MapStr) (*Classification, error) {

    err := mapstr.SetValueToStructByTags(cli, data)
    if nil != err {
        return nil, err
    }

    return cli, err
}

// ToMapStr to mapstr
func (cli *Classification) ToMapStr() mapstr.MapStr {
    return mapstr.SetValueToMapStrByTags(cli)
}
