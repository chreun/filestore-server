package util

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"io"
	"os"
	"time"
)

const TimeFormat = "2006-01-02 15:04:05"

func Md5File(file *os.File) (string, error) {
	hash := md5.New()
	_, err := io.Copy(hash, file)
	if err != nil {
		return "", errors.New("md5文件copy失败")
	}
	return hex.EncodeToString(hash.Sum(nil)), nil
}

func Md5String(str string) string {
	hash := md5.New()
	hash.Write([]byte(str))
	return hex.EncodeToString(hash.Sum(nil))
}

func GetTimestamp(datetime string) int64 {
	loc, _ := time.LoadLocation("Local") //获取时区
	tmp, _ := time.ParseInLocation(TimeFormat, datetime, loc)
	return tmp.Unix()
}

func CurTimestamp() int64 {
	return time.Now().Local().Unix()
}

func CurDateTime() string {
	return time.Now().Local().Format(TimeFormat)
}

func FileExist(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	return false
}

func HashPassword(password string) string  {
	hashStr, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashStr)
}

func VerifyPassword(password string, hashPassword string) bool  {
	err := bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password))
	if err == nil {
		return true
	}
	return false
}