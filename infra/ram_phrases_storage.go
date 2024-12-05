package infra

import "fmt"

var phrases_storage []string

type RamPhrasesStorage struct{}

func (self *RamPhrasesStorage) SavePhrase(phrase string) {
	phrases_storage = append(phrases_storage, phrase)
	fmt.Println(phrases_storage)
}

func CreateRamPhrasesStorage() *RamPhrasesStorage {
	return &RamPhrasesStorage{}
}
