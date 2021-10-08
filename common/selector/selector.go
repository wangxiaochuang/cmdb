package selector

import (
    "errors"
    "fmt"

    "github.com/wxc/cmdb/common"
    "github.com/wxc/cmdb/common/util"
)

type LabelAddOption struct {
    InstanceIDs []int64 `json:"instance_ids"`
    Labels      Labels  `json:"labels"`
}

type LabelAddRequest struct {
    Option    LabelAddOption `json:"option"`
    TableName string         `json:"table_name"`
}

type LabelRemoveOption struct {
    InstanceIDs []int64  `json:"instance_ids"`
    Keys        []string `json:"keys"`
}

type LabelRemoveRequest struct {
    Option    LabelRemoveOption `json:"option"`
    TableName string            `json:"table_name"`
}

type Operator string

const (
    DoesNotExist Operator = "!"
    Equals       Operator = "="
    In           Operator = "in"
    NotEquals    Operator = "!="
    NotIn        Operator = "notin"
    Exists       Operator = "exists"
)

var AvailableOperators = []Operator{
    DoesNotExist,
    Equals,
    In,
    NotEquals,
    NotIn,
    Exists,
}

type Selector struct {
    Key      string   `json:"key" field:"key" bson:"key"`
    Operator Operator `json:"operator" field:"operator" bson:"operator"`
    Values   []string `json:"values" field:"values" bson:"values"`
}

func (s *Selector) Validate() (string, error) {
    if util.InArray(s.Operator, AvailableOperators) == false {
        return "operator", fmt.Errorf("operator %s not available, available operators: %+v", s.Operator, AvailableOperators)
    }

    if (s.Operator == In || s.Operator == NotIn) && len(s.Values) == 0 {
        return "values", errors.New("values shouldn't be empty")
    }

    if (s.Operator == Exists || s.Operator == DoesNotExist) && len(s.Values) > 0 {
        return "values", errors.New("values shouldn be empty")
    }

    if (s.Operator == Equals || s.Operator == NotEquals) && len(s.Values) != 1 {
        return "values", errors.New("values field length for equal operation should exactly one")
    }

    if LabelNGKeyRule.MatchString(s.Key) == false {
        return "key", fmt.Errorf("key %s invalid", s.Key)
    }
    return "", nil
}

func (s *Selector) ToMgoFilter() (map[string]interface{}, error) {
    filter := make(map[string]interface{})
    field := "labels." + s.Key
    switch s.Operator {
    case In:
        filter = map[string]interface{}{
            field: map[string]interface{}{
                common.BKDBIN: s.Values,
            },
        }
    case NotIn:
        filter = map[string]interface{}{
            field: map[string]interface{}{
                common.BKDBNIN: s.Values,
            },
        }
    case DoesNotExist, Exists:
        filter = map[string]interface{}{
            field: map[string]interface{}{
                common.BKDBExists: s.Operator == Exists,
            },
        }
    case Equals:
        if len(s.Values) == 0 {
            return nil, errors.New("values empty")
        }
        firstValue := s.Values[0]
        filter = map[string]interface{}{
            field: firstValue,
        }
    case NotEquals:
        if len(s.Values) == 0 {
            return nil, errors.New("values empty")
        }
        firstValue := s.Values[0]
        filter = map[string]interface{}{
            field: map[string]interface{}{
                common.BKDBNE: firstValue,
            },
        }
    }
    return filter, nil
}

type Selectors []Selector

func (ss Selectors) Validate() (string, error) {
    for _, selector := range ss {
        if key, err := selector.Validate(); err != nil {
            return key, err
        }
    }
    return "", nil
}

func (ss Selectors) ToMgoFilter() (map[string]interface{}, error) {
    filters := make([]map[string]interface{}, 0)
    for _, selector := range ss {
        filter, err := selector.ToMgoFilter()
        if err != nil {
            return nil, err
        }
        filters = append(filters, filter)
    }
    if len(filters) == 0 {
        return make(map[string]interface{}), nil
    }
    if len(filters) == 1 {
        return filters[0], nil
    }
    return map[string]interface{}{
        common.BKDBAND: filters,
    }, nil
}
