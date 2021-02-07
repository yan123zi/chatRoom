package model

type User struct {
	Username    string `json:"username" gorm:"column:username"`
	Password    string `json:"password" gorm:"column:password"`
	Avatar      string `json:"avatar" gorm:"column:avatar"`
	Description string `json:"description" gorm:"column:description"`
}
