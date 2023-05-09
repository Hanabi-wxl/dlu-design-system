package service

import (
	"github.com/Hanabi-wxl/dlu-design-system/pkg/consts"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type LoginServiceImpl struct {
}

func (l LoginServiceImpl) CheckRole() {
}

type claims struct {
	UserId int64
	RoleId int8
	jwt.StandardClaims
}

// generateToken 根据username生成一个token
func generateToken(username string) string {
	// todo 根据用户名查用户

	token := newToken()
	return token
}

// newToken 根据信息创建token
func newToken() string {
	expiresTime := time.Now().Unix() + consts.DefaultExpiresTime
	userClaims := claims{
		UserId: 1,
		RoleId: 2,
		StandardClaims: jwt.StandardClaims{
			Audience:  "123",
			ExpiresAt: expiresTime,
			IssuedAt:  time.Now().Unix(),
			Issuer:    "dlu-design-system",
			NotBefore: time.Now().Unix(),
			Subject:   "token",
		},
	}
	var jwtSecret = []byte(consts.JwtSecret)
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, userClaims)
	if token, err := tokenClaims.SignedString(jwtSecret); err == nil {
		token = "Bearer " + token
		return token
	} else {
		println("generate token fail\n")
		return "fail"
	}
}

// CheckPassword 检验密码
func checkPassword(passwordInput, passwordSave string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(passwordInput), []byte(passwordSave))
	return err == nil
}
