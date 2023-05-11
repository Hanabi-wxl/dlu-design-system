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
	Number    string  `json:"number"`
	Name      string  `json:"name"`
	Phone     *string `json:"phone"`
	Email     *string `json:"email"`
	Gender    int8    `json:"gender"`
	CollegeID int64   `json:"college_id"`
	MajorID   int64   `json:"major_id"`
	ClassID   int64   `json:"class_id"`
	Password  string  `json:"password"`
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
