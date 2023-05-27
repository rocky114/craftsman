package storage

import (
	"context"
	"database/sql"

	"github.com/sirupsen/logrus"
)

func IsNotFoundAdmissionMajor(ctx context.Context, university, admissionTime string) bool {
	params := GetAdmissionMajorByUniversityAndTimeParams{
		University:    university,
		AdmissionTime: admissionTime,
	}

	_, err := GetQueries().GetAdmissionMajorByUniversityAndTime(ctx, params)

	if err == sql.ErrNoRows {
		return true
	}

	if err != nil {
		logrus.Errorf("GetAdmissionMajorByUniversityAndTime err: %v", err)
	}

	return false
}
