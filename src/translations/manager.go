package translations

type TranslationManager struct {
	translations map[string]LanguagePacket
}

func NewTranslationManager() *TranslationManager {
	return &TranslationManager{
		translations: make(map[string]LanguagePacket),
	}
}

func (t *TranslationManager) Add(language LanguagePacket) {
	t.translations[language.Code] = language
}

func (t *TranslationManager) Get(code string) LanguagePacket {
	return t.translations[code]
}

func (t *TranslationManager) GetAll() []LanguagePacket {
	var languages []LanguagePacket
	for _, language := range t.translations {
		languages = append(languages, language)
	}
	return languages
}

func (t *TranslationManager) LoadFromDirectory(directory string) {
	files, err := GetFiles(directory)
	if err != nil {
		return
	}
	for _, file := range files {
		language, err := LoadLanguage(file)
		if err != nil {
			continue
		}
		t.Add(language)
	}
}
