package db

type User struct {
    ID       int    `json:"id" gorm:"praimaly_key"`
    Name     string `json:"name"`
    Password string `json:"password"`
}

func GetUser(u *User) User {
	var user User
	db.Where(u).First(&user)
	return user
}

func CreateUser(u *User) {
	db.Create(u)
}