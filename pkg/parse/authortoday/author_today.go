package authortoday

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/paldraken/book_thief/pkg/fetch"
	"github.com/paldraken/book_thief/pkg/parse/types"
)

const (
	CHAPTER_BASE_URL = "https://author.today/reader/%s/chapter?id=%s&_=%s"
)

type rawChapters struct {
	Title string
	Url   string
}

type AT struct {
}

// https://author.today/work/210338
func (at *AT) Parse(workUrl string) (*types.ParsedBookInfo, error) {

	workId, err := at.workIdFromUrl(workUrl)
	if err != nil {
		return nil, err
	}
	ft := fetch.NewFetcher()

	doc, err := fetchBookInfo(workUrl, ft)
	if err != nil {
		return nil, err
	}

	pbi, err := at.bookInfo(doc)
	if err != nil {
		return nil, err
	}

	chaptersList := parseChaptersList(doc, workId)

	chaptersResp, err := fetchChapters(chaptersList, ft)
	if err != nil {
		return nil, err
	}

	chapters, err := chapters(chaptersResp, chaptersList)

	for _, ch := range chapters {
		fmt.Println(ch.Number, ch.Title, len(ch.Text))
	}

	fmt.Println(pbi)
	fmt.Println(err)
	return nil, nil
}

func (at *AT) workIdFromUrl(workUrl string) (string, error) {
	u, err := url.Parse(workUrl)
	if err != nil {
		return "", err
	}
	parts := strings.Split(u.Path, "/")
	return parts[len(parts)-1], nil
}

// Загрузить страницу с инфой о книге
func fetchBookInfo(workUrl string, ft fetch.Fetcher) (*goquery.Document, error) {
	req, err := http.NewRequest(http.MethodGet, workUrl, nil)
	if err != nil {
		return nil, err
	}

	resp, err := ft.Fetch(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	return doc, nil
}

// Загрузить главы с сайта
func fetchChapters(chprs []rawChapters, ft fetch.Fetcher) ([]*fetch.BatchResp, error) {
	requests := []*http.Request{}
	for _, ch := range chprs {
		req, _ := http.NewRequest(http.MethodGet, ch.Url, nil)
		requests = append(requests, req)
	}

	httpResps, _ := ft.FetchBatch(requests)

	for _, resp := range httpResps {
		if resp.Err != nil {
			return nil, resp.Err
		}
	}

	return httpResps, nil
}
