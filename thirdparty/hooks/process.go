import (
    "github.com/wxc/cmdb/common/http/rest"
    "github.com/wxc/cmdb/common/mapstr"
    "github.com/wxc/cmdb/storage/dal"
)

// SetVIPInfoForProcessHook if query fields contains vip info, set vip
 info for processes
func SetVIPInfoForProcessHook(kit *rest.Kit, processes []mapstr.MapStr, fields []string, table string, db dal.DB) (
    []mapstr.MapStr, error) {

    return processes, nil
}

// ParseVIPFieldsForProcessHook parse process vip fields for process
func ParseVIPFieldsForProcessHook(fields []string, table string) ([]string, []string) {

    return fields, make([]string, 0)
}

// UpdateProcessBindInfoHook if process need to update bind info, only update the specified fields
func UpdateProcessBindInfoHook(kit *rest.Kit, objID string, origin mapstr.MapStr, data mapstr.MapStr) error {
    return nil
}
