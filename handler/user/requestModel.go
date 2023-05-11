package handler

type IdRoleRequest struct {
	Id    int64 `uri:"id"`
	IsStu int8  `uri:"isStu"`
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
	CollegeId int  `uri:"collegeId"`
	IsStu     int8 `uri:"isStu"`
	Size      int  `uri:"size"`
	Num       int  `uri:"num"`
}

type MajorIdRequest struct {
	MajorId int  `uri:"majorId"`
	IsStu   int8 `uri:"isStu"`
	Size    int  `uri:"size"`
	Num     int  `uri:"num"`
}
