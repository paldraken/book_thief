package api

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

func (a *HttpApi) FetchBookChapter(workId, chapterId int) (*Chapter, error) {
	path := fmt.Sprintf("v1/work/%d/chapter/%d/text", workId, chapterId)
	body, err := a.makeRequest(path)
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

func (a *HttpApi) FetchBookChapters(workId int, chapterIds []int) ([]*Chapter, error) {
	if len(chapterIds) == 0 {
		return nil, nil
	}
	query := make([]string, len(chapterIds))
	for i := 0; i < len(chapterIds); i++ {
		query[i] = fmt.Sprintf("ids[%d]=%d", i, chapterIds[i])
	}

	path := fmt.Sprintf("v1/work/%d/chapter/many-texts?%s", workId, strings.Join(query, "&"))

	body, err := a.makeRequest(path)
	if err != nil {
		return nil, err
	}

	var chapter []*Chapter
	err = json.Unmarshal(body, &chapter)
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
