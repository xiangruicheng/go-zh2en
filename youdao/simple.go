package youdao

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type SimpleResponse struct {
	Result struct {
		Msg  string `json:"msg"`
		Code int    `json:"code"`
	} `json:"result"`
	Data struct {
		Entries []struct {
			Explain string `json:"explain"`
			Entry   string `json:"entry"`
		} `json:"entries"`
		Query    string `json:"query"`
		Language string `json:"language"`
		Type     string `json:"type"`
	} `json:"data"`
}

func CallSimple(q string) (*SimpleResponse, error) {
	response := new(SimpleResponse)
	url := "https://dict.youdao.com/suggest?num=5&ver=3.0&doctype=json&cache=false&le=en&q=" + q

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return response, nil
	}

	// 设置HTTP头部
	req.Header.Set("sec-ch-ua", `"Chromium";v="124", "Google Chrome";v="124", "Not-A.Brand";v="99"`)
	req.Header.Set("Accept", "application/json, text/plain, */*")
	req.Header.Set("Referer", "https://www.youdao.com/")
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/124.0.0.0 Safari/537.36")
	req.Header.Set("sec-ch-ua-platform", `"macOS"`)

	// 发送请求并获取响应
	resp, err := client.Do(req)
	if err != nil {
		return response, nil
	}
	defer resp.Body.Close()

	// 读取响应体
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return response, nil
	}

	err = json.Unmarshal([]byte(body), &response)
	if err != nil {
		return response, err
	}
	return response, nil
}
