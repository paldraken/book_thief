package authortoday

import (
	"fmt"
	"log"
	"net/url"
	"strconv"
	"strings"

	"github.com/paldraken/book_thief/pkg/parse/authortoday/api"
	"github.com/paldraken/book_thief/pkg/parse/types"
)

type AT struct {
	userToken string
}

// https://author.today/work/210338
func (at *AT) Parse(workUrl string) (*types.ParsedBookInfo, error) {

	workId, err := workIdFromUrl(workUrl)
	if err != nil {
		return nil, err
	}
	username := ""
	password := ""

	token, err := api.ObtainingAccessToken(username, password)
	if err != nil {
		return nil, err
	}
	at.userToken = token
	fmt.Println("at.userToken", token)

	curentUser, err := api.FetchCurrentUser(token)
	if err != nil {
		return nil, err
	}

	bookMeta, err := api.FetchBookMetaInfo(workId, token)
	if err != nil {
		return nil, err
	}

	pbi, err := at.bookInfo(bookMeta)
	if err != nil {
		return nil, err
	}

	bookChapters := []types.ParsedChapter{}
	for _, ch := range bookMeta.Chapters {
		fmt.Println(ch)
		chapter, err := api.FetchBookChapter(workId, ch.ID, token)
		if err != nil {
			return nil, err
		}
		text := decodeText(chapter.Key, chapter.Text, fmt.Sprintf("%d", curentUser.Id))
		bCh := types.ParsedChapter{}
		bCh.Number = ch.SortOrder
		bCh.Text = text
		bCh.Title = ch.Title

		bookChapters = append(bookChapters, bCh)
	}

	pbi.Chapters = bookChapters

	fmt.Println(curentUser.Id)

	if err != nil {
		log.Panic(err)
	}

	return pbi, nil
}

func workIdFromUrl(workUrl string) (int, error) {
	u, err := url.Parse(workUrl)
	if err != nil {
		return 0, err
	}
	parts := strings.Split(u.Path, "/")
	workIdStr := parts[len(parts)-1]

	workId, err := strconv.Atoi(workIdStr)
	if err != nil {
		return 0, err
	}
	return workId, nil
}
