package main

import (
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

//普通get请求
func getContent() {
	url := "https://www.liwenzhou.com/"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("GET failed")
		return
	}
	defer resp.Body.Close()
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("GET read failed")
		return
	}
	fmt.Println(string(content))
}

/*
1.先调用url解析文件
2.url参数设置
3.参数encode
*/
func getContentWithParams() {
	requestUrl := "http://test.wyw.com/test.php"
	data := &url.Values{}
	data.Set("name", "wyw")
	data.Set("age", "30")

	u, err := url.ParseRequestURI(requestUrl)
	if err != nil {
		fmt.Println("parse url error")
		return
	}
	u.RawQuery = data.Encode()
	fmt.Println(u.String())
	r, err := http.Get(u.String())
	if err != nil {
		fmt.Println("Get url error")
		return
	}
	defer r.Body.Close()
	c, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("Get read error")
		return
	}
	fmt.Println(string(c))
}

/**

 */
func postJson() {
	requestUrl := "https://u-openapi.yundasys.com/openapi-api/v1/accountOrder/createBmOrder"
	data := `{"appid":"999999","partner_id":"201700101001","secret":"123456789","orders":[{"order_serial_no":"YD498491290","khddh":"YD498491290","sender":{"name":"怡亲","mobile":"18642120118","province":"怡亲","city":"辽宁省","county":"朝阳县","address":"辽宁省朝阳市朝阳县二十家子镇柳城经济开发区","company":"","phone":""},"receiver":{"name":"钱茂武","mobile":"18001540785","province":"江苏","city":"南通","county":"海安市","address":"江苏 南通 海安市 塘苗路18号","company":"","phone":""},"order_type":"common","node_id":"350"}]}`
	r, err := http.Post(requestUrl, "application/json", strings.NewReader(data))
	if err != nil {
		fmt.Println("Post url error")
		return
	}
	defer r.Body.Close()
	c, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("Post read error")
		return
	}
	fmt.Println(string(c))
}

func postForm() {
	requestUrl := "http://test.wyw.com/test.php"
	data := &url.Values{}
	data.Set("name", "wyw")
	data.Set("age", "31")
	r, err := http.PostForm(requestUrl, *data)
	if err != nil {
		fmt.Println("Post url error")
		return
	}
	defer r.Body.Close()
	c, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("Post read error")
		return
	}
	fmt.Println(string(c))
}

func postClient() {
	requestUrl := "https://u-openapi.yundasys.com/openapi-api/v1/accountOrder/createBmOrder"
	data := `{"appid":"999999","partner_id":"201700101001","secret":"123456789","orders":[{"order_serial_no":"YD498491290","khddh":"YD498491290","sender":{"name":"怡亲","mobile":"18642120118","province":"怡亲","city":"辽宁省","county":"朝阳县","address":"辽宁省朝阳市朝阳县二十家子镇柳城经济开发区","company":"","phone":""},"receiver":{"name":"钱茂武","mobile":"18001540785","province":"江苏","city":"南通","county":"海安市","address":"江苏 南通 海安市 塘苗路18号","company":"","phone":""},"order_type":"common","node_id":"350"}]}`
	client := &http.Client{
		Timeout: time.Second * 5,
	}
	req, err := http.NewRequest("POST", requestUrl, strings.NewReader(data))
	if err != nil {
		fmt.Println("request error")
		return
	}
	signStr := data + "_04d4ad40eeec11e9bad2d962f53dda9d"
	sign := fmt.Sprintf("%x", md5.Sum([]byte(signStr)))
	req.Header.Add("Content-Type", "application/json;charset=utf-8")
	req.Header.Add("app-key", "999999")
	req.Header.Add("req-time", strconv.FormatInt(time.Now().Unix(), 10))
	req.Header.Add("sign", sign)
	resp, err := client.Do(req)
	defer resp.Body.Close()
	c, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("client read error")
		return
	}
	fmt.Println(string(c))
}

func main() {
	postClient()
}
