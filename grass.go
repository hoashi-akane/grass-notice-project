package main

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
	"io/ioutil"
	"log"
	"net/http"
	"slack"
	"time"
)

type githubEvent struct{
	CreatedAt string `json:"created_at"`
}

func main(){
	lambda.Start(handler)
}

func handler(){
	body,err := getResponse("https://api.github.com/users/hoashi-akane/events/public")
	if err != nil{
		log.Print("リクエスト・レスポンスエラー")
	}
	var hub []githubEvent
	err = json.Unmarshal(body, &hub)
	if err != nil{
		log.Print("エラー")
	}
	c := utcToJst(hub)
	resp, err := slack.GoSlack(c)
	if err != nil{log.Print("エラー")}
	if resp.StatusCode == 200{

	}
}
// 関数名(引数)(戻り値errも渡せる)
func utcToJst(hub []githubEvent)(c int) {
	loc, _ := time.LoadLocation("Asia/Tokyo")
	now := time.Now()
	now = now.Truncate( time.Hour ).Add( - time.Duration(now.Hour()) * time.Hour )
	// kは要素番号 vは内容 (key,value)
	for _, v := range hub {
		// JSTで作成（まだ表示される時間はUTC)
		t, e := time.ParseInLocation("2006-01-02T15:04:05Z", v.CreatedAt, loc)
		if e != nil { fmt.Print(e) }
		// 時間をJSTにする
		t = t.Add(9 * time.Hour)
		if t.After(now){
			c++
		}else {
			break
		}
	}
	return c
}

// リクエスト飛ばしてレスポンス受け取る
func getResponse(url string)(body []byte, err error){
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil{
	log.Print("エラー")
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
	log.Print("エラー")
	}
	body, err = ioutil.ReadAll(resp.Body)
	return body, err
}