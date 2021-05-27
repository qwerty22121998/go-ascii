package main

import (
	"github.com/qwerty22121998/go-ascii/service"
	"gocv.io/x/gocv"
	"os"
)

const ImgName = "img.jpg"

func main() {
	url := "https://photo2.tinhte.vn/data/attachment-files/2021/03/5373730_Golang.png"
	_, err := service.Download(url, ImgName)
	if err != nil {
		panic(err)
	}
	img := gocv.IMRead(ImgName, gocv.IMReadColor)
	//win := gocv.NewWindow("image")

	con := service.Converter{}

	res := con.Convert(img, 0.1)

	file, _ := os.Create("out.txt")

	for _, l := range res {
		//fmt.Println(string(l))
		file.WriteString(string(l) + "\n")
	}

	file.Close()

	//win.IMShow(con.Convert(img))
	//win.WaitKey(0)

	//
	//
	//win.IMShow(img)
	//win.WaitKey(0)
}
