package export

import (
	"fmt"

	"github.com/paldraken/book_thief/internal/export/fb2"
	"github.com/paldraken/book_thief/internal/parse/types"
)

const (
	FORMAT_FB2  = "FB2"
	FORMAT_EPUB = "EPUB"
	FORMAT_MOBI = "MOBY"
)

type InvalidFormat struct {
	Format string
}

func (e *InvalidFormat) Error() string {
	return fmt.Sprintf("Format %s not supported", e.Format)
}

type exporter interface {
	Export(book *types.BookData) ([]byte, error)
}

func ToFormat(book *types.BookData, format string) ([]byte, error) {
	m := map[string]func() exporter{
		FORMAT_FB2: func() exporter { return &fb2.FB2{} },
	}

	if format == "" {
		format = FORMAT_FB2
	}

	f, ok := m[format]

	if !ok {
		return nil, &InvalidFormat{Format: format}
	}

	res, err := f().Export(book)

	if err != nil {
		return nil, err
	}
	return res, nil
}
