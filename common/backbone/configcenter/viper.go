package configcenter

import (
    err "errors"
    "fmt"
    "os"
    "sync"

    "github.com/wxc/cmdb/storage/dal/mongo"
    "github.com/wxc/cmdb/storage/dal/redis"

    "github.com/spf13/viper"
)

var redisParser *viperParser
var mongodbParser *viperParser
var commonParser *viperParser
var extraParser *viperParser
var migrateParser *viperParser

var confLock sync.RWMutex

func checkDir(path string) error {
    info, err := os.Stat(path)
    if os.ErrNotExist == err {
        return fmt.Errorf("directory %s not exists", path)
    }
    if err != nil {
        return fmt.Errorf("stat directory %s faile, %s", path, err.Error())
    }
    if !info.IsDir() {
        return fmt.Errorf("%s is not directory", path)
    }

    return nil
}

func LoadConfigFromLocalFile(confPath string, handler *CCHandler) error {
    panic("in LoadConfigFromLocalFile")
    return nil
}

func Redis(prefix string) (redis.Config, error) {
    return redis.Config{}, err.New("can't find redis configuration")
}

func Mongo(prefix string) (mongo.Config, error) {
    return mongo.Config{}, err.New("can't find mongo configuration")
}

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
