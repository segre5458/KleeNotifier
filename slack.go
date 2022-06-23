package main

import "github.com/segre5458/KleeNotifier/slack"

func SendMessage() error {
	var message = slack.WebHook{
		Text: "hoge",
		Channel: "#segre-private-memo",
	}

	return Slack.Send(message)
}