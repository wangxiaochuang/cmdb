package mapstr

import "reflect"

func getZeroValue(valueType reflect.Type) interface{} {

    switch valueType.Kind() {
    case reflect.Ptr:
        return getZeroValue(valueType.Elem())
    case reflect.String:
        return ""
    case reflect.Int, reflect.Int16, reflect.Int8, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uint8:
        return 0
    }

    return nil
}
