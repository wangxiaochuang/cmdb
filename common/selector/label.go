package selector

import (
    "fmt"
    "regexp"
)

type Labels map[string]string

var (
    LabelNGKeyRule   = regexp.MustCompile(`^[a-zA-Z]([a-z0-9A-Z\-_.]*[a-z0-9A-Z])?$`)
    LabelNGValueRule = regexp.MustCompile(`^[a-z0-9A-Z]([a-z0-9A-Z\-_.]*[a-z0-9A-Z])?$`)
)

func (lng Labels) Validate() (string, error) {
    for key, value := range lng {
        // validate key
        if LabelNGKeyRule.MatchString(key) == false {
            return key, fmt.Errorf("key: %s format error", key)
        }
        if len(key) >= 64 {
            return key, fmt.Errorf("key: %s exceed max length 63", key)
        }

        // validate value
        field := fmt.Sprintf("%s:%s", key, value)
        if LabelNGValueRule.MatchString(value) == false {
            return field, fmt.Errorf("value: %s format error", field)
        }
        if len(value) >= 64 {
            return field, fmt.Errorf("value: %s exceed max length 63", field)
        }
    }
    return "", nil
}

func (lng Labels) AddLabel(l Labels) {
    for key, value := range l {
        lng[key] = value
    }
}

func (lng Labels) RemoveLabel(keys []string) {
    for _, key := range keys {
        delete(lng, key)
    }
}

type LabelInstance struct {
    Labels Labels `bson:"labels" json:"labels"`
}
