package configcenter

import (
    err "errors"
    "sync"

    "github.com/spf13/viper"
)

var redisParser *viperParser
var mongodbParser *viperParser
var commonParser *viperParser
var extraParser *viperParser
var migrateParser *viperParser

var confLock sync.RWMutex

func String(key string) (string, error) {
    confLock.RLock()
    defer confLock.RUnlock()

    if migrateParser != nil && migrateParser.isSet(key) {
        return migrateParser.getString(key), nil
    }
    if commonParser != nil && commonParser.isSet(key) {
        return commonParser.getString(key), nil
    }
    if extraParser != nil && extraParser.isSet(key) {
        return extraParser.getString(key), nil
    }
    return "", err.New("config not found")
}

type viperParser struct {
    parser *viper.Viper
}

func (vp *viperParser) getString(path string) string {
    return vp.parser.GetString(path)
}

func (vp *viperParser) getInt(path string) int {
    return vp.parser.GetInt(path)
}

func (vp *viperParser) getUint64(path string) uint64 {
    return vp.parser.GetUint64(path)
}

func (vp *viperParser) getBool(path string) bool {
    return vp.parser.GetBool(path)
}

func (vp *viperParser) isSet(path string) bool {
    return vp.parser.IsSet(path)
}

func (vp *viperParser) getInt64(path string) int64 {
    return vp.parser.GetInt64(path)
}
