package crawler

import (
	"database/sql"
	"errors"
	"fmt"
)

func NullString(str string) sql.NullString {
	return sql.NullString{String: str, Valid: true}
}

var collection = make(map[string]impl)

type impl interface {
	crawl() error
}

type university struct {
	name string
	code string
}

func (u *university) crawl() error {
	return errors.New("crawl interface not implemented")
}

func Crawl(code string) error {
	if crawler, ok := collection[code]; ok {
		return crawler.crawl()
	}

	return fmt.Errorf("can't find code: %s", code)
}
