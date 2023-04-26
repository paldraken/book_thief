package authortoday

import (
	"fmt"
	"hash/crc32"
	"regexp"
	"sync"

	"github.com/paldraken/book_thief/pkg/parse/types"
	"github.com/paldraken/book_thief/pkg/parse/utils"
)

func extractImagesFromText(chapters []*types.BookChapter) []string {
	result := []string{}
	for _, chapter := range chapters {

		re := regexp.MustCompile(`<img(?:[^>]*?\s+)?src="([^"]+)"`)
		srcValues := re.FindAllStringSubmatch(chapter.Text, -1)

		for _, matches := range srcValues {
			result = append(result, matches[1])
		}
	}
	return result
}

func getBookImages(book *types.BookData) []*types.BookImage {
	imageList := extractImagesFromText(book.Chapters)

	wg := sync.WaitGroup{}
	for _, url := range imageList {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			hash := crc32.ChecksumIEEE([]byte(url))
			id := fmt.Sprintf("%x", hash)
			img, err := utils.DownloadImage(url)
			sync.

		}(url)
	}
}
