package mymemory

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type mymemResp struct {
	ResponseData struct {
		TranslatedText string `json:"translatedText"`
	} `json:"responseData"`
}

func TranslEngToRus(str string) (string, error) {
	base := "https://api.mymemory.translated.net/get"

	// для точечного порядка элементов тк в api mymemory это важно
	query := "q=" + url.QueryEscape(str) + "&langpair=" + url.QueryEscape("en|ru")
	fullURL := base + "?" + query


	resp, err := http.Get(fullURL)

	if err != nil {
		return "", fmt.Errorf("error with mymemory translation [en -> ru]: %s %e", fullURL, err)
	}
	if resp.StatusCode != 200 {
		return "", fmt.Errorf("mymemory translation [en -> ru] with status code: %d", resp.StatusCode)
	}

	jsonResp, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("cant read body of response: %e", err)
	}
	var structResp mymemResp
	err = json.Unmarshal(jsonResp, &structResp)
	if err != nil {
		return "", fmt.Errorf("cant unmarshall json: %e", err)
	}

	return structResp.ResponseData.TranslatedText, nil
}
