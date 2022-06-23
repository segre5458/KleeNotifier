package slack

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type SlackHandler struct {
	webhook string
}

func New(webhook string) SlackHandler {
	return SlackHandler{webhook}
}

func (s SlackHandler) Send(web WebHook) error {
	if s.webhook == "" {
		return fmt.Errorf("webhook is not invalid")
	}
	bData, err := json.Marshal(web)
	if err != nil {
		return err
	}
	_, err = http.Post(s.webhook, "application/json", bytes.NewBuffer(bData))
	return err
}
