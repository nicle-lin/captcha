package main

import (
	"fmt"
	"github.com/otiai10/gosseract"
	"image"
	"net"
	"net/http"
	"time"
)

const (
	URL string = "http://localhost:9091"
)

func main() {
	decodeimage()

	img, _ := HttpGetImg(URL)
	client, err := gosseract.NewClient()
	if err != nil {
		fmt.Println(err)
		return
	}

	c := client.Image(img)
	captcha, err := c.Out()
	fmt.Println("captcha:", captcha)
	if err != nil {
		fmt.Println(err)
		return
	}

}

func decodeFromRemote() {

}

func decodeimage() {
	client, err := gosseract.NewClient()
	if err != nil {
		fmt.Println(err)
	}
	//c := client.Src("sample.png")
	c := client.Src("sample.png")
	out, err := c.Out()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(out)
}

func HttpGetImg(url string) (img image.Image, err error) {

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Printf("url:%s 填写有错误:%s\n", url, err.Error())
		return
	}
	req.Header.Set("Content-Type", "charset=UTF-8")
	req.Header.Set("Accept", "image/*,text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Host", req.Host)
	req.Header.Set("Referer", req.Host)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/42.0.2311.135 Safari/537.36")
	tr := &http.Transport{DisableKeepAlives: false,
		Dial: func(netw, addr string) (net.Conn, error) {
			c, err := net.DialTimeout(netw, addr, time.Second*30)
			if err != nil {
				return nil, err
			}
			c.SetDeadline(time.Now().Add(5 * time.Second))
			return c, nil
		}}

	client := &http.Client{Transport: tr}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("HTTP请求url:%s失败,err:%s\n", url, err.Error())
		return
	}
	defer resp.Body.Close()

	img, _, err = image.Decode(resp.Body)
	if err != nil {
		fmt.Printf("Image Decode->url:%s,error:%s\n", url, err.Error())
		return
	}
	return
}
