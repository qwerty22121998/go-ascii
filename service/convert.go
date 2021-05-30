package service

import (
	"bytes"
	"github.com/qwerty22121998/go-ascii/pkg/converter"
	"github.com/qwerty22121998/go-ascii/pkg/downloader"
	"image/png"
	"strings"
)

type ConvertService interface {
	FromUrlToString(url string, size uint) (string, error)
	FromUrlToImage(url string, size uint) ([]byte, error)
}

func NewConvertService() ConvertService {
	return &convertService{}
}

type convertService struct {
}

func (c *convertService) FromUrlToString(url string, size uint) (string, error) {
	reader, err := downloader.Download(url)
	if err != nil {
		return "", err
	}

	res, err := converter.Ascii().ConvertToAscii(reader, size)
	if err != nil {
		return "", err
	}
	str := make([]string, len(res))

	for i := range res {
		str[i] = string(res[i])
	}

	return strings.Join(str, "\n"), nil
}

func (c *convertService) FromUrlToImage(url string, size uint) ([]byte, error) {
	reader, err := downloader.Download(url)
	if err != nil {
		return nil, err
	}

	img, err := converter.Ascii().ConvertToImage(reader, size)
	if err != nil {
		return nil, err
	}

	res := new(bytes.Buffer)

	if err = png.Encode(res, img); err != nil {
		return nil, err
	}

	return res.Bytes(), nil
}
