package time

import (
    "strings"
    "time"

    "github.com/wxc/cmdb/common/json"
    "github.com/wxc/cmdb/common/util"

    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/bsontype"
    "go.mongodb.org/mongo-driver/x/bsonx"
)

type Time struct {
    time.Time `bson:",inline" json:",inline"`
}

func (t Time) MarshalJSON() ([]byte, error) {
    return json.Marshal(t.Time)
}

func (t *Time) UnmarshalJSON(data []byte) error {
    // Ignore null, like in the main JSON package.
    if len(data) == 0 || string(data) == "null" {
        return nil
    }

    dataStr := strings.Trim(string(data), "\"")

    timeType, isTime := util.IsTime(dataStr)
    if isTime {
        t.Time = util.Str2Time(dataStr, timeType)
        return nil
    }

    return json.Unmarshal(data, &t.Time)
}

// MarshalBSONValue implements bson.MarshalBSON interface
func (t Time) MarshalBSONValue() (bsontype.Type, []byte, error) {
    return bsonx.Time(t.Time).MarshalBSONValue()
}

func (t *Time) UnmarshalBSONValue(typo bsontype.Type, raw []byte) error {
    switch typo {
    case bsontype.DateTime:
        rv := bson.RawValue{Type: bsontype.DateTime, Value: raw}
        t.Time = rv.Time()
        return nil
    case bsontype.String:
        rawStr := bson.RawValue{Type: bsontype.String, Value: raw}
        return t.UnmarshalJSON([]byte(rawStr.String()))
    }

    return bson.Unmarshal(raw, &t.Time)
}
