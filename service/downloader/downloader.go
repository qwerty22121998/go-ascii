package downloader

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func Download(url string, name string) (*os.File, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("status code error : %v", resp.StatusCode)
	}
	file, err := os.Create(name)
	if err != nil {
		return nil, err
	}

	if _, err := io.Copy(file, resp.Body); err != nil {
		return nil, err
	}
	return file, nil

}
