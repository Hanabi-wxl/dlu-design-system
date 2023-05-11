package utils

import (
	"github.com/Hanabi-wxl/dlu-design-system/pkg/consts"
	"github.com/dgrijalva/jwt-go"
	"github.com/sirupsen/logrus"
	"time"
)

type Claims struct {
	UserId int64
	RoleId int64
	jwt.StandardClaims
}

// GenerateToken 根据username生成一个token
func GenerateToken(id, roleId int64) string {
	token := newToken(id, roleId)
	return token
}

// newToken 根据信息创建token
func newToken(id, roleId int64) string {
	expiresTime := time.Now().Unix() + consts.DefaultExpiresTime
	userClaims := Claims{
		UserId: id,
		RoleId: roleId,
		StandardClaims: jwt.StandardClaims{
			Audience:  "dlu-ie-college",
			ExpiresAt: expiresTime,
			IssuedAt:  time.Now().Unix(),
			Issuer:    "dlu-design-system",
			NotBefore: time.Now().Unix(),
			Subject:   "Authorization",
		},
	}
	var jwtSecret = []byte(consts.JwtSecret)
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, userClaims)
	if token, err := tokenClaims.SignedString(jwtSecret); err == nil {
		token = "Bearer " + token
		return token
	} else {
		logrus.Error("generate token fail")
		return "fail"
	}
}

// ParseToken 验证用户token
func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (i interface{}, e error) {
		return []byte(consts.JwtSecret), nil
	})
	if err == nil && tokenClaims != nil {
		if claim, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claim, nil
		}
	}
	return nil, err
}
