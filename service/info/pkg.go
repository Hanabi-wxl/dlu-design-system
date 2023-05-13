package service

import "sync"

var infoService *InfoService
var infoServiceOnce sync.Once

type InfoService struct {
	SchoolService
	CollegeService
	MajorService
	ClassService
	TitleService
	DegreeService
	SectionService
}

func GetInfoService() *InfoService {
	infoServiceOnce.Do(func() {
		infoService = &InfoService{}
		schoolService := new(SchoolServiceImpl)
		collegeService := new(CollegeServiceImpl)
		majorService := new(MajorServiceImpl)
		classService := new(ClassServiceImpl)
		titleService := new(TitleServiceImpl)
		degreeService := new(DegreeServiceImpl)
		sectionService := new(SectionServiceImpl)

		infoService.SchoolService = schoolService
		infoService.CollegeService = collegeService
		infoService.MajorService = majorService
		infoService.ClassService = classService
		infoService.TitleService = titleService
		infoService.DegreeService = degreeService
		infoService.SectionService = sectionService
	})
	return infoService
}
