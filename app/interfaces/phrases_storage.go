package interfaces

type PhrasesStorage interface {
	SavePhrase(phrase string)
	GetPhrase(target_substr string) string
}
