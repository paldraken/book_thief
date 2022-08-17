package authortoday

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/paldraken/book_thief/pkg/fetch"
	"github.com/paldraken/book_thief/pkg/parse/types"
)

func chapters(chaptersResp []*fetch.BatchResp, chaptersList []rawChapters) ([]types.ParsedChapter, error) {
	chapters := []types.ParsedChapter{}
	for i, resp := range chaptersResp {
		chText, err := parseChapter(resp.Resp)
		if err != nil {
			return nil, err
		}
		secret := resp.Resp.Header.Get("Reader-Secret")
		chText = decodeText(secret, chText, "")
		chapters = append(chapters, types.ParsedChapter{
			Number: i,
			Title:  chaptersList[i].Title,
			Text:   chText,
		})
	}
	return chapters, nil
}

func parseChapter(resp *http.Response) (string, error) {
	type jsonContent struct {
		Text string `json:"text"`
	}

	type jsonChapter struct {
		IsSuccessful bool        `json:"isSuccessful"`
		IsWarning    bool        `json:"isWarning"`
		Messages     string      `json:"messages"`
		Data         jsonContent `joson:"data"`
	}

	rawJson, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var contentJson jsonChapter
	if err := json.Unmarshal(rawJson, &contentJson); err != nil {
		return "", err
	}

	return contentJson.Data.Text, nil
}

// Алгоритм дешифровки текста главы
func decodeText(secret, text, userId string) string {

	magic := reverseString(secret) + "@_@" + userId

	counter := 0
	result := ""

	for _, c := range text {
		code1 := int(c)
		mIdx := int(float64(counter % len(magic)))
		code2 := int(magic[mIdx])
		newCh := code1 ^ code2
		counter++
		result = result + string(rune(newCh))
	}
	return result
}

func reverseString(str string) string {
	res := ""
	for i := len(str) - 1; i >= 0; i-- {
		res = res + string(str[i])
	}
	return res
}
