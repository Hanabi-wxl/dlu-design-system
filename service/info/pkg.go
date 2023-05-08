package service

type InfoService struct {
	SchoolService
	CollegeService
}

func GetInfoService() InfoService {
	var schoolService SchoolServiceImpl
	var collegeService CollegeServiceImpl
	infoService := InfoService{}
	infoService.SchoolService = schoolService
	infoService.CollegeService = collegeService
	return infoService
}
