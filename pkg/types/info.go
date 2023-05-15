package types

type ClassReq struct {
	ID      int64  `json:"id"`
	Name    string `json:"name"`
	MajorId int64  `json:"major_id"`
	Grade   int64  `json:"grade"`
}

type ClassQueryReq struct {
	PageRequest
	MajorId int64 `json:"major_id"`
	Year    int64 `json:"year"`
}

type CollegeReq struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type MajorReq struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	CollegeID int64  `json:"college_id"`
}
