package crawler

import (
	"errors"
	"fmt"
	"time"
)

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

func ContainAdmissionTime(admissionTime string) bool {
	currentTime := time.Now()
	years := []string{
		currentTime.AddDate(-2, 0, 0).Format("2006-01-02"),
		currentTime.AddDate(-1, 0, 0).Format("2006-01-02"),
	}

	for _, year := range years {
		if admissionTime == year {
			return true
		}
	}

	return false
}
