package chapters

import (
	"sync"

	"github.com/paldraken/book_thief/pkg/parse/authortoday/api"
)

func downloadChapters(workId int, chms []api.ChapterMeta, token string) ([]*dlChapterRes, error) {
	tasks := make(chan *chWorkerArgs)
	results := make(chan *dlChapterRes)

	atApi := api.NewHttpApi()

	var wg sync.WaitGroup
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go chapterWorker(&wg, tasks, results, atApi)
	}

	go func() {
		for _, chMeta := range chms {
			tasks <- &chWorkerArgs{
				workId: workId,
				chMeta: chMeta,
				token:  token,
			}
		}
		wg.Wait()
		close(tasks)
	}()

	result := make([]*dlChapterRes, len(chms))

	chError := &chaptersError{}

	for i := 0; i < len(chms); i++ {
		res := <-results
		if res.err != nil {
			result[i] = nil
			chError.chaperIds = append(chError.chaperIds, chms[i].ID)
			chError.lastError = res.err
		}
		result[i] = res
	}
	close(results)

	if chError.lastError != nil {
		return nil, chError
	}
	return result, nil
}

func chapterWorker(
	wg *sync.WaitGroup,
	tasks <-chan *chWorkerArgs,
	result chan<- *dlChapterRes,
	atApi api.Api,
) {
	defer wg.Done()
	for {
		task := <-tasks
		chaper, err := atApi.FetchBookChapter(task.workId, task.chMeta.ID, task.token)

		result <- &dlChapterRes{
			chMeta: task.chMeta,
			ch:     chaper,
			err:    err,
		}
	}
}
