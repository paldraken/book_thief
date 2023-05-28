package utils

import (
	"errors"
	"io"
	"net/http"
	"sync"
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

func DownloadImagesByWorkers(urls []string) ([]*Image, error) {
	urlCh := make(chan string)
	results := make(chan *Image)

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			downloadWorker(urlCh, results)
		}()
	}

	for _, url := range urls {
		urlCh <- url
	}
	close(urlCh)
	wg.Wait()
	close(results)

	images := []*Image{}
	for r := range results {
		images = append(images, r)
	}
	return images, nil
}

func downloadWorker(urlCh <-chan string, results chan<- *Image) {
	for url := range urlCh {
		img, err := DownloadImage(url)
		if err != nil {
			img = &Image{Url: url, Data: nil}
		}
		results <- img
	}
}
