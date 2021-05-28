package main

import (
	"github.com/qwerty22121998/go-ascii/service"
	"gocv.io/x/gocv"
	"os"
)

const ImgName = "img.jpg"

func main() {
	url := "https://media-cdn.laodong.vn/Storage/NewsPortal/2019/4/2/666434/11.jpg"
	_, err := service.Download(url, ImgName)
	if err != nil {
		panic(err)
	}
	img := gocv.IMRead(ImgName, gocv.IMReadColor)
	win := gocv.NewWindow("image")

	win.IMShow(img)

	con := service.DefaultConverter()

	res := con.ToASCII(img, 50)

	file, _ := os.Create("out.txt")

	for _, l := range res {
		//fmt.Println(string(l))
		file.WriteString(string(l) + "\n")
	}

	file.Close()

	out := con.Render(res)

	win.IMShow(out)

	gocv.IMWrite("banh.jpg", out)
	for {
		win.WaitKey(0)
	}

	//
	//
	//win.IMShow(img)
	//win.WaitKey(0)
}
