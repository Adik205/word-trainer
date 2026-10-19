package internal

import (
	"encoding/json"
	"fmt"
	"os"
)

type Data struct {
	version int
	nextId int
	words []Word
}

type Word struct {
	Id           int      `json:"id"`
	English      string   `json:"english"`
	Russian      []string `json:"russian"`
	Correct      int      `json:"correct"`
	Incorrect    int      `json:"incorrect"`
	NeedToRepeat bool     `json:"needToRepeat"`
}

// Загружает json-данные
func LoadData(jsonContent []byte) {
	var dictionaryEntry DictionaryEntry

	err := json.Unmarshal(jsonContent, &dictionaryEntry)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Ошибка разбора JSON:", err)
		os.Exit(1)
	}

	var englishWords []string
	for i := 0; i < len(d)

	

}
