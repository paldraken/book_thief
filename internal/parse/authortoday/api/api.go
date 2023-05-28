package api

const (
	DEFAULT_TOKEN = "guest"
	BASE_DOMAIN   = "https://api.author.today/"
)

type Api interface {
	FetchBookMetaInfo(workId int) (*BookMetaInfo, error)
	FetchBookChapter(workId, chapterId int) (*Chapter, error)
	FetchBookChapters(workId int, chapterIds []int) ([]*Chapter, error)
	FetchCurrentUser() (*CurrentUser, error)
}

type HttpApi struct {
	token string
}

func NewHttpApi(token string) Api {
	return &HttpApi{token: token}
}
