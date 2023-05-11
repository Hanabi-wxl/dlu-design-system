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
}

type UpdateRoleRequest struct {
	Id     int64 `json:"id"`
	RoleId int64 `json:"role_id"`
}

type UserLoginRequest struct {
	Number   string `json:"number"`
	Password string `json:"password"`
}
