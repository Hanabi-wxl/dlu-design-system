package service

import "sync"

var userService *UserService
var userServiceOnce sync.Once

type UserService struct {
	InfoService
	ManageService
	LoginService
	RoleService
}

func GetUserService() *UserService {
	userServiceOnce.Do(func() {
		userService = &UserService{}
		infoService := new(InfoServiceImpl)
		loginService := new(LoginServiceImpl)
		roleService := new(RoleServiceImpl)
		manageService := new(ManageServiceImpl)

		userService.InfoService = infoService
		userService.RoleService = roleService
		userService.LoginService = loginService
		userService.ManageService = manageService
	})
	return userService
}
