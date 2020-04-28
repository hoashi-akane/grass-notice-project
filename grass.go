package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type githubEvent struct{
	CreatedAt string `json:"created_at"`
}
func main(){
	req, err := http.NewRequest(http.MethodGet, "https://api.github.com/users/hoashi-akane/events/public", nil)
	if err != nil{
		fmt.Print("エラ")
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Print("エラー")
	}

	body, err := ioutil.ReadAll(resp.Body)
	var hub []githubEvent
	json_err := json.Unmarshal(body, &hub)
	if json_err != nil{
		fmt.Print("エラー")
	}
	for k, v := range hub{
		fmt.Print(k, v)
	}
}