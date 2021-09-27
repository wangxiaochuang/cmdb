package errors

// UnknownTheCodeStrf const define unknow code
const UnknownTheCodeStrf = "the error code is '%d', unknown meaning"

// UnknownTheLanguageStrf define unknow language
const UnknownTheLanguageStrf = "the language code is '%s', unknown meaning"

// defaultLanguage default language package name
const defaultLanguage = "default"

type ErrorCode map[string]string

type CCError interface {
    error
}

// ccError CC custom error define
type ccError struct {
    code     int
    callback func() string
}
