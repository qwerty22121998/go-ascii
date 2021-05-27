package service

import (
	"fmt"
	"gocv.io/x/gocv"
	"image"
)

type Converter struct {
}

var CharList = []rune("@%#*+=-:. ")



func GetChar(value uint8) rune {
	part := 255 / (len(CharList) - 1)
	idx := int(value) / part
	return CharList[idx]
}

func (c *Converter) Convert(img gocv.Mat, scale float64) [][]rune {

	gocv.Resize(img, &img, image.Point{}, scale, scale, gocv.InterpolationCubic)

	fmt.Println(img.Rows(), img.Cols())

	result := make([][]rune, img.Rows())

	gray := gocv.Zeros(img.Rows(), img.Cols(), img.Type())
	gocv.CvtColor(img, &gray, gocv.ColorRGBToGray)
	//
	for i := 0; i < gray.Rows(); i++ {
		result[i] = make([]rune, img.Cols())
		for j := 0; j < gray.Cols(); j++ {
			value := gray.GetUCharAt(i, j)
			result[i][j] = GetChar(value)
		}
	}

	return result
}
