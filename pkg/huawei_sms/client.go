package huaweisms

import (
	"bytes"
	"io"
	"net/http"
	"net/http/cookiejar"
	"strings"
)

// HuaweiClient 华为路由器客户端
type HuaweiClient struct {
	BaseURL  string
	Username string
	Password string
	Token    string
	Client   *http.Client
}

// NewHuaweiClient 创建新的客户端实例
func NewHuaweiClient(baseURL, username, password string) *HuaweiClient {
	jar, _ := cookiejar.New(nil)

	return &HuaweiClient{
		BaseURL:  strings.TrimRight(baseURL, "/"),
		Username: username,
		Password: password,
		Client: &http.Client{
			Jar: jar,
		},
	}
}

// get 执行 GET 请求
func (h *HuaweiClient) get(url string, headers map[string]string) ([]byte, *http.Response, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, nil, err
	}
	for k, v := range headers {
		req.Header.Set(k, v)
	}

	resp, err := h.Client.Do(req)
	if err != nil {
		return nil, nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	return body, resp, err
}

// post 执行 POST 请求
func (h *HuaweiClient) post(url string, data []byte, headers map[string]string) ([]byte, *http.Response, error) {
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(data))
	if err != nil {
		return nil, nil, err
	}
	for k, v := range headers {
		req.Header.Set(k, v)
	}

	resp, err := h.Client.Do(req)
	if err != nil {
		return nil, nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	return body, resp, err
}
