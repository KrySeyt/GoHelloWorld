package app

import (
	"fmt"
	"packname/app/interfaces"
)

type SayHello struct {
	phrasesStorage     interfaces.PhrasesStorage
	phrasesDictionary  interfaces.PhrasesDictionary
	transactionManager interfaces.TransactionManager
}

func (self *SayHello) Execute(name string) {
	self.transactionManager.Begin()

	phrase := fmt.Sprintf("Hello, %s!", name)
	
	self.phrasesStorage.SavePhrase(phrase)
	self.phrasesDictionary.AddPhrase(phrase)

	self.transactionManager.Commit()
}

func CreateSayHello(
	phrasesStorage interfaces.PhrasesStorage,
	phrasesDictionary  interfaces.PhrasesDictionary,
	transactionManager interfaces.TransactionManager,
) *SayHello {
	return &SayHello{
		phrasesStorage:     phrasesStorage,
		phrasesDictionary: phrasesDictionary,
		transactionManager: transactionManager,
	}
}
