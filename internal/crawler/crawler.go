package crawler

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/sirupsen/logrus"

	"github.com/rocky114/craftsman/internal/storage"
)

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

func Crawl(ctx context.Context, code string, admissionTime string) error {
	if crawler, ok := collection[code]; !ok {
		return fmt.Errorf("can't find code: %s", code)
	} else {
		params := storage.GetAdmissionMajorByUniversityAndTimeParams{
			University:    crawler.getUniversityName(),
			AdmissionTime: admissionTime,
		}
		if result, err := storage.GetQueries().GetAdmissionMajorByUniversityAndTime(ctx, params); err != nil && err != sql.ErrNoRows {
			return err
		} else {
			if result.AdmissionTime == admissionTime {
				logrus.Infof("%s university %s admission major data already exist", crawler.getUniversityName(), admissionTime)
				return nil
			}
		}

		crawler.setAdmissionTime(admissionTime)

		if err := crawler.crawl(ctx); err != nil {
			return err
		}

		if result, err := storage.GetQueries().GetAdmissionMajorByUniversityAndTime(ctx, params); err != nil {
			logrus.Errorf("get admission major university %s admission_time %s err: %v", crawler.getUniversityName(), admissionTime, err)
		} else {
			if result.AdmissionTime == admissionTime {
				params := storage.UpdateUniversityLastAdmissionTimeParams{
					LastAdmissionTime: admissionTime,
					Code:              code,
				}
				if err = storage.GetQueries().UpdateUniversityLastAdmissionTime(ctx, params); err != nil {
					logrus.Errorf("update university %s las_admission_time err: %v", crawler.getUniversityName(), err)
				}
			}
		}

		return nil
	}
}
