package utils

import (
	"github.com/Hanabi-wxl/dlu-design-system/pkg/consts"
	service "github.com/Hanabi-wxl/dlu-design-system/service/user"
	"github.com/dgrijalva/jwt-go"
)

// ParseToken 验证用户token
func ParseToken(token string) (*service.Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &service.Claims{}, func(token *jwt.Token) (i interface{}, e error) {
		return consts.JwtSecret, nil
	})
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*service.Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}

// ParseRoleId 解析角色Id
func ParseRoleId(token string) (*int64, error) {
	claims, err := ParseToken(token)
	if err != nil {
		return nil, err
	}
	return &claims.RoleId, nil
}
