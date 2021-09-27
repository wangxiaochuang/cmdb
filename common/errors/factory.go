package errors

func New(errCode int, errMsg string) CCErrorCoder {
    return &ccError{code: errCode, callback: func() string {
        return errMsg
    }}
}
