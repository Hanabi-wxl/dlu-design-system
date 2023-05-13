package service

var userService UserService

type UserService struct {
	InfoService
	ManageService
	LoginService
	RoleService
}

func GetUserService() func() UserService {
	return func() UserService {
		infoService := new(InfoServiceImpl)
		loginService := new(LoginServiceImpl)
		roleService := new(RoleServiceImpl)
		manageService := new(ManageServiceImpl)

		userService.InfoService = infoService
		userService.RoleService = roleService
		userService.LoginService = loginService
		userService.ManageService = manageService
		return userService
	}
}
