package youdao

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type ComplexResponse struct {
	Input string `json:"input"`
	Meta  struct {
		Input           string   `json:"input"`
		GuessLanguage   string   `json:"guessLanguage"`
		IsHasSimpleDict string   `json:"isHasSimpleDict"`
		Le              string   `json:"le"`
		Lang            string   `json:"lang"`
		Dicts           []string `json:"dicts"`
	} `json:"meta"`
	Fanyi struct {
		Voice string `json:"voice"`
		Input string `json:"input"`
		Type  string `json:"type"`
		Tran  string `json:"tran"`
	} `json:"fanyi"`
	Le   string `json:"le"`
	Lang string `json:"lang"`
}

func CallComplex(q string) (*ComplexResponse, error) {
	response := new(ComplexResponse)

	api := "https://dict.youdao.com/jsonapi_s?doctype=json&jsonversion=4"

	// 构建POST请求体
	postData := url.Values{}
	postData.Set("q", q) // 注意：这里我们发送了原始的字符串，不是URL编码的
	postData.Set("le", "en")
	postData.Set("t", "4")
	postData.Set("client", "web")
	postData.Set("sign", "51685965cf2d62f61fdac918eb019c4b") // 假设这是有效的签名
	postData.Set("keyfrom", "webdict")

	// 将url.Values编码为字符串
	postBody := bytes.NewBufferString(postData.Encode())

	// 创建HTTP客户端
	client := &http.Client{}

	// 创建HTTP请求
	req, err := http.NewRequest("POST", api, postBody)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return response, err
	}

	// 设置HTTP头部
	req.Header.Set("sec-ch-ua", `"Chromium";v="124", "Google Chrome";v="124", "Not-A.Brand";v="99"`)
	req.Header.Set("Accept", "application/json, text/plain, */*")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Referer", "https://www.youdao.com/")
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/124.0.0.0 Safari/537.36")
	req.Header.Set("sec-ch-ua-platform", `"macOS"`)

	// 发送请求并获取响应
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return response, err
	}
	defer resp.Body.Close()

	// 读取响应体
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return response, err
	}

	err = json.Unmarshal([]byte(body), &response)
	if err != nil {
		return response, err
	}
	return response, nil
}
