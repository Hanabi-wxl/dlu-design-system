package service

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/Hanabi-wxl/dlu-design-system/pkg/consts"
	"github.com/dgrijalva/jwt-go"
	"strconv"
	"time"
)

type LoginServiceImpl struct {
}

func (l LoginServiceImpl) CheckRole() {
}

type Claims struct {
	UserId int64
	RoleId int8
	jwt.StandardClaims
}

// GenerateToken 根据username生成一个token
func GenerateToken(username string) string {
	// todo 根据用户名查用户

	// todo 生成token
	token := NewToken()
	return token
}

// NewToken 根据信息创建token
func NewToken() string {
	expiresTime := time.Now().Unix() + consts.DefaultExpiresTime
	fmt.Printf("expiresTime: %v\n", expiresTime)
	fmt.Printf("id: %v\n", strconv.FormatInt(10, 10))
	claims := Claims{
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
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	if token, err := tokenClaims.SignedString(jwtSecret); err == nil {
		token = "Bearer " + token
		println("generate token success!\n")
		return token
	} else {
		println("generate token fail\n")
		return "fail"
	}
}

// EnCoder 密码加密
func EnCoder(password string) string {
	h := hmac.New(sha256.New, []byte(password))
	sha := hex.EncodeToString(h.Sum(nil))
	fmt.Println("Result: " + sha)
	return sha
}
