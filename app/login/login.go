package login

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

//定义返回结构体
type Result struct {
	Code int `json:"code"`
	Data struct {
		Uid          int    `json:"uid"`
		Nickname     string `json:"nickname"`
		Avatar       string `json:"avatar"`
		Gcid         string `json:"gcid"`
		Gcess        string `json:"gcess"`
		ServerId     string `json:"serverId"`
		Ticket       string `json:"ticket"`
		CookieString string `json:"cookieString"`
	} `json:"data"`
	Error struct {
		Code int    `json:"code"`
		Msg  string `json:"msg"`
	} `json:"error"`
	Extra struct {
		Cost      float32 `json:"cost"`
		RequestId string  `json:"requestId"`
	} `json:"extra"`
}

func init() {
	fmt.Println("检查是否登录")

	//初始化
	geekLoginInit()
}

func geekLoginInit() {
	res, _ := http.Get("https://account.geekbang.org/signin?redirect=https%3A%2F%2Ftime.geekbang.org%2F")
	defer res.Body.Close()
}

func GeekLogin() {
	//开始登录

	data := make(map[string]interface{})
	data["appid"] = "1"
	data["cellphone"] = "13266503566"
	data["country"] = "86"
	data["data"] = "1"
	data["password"] = "2121fdafdas"
	data["platform"] = "3"

	bytesData, _ := json.Marshal(data)

	print(string(bytesData))
	// header := map[string]string{
	// 	"Referer":    "https://account.geekbang.org/signin?redirect=https%3A%2F%2Ftime.geekbang.org%2F",
	// 	"Accept":     "application/json",
	// 	"Connection": "keep-alive",
	// }

	res, _ := http.Post("https://account.geekbang.org/account/ticket/login", "application/json", bytes.NewReader(bytesData))
	body, _ := ioutil.ReadAll(res.Body)
	fmt.Print(string(body))
}
