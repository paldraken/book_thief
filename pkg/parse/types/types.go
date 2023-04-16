package types

type Config struct {
	Username string
	Password string
}

const (
	Author     = "Author"
	CoAuthor   = "CoAuthor"
	Translator = "Translator"
)

type BookAuthor struct {
	FullName  string
	Firstname string
	Lastname  string
	NickName  string
	Email     string
	Type      string
}

type BookChapter struct {
	Number int
	Title  string
	Text   string
	Url    string
}

type BookData struct {
	Title      string
	Authors    []BookAuthor
	WorkForm   string
	Genres     []string
	Series     string
	SeriesN    string
	Tags       []string
	Date       string
	Annotation string
	Chapters   []BookChapter
	IsDone     bool
}
