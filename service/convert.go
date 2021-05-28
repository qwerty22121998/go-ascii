package service

import (
	"fmt"
	"gocv.io/x/gocv"
	"image"
	"image/color"
	"math"
)

type Converter struct {
	textColor     color.RGBA
	textThickness int
	textScale     float64
	multiplier    int
}

var CharList = []rune(" .:-=+*#X@")

func DefaultConverter() *Converter {
	return &Converter{
		textColor:     color.RGBA{R: 255, G: 255, B: 255},
		textThickness: 3,
		textScale:     1,
		multiplier:    30,
	}
}

func GetChar(value uint8) rune {
	part := 255 / (len(CharList) - 1)
	idx := int(value) / part
	return CharList[idx]
}

func getScale(mat gocv.Mat, maxSize int) float64 {
	return math.Min(float64(maxSize)/float64(mat.Rows()), float64(maxSize)/float64(mat.Cols()))
}

func (c *Converter) Render(img [][]rune) gocv.Mat {
	h := len(img)
	w := len(img[0])
	mat := gocv.NewMatWithSize(h*c.multiplier, w*c.multiplier, gocv.MatTypeCV8UC3)
	for y, l := range img {
		for x, v := range l {
			gocv.PutText(&mat, fmt.Sprintf("%c", v), image.Point{
				X: x * c.multiplier,
				Y: y * c.multiplier,
			}, gocv.FontHersheyComplexSmall, c.textScale, color.RGBA{R: 255, G: 255, B: 255}, c.textThickness)
		}
	}
	return mat
}

func (c *Converter) ToASCII(img gocv.Mat, maxSize int) [][]rune {
	scale := getScale(img, maxSize)
	gocv.Resize(img, &img, image.Point{}, scale, scale, gocv.InterpolationCubic)

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
