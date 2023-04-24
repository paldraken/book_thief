package utils

import (
	"errors"
	"io"
	"net/http"
	"time"
)

type Image struct {
	Url         string
	Data        []byte
	ContentType string
}

func DownloadImage(url string) (*Image, error) {
	client := http.Client{
		Timeout: 5 * time.Second,
	}

	resp, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.ContentLength > 5*1024*1024 {
		return nil, errors.New("image size is more than 5 MB")
	}

	buf := make([]byte, resp.ContentLength)
	_, err = io.ReadFull(resp.Body, buf)
	if err != nil {
		return nil, err
	}

	return &Image{
		Url:         url,
		Data:        buf,
		ContentType: resp.Header.Get("Content-Type"),
	}, nil
}
