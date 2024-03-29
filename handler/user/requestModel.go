package handler

type IdRoleRequest struct {
	Id    int64 `uri:"id"`
	IsStu int8  `uri:"isStu"`
}

type IdRequest struct {
	Id int64 `uri:"id"`
}

type RoleIdRequest struct {
	RoleId int64 `uri:"roleId"`
}

type NumberRequest struct {
	Number string `uri:"number"`
	IsStu  int8   `uri:"isStu"`
}

type UpdateRoleRequest struct {
	Id     int64 `json:"id"`
	RoleId int64 `json:"role_id"`
}

type UserLoginRequest struct {
	Number   string `json:"number"`
	Password string `json:"password"`
}

type CollegeIdRequest struct {
	CollegeId int64 `uri:"collegeId"`
	IsStu     int8  `uri:"isStu"`
	Size      int   `uri:"size"`
	Num       int   `uri:"num"`
}

type MajorIdRequest struct {
	MajorId int64 `uri:"majorId"`
	IsStu   int8  `uri:"isStu"`
	Size    int   `uri:"size"`
	Num     int   `uri:"num"`
}
