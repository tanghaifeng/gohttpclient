package gohttpclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func Get(url string, params map[string]interface{}, timeout time.Duration) (string, error) {
	if len(params) > 0 {
		url += sp(params)
	}
	client := &http.Client{Timeout: timeout}
	resp, err := client.Get(url)
	if err != nil {
		log.Fatal("http get request error")
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	return string(body), err
}

func Post(url string, contentType string, data interface{}, timeout time.Duration) (string, error) {
	client := &http.Client{Timeout: timeout}
	jsonStr, _ := json.Marshal(data)
	resp, err := client.Post(url, contentType, bytes.NewBuffer(jsonStr))
	if err != nil {
		log.Fatal("http post request error")
	}
	defer resp.Body.Close()
	result, err := ioutil.ReadAll(resp.Body)
	return string(result), err
}

func PostForm(urlStr string, params map[string][]string) (string, error) {
	resp, err := http.PostForm(urlStr, params)
	if err != nil {
		log.Fatal("http get request error")
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	return string(body), err
}

// sp
func sp(params map[string]interface{}) string {
	if len(params) == 0 {
		return ""
	}
	var paramStr = ""
	for k, v := range params {
		str := fmt.Sprintf("%v", v)
		paramStr += k + "=" + str + "&"
	}
	return "?" + paramStr[0:len(paramStr)-1]
}
