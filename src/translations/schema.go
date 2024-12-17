package translations

import (
	"gopkg.in/yaml.v3"
	"os"
)

type LanguagePacket struct {
	Code        string            `yaml:"code"`
	Name        string            `yaml:"name"`
	Native      string            `yaml:"native"`
	Periodicals Periodicals       `yaml:"periodicals"`
	Keybindings []Keybinding      `yaml:"key_bindings"`
	Interface   map[string]string `yaml:"interface"`
}

type Periodicals struct {
	Months       []string          `yaml:"months"`
	AbbrMonths   []string          `yaml:"abbreviated_months"`
	Weekdays     []string          `yaml:"weekdays"`
	AbbrWeekdays []string          `yaml:"abbreviated_weekdays"`
	StartDay     uint8             `yaml:"start_day"`
	Periods      map[string]string `yaml:"periods"`
	Formats      map[string]string `yaml:"formats"`
}

type Keybinding struct {
	Key    string `yaml:"key"`
	Action string `yaml:"action"`
}

type InterfaceText struct {
	Key   string `yaml:"key"`
	Value string `yaml:"value"`
}

func LoadLanguage(path string) (LanguagePacket, error) {
	var language LanguagePacket
	file, err := os.Open(path)
	if err != nil {
		return language, err
	}
	defer file.Close()
	decoder := yaml.NewDecoder(file)
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
