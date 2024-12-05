package app

import (
	"packname/app/interfaces"
)

type GetHello struct {
	phrasesStorage interfaces.PhrasesStorage
}

func (self *GetHello) Execute(name string) string {
	return self.phrasesStorage.GetPhrase(name)
}

func CreateGetHello(
	phrasesStorage interfaces.PhrasesStorage,
) *GetHello {
	return &GetHello{
		phrasesStorage: phrasesStorage,
	}
}
