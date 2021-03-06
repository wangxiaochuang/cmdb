package metadata

import (
	"github.com/wxc/cmdb/common"
	"github.com/wxc/cmdb/common/errors"
	"github.com/wxc/cmdb/common/mapstr"
)

// SetInst contains partial fields of a real set
type SetInst struct {
	BizID         int64  `bson:"bk_biz_id" json:"bk_biz_id" mapstructure:"bk_biz_id"`
	SetID         int64  `bson:"bk_set_id" json:"bk_set_id" mapstructure:"bk_set_id"`
	SetName       string `bson:"bk_set_name" json:"bk_set_name" mapstructure:"bk_set_name"`
	SetStatus     string `bson:"bk_service_status" json:"bk_service_status" mapstructure:"bk_service_status"`
	SetEnv        string `bson:"bk_set_env" json:"bk_set_env" mapstructure:"bk_set_env"`
	SetTemplateID int64  `bson:"set_template_id" json:"set_template_id" mapstructure:"set_template_id"`
	ParentID      int64  `bson:"bk_parent_id" json:"bk_parent_id" mapstructure:"bk_parent_id"`

	Creator         string `field:"creator" json:"creator,omitempty" bson:"creator" mapstructure:"creator"`
	CreateTime      Time   `field:"create_time" json:"create_time,omitempty" bson:"create_time" mapstructure:"create_time"`
	LastTime        Time   `field:"last_time" json:"last_time,omitempty" bson:"last_time" mapstructure:"last_time"`
	SupplierAccount string `field:"bk_supplier_account" json:"bk_supplier_account,omitempty" bson:"bk_supplier_account" mapstructure:"bk_supplier_account"`
}

// ModuleInst contains partial fields of a real module
type ModuleInst struct {
	BizID             int64  `bson:"bk_biz_id" json:"bk_biz_id" field:"bk_biz_id" mapstructure:"bk_biz_id"`
	SetID             int64  `bson:"bk_set_id" json:"bk_set_id" field:"bk_set_id" mapstructure:"bk_set_id"`
	ModuleID          int64  `bson:"bk_module_id" json:"bk_module_id" field:"bk_module_id" mapstructure:"bk_module_id"`
	ModuleName        string `bson:"bk_module_name" json:"bk_module_name" field:"bk_module_name" mapstructure:"bk_module_name"`
	SupplierAccount   string `bson:"bk_supplier_account" json:"bk_supplier_account" field:"bk_supplier_account" mapstructure:"bk_supplier_account"`
	ServiceCategoryID int64  `bson:"service_category_id" json:"service_category_id" field:"service_category_id" mapstructure:"service_category_id"`
	ServiceTemplateID int64  `bson:"service_template_id" json:"service_template_id" field:"service_template_id" mapstructure:"service_template_id"`
	ParentID          int64  `bson:"bk_parent_id" json:"bk_parent_id" field:"bk_parent_id" mapstructure:"bk_parent_id"`
	SetTemplateID     int64  `bson:"set_template_id" json:"set_template_id" field:"set_template_id" mapstructure:"set_template_id"`
	Default           int64  `bson:"default" json:"default" field:"default" mapstructure:"default"`
	HostApplyEnabled  bool   `bson:"host_apply_enabled" json:"host_apply_enabled" field:"host_apply_enabled" mapstructure:"host_apply_enabled"`
}

type BizInst struct {
	BizID           int64  `bson:"bk_biz_id" mapstructure:"bk_biz_id"`
	BizName         string `bson:"bk_biz_name" mapstructure:"bk_biz_name"`
	SupplierAccount string `bson:"bk_supplier_account" mapstructure:"bk_supplier_account"`
}

type BizBasicInfo struct {
	BizID   int64  `bson:"bk_biz_id" json:"bk_biz_id" field:"bk_biz_id" mapstructure:"bk_biz_id"`
	BizName string `bson:"bk_biz_name" json:"bk_biz_name" field:"bk_biz_name" mapstructure:"bk_biz_name"`
}

type CloudInst struct {
	CloudID   int64  `bson:"bk_cloud_id" json:"bk_cloud_id"`
	CloudName string `bson:"bk_cloud_name" json:"bk_cloud_name"`
}

type ProcessInst struct {
	ProcessID       int64  `json:"bk_process_id" bson:"bk_process_id"`               // ????????????
	ProcessName     string `json:"bk_process_name" bson:"bk_process_name"`           // ????????????
	BindIP          string `json:"bind_ip" bson:"bind_ip"`                           // ??????IP, ??????: [{ID: "1", Name: "127.0.0.1"}, {ID: "2", Name: "0.0.0.0"}, {ID: "3", Name: "????????????IP"}, {ID: "4", Name: "????????????IP"}]
	PORT            string `json:"port" bson:"port"`                                 // ??????, ???????????????"8080", ?????????????????????"8080-8089", ????????????????????????"8080-8089,8199"
	PROTOCOL        string `json:"protocol" bson:"protocol"`                         // ??????, ??????: [{ID: "1", Name: "TCP"}, {ID: "2", Name: "UDP"}],
	FuncName        string `json:"bk_func_name" bson:"bk_func_name"`                 // ????????????
	StartParamRegex string `json:"bk_start_param_regex" bson:"bk_start_param_regex"` // ????????????????????????
}

type HostIdentifier struct {
	HostID          int64                       `json:"bk_host_id" bson:"bk_host_id"`
	CloudID         int64                       `json:"bk_cloud_id" bson:"bk_cloud_id"`
	InnerIP         StringArrayToString         `json:"bk_host_innerip" bson:"bk_host_innerip"`
	OSType          string                      `json:"bk_os_type" bson:"bk_os_type"`
	SupplierAccount string                      `json:"bk_supplier_account" bson:"bk_supplier_account"`
	HostIdentModule map[string]*HostIdentModule `json:"associations" bson:"associations"`
	Process         []HostIdentProcess          `json:"process" bson:"process"`
}

