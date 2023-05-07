package scraper

import "database/sql"

func NullString(str string) sql.NullString {
	return sql.NullString{String: str, Valid: true}
}

type crawler struct {
	university string
}
