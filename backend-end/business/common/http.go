package common

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"time"
)

var (
	ErrInvalidResponse = errors.New("invalid response")
	ErrRequestFailed   = errors.New("request failed")
)

type Client struct {
	client    *http.Client
	baseURL   string
	headers   map[string]string
	transport *http.Transport
}

// NewClient 创建新的HTTP客户端
func NewClient() *Client {
	return &Client{
		client: &http.Client{
			Timeout: 30 * time.Second,
		},
		headers:   make(map[string]string),
		transport: &http.Transport{},
	}
}

// SetTimeout 设置请求超时时间
func (c *Client) SetTimeout(timeout time.Duration) {
	c.client.Timeout = timeout
}

// SetBaseURL 设置基础URL
func (c *Client) SetBaseURL(baseURL string) {
	c.baseURL = baseURL
}

// AddHeader 添加公共请求头
func (c *Client) AddHeader(key, value string) {
	c.headers[key] = value
}

// SkipTLSVerify 跳过TLS验证（仅测试环境使用）
func (c *Client) SkipTLSVerify(skip bool) {
	c.transport.TLSClientConfig.InsecureSkipVerify = skip
	c.client.Transport = c.transport
}

// Get 发送GET请求
func (c *Client) Get(path string, queryParams map[string]string) ([]byte, int, error) {
	reqURL := c.buildURL(path)
	reqURL = c.addQueryParams(reqURL, queryParams)
	return c.sendRequest("GET", reqURL, nil, "")
}

// Post 发送表单POST请求
func (c *Client) Post(path string, formData map[string]string) ([]byte, int, error) {
	reqURL := c.buildURL(path)
	data := url.Values{}
	for k, v := range formData {
		data.Set(k, v)
	}
	return c.sendRequest("POST", reqURL, strings.NewReader(data.Encode()), "application/x-www-form-urlencoded")
}

// PostJSON 发送JSON POST请求
func (c *Client) PostJSON(path string, payload interface{}) ([]byte, int, error) {
	reqURL := c.buildURL(path)
	jsonData, err := json.Marshal(payload)
	if err != nil {
		return nil, 0, fmt.Errorf("json marshal error: %w", err)
	}
	return c.sendRequest("POST", reqURL, bytes.NewReader(jsonData), "application/json")
}

// UploadFile 上传文件
func (c *Client) UploadFile(path, filePath, fieldName string, extraParams map[string]string) ([]byte, int, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, 0, fmt.Errorf("open file error: %w", err)
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	// 添加文件
	part, err := writer.CreateFormFile(fieldName, filepath.Base(filePath))
	if err != nil {
		return nil, 0, fmt.Errorf("create form file error: %w", err)
	}
	_, err = io.Copy(part, file)
	if err != nil {
		return nil, 0, fmt.Errorf("copy file error: %w", err)
	}

	// 添加额外参数
	for k, v := range extraParams {
		_ = writer.WriteField(k, v)
	}

	err = writer.Close()
	if err != nil {
		return nil, 0, fmt.Errorf("close writer error: %w", err)
	}

	reqURL := c.buildURL(path)
	return c.sendRequest("POST", reqURL, body, writer.FormDataContentType())
}

// 发送请求通用方法
func (c *Client) sendRequest(method, reqURL string, body io.Reader, contentType string) ([]byte, int, error) {
	req, err := http.NewRequest(method, reqURL, body)
	if err != nil {
		return nil, 0, fmt.Errorf("create request error: %w", err)
	}

	// 设置公共请求头
	for k, v := range c.headers {
		req.Header.Set(k, v)
	}

	// 设置内容类型
	if contentType != "" {
		req.Header.Set("Content-Type", contentType)
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, 0, fmt.Errorf("do request error: %w", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, resp.StatusCode, fmt.Errorf("read response error: %w", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return respBody, resp.StatusCode, fmt.Errorf("%w: status code %d", ErrRequestFailed, resp.StatusCode)
	}

	return respBody, resp.StatusCode, nil
}

// 构建完整URL
func (c *Client) buildURL(path string) string {
	if c.baseURL == "" {
		return path
	}
	return strings.TrimRight(c.baseURL, "/") + "/" + strings.TrimLeft(path, "/")
}

// 添加查询参数
func (c *Client) addQueryParams(rawURL string, params map[string]string) string {
	if len(params) == 0 {
		return rawURL
	}

	u, err := url.Parse(rawURL)
	if err != nil {
		return rawURL
	}

	q := u.Query()
	for k, v := range params {
		q.Add(k, v)
	}
	u.RawQuery = q.Encode()

	return u.String()
}

// ParseJSONResponse 解析JSON响应到结构体
func ParseJSONResponse(data []byte, v interface{}) error {
	if err := json.Unmarshal(data, v); err != nil {
		return fmt.Errorf("json unmarshal error: %w", err)
	}
	return nil
}
