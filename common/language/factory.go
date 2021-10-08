package language

var EmptyLanguageSetting = map[string]LanguageMap{}

func NewFromCtx(lang map[string]LanguageMap) CCLanguageIf {
    tmp := &ccLanguageHelper{}
    tmp.Load(lang)
    return tmp
}
