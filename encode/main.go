package main

import (
	"fmt"
	"github.com/afocus/captcha"
	"image/color"
	"image/png"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", showImg)
	log.Println("listening on 0.0.0.0:9091....")
	err := http.ListenAndServe(":9091", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func showImg(w http.ResponseWriter, r *http.Request) {
	//一些默认的设置
	cap := captcha.New()

	//当然可以通过 AddFont()来追加
	cap.SetFont("Tahoma.ttf")

	cap.SetSize(60, 30)

	cap.SetDisturbance(captcha.NORMAL)

	cap.SetFrontColor(color.Gray{3})

	cap.SetBkgColor(color.White)

	img := cap.CreateCustom("1234ab")

	w.Header().Set("Content-Type", "image/png")
	png.Encode(w, img)
	fmt.Printf("come from %s...\n", r.Host)
}
