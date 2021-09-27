package language

// DefaultCCLanguageIf defines default language interface
type DefaultCCLanguageIf interface {
    // Language returns an content with key
    Language(key string) string
    // Errorf returns an content with key
    Languagef(key string, args ...interface{}) string
}

// CCLanguageIf defines error information conversion
type CCLanguageIf interface {
    // CreateDefaultCCLanguageIf create new language error interface instance
    CreateDefaultCCLanguageIf(language string) DefaultCCLanguageIf
    // Language returns an content with key
    Language(language, key string) string
    // Languagef returns an content with key
    Languagef(language, key string, args ...interface{}) string

    Load(res map[string]LanguageMap)
}
