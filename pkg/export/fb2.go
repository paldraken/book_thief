package export

import (
	"encoding/xml"
	"time"

	"github.com/paldraken/book_thief/pkg/parse/types"
)

type fb2Description struct {
	TitleInfo    *fb2TitleInfo    `xml:"title-info"`
	DocumentInfo *fb2DocumentInfo `xml:"document-info"`
}

type fb2DocumentInfo struct {
	Author      *fb2Author `xml:"author"`
	ProgramUsed string     `xml:"program-used"`
	Date        string     `xml:"date"`
	SrcUrl      string     `xml:"src-url"`
}

type fb2TitleInfo struct {
	BookTitle  string     `xml:"book-title"`
	Author     *fb2Author `xml:"author"`
	Annotation string     `xml:"annotation"`
	Date       string     `xml:"date"`
	Lang       string     `xml:"lang"`
}

type fb2Author struct {
	FirstName string `xml:"first-name"`
	LastName  string `xml:"last-name"`
	Nickname  string `xml:"nickname"`
}

type fb2Section struct {
	Title   string `xml:"title"`
	Content string `xml:",innerxml"`
}

type fb2Body struct {
	Title    string        `xml:"title"`
	Sections []*fb2Section `xml:"section"`
}

type fb2fictionBook struct {
	XMLName     xml.Name       `xml:"FictionBook"`
	XMLlns      string         `xml:"xmlns,attr"`
	XMLlnsl     string         `xml:"xmlns:l,attr"`
	Description fb2Description `xml:"description"`
	Body        fb2Body        `xml:"body"`
}

type fb2 struct{}

func (f *fb2) Export(book *types.BookData) ([]byte, error) {
	res := &fb2fictionBook{
		XMLlns:  "http://www.gribuser.ru/xml/fictionbook/2.0",
		XMLlnsl: "http://www.w3.org/1999/xlink",
	}

	res.Description = description(book)
	res.Body = body(book)

	out, _ := xml.MarshalIndent(res, " ", "  ")
	out = append([]byte(xml.Header), out...)

	return out, nil
}

func description(book *types.BookData) fb2Description {
	var afn string
	var aln string
	if len(book.Authors) > 0 {
		afn = book.Authors[0].Firstname
		aln = book.Authors[0].Lastname
	}

	return fb2Description{
		TitleInfo: &fb2TitleInfo{
			Author: &fb2Author{
				FirstName: afn,
				LastName:  aln,
			},
			BookTitle:  book.Title,
			Annotation: book.Annotation,
			Date:       book.Date,
			Lang:       "ru",
		},
		DocumentInfo: &fb2DocumentInfo{
			Author: &fb2Author{
				Nickname: "book thief",
			},
			ProgramUsed: "At book thief",
			SrcUrl:      "http://author.today/",
			Date:        time.Now().Format(time.RFC822),
		},
	}
}

func body(book *types.BookData) fb2Body {
	body := fb2Body{
		Title:    book.Title,
		Sections: []*fb2Section{},
	}

	for _, ch := range book.Chapters {
		section := fb2Section{Title: ch.Title, Content: "\n" + ch.Text}

		body.Sections = append(body.Sections, &section)
	}

	return body
}
