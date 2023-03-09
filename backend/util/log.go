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
	webhookURL := "https://hooks.slack.com/services/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"
	resp, err := http.PostForm(webhookURL, url.Values{"payload": {string(p)}})
	if err != nil {
			return err
	}
	defer resp.Body.Close()
	return nil
}