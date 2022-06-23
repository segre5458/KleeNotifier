package main

import (
	"os"

	"github.com/segre5458/KleeNotifier/slack"
)

var Slack slack.SlackHandler

func init(){
	var hook = os.Getenv("KleeHookURL")
	if hook == ""{
		panic("Hook URL is not invalid")
	}
	Slack = slack.New(hook)
}

func main() {
	SendMessage()
}