package login

import (
	"github.com/mmzou/geektime-dl/requester"
)

//Client login client
type Client struct {
	*requester.HTTPClient
}

// Result 从百度服务器解析的数据结构
type Result struct {
	Code int `json:"code"`
	Data struct {
		UID          int    `json:"uid"`
		Name         string `json:"nickname"`
		Avatar       string `json:"avatar"`
		GCID         string `json:"gcid"`
		GCESS        string `json:"gcess"`
		ServerID     string `json:"serverId"`
		Ticket       string `json:"ticket"`
		CookieString string `json:"cookieString"`
	} `json:"data"`
	Error struct {
		Code int    `json:"code"`
		Msg  string `json:"msg"`
	} `json:"error"`
	Extra struct {
		Cost      float64 `json:"cost"`
		RequestID string  `json:"request-id"`
	} `json:"extra"`
}

func NewLoginClient() *Client {
	c := &Client{
		HTTPClient: requester.NewHTTPClient(),
	}
	c.InitLoginPage()

	return c
}

//InitLoginPage init
func (c *Client) InitLoginPage() {
	res, _ := c.Get("https://account.geekbang.org/signin?redirect=https%3A%2F%2Ftime.geekbang.org%2F")
	defer res.Body.Close()
}
