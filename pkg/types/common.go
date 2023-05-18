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

// NumberMajorRequest
// @Description: 账号专业角色
type NumberMajorRequest struct {
	Number  string `uri:"number" binding:"required"`
	MajorId int64  `uri:"majorId" binding:"required"`
	IsStu   int8   `uri:"isStu" binding:"required,max=2,min=1"`
}

// PageResp
// @Description: 分页响应
type PageResp struct {
	ItemTotal int64 `json:"item_total"` // 总元素数
	PageTotal int64 `json:"page_total"` // 总页数
	Array     any   `json:"array"`
}

// CollegeIdReq
// @Description: 学院id
type CollegeIdReq struct {
	CollegeId int64
}
