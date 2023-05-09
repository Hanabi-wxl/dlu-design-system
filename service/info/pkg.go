package service

var infoService InfoService

type InfoService struct {
	SchoolService
	CollegeService
	MajorService
	ClassService
	TitleService
	DegreeService
	SectionService
}

func GetInfoService() func() InfoService {
	return func() InfoService {
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
		return infoService
	}
}
