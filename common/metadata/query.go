package metadata

import (
	"strings"

	"github.com/wxc/cmdb/common"
	"github.com/wxc/cmdb/common/mapstr"
)

// Deprecated: SearchLimit sub condition
type SearchLimit struct {
	Offset int64 `json:"start" field:"start"`
	Limit  int64 `json:"limit" field:"limit"`
}

// SearchSort sub condition
type SearchSort struct {
	IsDsc bool   `json:"is_dsc"`
	Field string `json:"field"`
}

// QueryCondition the common query condition definition
type QueryCondition struct {
	Fields    []string      `json:"fields"`
	Page      BasePage      `json:"page"`
	Condition mapstr.MapStr `json:"condition"`
	// 非必填，只能用来查时间，且与Condition是与关系
	TimeCondition  *TimeCondition `json:"time_condition,omitempty"`
	DisableCounter bool           `json:"disable_counter"`
}

// IsIllegal  limit is illegal, if limit = 0; change to default page size
func (qc *QueryCondition) IsIllegal() bool {
	if qc.Page.Limit == 0 {
		qc.Page.Limit = common.BKDefaultLimit
	}
	if qc.Page.Limit > common.BKMaxPageSize && qc.Page.Limit != common.BKNoLimit {
		return true
	}
	return false
}

// QueryResult common query result
type QueryResult struct {
	Count uint64          `json:"count"`
	Info  []mapstr.MapStr `json:"info"`
}

type QueryConditionResult ResponseInstData

// SearchSortParse SearchSort parse interface
type SearchSortParse interface {
	String(sort string) SearchSortParse
	Field(field string, isDesc bool) SearchSortParse
	Set(ssArr []SearchSort) SearchSortParse
	ToMongo() string
	ToSearchSortArr() []SearchSort
}

// searchSortParse SearchSort parse struct
type searchSortParse struct {
	data []SearchSort
}

func NewSearchSortParse() SearchSortParse {
	return &searchSortParse{}
}

//  String convert string sort to cc SearchSort struct array
func (ss *searchSortParse) String(sort string) SearchSortParse {
	if sort == "" {
		return ss
	}
	sortArr := strings.Split(sort, ",")
	for _, sortItem := range sortArr {
		sortItemArr := strings.Split(sortItem, ":")
		ssInst := SearchSort{
			Field: sortItemArr[0],
		}
		if len(sortItemArr) > 1 && strings.TrimSpace(sortItemArr[1]) == "-1" {
			ssInst.IsDsc = true

		}
		ss.data = append(ss.data, ssInst)
	}
	return ss
}

//  Field   cc SearchSort struct array
func (ss *searchSortParse) Field(field string, isDesc bool) SearchSortParse {

	ssInst := SearchSort{
		Field: field,
		IsDsc: isDesc,
	}
	ss.data = append(ss.data, ssInst)
	return ss
}

func (ss *searchSortParse) Set(ssArr []SearchSort) SearchSortParse {
	ss.data = append(ss.data, ssArr...)
	return ss
}

// ToSearchSortArr cc SearchSort struct to mongodb sort filed
func (ss *searchSortParse) ToSearchSortArr() []SearchSort {
	return ss.data
}

// searchSortParse cc SearchSort struct to mongodb sort filed
func (ss *searchSortParse) ToMongo() string {
	var orderByArr []string
	for _, item := range ss.data {
		if item.IsDsc {
			orderByArr = append(orderByArr, item.Field+":-1")
		} else {
			orderByArr = append(orderByArr, item.Field+":1")
		}
	}
	return strings.Join(orderByArr, ",")
}

// IsIllegal  limit is illegal
func (page SearchLimit) IsIllegal() bool {
	if page.Limit > common.BKMaxPageSize && page.Limit != common.BKNoLimit ||
		page.Limit == 0 {
		return true
	}
	return false
}
