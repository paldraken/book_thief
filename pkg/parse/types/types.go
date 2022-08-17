package types

type ParsedAuthor struct {
	Firstname string
	Lastname  string
	NickName  string
	Email     string
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
