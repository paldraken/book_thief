package fb2

import (
	"encoding/base64"
	"encoding/xml"
	"fmt"
	"strings"
	"time"

	"github.com/paldraken/book_thief/pkg/export/prepare"
	"github.com/paldraken/book_thief/pkg/parse/types"
)

type FB2 struct{}

func (f *FB2) Export(book *types.BookData) ([]byte, error) {
	res := &fb2fictionBook{
		XMLlns:  "http://www.gribuser.ru/xml/fictionbook/2.0",
		XMLlnsl: "http://www.w3.org/1999/xlink",
	}

	res.Description = description(book)
	res.Body = body(book)
	res.Binary = images(book)

	out, _ := xml.MarshalIndent(res, " ", "  ")
	out = append([]byte(xml.Header), out...)

	contentStr := string(out)
	for _, img := range book.Images {
		placeholder := fmt.Sprintf("#__bt_binary__#%s#", img.Id)
		replace := fmt.Sprintf("<image l:href=\"#%s\"></image>", img.Id)
		contentStr = strings.Replace(contentStr, placeholder, replace, 1)
	}

	return []byte(contentStr), nil
}

func images(book *types.BookData) []*fb2Binary {
	result := []*fb2Binary{}
	for _, img := range book.Images {
		result = append(result, &fb2Binary{
			Data:        base64.StdEncoding.EncodeToString(img.Data),
			Id:          img.Id,
			ContentType: img.ContentType,
		})
	}
	return result
}

func description(book *types.BookData) fb2Description {
	var afn, aln, nn string
	if len(book.Authors) > 0 {
		afn = book.Authors[0].Firstname
		aln = book.Authors[0].Lastname
		nn = book.Authors[0].NickName

		if len(aln) == 0 && len(afn) == 0 {
			aln = book.Authors[0].FullName
		}
	}

	titleInfo := &fb2TitleInfo{
		Author: &fb2Author{
			FirstName: afn,
			LastName:  aln,
			Nickname:  nn,
		},
		BookTitle:  book.Title,
		Annotation: book.Annotation,
		Date:       book.Date,
		Lang:       "ru",
	}

	if book.CoverId != "" {
		titleInfo.Coverpage = &fb2CoverPage{
			Image: &fb2Image{
				LHref: "#" + book.CoverId,
			},
		}
	}

	return fb2Description{
		TitleInfo: titleInfo,
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

		text := prepare.SanitizeForFB2(ch.Text)

		section := fb2Section{
			Title: struct {
				P string `xml:"p"`
			}{
				P: ch.Title,
			},
			Content: "\n" + text,
		}
		body.Sections = append(body.Sections, &section)
	}

	return body
}
