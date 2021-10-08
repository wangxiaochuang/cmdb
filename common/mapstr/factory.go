package mapstr

import (
    "encoding/json"
    "fmt"
    "reflect"
)

// New create a new MapStr instance
func New() MapStr {
    return MapStr{}
}

// NewArray create MapStr array
func NewArray() []MapStr {
    return []MapStr{}
}

// NewArrayFromMapStr create a new array from mapstr array
func NewArrayFromMapStr(datas []MapStr) []MapStr {
    results := []MapStr{}
    for _, item := range datas {
        results = append(results, item)
    }
    return results
}

func NewFromInterface(data interface{}) (MapStr, error) {

    switch tmp := data.(type) {
    default:
        return convertInterfaceIntoMapStrByReflection(data, "")
    case nil:
        return MapStr{}, nil
    case MapStr:
        return tmp, nil
    case []byte:
        result := New()
        if 0 == len(tmp) {
            return result, nil
        }
        err := json.Unmarshal(tmp, &result)
        return result, err
    case string:
        result := New()
        if 0 == len(tmp) {
            return result, nil
        }
        err := json.Unmarshal([]byte(tmp), &result)
        return result, err
    case *map[string]interface{}:
        return MapStr(*tmp), nil
    case map[string]string:
        result := New()
        for key, val := range tmp {
            result.Set(key, val)
        }
        return result, nil
    case map[string]interface{}:
        return MapStr(tmp), nil
    }
}

// NewFromMap create a new MapStr from map[string]interface{} type
func NewFromMap(data map[string]interface{}) MapStr {
    return MapStr(data)
}

func NewFromStruct(targetStruct interface{}, tagName string) MapStr {
    return SetValueToMapStrByTagsWithTagName(targetStruct, tagName)
}

// NewArrayFromInterface create a new array from interface
func NewArrayFromInterface(datas []map[string]interface{}) []MapStr {
    results := []MapStr{}
    for _, item := range datas {
        results = append(results, item)
    }
    return results
}

// SetValueToMapStrByTags  convert a struct to MapStr by tags default tag name is field
func SetValueToMapStrByTags(source interface{}) MapStr {
    return SetValueToMapStrByTagsWithTagName(source, "field")
}

func SetValueToMapStrByTagsWithTagName(source interface{}, tagName string) MapStr {
    values := MapStr{}

    if nil == source {
        return values
    }

    targetType := getTypeElem(reflect.TypeOf(source))
    targetValue := getValueElem(reflect.ValueOf(source))

    setMapStrByStruct(targetType, targetValue, values, tagName)

    return values
}

// SetValueToStructByTags set the struct object field value by tags, default tag name is field
func SetValueToStructByTags(target interface{}, values MapStr) error {
    return SetValueToStructByTagsWithTagName(target, values, "field")
}

// SetValueToStructByTagsWithTagName set the struct object field value by tags
func SetValueToStructByTagsWithTagName(target interface{}, values MapStr, tagName string) error {

    targetType := reflect.TypeOf(target)
    targetValue := reflect.ValueOf(target)

    return setStructByMapStr(targetType, targetValue, values, tagName)
}

func convertInterfaceIntoMapStrByReflection(target interface{}, tagName string) (MapStr, error) {

    value := reflect.ValueOf(target)
    switch value.Kind() {
    case reflect.Map:
        return dealMap(value, tagName)
    case reflect.Struct:
        return dealStruct(value.Type(), value, tagName)
    }

    return nil, fmt.Errorf("no support the kind(%s)", value.Kind())
}
