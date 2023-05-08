package service

type InfoService struct {
	SchoolService
}

func GetInfoService() InfoService {
	var schoolService SchoolServiceImpl
	infoService := InfoService{}
	infoService.SchoolService = schoolService
	return infoService
}
