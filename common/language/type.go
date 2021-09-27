package language

// UnknownTheCodeStrf const define unknow code
const UnknownTheKeyStrf = "the key is '%v', unknown meaning"

// UnknownTheLanguageStrf define unknow language
const UnknownTheLanguageStrf = "the language code is '%s', unknown meaning"

// defaultLanguage default language package name
const defaultLanguage = "default"

// LanguageMap  mapping
type LanguageMap map[string]string

// ccError  CC custom error  defind
type ccLanguage struct {
    key      string
    callback func() string
}
