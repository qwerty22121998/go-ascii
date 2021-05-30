package main

import (
	"github.com/qwerty22121998/go-ascii/service/converter"
	"os"
)

const ImgName = "banh.jpg"

func main() {
	c := converter.Ascii(50)

	file, _ := os.Open("banh.jpg")
	defer file.Close()

	out, _ := os.Create("out.png")
	defer out.Close()

	img, _ := c.Load(file)

	gray := c.ToGray(img)

	r := c.ToRune(gray)

	as := c.ToImage(r)

	c.Write(as,out)
}
