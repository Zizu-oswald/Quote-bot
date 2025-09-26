package mymemory

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type mymemResp struct {
	ResponseData struct {
		TranslatedText string `json:"translatedText"`
	} `json:"responseData"`
}

func TranslEngToRus(str string) (string, error){
	str = strings.ReplaceAll(str, " ", "%20")
	url := "https://api.mymemory.translated.net/get?q=%s&langpair=en|ru"
	resp, err := http.Get(fmt.Sprintf(url, str))
	if err != nil {
	  return "", fmt.Errorf("error with mymemory translation [en -> ru]: %e", err)
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