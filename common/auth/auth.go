package auth

import (
    "strconv"
    "sync"

    "github.com/wxc/cmdb/common/blog"
)

var EnableAuth = "true"
var enableAuth = true
var EnableAuthFlag *authValue
var once = sync.Once{}

type authValue struct{}

func (a *authValue) String() string {
    return strconv.FormatBool(enableAuth)
}

func (a *authValue) Set(s string) error {
    v, err := strconv.ParseBool(s)
    if err != nil {
            return err
    }
    setEnableAuth(v)
    return nil
}

func (a *authValue) Type() string {
    return "bool"
}

func init() {
    var err error
    enableAuth, err = strconv.ParseBool(EnableAuth)
    if err != nil {
            blog.Fatalf("[auth] enableAuth %s configuration invalid", EnableAuth)
    }
}

func setEnableAuth(enable bool) {
    once.Do(func() {
        enableAuth = enable
    })
}

func EnableAuthorize() bool {
    return enableAuth
}
