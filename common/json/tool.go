package json

import (
    "bytes"

    "github.com/tidwall/gjson"
)

func CutJsonDataWithFields(jsonData *string, fields []string) *string {
    if jsonData == nil {
        empty := ""
        return &empty
    }
    if len(fields) == 0 || *jsonData == "" {
        return jsonData
    }
    elements := gjson.GetMany(*jsonData, fields...)
    last := len(fields) - 1
    jsonBuffer := bytes.Buffer{}
    jsonBuffer.Write([]byte{'{'})
    for idx, field := range fields {
        jsonBuffer.Write([]byte{'"'})
        jsonBuffer.Write([]byte(field))
        jsonBuffer.Write([]byte{'"'})
        jsonBuffer.Write([]byte{':'})
        if elements[idx].Raw == "" {
            jsonBuffer.Write([]byte("null"))
        } else {
            jsonBuffer.Write([]byte(elements[idx].Raw))
        }
        if idx != last {
            jsonBuffer.Write([]byte{','})
        }
    }
    jsonBuffer.Write([]byte{'}'})
    cutOff := jsonBuffer.String()
    return &cutOff
}
