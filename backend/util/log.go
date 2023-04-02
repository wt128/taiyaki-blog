package util

import (
	"encoding/json"
	"net/http"
	"net/url"
	"os"

	"github.com/joho/godotenv"
)

type payload struct {
	Text string `json:"text"`
}

func ErrorNotice(errMessage error) (err error) {
	p, err := json.Marshal(payload{Text: (errMessage.Error())})
	if err != nil {
		panic(err)
	}
	// envファイルのパスを渡す。何も渡さないと、どうディレクトリにある、.envファイルを探す
	godotenv.Load()
	webhookURL := os.Getenv("SLACK_WEBHOOK")
	if err != nil {
		panic("Error loading .env file")
	}
	resp, err := http.PostForm(webhookURL, url.Values{"payload": {string(p)}})
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}