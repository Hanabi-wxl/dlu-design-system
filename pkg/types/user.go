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