type HostIdentProcess struct {
	ProcessID   int64  `json:"bk_process_id" bson:"bk_process_id"`     // ????????????
	ProcessName string `json:"bk_process_name" bson:"bk_process_name"` // ????????????
	// deprecated  ??????????????????????????????
	BindIP string `json:"bind_ip" bson:"bind_ip"` // ??????IP, ??????: [{ID: "1", Name: "127.0.0.1"}, {ID: "2", Name: "0.0.0.0"}, {ID: "3", Name: "????????????IP"}, {ID: "4", Name: "????????????IP"}]
	// deprecated  ??????????????????????????????
	Port string `json:"port" bson:"port"` // ??????, ???????????????"8080", ?????????????????????"8080-8089", ????????????????????????"8080-8089,8199"
	// deprecated  ??????????????????????????????
	Protocol        string `json:"protocol" bson:"protocol"`                         // ??????, ??????: [{ID: "1", Name: "TCP"}, {ID: "2", Name: "UDP"}],
	FuncName        string `json:"bk_func_name" bson:"bk_func_name"`                 // ????????????
	StartParamRegex string `json:"bk_start_param_regex" bson:"bk_start_param_regex"` // ????????????????????????
	// deprecated  ??????????????????????????????
	PortEnable bool `field:"bk_enable_port" json:"bk_enable_port" bson:"bk_enable_port"`
	// BindInfo ??????????????????
	BindInfo []ProcBindInfo `field:"bind_info" json:"bind_info" bson:"bind_info"`
}

// HostIdentModule HostIdentifier module define
type HostIdentModule struct {
	BizID    int64  `json:"bk_biz_id"`    // ??????ID
	SetID    int64  `json:"bk_set_id"`    // ????????????(bk_set_id)
	ModuleID int64  `json:"bk_module_id"` // ????????????(bk_module_id)
	Layer    *Layer `json:"layer"`        // ???????????????
}

type Layer struct {
	InstID   int64  `json:"bk_inst_id"`
	InstName string `json:"bk_inst_name"`
	ObjID    string `json:"bk_obj_id"`
	Child    *Layer `json:"child"`
}

type MainlineInstInfo struct {
	InstID   int64  `json:"bk_inst_id" bson:"bk_inst_id"`
	InstName string `json:"bk_inst_name" bson:"bk_inst_name"`
	ObjID    string `json:"bk_obj_id" bson:"bk_obj_id"`
	ParentID int64  `json:"bk_parent_id" bson:"bk_parent_id"`
}

// SearchIdentifierParam defines the param
type SearchIdentifierParam struct {
	IP   IPParam `json:"ip"`
	Page BasePage
}

// SearchHostIdentifierParam ???????????????????????????
type SearchHostIdentifierParam struct {
	HostIDs []int64 `json:"host_ids"`
}

type IPParam struct {
	Data    []string `json:"data"`
	CloudID *int64   `json:"bk_cloud_id"`
}

type SearchHostIdentifierResult struct {
	BaseResp `json:",inline"`
	Data     SearchHostIdentifierData `json:"data"`
}

// SearchHostIdentifierData host identifier detail
type SearchHostIdentifierData struct {
	Count int              `json:"count"`
	Info  []HostIdentifier `json:"info"`
}

// SearchInstsNamesOption search instances names option
type SearchInstsNamesOption struct {
	ObjID string `json:"bk_obj_id"`
	BizID int64  `json:"bk_biz_id"`
	Name  string `json:"name"`
}

var ObjsForSearchName = map[string]bool{
	common.BKInnerObjIDSet:    true,
	common.BKInnerObjIDModule: true,
}

// Validate verify the SearchInstsNamesOption
func (o *SearchInstsNamesOption) Validate() (rawError errors.RawErrorInfo) {
	if _, ok := ObjsForSearchName[o.ObjID]; !ok {
		return errors.RawErrorInfo{
			ErrCode: common.CCErrCommParamsInvalid,
			Args:    []interface{}{"bk_obj_id"},
		}
	}

	if o.BizID <= 0 {
		return errors.RawErrorInfo{
			ErrCode: common.CCErrCommParamsInvalid,
			Args:    []interface{}{"bk_biz_id"},
		}
	}

	if o.Name == "" {
		return errors.RawErrorInfo{
			ErrCode: common.CCErrCommParamsInvalid,
			Args:    []interface{}{"name"},
		}
	}

	return errors.RawErrorInfo{}
}

//CreateManyCommInst parameters for creating multiple instances
type CreateManyCommInst struct {
	ObjID   string          `json:"bk_obj_id"`
	Details []mapstr.MapStr `json:"details"`
}

//CreateManyCommInstResult result of creating multiple instances
type CreateManyCommInstResult struct {
	BaseResp `json:",inline"`
	Data     CreateManyCommInstResultDetail `json:"data"`
}

//CreateManyCommInstResultDetail details of CreateManyInstancesResult
type CreateManyCommInstResultDetail struct {
	SuccessCreated map[int64]int64  `json:"success_created"`
	Error          map[int64]string `json:"error_msg"`
}

func NewManyCommInstResultDetail() *CreateManyCommInstResultDetail {
	return &CreateManyCommInstResultDetail{
		SuccessCreated: make(map[int64]int64, 0),
		Error:          make(map[int64]string, 0),
	}
}
