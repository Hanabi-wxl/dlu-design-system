package types

// PageRequest
// @Description: 分页请求
type PageRequest struct {
	Size int `uri:"size" binding:"required"`
	Num  int `uri:"num" binding:"required"`
}

// IdRequest
// @Description: 用户id
type IdRequest struct {
	Id int64 `uri:"id" binding:"required"`
}

// IdRoleRequest
// @Description: 用户id-用户角色
type IdRoleRequest struct {
	Id    int64 `uri:"id" binding:"required"`
	IsStu int8  `uri:"isStu" binding:"required,max=2,min=1"`
}

// RoleIdRequest
// @Description: 角色id
type RoleIdRequest struct {
	RoleId int64 `uri:"roleId" binding:"required,max=7,min=1"`
}

// NumberRequest
// @Description: 账号-角色
type NumberRequest struct {
	Number string `uri:"number" binding:"required"`
	IsStu  int8   `uri:"isStu" binding:"required,max=2,min=1"`
}
