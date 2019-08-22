package model

import (
	"filestore-server/db"
	"fmt"
	"time"
)

type User struct {
	ID      uint       `gorm:"primary_key"`
	Cellphone  string   `gorm:"type:varchar(20);not null;unique_index"`
	Username   string   `gorm:"type:varchar(64);not null;"`
	Password string     `gorm:"type:varchar(64);not null;"`
	Email string         `gorm:"type:varchar(64);not null;"`
	RegisterAt time.Time   `gorm:"type:datetime;not null;"`
	LastLogin time.Time    `gorm:"type:datetime;not null;"`
	Profile string       `gorm:"type:varchar(512);not null;"`
	State byte   	 `gorm:"type:tinyint(4);not null;"`
}

func GetUserByPhone(cellphone string) *User {
	user := new(User)
	if err := db.Conn().Where("cellphone = ?", cellphone).First(&user).Error; err != nil {
		fmt.Println(err)
	}
	return user
}


