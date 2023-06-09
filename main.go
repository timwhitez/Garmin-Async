package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"
)

// 请务必先在网站上已成功添加佳明国内与国际账号
// 填写在dailysync.vyzt.dev操作同步时抓包获取到的userid与token

var UserId = ""
var Cookie = ""

func main() {
	//使用方式 ./garminAsync.exe
	//或者后面跟上请求间隔时间(小时) ./garminAsync.exe 6
	freq := 12
	if len(os.Args) == 2 {
		tmp, err := strconv.Atoi(os.Args[1])
		if err == nil {
			freq = tmp
		}
	}
	for {
		sendFunc()
		time.Sleep(time.Duration(freq) * time.Hour)
	}
}

func sendFunc() {
	url := "https://dailysync.vyzt.dev/api/garmin/syncWorkoutData"
	method := "POST"

	payload := []byte(`{"userId":"` + UserId + `"}`)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, bytes.NewBuffer(payload))

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Host", "dailysync.vyzt.dev")
	req.Header.Add("Accept", "*/*")
	req.Header.Add("Accept-Encoding", "identity")
	req.Header.Add("Accept-Language", "zh-CN,zh;q=0.9")
	req.Header.Add("Content-Length", "45")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Cookie", Cookie)
	req.Header.Add("Origin", "https://dailysync.vyzt.dev")
	req.Header.Add("Referer", "https://dailysync.vyzt.dev/dashboard")
	req.Header.Add("Sec-Fetch-Dest", "empty")
	req.Header.Add("Sec-Fetch-Mode", "cors")
	req.Header.Add("Sec-Fetch-Site", "same-origin")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/536.00 (KHTML, like Gecko) Chrome/110.0.0.0 Safari/536.00")
	req.Header.Add("sec-ch-ua", "\"Google Chrome\";v=\"113\", \"Chromium\";v=\"113\", \"Not-A.Brand\";v=\"24\"")
	req.Header.Add("sec-ch-ua-mobile", "?0")
	req.Header.Add("sec-ch-ua-platform", "\"Windows\"")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(body))
}
