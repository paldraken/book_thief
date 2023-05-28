package authortoday

import (
	"fmt"

	"github.com/paldraken/book_thief/internal/parse/authortoday/api"
	"github.com/paldraken/book_thief/internal/parse/types"
)

func (at *AT) bookInfo(bookMeta *api.BookMetaInfo) (*types.BookData, error) {
	var pbi types.BookData

	// @toDo check bookDetail.State for deleted books

	pbi.Title = bookMeta.Title
	pbi.IsDone = bookMeta.IsFinished
	pbi.WorkForm = parseWorkForm(bookMeta)
	pbi.Series = bookMeta.SeriesTitle
	pbi.SeriesN = fmt.Sprintf("%d", bookMeta.SeriesOrder)
	pbi.Date = parseDate(bookMeta)
	pbi.Annotation = bookMeta.Annotation
	pbi.Tags = bookMeta.Tags

	pbi.Authors = parseAuthors(bookMeta)

	return &pbi, nil
}

func parseAuthors(bmi *api.BookMetaInfo) []types.BookAuthor {
	result := make([]types.BookAuthor, 1)

	if bmi.OriginalAuthor != "" { // Translte
		originalAuthor := types.BookAuthor{
			FullName: bmi.OriginalAuthor,
			Type:     types.Author,
		}
		result = append(result, originalAuthor)

		if bmi.Translator != "" {
			author := types.BookAuthor{
				FullName: bmi.Translator,
				Type:     types.Translator,
			}
			result = append(result, author)
		}

		result = append(result, originalAuthor)
	} else {
		author := types.BookAuthor{
			FullName: bmi.AuthorFIO,
			NickName: bmi.AuthorUserName,
			Type:     types.Author,
		}
		result = append(result, author)

		if bmi.CoAuthorConfirmed {
			author := types.BookAuthor{
				FullName: bmi.CoAuthorFIO,
				NickName: bmi.CoAuthorUserName,
				Type:     types.CoAuthor,
			}
			result = append(result, author)
		}

		if bmi.SecondCoAuthorConfirmed {
			author := types.BookAuthor{
				FullName: bmi.SecondCoAuthorFIO,
				NickName: bmi.SecondCoAuthorUserName,
				Type:     types.CoAuthor,
			}
			result = append(result, author)
		}

	}
	return result
}

func parseWorkForm(bmi *api.BookMetaInfo) string {
	switch bmi.WorkForm {
	case 1:
		return "Story"
	case 2:
		return "Novel"
	case 3:
		return "StoryBook"
	case 4:
		return "Poetry"
	case 5:
		return "Translation"
	case 6:
		return "Tale"
	default:
		return "Any"
	}
}

func parseDate(bmi *api.BookMetaInfo) string {
	if bmi.IsFinished {
		return bmi.FinishTime.String()
	} else {
		return bmi.LastUpdateTime.String()
	}
}
