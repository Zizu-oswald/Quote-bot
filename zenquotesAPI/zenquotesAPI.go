package zenquotesapi

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type quoteJSON []struct {
	Q string `json:"q"`
	A string `json:"a"`
}

type Quote struct {
	Quote  string
	Author string
}

func (q *Quote) IntoString(lang string) string {
	switch lang {
	case "en":
		return fmt.Sprintf("Quote: %s\nAuthor: %s", q.Quote, q.Author)
	case "ru":
		// runes := []rune(q.Author)
		// q.Author = string(runes[1:])
		return fmt.Sprintf("Цитата: %s\nАвтор: %s", q.Quote, q.Author)
	}
	return ""
}

func GetRandomQuote() (Quote, error) {
	url := "https://zenquotes.io/api/random"
	resp, err := http.Get(url)
	if err != nil || resp.StatusCode != 200 {
		return Quote{}, err
	}

	var quote quoteJSON
	err = json.NewDecoder(resp.Body).Decode(&quote)
	if err != nil {
		return Quote{}, err
	}
	return Quote{Quote: quote[0].Q, Author: quote[0].A}, nil
}
