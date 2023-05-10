package service

var userService UserService

type UserService struct {
	InfoService
	ManageService
	LoginService
	RoleService
}

type UserRequest struct {
	ID        int64   `json:"id"`
	Number    string  `json:"number" binding:"min=8,max=8,required"`
	Name      string  `json:"name" binding:"required,max=16"`
	Phone     *string `json:"phone" binding:"required"`
	Email     *string `json:"email" binding:"required"`
	Gender    int8    `json:"gender" binding:"required"`
	CollegeID int64   `json:"college_id" binding:"required"`
	MajorID   int64   `json:"major_id" binding:"required"`
	ClassID   int64   `json:"class_id"`
	Password  string  `json:"password" binding:"required"`
	TitleID   int64   `json:"title_id"`
	DegreeID  int64   `json:"degree_id"`
	OfficeID  int64   `json:"office_id"`
	IsStu     int8    `json:"is_stu"`
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
