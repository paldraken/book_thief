package fb2

import "encoding/xml"

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
	BookTitle  string        `xml:"book-title"`
	Author     *fb2Author    `xml:"author"`
	Annotation string        `xml:"annotation"`
	Date       string        `xml:"date"`
	Lang       string        `xml:"lang"`
	Email      string        `xml:"email"`
	Coverpage  *fb2CoverPage `xml:"coverpage"`
}

type fb2Author struct {
	FirstName string `xml:"first-name"`
	LastName  string `xml:"last-name"`
	Nickname  string `xml:"nickname"`
}

type fb2Section struct {
	XMLName xml.Name `xml:"section"`

	Title struct {
		P string `xml:"p"`
	} `xml:"title"`
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
	Binary      []*fb2Binary   `xml:"binary"`
}

type fb2Binary struct {
	Id          string `xml:"id,attr"`
	ContentType string `xml:"content-type:l,attr"`
	Data        string `xml:",innerxml"`
}

type fb2CoverPage struct {
	Image *fb2Image `xml:"image"`
}

type fb2Image struct {
	LHref string `xml:"l:href,attr"`
}
