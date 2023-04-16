package api

const (
	DEFAULT_TOKEN = "guest"
	BASE_DOMAIN   = "https://api.author.today/"
)

type Api interface {
	ObtainingAccessToken(login, password string) (string, error)
	FetchBookMetaInfo(workId int, userToken string) (*BookMetaInfo, error)
	FetchBookChapter(workId, chapterId int, userToken string) (*Chapter, error)
	FetchCurrentUser(userToken string) (*CurrentUser, error)
}

type HttpApi struct{}

func NewHttpApi() Api {
	return &HttpApi{}
}
