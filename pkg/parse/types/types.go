package types

const (
	Author     = "Author"
	CoAuthor   = "CoAuthor"
	Translator = "Translator"
)

type ParsedAuthor struct {
	FullName  string
	Firstname string
	Lastname  string
	NickName  string
	Email     string
	Type      string
}

type ParsedChapter struct {
	Number int
	Title  string
	Text   string
	Url    string
}

type ParsedBookInfo struct {
	Title      string
	Authors    []ParsedAuthor
	WorkForm   string
	Genres     []string
	Series     string
	SeriesN    string
	Tags       []string
	Date       string
	Annotation string
	Chapters   []ParsedChapter
	IsDone     bool
}
