package model

type Content struct {
	Id     int64  `gorm:"primaryKey;column:id"`
	Title  string `gorm:"column:title"`
	Desc   string `gorm:"column:desc"`
	UserId int64  `gorm:"column:user_id"`
}
