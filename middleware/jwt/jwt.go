package jwt

import (
	"github.com/Hanabi-wxl/dlu-design-system/pkg/errno"
	"github.com/Hanabi-wxl/dlu-design-system/pkg/result"
	"github.com/Hanabi-wxl/dlu-design-system/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"strings"
)

// Auth
// @Description: 检验是否携带jwt
// @auth sinre 2023-05-12 22:15:34
// @return gin.HandlerFunc
func Auth() gin.HandlerFunc {
	return func(context *gin.Context) {
		logrus.Info("==============开始认证===========")
		auth := context.Request.Header.Get("Authorization")
		if len(auth) == 0 {
			context.Abort()
			context.JSON(http.StatusUnauthorized, result.NewResult(errno.UnAuthorizationErrCode, errno.UnAuthorizationErrMsg, nil))
		} else {
			auth = strings.Fields(auth)[1]
			claim, err := utils.ParseToken(auth)
			if err != nil {
				context.Abort()
				context.JSON(http.StatusUnauthorized, result.NewResult(errno.UnAuthorizationErrCode, errno.UnAuthorizationErrMsg, nil))
			} else {
				context.Set("userId", claim.Id)
				context.Set("roleId", claim.RoleId)
				context.Set("number", claim.Number)
				logrus.Info("==============认证成功===========")
				context.Next()
			}
		}
	}
}

// NeedAuth
// @Description: 要求最低权限
// @auth sinre 2023-05-12 22:15:50
// @param roleId
// @return gin.HandlerFunc
func NeedAuth(roleId int64) gin.HandlerFunc {
	return func(context *gin.Context) {
		logrus.Info("==============开始认证===========")
		auth := context.Request.Header.Get("Authorization")
		if len(auth) == 0 {
			context.Abort()
			context.JSON(http.StatusUnauthorized, result.NewResult(errno.UnAuthorizationErrCode, errno.UnAuthorizationErrMsg, nil))
		} else {
			auth = strings.Fields(auth)[1]
			claim, err := utils.ParseToken(auth)
			if err != nil {
				context.Abort()
				context.JSON(http.StatusUnauthorized, result.NewResult(errno.UnAuthorizationErrCode, errno.UnAuthorizationErrMsg, nil))
			} else if roleId > claim.RoleId {
				context.Abort()
				context.JSON(http.StatusUnauthorized, result.NewResult(errno.UnAuthorizationErrCode, errno.UnAuthorizationErrMsg, nil))
			} else {
				context.Set("userId", claim.Id)
				context.Set("roleId", claim.RoleId)
				context.Set("number", claim.Number)
				logrus.Info("==============认证成功===========")
				context.Next()
			}
		}
	}
}

// MustAuth
// @Description: 必须为此权限
// @auth sinre 2023-05-12 22:16:05
// @param roleId
// @return gin.HandlerFunc
func MustAuth(roleId int64) gin.HandlerFunc {
	return func(context *gin.Context) {
		logrus.Info("==============开始认证===========")
		auth := context.Request.Header.Get("Authorization")
		if len(auth) == 0 {
			context.Abort()
			context.JSON(http.StatusUnauthorized, result.NewResult(errno.UnAuthorizationErrCode, errno.UnAuthorizationErrMsg, nil))
		} else {
			auth = strings.Fields(auth)[1]
			claim, err := utils.ParseToken(auth)
			if err != nil {
				context.Abort()
				context.JSON(http.StatusUnauthorized, result.NewResult(errno.UnAuthorizationErrCode, errno.UnAuthorizationErrMsg, nil))
			} else if roleId == claim.RoleId {
				context.Abort()
				context.JSON(http.StatusUnauthorized, result.NewResult(errno.UnAuthorizationErrCode, errno.UnAuthorizationErrMsg, nil))
			} else {
				context.Set("userId", claim.Id)
				context.Set("roleId", claim.RoleId)
				context.Set("number", claim.Number)
				logrus.Info("==============认证成功===========")
				context.Next()
			}
		}
	}
}
