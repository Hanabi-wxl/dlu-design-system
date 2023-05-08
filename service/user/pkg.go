package service

type UserService struct {
	UserInfoService
	LoginService
	RoleService
}

func GetUserService() UserService {
	var userInfoService UserInfoServiceImpl
	var roleService RoleServiceImpl
	var loginService LoginServiceImpl
	userService := UserService{}
	userService.UserInfoService = userInfoService
	userService.RoleService = roleService
	userService.LoginService = loginService
	return userService
}
