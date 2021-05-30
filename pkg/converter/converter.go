package converter

import (
	"fmt"
	"github.com/nfnt/resize"
	"golang.org/x/image/font"
	"golang.org/x/image/font/inconsolata"
	"golang.org/x/image/math/fixed"
	"image"
	"image/color"
	_ "image/gif"
	_ "image/jpeg"
	"image/png"
	_ "image/png"
	"io"
	"math"
	"os"
)

type NativeConverter struct {
	Char []rune
	font font.Face
}

func Ascii() *NativeConverter {
	f := *inconsolata.Regular8x16
	f.Left = 0
	f.Advance = f.Ascent

	return &NativeConverter{
		Char: []rune(` .:-=+*#%@`),
		font: &f,
	}
}

func (c *NativeConverter) ConvertToAscii(reader io.Reader, size uint) ([][]rune, error) {
	img, err := c.Load(reader)
	if err != nil {
		return nil, err
	}
	gray := c.ToGray(img)

	return c.ToRune(gray, size), nil
}

func (c *NativeConverter) ConvertToImage(reader io.Reader, size uint) (image.Image, error) {
	img, err := c.Load(reader)
	if err != nil {
		return nil, err
	}

	gray := c.ToGray(img)

	r := c.ToRune(gray, size)

	return c.ToImage(r), nil
}

func (c *NativeConverter) Load(file io.Reader) (image.Image, error) {
	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}
	return img, nil
}

func (c *NativeConverter) ToGray(img image.Image) *image.Gray {
	bound := img.Bounds()
	gray := image.NewGray(bound)
	for x := bound.Min.X; x < bound.Max.X; x++ {
		for y := bound.Min.Y; y < bound.Max.Y; y++ {
			gray.Set(x, y, img.At(x, y))
		}
	}
	return gray
}

func (c *NativeConverter) getChar(value uint8) rune {
	part := math.Ceil(256 / float64(len(c.Char)))
	idx := int(float64(value) / part)
	return c.Char[idx]
}

func (c *NativeConverter) ToImage(r [][]rune) image.Image {
	//metric := c.font.Metrics()
	factor := int(c.font.Metrics().Ascent >> 6)
	h := factor * len(r)
	w := factor * len(r[0])
	//h := int(metric.Height) // c.font.Ascent + c.font.Descent
	//w := int(metric.Height)

	img := image.NewGray(image.Rect(0, 0, w, h))
	fmt.Println(img.Rect)

	drawer := &font.Drawer{
		Dst:  img,
		Src:  image.NewUniform(color.White),
		Face: c.font,
	}

	for i, l := range r {
		drawer.Dot = fixed.Point26_6{
			X: 0,
			Y: c.font.Metrics().XHeight * fixed.Int26_6(i+1),
		}

		drawer.DrawString(string(l))
	}

	return img

}

func (c *NativeConverter) ToRune(img *image.Gray, size uint) [][]rune {
	img = resize.Thumbnail(size, size, img, resize.NearestNeighbor).(*image.Gray)
	bound := img.Bounds()

	result := make([][]rune, bound.Dy())
	for i := 0; i < len(result); i++ {
		result[i] = make([]rune, bound.Dx())
	}

	for x := bound.Min.X; x < bound.Max.X; x++ {
		for y := bound.Min.Y; y < bound.Max.Y; y++ {
			result[y][x] = c.getChar(img.GrayAt(x, y).Y)
		}
	}
	return result
}

func (c *NativeConverter) Write(img image.Image, file *os.File) error {
	return png.Encode(file, img)
}
