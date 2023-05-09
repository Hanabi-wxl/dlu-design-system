package service

var infoService InfoService

type InfoService struct {
	SchoolService
	CollegeService
}

func GetInfoService() func() InfoService {
	return func() InfoService {
		schoolService := new(SchoolServiceImpl)
		collegeService := new(CollegeServiceImpl)
		infoService.SchoolService = schoolService
		infoService.CollegeService = collegeService
		return infoService
	}
}
