package chapters

import (
	"sort"

	"github.com/paldraken/book_thief/pkg/parse/authortoday/api"
	"github.com/paldraken/book_thief/pkg/parse/types"
)

type chWorkerArgs struct {
	workId int
	chMeta api.ChapterMeta
	token  string
}

type dlChapterRes struct {
	chMeta api.ChapterMeta
	ch     *api.Chapter
	err    error
}

type chaptersError struct {
	chaperIds []int
	lastError error
}

func (e *chaptersError) Error() string {
	return e.lastError.Error()
}

func Get(token string, bm *api.BookMetaInfo, userId string) ([]types.BookChapter, error) {

	workId := bm.ID
	dlChapters, err := downloadChapters(workId, bm.Chapters, token)
	if err != nil {
		return nil, err
	}

	result := prepareResult(bm, dlChapters, userId)

	sort.Slice(result, func(i, j int) bool {
		return result[i].Number < result[j].Number
	})

	return result, nil
}

func prepareResult(bm *api.BookMetaInfo, dlChapters []*dlChapterRes, userId string) []types.BookChapter {

	result := []types.BookChapter{}

	decodeChan := make(chan types.BookChapter, len(bm.Chapters))

	for _, dlCh := range dlChapters {
		go func(dlCh *dlChapterRes) {
			ch := dlCh.ch
			chMeta := dlCh.chMeta

			text := decodeText(ch.Key, ch.Text, userId)

			bCh := types.BookChapter{}
			bCh.Number = chMeta.SortOrder
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
