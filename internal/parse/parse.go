package parse

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/paldraken/book_thief/internal/parse/authortoday"
	"github.com/paldraken/book_thief/internal/parse/types"
)

type parse interface {
	Parse(url string) (*types.BookData, error)
}

type LibNotSupported struct {
	Domain string
}

func (e *LibNotSupported) Error() string {
	return fmt.Sprintf("Library %s not supported", e.Domain)
}

func Newparse(url string) (parse, error) {
	m := map[string]func() parse{
		"author.today": func() parse { return &authortoday.AT{} },
	}

	domain, err := parseDomainFromUrl(url)
	if err != nil {
		return nil, err
	}

	f, ok := m[domain]
	if !ok {
		return nil, &LibNotSupported{Domain: domain}
	}
	return f(), nil
}

func parseDomainFromUrl(libUrl string) (string, error) {
	u, err := url.Parse(libUrl)
	if err != nil {
		return "", err
	}
	parts := strings.Split(u.Hostname(), ".")
	domain := parts[len(parts)-2] + "." + parts[len(parts)-1]
	return domain, nil
}
