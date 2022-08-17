package fetch

import (
	"fmt"
	"net/http"
	"sort"
	"time"
)

type myFetcher struct {
	client *http.Client
}

func NewMyFetcher() Fetcher {
	var myFetcher myFetcher

	client := &http.Client{
		Timeout: 3 * time.Second,
	}
	myFetcher.client = client

	return &myFetcher
}

func (p *myFetcher) Fetch(req *http.Request) (*http.Response, error) {
	return p.client.Do(req)
}

func (p *myFetcher) FetchBatch(reqs []*http.Request) (answers []*BatchResp, allOk bool) {

	fCh := make(chan bool, 2)
	dCh := make(chan *BatchResp)

	defer func() {
		close(dCh)
		close(fCh)
	}()

	for i, req := range reqs {
		httpRes := &BatchResp{N: i}
		go (func(httpRes *BatchResp, req *http.Request) {
			fCh <- true
			fmt.Println(httpRes.N, req.URL)
			resp, err := p.client.Do(req)
			if err != nil {
				httpRes.Err = err
			} else {
				httpRes.Resp = resp
			}
			dCh <- httpRes
			time.Sleep(50 * time.Millisecond)
			<-fCh
		})(httpRes, req)
	}

	errCount := 0
	for i := 0; i < len(reqs); i++ {
		resp := <-dCh
		answers = append(answers, resp)
		if resp.Err != nil {
			errCount++
		}
	}

	sort.Slice(answers[:], func(i, j int) bool { return answers[i].N < answers[j].N })

	return answers, errCount == 0
}
