package model

type Member struct {
	IdModel
	Name      string `json:"name"`
	Password  string `json:"password"`
	Nickname  string `json:"nickname"`
	Email     string `json:"email"`
	Telephone string `json:"telephone"`
	Status    string `json:"status" gorm:"default:active"`
	IsAdmin   uint8  `json:"is_admin" gorm:"default:0"`
	IsSuper   uint8  `json:"is_super" gorm:"default:0"`
	TimeModel
}

func (Member) TableName() string {
	return "member"
}
