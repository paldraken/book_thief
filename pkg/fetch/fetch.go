package fetch

import "net/http"

type Fetcher interface {
	Fetch(req *http.Request) (*http.Response, error)
	FetchBatch(requests []*http.Request) (answers []*BatchResp, allOk bool)
}

type BatchResp struct {
	N    int
	Resp *http.Response
	Err  error
}

func NewFetcher() Fetcher {
	return NewMyFetcher()
}
