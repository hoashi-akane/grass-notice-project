package slack

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type Slack struct{
	Channel string `json:"channel"`
	Username string `json:"username"`
	Text string `json:"text"`
	IconEmoji string `json:"icon_emoji"`
	IconURL string `json:"icon_url"`
}

func GoSlack(c int)(resp *http.Response, err error){
	var text string
 	if c == 0{
 		text = "草はやせ？？？ :gopher:"
	}else{
		text = fmt.Sprintf("今日は%d回芝に貢献しました。:kusa:", c)
	}
	s := Slack{
		Channel: "",
		Username: "芝生BOT",
		Text: text,
		IconEmoji: ":gopher",
		IconURL: "",
	}
	jsonparam, _ := json.Marshal(s)
	resp, err = http.PostForm(WEBHOOKURL, url.Values{"payload": {string(jsonparam)}},)
	defer resp.Body.Close()
	return resp, err
}