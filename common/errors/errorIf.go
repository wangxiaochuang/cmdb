package errors

type CCErrorCoder interface {
    error
    GetCode() int
}

type DefaultCCErrorIf interface {
    // Error returns an error with error code
    Error(errCode int) error
    // Errorf returns an error with error code
    Errorf(errCode int, args ...interface{}) error

    // CCError returns an error with error code
    CCError(errCode int) CCErrorCoder
    // CCErrorf returns an error with error code
    CCErrorf(errCode int, args ...interface{}) CCErrorCoder

    // New create a new error with error code and message
    New(errorCode int, msg string) error
}

type CCErrorIf interface {
    // CreateDefaultCCErrorIf create new language error interface instance
    CreateDefaultCCErrorIf(language string) DefaultCCErrorIf
    // Error returns an error for specific language
    Error(language string, errCode int) error
    // Errorf Errorf returns an error with args for specific language
    Errorf(language string, errCode int, args ...interface{}) error

    Load(res map[string]ErrorCode)
}

func NewFromStdError(err error, defaultErrCode int) CCErrorCoder {
    ccErr, ok := err.(CCErrorCoder)
    if ok == true {
            return ccErr
    }
    return New(defaultErrCode, err.Error())
}

// globalCCError 代表从zk中读取到的error配置，
// 结合 utils.GetDefaultCCError 使用即可实现国际化
// 设计背景：用于减少不必要的参数传递
var globalCCError CCErrorIf

// GetGlobalCCError return CCErrorIf from zk, please use SetGlobalCCError initialize it before use
// or check nil
func GetGlobalCCError() CCErrorIf {
    return globalCCError
}

func SetGlobalCCError(ccError CCErrorIf) {
    globalCCError = ccError
}

var (
    // 1199074 is CCErrCommGlobalCCErrorNotInitialized actually
    GlobalCCErrorNotInitialized = New(1199074, "global cc error not initialized")
)

type RawErrorInfo struct {
    ErrCode int
    Args    []interface{}
}

func (rei *RawErrorInfo) ToCCError(ccErrorIF DefaultCCErrorIf) CCErrorCoder {
    if rei.ErrCode == 0 {
        return nil
    }
    if len(rei.Args) == 0 {
        return ccErrorIF.CCError(rei.ErrCode)
    }
    return ccErrorIF.CCErrorf(rei.ErrCode, rei.Args...)
}
