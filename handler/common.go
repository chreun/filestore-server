package handler

import (
	"github.com/gin-gonic/gin"
)

func ResFmt(code int, msg string, data interface{}) interface{} {
	return gin.H{"code": code, "msg": msg, "data": data}
}

const AllOk  = 0

const UsernameError = 101
const PasswordError = 102
const PasswordRepeatError = 103
const UserNotExist = 104
const PasswordNotEqual = 105
const InsertUserError = 106
const UserExisted = 107
