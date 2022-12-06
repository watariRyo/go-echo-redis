package model

type User struct {
	ID       int    `json:"id" gorm:"primaly_key"`
	Name     string `json:"name"`
	Password string `json:"password"`
}
