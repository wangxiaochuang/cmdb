package mapstr

import (
    "bytes"
    "encoding/json"
)

// DecodeFromMapStr convert input into json, then decode json into data
// 接口背景：mapstr 直接解析结构体实现的不完整，有很多坑点，已知问题：结构体中指针类型会导致 mapstr 解析结构体异常。
// 新的问题：mapstr 转json时数据会丢失
func DecodeFromMapStr(data interface{}, input MapStr) error {
    inputBytes, err := json.Marshal(input)
    if err != nil {
        return err
    }
    d := json.NewDecoder(bytes.NewReader(inputBytes))
    d.UseNumber()
    err = d.Decode(data)
    return err
}
