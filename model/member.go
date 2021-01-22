package model

type Member struct {
	IdModel
	Username  string `json:"username"`
	Password  string `json:"password"`
	Nickname  string `json:"nickname" gorm:"default:null"`
	Avatar    string `json:"avatar" gorm:"default:null"`
	Email     string `json:"email" gorm:"default:''"`
	Telephone string `json:"telephone" gorm:"default:''"`
	Status    string `json:"status" gorm:"default:active"`
	IsAdmin   uint8  `json:"is_admin" gorm:"default:0"`
	IsSuper   uint8  `json:"is_super" gorm:"default:0"`
	TimeModel
}

func (Member) TableName() string {
	return "member"
}
