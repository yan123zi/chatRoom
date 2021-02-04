package model

type User struct {
	Username    string `json:"username" gorm:"column:"`
	Password    string `json:"password"`
	Avatar      string `json:"avatar"`
	Description string `json:"description"`
}
