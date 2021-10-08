package mapstr

import "reflect"

// GetTags parse a object and get the all tags
func GetTags(target interface{}, tagName string) []string {

    targetType := reflect.TypeOf(target)
    switch targetType.Kind() {
    default:
        break
    case reflect.Ptr:
        targetType = targetType.Elem()
    }

    numField := targetType.NumField()
    tags := make([]string, 0)
    for i := 0; i < numField; i++ {
        structField := targetType.Field(i)
        if tag, ok := structField.Tag.Lookup("field"); ok {
            tags = append(tags, tag)
        }
    }
    return tags

}
