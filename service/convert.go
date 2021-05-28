package service

import (
	"fmt"
	"gocv.io/x/gocv"
	"image"
	"math"
)

type Converter struct {
}

var CharList = []rune(" .:-=+*#%@")

func GetChar(value uint8) rune {
	part := 255 / (len(CharList) - 1)
	idx := int(value) / part
	return CharList[idx]
}

func GetScale(mat gocv.Mat, maxSize int) float64 {
	return 1.0 / math.Ceil(math.Max(float64(mat.Rows()/maxSize), float64(mat.Cols()/maxSize)))
}

func (c *Converter) Convert(img gocv.Mat) [][]rune {

	scale := GetScale(img, 100)
	fmt.Println("Scale ratio", scale)
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
