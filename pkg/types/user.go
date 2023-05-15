package types

// UpdateRoleRequest
// @Description: 更新角色请求
type UpdateRoleRequest struct {
	Id     int64 `json:"id" binding:"required"`
	RoleId int64 `json:"role_id" binding:"required,max=7,min=1"`
}

// UserLoginRequest
// @Description: 用户登陆请求
type UserLoginRequest struct {
	Number   string `json:"number" binding:"required,max=8,min=8"`
	Password string `json:"password" binding:"required,max=16,min=8"`
}

// AddUserRequest
// @Description: 用户添加更新请求
type AddUserRequest struct {
	Number    string  `json:"number" binding:"required"`
	Name      string  `json:"name" binding:"required"`
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
	IsStu     int8    `json:"is_stu" binding:"required"`
}

// UpdateUserRequest
// @Description: 用户添加更新请求
type UpdateUserRequest struct {
	ID        int64   `json:"id" binding:"required"`
	Number    string  `json:"number" binding:"required"`
	Name      string  `json:"name" binding:"required"`
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
	IsStu     int8    `json:"is_stu" binding:"required"`
}

// UserResp
// @Description: 用户请求
type UserResp struct {
	ID          int64   `json:"id"`
	Number      string  `json:"number"`
	Name        string  `json:"name"`
	Phone       *string `json:"phone"`
	Email       *string `json:"email"`
	CollegeID   int64   `json:"college_id"`
	CollegeName string  `json:"college_name"`
	MajorID     int64   `json:"major_id"`
	MajorName   string  `json:"major_name"`
	IsStu       int8    `json:"is_stu"`
	Student
	Teacher
}
type Student struct {
	ClassID   int64  `json:"class_id"`
	ClassName string `json:"class_name"`
}
type Teacher struct {
	RoleId     int    `json:"role_id"`
	TitleID    int64  `json:"title_id"`
	TitleName  string `json:"title_name"`
	DegreeID   int64  `json:"degree_id"`
	DegreeName string `json:"degree_name"`
	OfficeID   int64  `json:"office_id"`
	OfficeName string `json:"office_name"`
}

// RoleReq
// @Description: 角色
type RoleReq struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

// RoleIdReq
// @Description: 角色id
type RoleIdReq struct {
	RoleId int64 `uri:"roleId"`
}

// UserQueryByMajorReq
// @Description: 根据专业查询用户
type UserQueryByMajorReq struct {
	PageRequest
	MajorId int64 `uri:"majorId"`
	IsStu   int8  `uri:"isStu"`
}

// UserQueryByCollegeReq
// @Description: 根据学院查询用户
type UserQueryByCollegeReq struct {
	PageRequest
	CollegeId int64 `uri:"collegeId"`
	IsStu     int8  `uri:"isStu"`
}
