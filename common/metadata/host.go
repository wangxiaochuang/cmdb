package metadata

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/wxc/cmdb/common"
	"github.com/wxc/cmdb/common/blog"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/bsoncodec"
	"go.mongodb.org/mongo-driver/bson/bsonrw"
	"go.mongodb.org/mongo-driver/bson/bsontype"
	"go.mongodb.org/mongo-driver/x/bsonx"
	"go.mongodb.org/mongo-driver/x/bsonx/bsoncore"
)

// host map with string type ip and operator, can only get host from db with this map
type HostMapStr map[string]interface{}

func (h *HostMapStr) UnmarshalBSON(b []byte) error {
	if h == nil {
		return bsonx.ErrNilDocument
	}
	elements, err := bsoncore.Document(b).Elements()
	if err != nil {
		return err
	}

	if *h == nil {
		*h = map[string]interface{}{}
	}
	for _, element := range elements {
		rawValue := element.Value()
		switch element.Key() {
		case common.BKHostInnerIPField:
			innerIP, err := parseBsonStringArrayValueToString(rawValue)
			if err != nil {
				return err
			}
			(*h)[common.BKHostInnerIPField] = string(innerIP)
		case common.BKHostOuterIPField:
			outerIP, err := parseBsonStringArrayValueToString(rawValue)
			if err != nil {
				return err
			}
			(*h)[common.BKHostOuterIPField] = string(outerIP)
		case common.BKOperatorField:
			operator, err := parseBsonStringArrayValueToString(rawValue)
			if err != nil {
				return err
			}
			(*h)[common.BKOperatorField] = string(operator)
		case common.BKBakOperatorField:
			bakOperator, err := parseBsonStringArrayValueToString(rawValue)
			if err != nil {
				return err
			}
			(*h)[common.BKBakOperatorField] = string(bakOperator)
		default:
			dc := bsoncodec.DecodeContext{Registry: bson.DefaultRegistry}
			vr := bsonrw.NewBSONValueReader(rawValue.Type, rawValue.Data)
			decoder, err := bson.NewDecoderWithContext(dc, vr)
			if err != nil {
				return err
			}
			value := new(interface{})
			err = decoder.Decode(value)
			if err != nil {
				return err
			}
			(*h)[element.Key()] = *value
		}
	}
	return nil
}

func parseBsonStringArrayValueToString(value bsoncore.Value) ([]byte, error) {
	switch value.Type {
	case bsontype.Array:
		rawArray, rem, ok := bsoncore.ReadArray(value.Data)
		if !ok {
			return nil, bsoncore.NewInsufficientBytesError(value.Data, rem)
		}
		array, err := rawArray.Values()
		if err != nil {
			return nil, err
		}
		var buf bytes.Buffer
		arrayLen := len(array)
		for index, arrayValue := range array {
			if arrayValue.Type != bsontype.String {
				return nil, fmt.Errorf("invalid BSON type %v", arrayValue.Type)
			}
			str, rem, ok := bsoncore.ReadString(arrayValue.Data)
			if !ok {
				return nil, bsoncore.NewInsufficientBytesError(arrayValue.Data, rem)
			}
			buf.WriteString(str)
			if index != arrayLen-1 {
				buf.WriteByte(',')
			}
		}
		return buf.Bytes(), nil
	case bsontype.Null:
		return []byte{}, nil
	default:
		return nil, fmt.Errorf("invalid BSON type %v", value.Type)
	}
}

// special field whose string array value is parsed into string value from db
type StringArrayToString string

func (s *StringArrayToString) UnmarshalBSONValue(typo bsontype.Type, raw []byte) error {
	if s == nil {
		return bsonx.ErrNilDocument
	}
	value := bsoncore.Value{
		Type: typo,
		Data: raw,
	}
	str, err := parseBsonStringArrayValueToString(value)
	if err != nil {
		return err
	}
	*s = StringArrayToString(str)
	return err
}

var specialFields = []string{common.BKHostInnerIPField, common.BKHostOuterIPField, common.BKOperatorField, common.BKBakOperatorField}

// convert host ip and operator fields value from string to array
// NOTICE: if host special value is empty, convert it to null to trespass the unique check, **do not change this logic**
func ConvertHostSpecialStringToArray(host map[string]interface{}) map[string]interface{} {
	for _, field := range specialFields {
		value, ok := host[field]
		if !ok {
			continue
		}
		switch v := value.(type) {
		case string:
			v = strings.TrimSpace(v)
			v = strings.Trim(v, ",")
			if len(v) == 0 {
				host[field] = nil
			} else {
				host[field] = strings.Split(v, ",")
			}
		case []string:
			if len(v) == 0 {
				host[field] = nil
			}
		case []interface{}:
			if len(v) == 0 {
				host[field] = nil
			} else {
				blog.Errorf("host %s type invalid, value %v", field, host[field])
			}
		case nil:
		default:
			blog.Errorf("host %s type invalid, value %v", field, host[field])
		}
	}
	return host
}
