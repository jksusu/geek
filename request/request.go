package request

var (
	// UserAgent 浏览器标识
	UserAgent = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.92 Safari/537.36"
)

func (h *HttpClient) Request(mothod string, host string, data interface{}, header map[string]string) ([]byte, error) {
}
