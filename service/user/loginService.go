package service

import "github.com/Hanabi-wxl/dlu-design-system/pkg/errno"

type LoginService interface {
	// CheckLoginRole
	// @Description: 判断登陆的用户角色
	// @param number 学号工号
	// @return bool true为教师 false为学生
	// @return *errno.Errno 错误
	CheckLoginRole(number string) (map[string]bool, *errno.Errno)

	// StudentLogin
	// @Description: 学生登陆
	// @param number 学号
	// @param password 密码
	// @return map[string]interface{} 响应
	// @return *errno.Errno 异常
	StudentLogin(number, password string) (map[string]interface{}, *errno.Errno)

	// TeacherLogin
	// @Description: 教师登陆
	// @param number 工号
	// @param password 密码
	// @return map[string]interface{} 响应
	// @return *errno.Errno 异常
	TeacherLogin(number, password string) (map[string]interface{}, *errno.Errno)
}
