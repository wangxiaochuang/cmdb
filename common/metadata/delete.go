package metadata

import "github.com/wxc/cmdb/common/mapstr"

// DeleteOption common delete condition options
type DeleteOption struct {
	Condition mapstr.MapStr `json:"condition"`
}

// DeletedCountResult delete  api http response return result struct
type DeletedOptionResult struct {
	BaseResp `json:",inline"`
	Data     DeletedCount `json:"data"`
}
