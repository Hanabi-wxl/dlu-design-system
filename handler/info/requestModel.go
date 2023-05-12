package handler

// PageRequest
// @Description: 定义绑定数据类型
type PageRequest struct {
	Size int `uri:"size"`
	Num  int `uri:"num"`
}

type IdRequest struct {
	Id string `uri:"id"`
}
