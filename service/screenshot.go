package service

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	url2 "net/url"
	"os"
	"time"
)

func Screenshot(token string) string {

	url := url2.QueryEscape("https://vt.sm.cn/api/NCoVInfo/riskArea#/risk-areas")
	width := 1280
	height := 800
	fullPage := 1

	// 构造URL
	query := "https://www.screenshotmaster.com/api/v1/screenshot"
	query += fmt.Sprintf("?token=%s&url=%s&width=%d&height=%d&full_page=%d",
		token, url, width, height, fullPage)

	// 调用API
	resp, err := http.Get(query)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// 检查是否调用成功
	if resp.StatusCode != 200 {
		errorBody, _ := ioutil.ReadAll(resp.Body)
		panic(fmt.Errorf("error while calling api %s", errorBody))
	}

	filePath := "./" + time.Now().Format("2006-01-02") + ".png"

	// 保存截图
	file, err := os.Create(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		panic(err)
	}
	return filePath
}

// 获取文件的md5码
func GetFileMd5(path string) string {
	// 文件全路径名
	pFile, err := os.Open(path)
	if err != nil {
		_ = fmt.Errorf("打开文件失败，filename=%v, err=%v", path, err)
		return ""
	}
	defer pFile.Close()
	md5h := md5.New()
	io.Copy(md5h, pFile)
	return hex.EncodeToString(md5h.Sum(nil))
}

// 获取文件的md5码
func GetPicBase64(path string) string {

	fileByte, _ := ioutil.ReadFile(path)
	return base64.StdEncoding.EncodeToString(fileByte)
}
