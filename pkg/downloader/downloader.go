package downloader

import (
	"fmt"
	"io"
	"net/http"
)

func Download(url string) (io.Reader, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("status code error : %v", resp.StatusCode)
	}

	return resp.Body, nil
}
