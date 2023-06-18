package crawler

import (
	"context"
	"fmt"

	"github.com/sirupsen/logrus"

	"github.com/rocky114/craftsman/internal/storage"
)

const (
	userAgent = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/93.0.4577.82 Safari/537.36 Edg/93.0.961.52"
)

var currentIndex int

var collection = make(map[string]impl)

type impl interface {
	crawl(ctx context.Context) error
	setAdmissionTime(admissionTime string)
	getUniversityName() string
}

type university struct {
	name          string
	code          string
	admissionTime string
}

func (u *university) getUniversityName() string {
	return u.name
}

func (u *university) containAdmissionTime(admissionTime string) bool {
	if u.admissionTime == admissionTime {
		return true
	}

	return false
}

func (u *university) setAdmissionTime(admissionTime string) {
	u.admissionTime = admissionTime
}

func (u *university) crawl(ctx context.Context) error {
	return fmt.Errorf("university: %s not implment crawl interface", u.name)
}

func Crawl(ctx context.Context, code string, admissionTime string) (err error) {
	if crawler, ok := collection[code]; !ok {
		return fmt.Errorf("can't find code: %s", code)
	} else {
		logrus.Infof("crawl university: %s, time: %s running", crawler.getUniversityName(), admissionTime)
		defer func() {
			logrus.Infof("crawl university: %s, time: %s finishing", crawler.getUniversityName(), admissionTime)
		}()

		crawler.setAdmissionTime(admissionTime)

		if err = crawler.crawl(ctx); err != nil {
			return err
		}

		params := storage.UpdateUniversityLastAdmissionTimeParams{
			LastAdmissionTime: admissionTime,
			Code:              code,
		}
		if err = storage.GetQueries().UpdateUniversityLastAdmissionTime(ctx, params); err != nil {
			logrus.Errorf("UpdateUniversityLastAdmissionTime university %s err: %v", crawler.getUniversityName(), err)
		}

		return nil
	}
}
