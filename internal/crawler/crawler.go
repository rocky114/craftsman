package crawler

import (
	"context"
	"fmt"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/rocky114/craftsman/internal/storage"
)

var collection = make(map[string]impl)

type impl interface {
	crawl(ctx context.Context) error
	getLastAdmissionTime() string
	getUniversityName() string
}

type university struct {
	name              string
	code              string
	lastAdmissionTime string
}

func (u *university) getUniversityName() string {
	return u.name
}

func (u *university) getLastAdmissionTime() string {
	return u.lastAdmissionTime
}

func (u *university) crawl(ctx context.Context) error {
	return fmt.Errorf("university: %s not implment crawl interface", u.name)
}

func Crawl(ctx context.Context, code string, admissionTime string) error {
	if crawler, ok := collection[code]; !ok {
		return fmt.Errorf("can't find code: %s", code)
	} else {
		if lastAdmissionTime, err := storage.GetQueries().GetAdmissionTimeByUniversityName(ctx, crawler.getUniversityName()); err != nil {
			return err
		} else {
			if lastAdmissionTime == admissionTime {
				logrus.Infof("%s university %s admission major data already exist", crawler.getUniversityName(), lastAdmissionTime)
				return nil
			}
		}

		if err := crawler.crawl(ctx); err != nil {
			return err
		}

		return nil
	}
}

func containAdmissionTime(admissionTime string) bool {
	currentTime := time.Now()
	years := []string{
		currentTime.AddDate(-1, 0, 0).Format("2006"),
	}

	for _, year := range years {
		if admissionTime == year {
			return true
		}
	}

	return false
}
