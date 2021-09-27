package blog

import (
    "encoding/json"
    "flag"
    "fmt"
    "log"
    "reflect"
    "sync"
    "time"

    "github.com/wxc/cmdb/common/blog/glog"
)

func init() {
    flag.Set("logtostderr", "true")
}

type GlogWriter struct{}

// Write implements the io.Writer interface.
func (writer GlogWriter) Write(data []byte) (n int, err error) {
    glog.InfoDepth(1, string(data))
    return len(data), nil
}

// Output for mgo logger
func (writer GlogWriter) Output(calldepth int, s string) error {
    glog.InfoDepth(calldepth, s)
    return nil
}

func (writer GlogWriter) Print(v ...interface{}) {
    glog.InfoDepth(1, v...)
}
func (writer GlogWriter) Printf(format string, v ...interface{}) {
    glog.InfoDepthf(1, format, v...)
}

func (writer GlogWriter) Println(v ...interface{}) {
    glog.InfoDepth(1, v...)
}

var once sync.Once

func InitLogs() {
    once.Do(func() {
        log.SetOutput(GlogWriter{})
        log.SetFlags(0)
        go func() {
            d := time.Duration(5 * time.Second)
            tick := time.Tick(d)
            for {
                select {
                case <-tick:
                    glog.Flush()
                }
            }
        }()
    })
}

func CloseLogs() {
    glog.Flush()
}

var (
    Info        = glog.Infof
    Infof       = glog.Infof
    InfofDepthf = glog.InfoDepthf

    Warn  = glog.Warningf
    Warnf = glog.Warningf

    Error        = glog.Errorf
    Errorf       = glog.Errorf
    ErrorfDepthf = glog.ErrorfDepthf

    Fatal  = glog.Fatal
    Fatalf = glog.Fatalf

    V = glog.V
)

func Debug(args ...interface{}) {
    if format, ok := (args[0]).(string); ok {
        glog.InfoDepthf(1, format, args[1:]...)
    } else {
        glog.InfoDepth(1, args)
    }
}

func InfoJSON(format string, args ...interface{}) {
    params := []interface{}{}
    for _, arg := range args {
        if f, ok := arg.(errorFunc); ok {
            params = append(params, f.Error())
            continue
        }
        if f, ok := arg.(stringFunc); ok {
            params = append(params, f.String())
            continue
        }

        if arg == nil {
            params = append(params, []byte("null"))
            continue
        }

        kind := reflect.TypeOf(arg).Kind()
        if kind == reflect.Ptr {
            kind = reflect.TypeOf(arg).Elem().Kind()
        }
        if kind == reflect.Struct || kind == reflect.Interface ||
            kind == reflect.Array || kind == reflect.Map || kind == reflect.Slice {
            out, err := json.Marshal(arg)
            if err != nil {
                params = append(params, arg)
            } else {
                params = append(params, out)
            }
            continue
        }

        params = append(params, arg)
    }
    glog.InfoDepthf(1, format, params...)
}

func ErrorJSON(format string, args ...interface{}) {
    params := []interface{}{}
    for _, arg := range args {
        if f, ok := arg.(errorFunc); ok {
            params = append(params, f.Error())
            continue
        }
        if f, ok := arg.(stringFunc); ok {
            params = append(params, f.String())
            continue
        }
        out, err := json.Marshal(arg)
        if err != nil {
            params = append(params, err.Error())
        }
        params = append(params, out)
    }
    glog.ErrorDepth(1, fmt.Sprintf(format, params...))
}

func WarnJSON(format string, args ...interface{}) {

    params := []interface{}{}
    for _, arg := range args {
        if f, ok := arg.(errorFunc); ok {
            params = append(params, f.Error())
            continue
        }
        if f, ok := arg.(stringFunc); ok {
            params = append(params, f.String())
            continue
        }

        if arg == nil {
            params = append(params, []byte("null"))
            continue
        }

        kind := reflect.TypeOf(arg).Kind()
        if kind == reflect.Ptr {
            kind = reflect.TypeOf(arg).Elem().Kind()
        }
        if kind == reflect.Struct || kind == reflect.Interface ||
            kind == reflect.Array || kind == reflect.Map || kind == reflect.Slice {
            out, err := json.Marshal(arg)
            if err != nil {
                params = append(params, arg)
            } else {
                params = append(params, out)
            }
            continue
        }

        params = append(params, arg)
    }
    glog.WarningDepth(1, fmt.Sprintf(format, params...))
}

type errorFunc interface {
    Error() string
}
type stringFunc interface {
    String() string
}

func SetV(level int32) {
    glog.SetV(glog.Level(level))
}

func GetV() int32 {
    return int32(glog.GetV())
}
