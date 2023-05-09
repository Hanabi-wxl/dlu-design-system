package service

var userService UserService

type UserService struct {
	UserInfoService
	LoginService
	RoleService
}

func GetUserService() func() UserService {
	return func() UserService {
		userInfoService := new(UserInfoServiceImpl)
		loginService := new(LoginServiceImpl)
		roleService := new(RoleServiceImpl)
		userService.UserInfoService = userInfoService
		userService.RoleService = roleService
		userService.LoginService = loginService
		return userService
	}
}
