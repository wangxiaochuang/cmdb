package json

import (
    "strings"

    jsoniter "github.com/json-iterator/go"
)

var iteratorJson = jsoniter.Config{
    EscapeHTML:             true,
    SortMapKeys:            true,
    ValidateJsonRawMessage: true,
    UseNumber:              true,
}.Froze()

func MarshalToString(v interface{}) (string, error) {
    return iteratorJson.MarshalToString(v)
}

func Marshal(v interface{}) ([]byte, error) {
    return iteratorJson.Marshal(v)
}

func MarshalIndent(v interface{}, prefix, indent string) ([]byte, error) {
    return iteratorJson.MarshalIndent(v, prefix, indent)
}

func UnmarshalFromString(str string, v interface{}) error {
    return iteratorJson.UnmarshalFromString(str, v)
}

func Unmarshal(data []byte, v interface{}) error {
    return iteratorJson.Unmarshal(data, v)
}

func UnmarshalArray(items []string, result interface{}) error {
    strArrJSON := "[" + strings.Join(items, ",") + "]"
    return iteratorJson.Unmarshal([]byte(strArrJSON), result)
}
