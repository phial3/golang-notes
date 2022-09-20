package utils

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type JsonHttpClientHelper struct {
	client *http.Client
}

func NewHttpRequest() *JsonHttpClientHelper {
	return &JsonHttpClientHelper{
		client: &http.Client{},
	}
}

func NewHttpRequestFromClient(client *http.Client) *JsonHttpClientHelper {
	return &JsonHttpClientHelper{
		client: client,
	}
}

// Get
//
//	@Description: 发起Get请求
//	@receiver clientHelper
//	@param urlPath url请求绝对路径 https://www.***.com/users
//	@param headers http请求头
//	@param params  url中附带的参数信息，最后会拼接到绝对路径上 如： https://www.***.com/users?name=u&age=18
//	@return string 请求结果的字符串，推荐为 json格式
//	@return error  请求发生的错误，需要捕获
func (clientHelper *JsonHttpClientHelper) Get(urlPath string, headers, params map[string]string) (string, error) {
	//init url params
	requestUrl, err := makeRequestUrl(urlPath, params)
	if err != nil {
		return "", err
	}

	req, err := makeRequest(requestUrl, headers, "GET", nil)
	if err != nil {
		return "", err
	}

	response, err := clientHelper.client.Do(req)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	//read from response
	b, err := io.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	return string(b), nil
}

func makeRequest(requestUrl *url.URL, headers map[string]string, method string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(method, requestUrl.String(), body)
	if err != nil {
		return nil, err
	}

	//set header
	for key, value := range headers {
		req.Header.Add(key, value)
	}

	return req, nil
}

func makeRequestUrl(urlPath string, params map[string]string) (*url.URL, error) {
	urlParams := url.Values{}
	for key, value := range params {
		urlParams.Set(key, value)
	}
	requestURI, err := url.ParseRequestURI(urlPath)
	if err != nil {
		return nil, err
	}
	requestURI.RawQuery = urlParams.Encode()

	return requestURI, nil
}

// Post
//
//	@Description: 发起POST请求
//	@receiver clientHelper
//	@param urlPath url请求绝对路径 https://www.***.com/users
//	@param data 请求时需要传输的body
//	@param headers http请求头 Content-Type只能为 application/x-www-form-urlencoded 或 application/json
//	@param urlParams  url中附带的参数信息，最后会拼接到绝对路径上 如： https://www.***.com/users?name=u&age=18
//	@param bodyParams  Content-Type为application/x-www-form-urlencoded 特殊片段
//	@return string 请求结果的字符串，推荐为 json格式
//	@return error  请求发生的错误，需要捕获
func (clientHelper *JsonHttpClientHelper) Post(urlPath string, data interface{}, headers, urlParams, bodyParams map[string]string) (string, error) {
	//init url urlParams
	requestUrl, err := makeRequestUrl(urlPath, urlParams)
	if err != nil {
		return "", err
	}

	var req *http.Request
	if isJsonHttp(headers) {
		bodyJsonStr, err := convertDataToJson(data)
		if err != nil {
			return "", err
		}
		req, err = makeRequest(requestUrl, headers, "POST", strings.NewReader(bodyJsonStr))
	} else {
		bodyUrlValues := makeRequestBodyForFormUpload(bodyParams)
		req, err = makeRequest(requestUrl, headers, "POST", strings.NewReader(bodyUrlValues.Encode()))
		if err != nil {
			return "", err
		}
	}

	response, err := clientHelper.client.Do(req)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	//read from response
	b, err := io.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	return string(b), nil
}

func makeRequestBodyForFormUpload(maps map[string]string) *url.Values {
	values := url.Values{}
	if maps != nil {
		for key, value := range maps {
			values.Set(key, value)
		}
	}

	return &values
}

func isJsonHttp(headers map[string]string) bool {
	v, ok := headers["Content-Type"]
	if ok && v == "application/json" {
		return true
	}

	if !ok {
		return true
	}

	return false
}

func convertDataToJson(data interface{}) (string, error) {
	bodyJsonByte, err := json.Marshal(data)
	if err != nil {
		return "", err
	}

	bodyJsonStr := string(bodyJsonByte)
	if bodyJsonStr == "" || bodyJsonStr == "null" || bodyJsonStr == "NULL" || bodyJsonStr == "nil" {
		bodyJsonStr = "{}"
	}
	return bodyJsonStr, nil
}
