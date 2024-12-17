package translations

import (
	"Nextlaunch/src/config"
	"Nextlaunch/src/logging"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
)

type TranslationManager struct {
	*logging.Logger
	translations map[string]LanguagePacket
	code         string
}

var Manager *TranslationManager
var translationDir string

func init() {
	Manager = NewTranslationManager()

	Manager.Errorf("\x1b[41m\x1b[30mThis is a test message\x1b[0m")

	appDir, err := os.UserConfigDir()
	if err != nil {
		log.Fatal(err)
	}

	translationDir = appDir + "/translations"

}

func NewTranslationManager() *TranslationManager {
	if Manager == nil {

		logger := logging.NewLogger("Translations")

		Manager = &TranslationManager{
			Logger:       logger,
			translations: make(map[string]LanguagePacket),
			code:         "en-US",
		}
	}

	return Manager
}

func (t *TranslationManager) Synchronize() {
	// get the latest translation files from the API
	res, err := http.Get("https://api.nextlaunch.org/translations/" + config.Version)
	if err != nil {
		panic(err)
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err)
		}
	}(res.Body)

	var translations Translations

	err = json.NewDecoder(res.Body).Decode(&translations)
	if err != nil {
		panic(err)
	}

	for _, translation := range translations.Files {
		t.Debugf("Checking for translation: %s", translation)
		getFileSha256(translation.Code)

	}
}

func (t *TranslationManager) Add(language LanguagePacket) {
	t.translations[language.Code] = language
}

func (t *TranslationManager) Get(code string) LanguagePacket {
	return t.translations[code]
}

func (t *TranslationManager) SetCode(code string) {
	t.code = code
}

func (t *TranslationManager) GetCode() string {
	return t.code
}

func (t *TranslationManager) GetCurrent() LanguagePacket {
	return t.translations[t.code]
}

func (t *TranslationManager) Translate(key string) string {
	//current := t.GetCurrent()
	//if current.Keys[key] != nil {
	//	return current.Keys[key]
	//}
	return key
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

func getFileSha256(path string) string {
	hasher := sha256.New()

	file, err := os.Open(path) // Open the binary file
	if err != nil {
		println("Error opening binary file")
		return ""
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
		}
	}(file)

	_, err = io.Copy(hasher, file)
	if err != nil {
		return ""
	}

	return hex.EncodeToString(hasher.Sum(nil))
}
