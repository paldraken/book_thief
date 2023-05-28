package authortoday

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"

	"github.com/paldraken/book_thief/internal/parse/authortoday/api"
	"github.com/paldraken/book_thief/internal/parse/authortoday/chapters"
	"github.com/paldraken/book_thief/internal/parse/types"
	"github.com/paldraken/book_thief/internal/parse/utils"
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

	token, err := api.ObtainingAccessToken(username, password)
	if err != nil {
		return nil, err
	}

	atApi := api.NewHttpApi(token)

	curentUser, err := atApi.FetchCurrentUser()
	if err != nil {
		return nil, err
	}

	bookMeta, err := atApi.FetchBookMetaInfo(workId)
	if err != nil {
		return nil, err
	}

	pbi, err := at.bookInfo(bookMeta)
	if err != nil {
		return nil, err
	}

	bookChapters, err := chapters.Get(atApi, bookMeta, fmt.Sprintf("%d", curentUser.Id))
	if err != nil {
		return nil, err
	}
	pbi.Chapters = bookChapters

	if bookMeta.CoverURL != "" {
		cover, err := fetchImage(bookMeta.CoverURL, "cover")
		if err == nil {
			pbi.CoverId = "cover"
			pbi.Images = append(pbi.Images, cover)
		}
	}

	otherImages := bookImages(bookChapters)
	pbi.Images = append(pbi.Images, otherImages...)

	return pbi, nil
}

func fetchImage(url, imageId string) (*types.BookImage, error) {
	image, err := utils.DownloadImage(url)
	if err != nil {
		return nil, err
	}
	return &types.BookImage{
		Id:          imageId,
		Data:        image.Data,
		ContentType: image.ContentType,
	}, nil
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
