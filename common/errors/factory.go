package errors

var EmptyErrorsSetting = map[string]ErrorCode{}

// NewFactory create new CCErrorIf instance,
// dir is directory of errors description resource
func NewFactory(dir string) (CCErrorIf, error) {

    tmp := &ccErrorHelper{errCode: make(map[string]ErrorCode)}

    errcode, err := LoadErrorResourceFromDir(dir)
    if nil != err {
        //blog.Errorf("failed to load the error resource, error info is %s", err.Error())
        return nil, err
    }
    tmp.Load(errcode)

    return tmp, nil
}

func NewFromCtx(errcode map[string]ErrorCode) CCErrorIf {
    tmp := &ccErrorHelper{}
    tmp.Load(errcode)
    return tmp
}

func New(errCode int, errMsg string) CCErrorCoder {
    return &ccError{code: errCode, callback: func() string {
        return errMsg
    }}
}
