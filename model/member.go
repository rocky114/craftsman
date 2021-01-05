package model

type Member struct {
	CommonColumn
	Name      string `json:"name"`
	Password  string `json:"name"`
	Nickname  string `json:"nickname"`
	Email     string `json:"email"`
	Telephone string `json:"telephone"`
	Status    string `json:"status"`
	IsAdmin   uint8  `json:"is_admin"`
	IsSuper   uint8  `json:"is_super"`
}

func (Member) TableName() string {
	return "member"
}
