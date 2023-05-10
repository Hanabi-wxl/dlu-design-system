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

func Auth(roleId int64) gin.HandlerFunc {
	return func(context *gin.Context) {
		logrus.Info("==============开始认证===========")
		auth := context.Request.Header.Get("Authorization")
		if len(auth) == 0 {
			context.Abort()
			context.JSON(http.StatusUnauthorized, result.NewResult(errno.UnAuthorizationErrCode, errno.UnAuthorizationErrMsg, nil))
		}
		auth = strings.Fields(auth)[1]
		token, err := utils.ParseToken(auth)
		if err != nil {
			context.Abort()
			context.JSON(http.StatusUnauthorized, result.NewResult(errno.UnAuthorizationErrCode, errno.UnAuthorizationErrMsg, nil))
		}
		if roleId > token.RoleId {
			context.Abort()
			context.JSON(http.StatusUnauthorized, result.NewResult(errno.UnAuthorizationErrCode, errno.UnAuthorizationErrMsg, nil))
		}
		context.Set("userId", token.Id)
		logrus.Info("==============认证成功===========")
		context.Next()
	}
}

func MustAuth(roleId int64) gin.HandlerFunc {
	return func(context *gin.Context) {
		logrus.Info("==============开始认证===========")
		auth := context.Request.Header.Get("Authorization")
		if len(auth) == 0 {
			context.Abort()
			context.JSON(http.StatusUnauthorized, result.NewResult(errno.UnAuthorizationErrCode, errno.UnAuthorizationErrMsg, nil))
		}
		auth = strings.Fields(auth)[1]
		token, err := utils.ParseToken(auth)
		if err != nil {
			context.Abort()
			context.JSON(http.StatusUnauthorized, result.NewResult(errno.UnAuthorizationErrCode, errno.UnAuthorizationErrMsg, nil))
		}
		if roleId != token.RoleId {
			context.Abort()
			context.JSON(http.StatusUnauthorized, result.NewResult(errno.UnAuthorizationErrCode, errno.UnAuthorizationErrMsg, nil))
		}
		context.Set("userId", token.Id)
		logrus.Info("==============认证成功===========")
		context.Next()
	}
}
