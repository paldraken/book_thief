package authortoday

import (
	"fmt"
	"log"
	"net/url"
	"strconv"
	"strings"

	"github.com/paldraken/book_thief/pkg/parse/authortoday/api"
	"github.com/paldraken/book_thief/pkg/parse/authortoday/chapters"
	"github.com/paldraken/book_thief/pkg/parse/types"
)

type AT struct {
}

func (at *AT) Parse(workUrl string, config *types.Config) (*types.BookData, error) {

	workId, err := workIdFromUrl(workUrl)
	if err != nil {
		return nil, err
	}
	username := config.Username
	password := config.Password

	atApi := api.NewHttpApi()

	token, err := atApi.ObtainingAccessToken(username, password)
	if err != nil {
		return nil, err
	}

	curentUser, err := atApi.FetchCurrentUser(token)
	if err != nil {
		return nil, err
	}

	bookMeta, err := atApi.FetchBookMetaInfo(workId, token)
	if err != nil {
		return nil, err
	}

	pbi, err := at.bookInfo(bookMeta)
	if err != nil {
		return nil, err
	}

	bookChapters, err := chapters.Get(token, bookMeta, fmt.Sprintf("%d", curentUser.Id))

	pbi.Chapters = bookChapters

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
