package Service

import "sync"

type SubjectService struct {
	ApplyService
	ApproveService
}

var subjectService *SubjectService
var subjectServiceOnce sync.Once

func GetSubjectService() *SubjectService {
	subjectServiceOnce.Do(func() {
		subjectService = &SubjectService{}

		applyService := new(ApplyServiceImpl)
		approveService := new(ApproveServiceImpl)

		subjectService.ApplyService = applyService
		subjectService.ApproveService = approveService
	})
	return subjectService
}
