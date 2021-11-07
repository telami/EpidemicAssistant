package service

import (
	"EpidemicAssistant/encrypt"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

const overAllUrl = "https://lab.isaaclin.cn/nCoV/api/overall?latest=1"

const provinceUrl = "https://lab.isaaclin.cn/nCoV/api/area?latest=1&province=%s"

const riskGradeUrl = "http://103.66.32.242:8005/zwfwMovePortal/interface/interfaceJson"

// 全国疫情数据
func OverAllInfo() string {

	res, err := http.Get(overAllUrl)
	if err != nil {
		return ""
	}
	result, err := ioutil.ReadAll(res.Body)
	_ = res.Body.Close()
	if err != nil {
		return ""
	}
	return string(result)
}

// 指定省疫情数据
func GetProvinceInfo(provinceName string) string {

	//url encode
	provinceName = url.QueryEscape(provinceName)

	res, err := http.Get(fmt.Sprintf(provinceUrl, provinceName))
	if err != nil {
		return ""
	}
	result, err := ioutil.ReadAll(res.Body)
	_ = res.Body.Close()
	if err != nil {
		return ""
	}
	return string(result)
}

// 指定省疫情数据
func GetRiskGradeInfo() string {
	timestamp := strconv.FormatInt(time.Now().UnixNano()/1e9, 10)

	message := []byte(timestamp + "23y0ufFl5YxIyGrI8hWRUZmKkvtSjLQA" + "123456789abcdefg" + timestamp)
	signatureHeader := encrypt.GetSHA256HashCode(message)

	xwifTimestamp := encrypt.GetSHA256HashCode([]byte(timestamp + "fTN2pfuisxTavbTuYVSsNJHetwq5bJvCQkjjtiLM2dCratiA" + timestamp))

	song := make(map[string]interface{})
	song["appId"] = "NcApplication"
	song["paasHeader"] = "zdww"
	song["nonceHeader"] = "123456789abcdefg"
	song["timestampHeader"] = timestamp
	song["key"] = "3C502C97ABDA40D0A60FBEE50FAAD1DA"
	song["signatureHeader"] = signatureHeader
	bytesData, err := json.Marshal(song)
	if err != nil {
		return ""
	}
	reader := bytes.NewReader(bytesData)
	request, err := http.NewRequest("POST", riskGradeUrl, reader)
	if err != nil {
		return ""
	}
	request.Header.Add("x-wif-nonce", "QkjjtiLM2dCratiA")
	request.Header.Add("x-wif-paasid", "smt-application")
	request.Header.Add("x-wif-signature", xwifTimestamp)
	request.Header.Add("x-wif-timestamp", timestamp)
	request.Header.Add("Content-Type", "application/json;charset=UTF-8")
	client := http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		return ""
	}
	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return ""
	}
	_ = resp.Body.Close()
	return string(respBytes)
}
