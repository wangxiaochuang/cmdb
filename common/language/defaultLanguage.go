package language

// ccDefaultLanguageHelper regular language code helper
type ccDefaultLanguageHelper struct {
    languageType string
    languageStr  func(language, key string) string
    languageStrf func(language, key string, args ...interface{}) string
}

func (cli *ccDefaultLanguageHelper) Language(key string) string {
    ret := cli.languageStr(cli.languageType, key)
    return ret
}

func (cli *ccDefaultLanguageHelper) Languagef(key string, args ...interface{}) string {
    return cli.languageStrf(cli.languageType, key, args...)
}
