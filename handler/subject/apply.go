package handler

import (
	"github.com/Hanabi-wxl/dlu-design-system/dal/model"
	"github.com/Hanabi-wxl/dlu-design-system/middleware/jwt"
	"github.com/Hanabi-wxl/dlu-design-system/pkg/result"
	"github.com/Hanabi-wxl/dlu-design-system/pkg/types"
	Service "github.com/Hanabi-wxl/dlu-design-system/service/subject"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"net/http"
	"strconv"
)

const ApplySubjectCountErrMsg = "报题数量超过限制"

// AddSubject
//
//	@Description: 新增报题
//	@param c
func AddSubject(c *gin.Context) {
	var subjectRequest types.SubjectRequest
	if err := c.ShouldBindJSON(&subjectRequest); err != nil {
		c.JSON(http.StatusBadRequest, result.NewFailedResult(err.Error()))
		return
	}
	info := jwt.GetUserInfo(c.Request.Context())
	if info.RoleId != subjectRequest.RoleID {
		// 判断申报人id与jwt中的角色id是否一致
		c.JSON(http.StatusBadRequest, result.Failed())
		return
	}
	if info.RoleId == 1 {
		//   学生报题
		count, err := Service.GetSubjectService().CountSubjectByStudent(subjectRequest.StudentID)
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}
		if count > 1 {
			c.JSON(http.StatusBadRequest, result.NewFailedResult(ApplySubjectCountErrMsg))
			return
		}
	} else {
		//  老师报题
		count, err := Service.GetSubjectService().CountSubject(subjectRequest.FirstTeacherID)
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}
		// 老师报题数量超过12条报题失败
		if count > 13 {
			c.JSON(http.StatusBadRequest, result.NewFailedResult(ApplySubjectCountErrMsg))
			return
		}
	}
	var subject model.Subject
	copier.Copy(&subject, &subjectRequest)
	//subject := types.SubjectRequestTosubject(subjectRequest)
	majorIds := subjectRequest.MajorID
	err2 := Service.GetSubjectService().AddSubject(&subject, majorIds)
	if err2 != nil {
		c.JSON(http.StatusBadRequest, err2)
	} else {
		c.JSON(http.StatusOK, result.Ok())
	}

}

// UpdateSubject
//
//	@Description: 修改报题
//	@param c
func UpdateSubject(c *gin.Context) {
	var subjectRequest types.SubjectRequest
	userInfo := jwt.GetUserInfo(c.Request.Context())
	if err := c.ShouldBindJSON(&subjectRequest); err != nil {
		c.JSON(http.StatusBadRequest, result.NewFailedResult(err.Error()))
		return
	}
	var subject model.Subject
	copier.Copy(&subject, &subjectRequest)
	// 获取题目信息
	subject_old, err := Service.GetSubjectService().GetSubjectById(subject.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	if subject_old == nil {
		c.JSON(http.StatusBadRequest, result.Failed())
		return
	}
	// 学生修改报题
	if userInfo.RoleId == 1 {
		// 判断当前题目是否属于该学生且报题人是否是学生
		if userInfo.ID != *subject_old.StudentID || *subject_old.RoleID != 1 {
			// 当前学生没有权限无法修改
			c.JSON(http.StatusBadRequest, result.Failed())
			return
		}
	} else {
		// 教师或者管理员修改报题
		// 判断该题目是否属于该老师    TODO 目前仅第一指导老师能修改题目
		if userInfo.ID != subject_old.FirstTeacherID {
			c.JSON(http.StatusBadRequest, result.Failed())
			return
		}
	}
	// 修改报题
	majorIds := subjectRequest.MajorID
	errno := Service.GetSubjectService().UpdateSubject(&subject, majorIds)
	if errno != nil {
		c.JSON(http.StatusBadRequest, errno)
	} else {
		c.JSON(http.StatusOK, result.Ok())
	}
}

// DeleteSubject
//
//	@Description: 删除报题
//	@param c
func DeleteSubject(c *gin.Context) {
	var subjectId types.SubjectIdRequest
	if err := c.ShouldBindUri(&subjectId); err != nil {
		c.JSON(http.StatusBadRequest, result.NewFailedResult(err.Error()))
		return
	}
	subject, err := Service.GetSubjectService().GetSubjectById(subjectId.SubjectId)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	// 判断角色身份
	userInfo := jwt.GetUserInfo(c.Request.Context())
	if userInfo.RoleId == 1 && (*subject.RoleID != 1 || *subject.StudentID != userInfo.ID) {
		// 当前为学生    只能删除学生报题且为自己的题目
		c.JSON(http.StatusBadRequest, result.Failed())
		return
	}
	if userInfo.RoleId == 2 && subject.FirstTeacherID != userInfo.ID {
		// 老师只能删除自己的报题
		c.JSON(http.StatusBadRequest, result.Failed())
		return
	}
	err2 := Service.GetSubjectService().Delete(subjectId.SubjectId)
	if err2 != nil {
		c.JSON(http.StatusBadRequest, err2)
		return
	} else {
		c.JSON(http.StatusOK, result.Ok())
	}
}

// GetOrigins
//
//	@Description: 查询题目来源
//	@param c
func GetOrigins(c *gin.Context) {
	origins, err := Service.GetSubjectService().GetOrigins()
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	} else {
		c.JSON(http.StatusOK, result.NewOkResult(origins))
	}
}

// GetTypes
//
//	@Description: 查询题目所有类型
//	@param c
func GetTypes(c *gin.Context) {
	subjecttypes, err := Service.GetSubjectService().GetTypes()
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	} else {
		c.JSON(http.StatusOK, result.NewOkResult(subjecttypes))
	}
}

// GetSelfSubjectsByYear
//
//	@Description: 获取自己题目
//	@param c
func GetSelfSubjectsByYear(c *gin.Context) {
	var year types.Year
	userInfo := jwt.GetUserInfo(c.Request.Context())
	if err := c.ShouldBindUri(&year); err != nil {
		c.JSON(http.StatusBadRequest, result.NewFailedResult(err.Error()))
		return
	}
	yearDate, _ := strconv.ParseInt(year.Year, 10, 64)
	list, err := Service.GetSubjectService().GetSelfSubjectsByYear(yearDate, userInfo)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	} else {
		c.JSON(http.StatusOK, result.NewOkResult(list))
	}
}
