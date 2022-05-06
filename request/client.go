package request

import (
	"net/http"
	"net/http/cookiejar"
	"time"
)

type HttpClient struct {
	http.Client
	Agent string
}

func NewClient() *HttpClient {
	c := &HttpClient{
		Client: http.Client{
			Timeout: 10 * time.Second,
		},
	}
	return c
}

func (h *HttpClient) SetAgent(ua string) {
	h.Agent = ua
}

func (h *HttpClient) SetCookie(jar http.CookieJar) {
	h.Client.Jar = jar
}

func (h *HttpClient) ResetCookie() {
	h.Jar, _ = cookiejar.New(nil)
}
