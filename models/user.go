package models

import (
	"aschool/conn"
	"crypto/md5"
	"encoding/hex"
)

type User struct {
	Id         int    `json:"id"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	Email      string `json:"email"`
	RoleId     int    `json:"roleId"`
	LoginCount int    `json:"login_count"`
	LastTime   int64  `json:"last_time"`
	LastIp     string `json:"last_ip"`
	State      int    `json:"state"`
	Created    int64  `json:"created"`
	Updated    int64  `json:"updated"`
}

const secret = "wenjie.blog.csdn.net"

// 加密密码
func EncryptPassword(data []byte) (result string) {
	h := md5.New()
	h.Write([]byte(secret))
	return hex.EncodeToString(h.Sum(data))
}

// 登录
func Login(userName string, password string) (user []*User, err error) {

	// 生成加密密码
	db := conn.DB
	db = db.Where("username = ?", userName)
	db = db.Where("password = ?", EncryptPassword([]byte(password)))
	if err = db.Find(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

// 注册
func CreateUser(user *User) (err error) {
	err = conn.DB.Create(&user).Error
	return
}
