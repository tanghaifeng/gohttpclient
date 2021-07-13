package gohttpclient

import (
	"testing"
	"time"
)

func TestGet(t *testing.T) {
	p := map[string]interface{}{
		"age":  30,
		"name": "Tim",
	}
	resp, _ := Get("http://shouce.jb51.net/gopl-zh/ch11/ch11-01.html", p, 1*time.Second)
	t.Log(resp)
}
