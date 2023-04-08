package api

import (
	"encoding/json"
	"fmt"
	"time"
)

func FetchBookChapter(workId, chapterId int, userToken string) (*Chapter, error) {
	path := fmt.Sprintf("v1/work/%d/chapter/%d/text", workId, chapterId)
	body, err := makeRequest(path, userToken)
	if err != nil {
		return nil, err
	}

	chapter := &Chapter{}
	err = json.Unmarshal(body, chapter)
	if err != nil {
		return nil, err
	}
	return chapter, nil
}

type Chapter struct {
	ID                   int       `json:"id"`
	Title                string    `json:"title"`
	IsDraft              bool      `json:"isDraft"`
	PublishTime          time.Time `json:"publishTime"`
	LastModificationTime time.Time `json:"lastModificationTime"`
	Text                 string    `json:"text"`
	Key                  string    `json:"key"`
}
