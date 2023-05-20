package crawler

import (
	"context"
	"fmt"
	"time"

	"github.com/rocky114/craftsman/internal/storage"
)

var collection = make(map[string]impl)

type impl interface {
	crawl(ctx context.Context) error
	getLastAdmissionTime() string
}

type university struct {
	name              string
	code              string
	lastAdmissionTime string
}

func (u *university) getLastAdmissionTime() string {
	return u.lastAdmissionTime
}

func (u *university) crawl(ctx context.Context) error {
	return fmt.Errorf("university: %s not implment crawl interface", u.name)
}

func Crawl(ctx context.Context, code string) error {
	if crawler, ok := collection[code]; !ok {
		return fmt.Errorf("can't find code: %s", code)
	} else {
		if lastAdmissionTime, err := storage.GetQueries().GetUniversityLastAdmissionTime(ctx, code); err != nil {
			return err
		} else {
			if lastAdmissionTime == time.Now().AddDate(-1, 0, 0).Format("2006") {
				return nil
			}
		}

		if err := crawler.crawl(ctx); err != nil {
			return err
		}

		param := storage.UpdateUniversityLastAdmissionTimeParams{
			LastAdmissionTime: crawler.getLastAdmissionTime(),
			Code:              code,
		}

		return storage.GetQueries().UpdateUniversityLastAdmissionTime(ctx, param)
	}
}

func containAdmissionTime(admissionTime string) bool {
	currentTime := time.Now()
	years := []string{
		currentTime.AddDate(-2, 0, 0).Format("2006"),
		currentTime.AddDate(-1, 0, 0).Format("2006"),
	}

	for _, year := range years {
		if admissionTime == year {
			return true
		}
	}

	return false
}
