/*
   Create: 2024/1/15
   Project: ApolloCLI
   Github: https://github.com/landers1037
   Copyright Renj
*/

package http

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	url2 "net/url"
)

// 生成统一的client
const (
	GET    = "GET"
	POST   = "POST"
	DELETE = "DELETE"
)

const (
	ApolloXApi = "http://127.0.0.1:9090"
)

func Get(url string, query map[string]string) (res []byte, err error) {
	return Send(GET, url, nil, query)
}

func Post(url string, body []byte, query map[string]string) (res []byte, err error) {
	return Send(POST, url, body, query)
}

func Send(method, url string, body []byte, query map[string]string) (res []byte, err error) {
	var tempUrl = fmt.Sprintf("%s%s", ApolloXApi, url)
	if query != nil {
		params := url2.Values{}
		for key, val := range query {
			params.Add(key, val)
		}
		tempUrl = fmt.Sprintf("%s?%s", tempUrl, params.Encode())
	}

	req, err := http.NewRequest(method, tempUrl, bytes.NewBuffer(body))
	if err != nil {
		panic(err)
	}

	// 设置HTTP请求头
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set(XLocal, YES)
	client := &http.Client{}
	data, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer data.Body.Close()

	return io.ReadAll(data.Body)
}
