package chapters

import (
	"sort"

	"github.com/paldraken/book_thief/pkg/parse/authortoday/api"
	"github.com/paldraken/book_thief/pkg/parse/types"
)

type dlChapter struct {
	*api.Chapter
	SortOrder int
}

func Get(atApi api.Api, bm *api.BookMetaInfo, userId string) ([]types.BookChapter, error) {

	mSort := make(map[int]int)

	workId := bm.ID
	var chaperIds = make([]int, len(bm.Chapters))
	for i, ch := range bm.Chapters {
		chaperIds[i] = ch.ID
		mSort[ch.ID] = ch.SortOrder
	}

	chunks := chunkInts(chaperIds, 100)

	var chapters []*dlChapter
	for _, chunk := range chunks {
		chs, err := atApi.FetchBookChapters(workId, chunk)

		if err != nil {
			return nil, err
		}

		for _, ch := range chs {
			chapters = append(chapters, &dlChapter{ch, mSort[ch.ID]})
		}
	}

	result := prepareResult(bm, chapters, userId)

	sort.Slice(result, func(i, j int) bool {
		return result[i].Number < result[j].Number
	})

	return result, nil
}

func chunkInts(ints []int, chunkSize int) [][]int {
	var chunks [][]int
	for i := 0; i < len(ints); i += chunkSize {
		end := i + chunkSize
		if end > len(ints) {
			end = len(ints)
		}
		chunks = append(chunks, ints[i:end])
	}
	return chunks
}

func prepareResult(bm *api.BookMetaInfo, dlChapters []*dlChapter, userId string) []types.BookChapter {

	result := []types.BookChapter{}

	decodeChan := make(chan types.BookChapter, len(bm.Chapters))

	for _, dlCh := range dlChapters {
		go func(ch *dlChapter) {

			text := decodeText(ch.Key, ch.Text, userId)
			bCh := types.BookChapter{}
			bCh.Number = ch.SortOrder
			bCh.Text = text
			bCh.Title = ch.Title
			decodeChan <- bCh
		}(dlCh)
	}

	for i := 0; i < len(bm.Chapters); i++ {
		result = append(result, <-decodeChan)
	}
	close(decodeChan)
	return result
}
