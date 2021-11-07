package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

const webHookUrl = "https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=%s"

//发送markdown文本
func SendMarkdownMessage(info string, key string) {
	content := make(map[string]interface{})
	content["content"] = info

	data := make(map[string]interface{})
	data["msgtype"] = "markdown"
	data["markdown"] = content

	bytesData, err := json.Marshal(data)
	if err != nil {
		return
	}
	reader := bytes.NewReader(bytesData)
	request, err := http.NewRequest("POST", fmt.Sprintf(webHookUrl, key), reader)
	if err != nil {
		return
	}
	client := http.Client{}
	_, _ = client.Do(request)
}

//发送图片
func SendImageMessage(base64 string, md5 string, key string) {
	content := make(map[string]interface{})
	content["base64"] = base64
	content["md5"] = md5

	data := make(map[string]interface{})
	data["msgtype"] = "image"
	data["image"] = content

	bytesData, err := json.Marshal(data)
	if err != nil {
		return
	}
	reader := bytes.NewReader(bytesData)
	request, err := http.NewRequest("POST", fmt.Sprintf(webHookUrl, key), reader)
	if err != nil {
		return
	}
	client := http.Client{}
	_, _ = client.Do(request)
}
