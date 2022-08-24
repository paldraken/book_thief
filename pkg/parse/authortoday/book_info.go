package authortoday

import (
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/paldraken/book_thief/pkg/parse/types"
)

func (at *AT) bookInfo(doc *goquery.Document) (*types.ParsedBookInfo, error) {
	var pbi types.ParsedBookInfo

	pbi.Title = parseTtitle(doc)
	pbi.Authors = parseAuthors(doc)
	pbi.WorkForm, pbi.Genres = parseWorkFormAndGenres(doc)
	pbi.Series, pbi.SeriesN = parseSeries(doc)
	pbi.Tags = parseTags(doc)
	pbi.Date = parseDate(doc)
	pbi.Annotation = parseAnnotation(doc)
	pbi.IsDone = parseIsDone(doc)

	return &pbi, nil
}

func parseTtitle(doc *goquery.Document) string {
	return strings.TrimSpace(doc.Find(".content .book-meta-panel .book-title").Text())
}

func parseAuthors(doc *goquery.Document) []types.ParsedAuthor {
	authors := []types.ParsedAuthor{}
	doc.Find(".content .book-meta-panel .book-authors a").
		EachWithBreak(func(i int, s *goquery.Selection) bool {
			var a types.ParsedAuthor
			str := s.Text()
			arr := strings.Split(str, " ")
			switch {
			case len(arr) == 1:
				a.Firstname = arr[0]
			case len(arr) > 1:
				a.Firstname = arr[0]
				a.Lastname = arr[1]
			}
			authors = append(authors, a)
			return true
		})

	return authors
}

func parseWorkFormAndGenres(doc *goquery.Document) (string, []string) {
	workForm := ""
	genres := []string{}
	doc.Find(".content .book-meta-panel .book-genres a").
		Each(func(i int, s *goquery.Selection) {
			if i == 0 {
				workForm = s.Text()
			} else {
				genres = append(genres, s.Text())
			}
		})
	return workForm, genres
}

func parseSeries(doc *goquery.Document) (string, string) {
	var ser string
	var serN string
	doc.Find(".content .book-meta-panel .book-genres").Siblings().
		Each(func(i int, s *goquery.Selection) {
			check := strings.TrimSpace(s.Find(".text-muted").First().Text())
			if check != "Цикл:" {
				return
			}
			ser = strings.TrimSpace(s.Find("a").First().Text())
			serN = strings.TrimSpace(s.Find("span").Last().Text())
			if serN != "" {
				var rx = regexp.MustCompile(`[^0-9]+`)
				serN = rx.ReplaceAllString(serN, "")
			}
		})
	return ser, serN
}

func parseTags(doc *goquery.Document) []string {
	tags := []string{}
	doc.Find(".content .book-meta-panel .tags a").
		Each(func(i int, s *goquery.Selection) {
			tag := strings.TrimSpace(s.Text())
			if len(tag) > 0 {
				tags = append(tags, tag)
			}
		})
	return tags
}

func parseDate(doc *goquery.Document) string {
	return doc.Find(".content .book-genres+div span:nth-child(3) .hint-top").
		AttrOr("data-time", "")
}

func parseAnnotation(doc *goquery.Document) string {
	return doc.Find("#tab-annotation .annotation").Text()
}

func parseIsDone(doc *goquery.Document) bool {
	el := doc.Find(".book-meta-panel .book-status-icon")
	return el.HasClass("icon-check")
}

func parseChaptersList(doc *goquery.Document, workId string) []rawChapters {

	timestamp := fmt.Sprintf("%d", time.Now().UnixMilli())

	chapters := []rawChapters{}
	doc.Find("#tab-chapters .table-of-content li a").Each(func(i int, s *goquery.Selection) {
		if href, ok := s.Attr("href"); ok {
			tmp := strings.Split(href, "/")
			chapterId := tmp[len(tmp)-1]
			url := fmt.Sprintf(CHAPTER_BASE_URL, workId, chapterId, timestamp)
			chapters = append(chapters, rawChapters{
				strings.TrimSpace(s.Text()),
				url,
			})
		}
	})
	return chapters
}
