package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type githubEvent struct{
	CreatedAt string `json:"created_at"`
}

func main(){
	req, err := http.NewRequest(http.MethodGet, "https://api.github.com/users/hoashi-akane/events/public", nil)
	if err != nil{
		fmt.Print("エラー")
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Print("エラー")
	}

	body, err := ioutil.ReadAll(resp.Body)
	var hub []githubEvent
	err = json.Unmarshal(body, &hub)
	if err != nil{
		fmt.Print("エラー")
	}
	c := utcToJst(hub)
	if c == 0{
		fmt.Print("なし")
	}else{
		fmt.Print(c)
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