package main

import (
	"github.com/qwerty22121998/go-ascii/service"
	"gocv.io/x/gocv"
	"os"
)

const ImgName = "img.jpg"

func main() {
	url := "https://ca.slack-edge.com/TBFDUP13L-U01NDR1SKL2-36b20d2cb483-512"
	_, err := service.Download(url, ImgName)
	if err != nil {
		panic(err)
	}
	img := gocv.IMRead(ImgName, gocv.IMReadColor)
	//win := gocv.NewWindow("image")

	con := service.Converter{}

	res := con.Convert(img)

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
