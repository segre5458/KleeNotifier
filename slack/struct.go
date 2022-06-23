package slack

type WebHook struct {
	Text       string `json:"text"`
	UserName   string `json:"username,omitempty"`
	Icon_emoji string `json:"icon_emoji,omitempty"`
	Icon_url   string `json:"icon_url,omitempty"`
	Channel    string `json:"channel,omitempty"`
}
