package util

import (
	"encoding/json"
	"net/http"
	"net/url"
)

type payload struct {
	Text string `json:"text"`
}

func ErrorNotice(errMessage error) (err error){
	p, err := json.Marshal(payload{Text: (errMessage.Error())})
	if err != nil {
			panic(err)
	}
	webhookURL := env('SLACK_WEBHOOK')
	resp, err := http.PostForm(webhookURL, url.Values{"payload": {string(p)}})
	if err != nil {
			return err
	}
	defer resp.Body.Close()
	return nil
}