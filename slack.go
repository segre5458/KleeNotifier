package main

import (
	"time"

	"github.com/segre5458/KleeNotifier/slack"
)

func SendMessage() error {
	var text = GetText(time.Now()) 
	var message = slack.WebHook{
		Text: text,
		Icon_emoji: ":smile:",
	}

	return Slack.Send(message)
}