package model

import (
	"filestore-server/db"
	"time"
)

type UserToken struct {
	UserId    uint    `gorm:"primary_key"`
	Token  string   `gorm:"type:varchar(32);not null;unique_index"`
	UpdateAt time.Time   `gorm:"type:datetime;not null;"`
}

func SetToken(userId uint, token string) {
	ut := new(UserToken)
	db.Conn().Where("user_id = ?", userId).First(&ut)
	ut.UserId = userId
	ut.Token = token
	ut.UpdateAt = time.Now()
	db.Conn().Save(&ut)
}

func GetUserId(token string) uint {
	tk := new(UserToken)
	db.Conn().Where("token = ?", token).First(&tk)
	return tk.UserId
}