package infra

import (
	"fmt"
	"net/url"

	"github.com/go-resty/resty/v2"
)

type HttpPhrasesDictionary struct {
	client         *resty.Client
	dictionaryHost string
}

func (self *HttpPhrasesDictionary) AddPhrase(phrase string) {
	url, err := url.JoinPath(self.dictionaryHost, "/post")
	if err != nil {
		panic(err)
	}

	response, err := self.client.R().Post(url)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Phrases dictionary response status code: %d. URL: %s", response.StatusCode(), url)
}

func CreateHttpPhrasesDictionary(client *resty.Client, dictionaryHost string) *HttpPhrasesDictionary {
	return &HttpPhrasesDictionary{
		client:         client,
		dictionaryHost: dictionaryHost,
	}
}
