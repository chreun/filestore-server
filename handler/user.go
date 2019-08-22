package handler

import (
	"filestore-server/db"
	"filestore-server/model"
	"filestore-server/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
	"regexp"
	"strconv"
	"time"
)


func DoLoginHandler(c *gin.Context)  {
	username :=  c.PostForm("username")
	password :=  c.PostForm("password")
	if code, errMsg := doCheckParam(username, password); code > 0 {
		c.JSON(http.StatusOK, ResFmt(code, errMsg, nil))
		return
	}
	user := model.GetUserByPhone(username)
	if user.ID == 0 {
		c.JSON(http.StatusOK, ResFmt(UserNotExist, "用户不存在", nil))
		return
	}
	if !util.VerifyPassword(password, user.Password) {
		c.JSON(http.StatusOK, ResFmt(PasswordNotEqual, "密码错误", nil))
		return
	}
	token := util.Md5String(util.CurDateTime() + username + strconv.Itoa(rand.Int()))
	model.SetToken(user.ID, token)
	data := map[string]string{"token":token}
	c.JSON(http.StatusOK, ResFmt(AllOk, "", data))
}

func checkUserName(username string) bool  {
	match, _ := regexp.MatchString(`^1([38][0-9]|14[57]|5[^4])\d{8}$`, username)
	return match
}

func doCheckParam(username string, password string) (int, string) {
	if !checkUserName(username) {
		return UsernameError, "账户格式错误,请填写正确手机号"
	}
	if len(password) < 6 {
		return PasswordError, "密码长度小于6位"
	}
	return 0, ""
}


func DoRegisterHandler(c *gin.Context)  {
	username :=  c.PostForm("username")
	password :=  c.PostForm("password")
	repeatPassword :=  c.PostForm("repeat_password")
	if code, errMsg := doCheckParam(username, password); code > 0 {
		c.JSON(http.StatusOK, ResFmt(code, errMsg, nil))
		return
	}
	dbUser := model.GetUserByPhone(username)
	if dbUser.ID > 0 {
		c.JSON(http.StatusOK, ResFmt(UserExisted, "用户已存在,请登录", nil))
		return
	}
	if password != repeatPassword {
		c.JSON(http.StatusOK, ResFmt(PasswordRepeatError, "两次密码不一致", nil))
		return
	}
	user := model.User{
		Username: "用户" + username[len(username) - 4:],
		Cellphone: username,
		Password: util.HashPassword(password),
		RegisterAt: time.Now(),
		LastLogin: time.Now(),
	}
	if err := db.Conn().Create(&user).Error;err != nil {
		fmt.Println(err)
		c.JSON(http.StatusOK, ResFmt(InsertUserError, "插入用户失败", nil))
		return
	}
	c.JSON(http.StatusOK, ResFmt(AllOk, "", nil))

}