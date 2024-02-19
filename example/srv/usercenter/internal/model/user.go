package model

import (
	"time"

	"github.com/ryon-wen/own-utils/own"
	"gorm.io/gorm"
)

var db *gorm.DB

func RegisterModel(sql *own.MySQL) {
	db = own.InitModel(sql)
	db.Table("users")
}

type User struct {
	Id        int64
	Nickname  string
	Mobile    string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

func CreateUser(user *User) error {
	return db.Omit("deleted_at").Create(&user).Error
}

func UpdateUser(user *User) error {
	return db.Omit("id", "created_at", "deleted_at").Updates(&user).Error
}

func GetUser(id int64) (*User, error) {
	var u User
	res := db.Where("id = ?", id).First(&u)
	return &u, res.Error
}

func DeleteUser(id int64) error {
	return db.Where("id = ?", id).Update("deleted_at", time.Now()).Error
}
