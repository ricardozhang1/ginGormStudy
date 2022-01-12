package utils

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

/*
处理密码加密
*/

// 对明文密码进行简单加密
func HashPassword(password string) (string, error) {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("failed to hash password, err: %v", err)
	}
	return string(hashPassword), nil
}

// 用于登录时 明文密码 与 登录密码进行验证
func CheckPassword(password string, hashPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password))
}

