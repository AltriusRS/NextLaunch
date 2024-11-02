package translations

import (
	"github.com/pelletier/go-toml/v2"
	"os"
)

type LanguagePacket struct {
	Code   string            `toml:"code" comment:"Language code - These are defined in ISO 639-1 and should be lowercase (eg: en, de, es, etc.) followed by an underscore and the country code in line with ISO 3166-1 Alpha 3 codes (eg: en_USA, en_GBR, es_ESP, es_MEX, etc.)"`
	Name   string            `toml:"name" comment:"Language name (eg: British English, American English, German, Spanish (Spain), etc.)"`
	Native string            `toml:"native" comment:"Language native name (eg: English, English for idiots, Deutsch, Español, etc.)"`
	Keys   map[string]string `json:"keys"`
}

func LoadLanguage(path string) (LanguagePacket, error) {
	var language LanguagePacket
	file, err := os.Open(path)
	if err != nil {
		return language, err
	}
	defer file.Close()
	decoder := toml.NewDecoder(file)
	err = decoder.Decode(&language)
	if err != nil {
		return language, err
	}
	return language, nil
}

func GetFiles(directory string) ([]string, error) {
	files, err := os.ReadDir(directory)
	if err != nil {
		return nil, err
	}
	var result []string
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		result = append(result, file.Name())
	}
	return result, nil
}
