package types

import "time"

type LogDateReq struct {
	Start time.Time `uri:"startId"` // 开始时间 毫秒值
	End   time.Time `uri:"end"`     // 结束时间 毫秒值
	Size  int       `uri:"size"`
	Num   int       `uri:"num"`
}

type LogStateReq struct {
	StateId int8 `uri:"stateId"` // 状态 1正常 2错误
	Size    int  `uri:"size"`
	Num     int  `uri:"num"`
}

type LogMethodReq struct {
	MethodId int8 `uri:"methodId"` // 1增 2删 3改 4查
	Size     int  `uri:"size"`
	Num      int  `uri:"num"`
}
