package authortoday

import (
	"fmt"
	"hash/crc32"
	"regexp"
	"strings"
	"sync"

	"github.com/paldraken/book_thief/internal/parse/types"
	"github.com/paldraken/book_thief/internal/parse/utils"
)

type imgInfo struct {
	chapterN  int
	origTag   string
	url       string
	btReplace string
}

func bookImages(chapters []*types.BookChapter) []*types.BookImage {
	imagesInfo := extractImgTagsFromChapters(chapters)
	return getBookImages(imagesInfo)
}

func extractImgTagsFromChapters(chapters []*types.BookChapter) []*imgInfo {
	result := []*imgInfo{}
	for _, chapter := range chapters {

		srcRegex := regexp.MustCompile(`<img[^>]+src="([^"]+)"`)
		imgRegex := regexp.MustCompile(`<img[^>]+\>`)
		matches := imgRegex.FindAllString(chapter.Text, -1)
		for _, imgTag := range matches {

			srcMatch := srcRegex.FindStringSubmatch(imgTag)

			if len(srcMatch) != 2 {
				continue
			}
			src := srcMatch[1]
			url := addATDomain(src)
			hash := hashFromUrl(url)

			replacement := fmt.Sprintf("#__bt_binary__#%s#", hash)

			it := &imgInfo{
				chapterN:  chapter.Number,
				origTag:   imgTag,
				url:       url,
				btReplace: replacement,
			}
			result = append(result, it)
			replaceImgToPlaceolder(chapter, it)
		}
	}
	return result
}

func replaceImgToPlaceolder(chapter *types.BookChapter, info *imgInfo) {
	chapter.Text = strings.Replace(chapter.Text, info.origTag, info.btReplace, 1)
}

func getBookImages(imgesInfo []*imgInfo) []*types.BookImage {

	result := make([]*types.BookImage, len(imgesInfo))

	wg := sync.WaitGroup{}
	for i, info := range imgesInfo {
		wg.Add(1)
		go func(info *imgInfo, i int) {
			defer wg.Done()

			id := hashFromUrl(info.url)
			img, err := utils.DownloadImage(info.url)

			var data []byte
			var contentType string

			if err != nil {
				data = nil
				contentType = "image unavailable: " + info.url
				fmt.Println("Download error:", err)
			} else {
				data = img.Data
				contentType = img.ContentType
			}
			result[i] = &types.BookImage{
				Id:          id,
				Data:        data,
				ContentType: contentType,
			}
		}(info, i)
	}
	wg.Wait()

	return result
}

func addATDomain(url string) string {
	match, _ := regexp.MatchString(`^(http|https):\/\/[a-zA-Z0-9\-\.]+\.[a-zA-Z]{2,}(?:\/[\w\-\.]*)*$`, url)
	if match {
		return url
	} else {
		return "https://cm.author.today" + url
	}
}

func hashFromUrl(url string) string {
	hash := crc32.ChecksumIEEE([]byte(url))
	return fmt.Sprintf("%x", hash)
}
