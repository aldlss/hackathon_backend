package model

type User struct {
	Id       int64  `gorm:"primaryKey;column:id"`
	Username string `gorm:"unique;column:username"`
	Password string `gorm:"column:password"`
}
