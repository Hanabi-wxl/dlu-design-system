package service

import (
	"github.com/Hanabi-wxl/dlu-design-system/dal/db"
	"github.com/Hanabi-wxl/dlu-design-system/pkg/errno"
	"github.com/Hanabi-wxl/dlu-design-system/pkg/utils"
	"golang.org/x/crypto/bcrypt"
	"sync"
)

type LoginServiceImpl struct {
}

func (l LoginServiceImpl) CheckLoginRole(number string) (map[string]bool, *errno.Errno) {
	var teacher int64
	var student int64
	res := make(map[string]bool, 1)
	var err error
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		teacher, err = db.Teacher.Where(db.Teacher.Number.Eq(number)).Count()
		wg.Done()
	}()
	go func() {
		student, err = db.Student.Where(db.Student.Number.Eq(number)).Count()
		wg.Done()
	}()
	wg.Wait()
	if err != nil {
		return nil, errno.NewErrno(errno.DbErrorCode, err.Error())
	}
	if teacher == 0 && student == 0 {
		return nil, errno.NewErrno(errno.NumberNotFoundErrCode, errno.NumberNotFoundErrMsg)
	} else if teacher == 0 {
		res["isStudent"] = true
		return res, nil
	} else if student == 0 {
		res["isStudent"] = false
		return res, nil
	} else {
		return nil, errno.NewErrno(errno.NumberDuplicateErrCode, errno.NumberDuplicateErrMsg)
	}
}

func (l LoginServiceImpl) StudentLogin(number, password string) (map[string]interface{}, *errno.Errno) {
	res := make(map[string]interface{}, 3)
	if student, err := db.Student.Where(db.Student.Number.Eq(number)).
		Select(db.Student.ID, db.Student.Number, db.Student.Password).Take(); err != nil {
		return nil, errno.NewErrno(errno.DbErrorCode, errno.NumberNotFoundErrMsg)
	} else {
		if checkPassword(password, student.Password) {
			token := utils.GenerateToken(student.ID, 1, number)
			res["token"] = token
			return res, nil
		} else {
			return nil, errno.NewErrno(errno.PasswordIncorrectErrCode, errno.PasswordIncorrectErrMsg)
		}
	}
}

func (l LoginServiceImpl) TeacherLogin(number, password string) (map[string]interface{}, *errno.Errno) {
	res := make(map[string]interface{}, 1)
	if teacher, err := db.Teacher.Where(db.Teacher.Number.Eq(number)).
		Select(db.Teacher.ID, db.Teacher.Number, db.Teacher.Password, db.Teacher.RoleID).Take(); err != nil {
		return nil, errno.NewErrno(errno.DbErrorCode, errno.NumberNotFoundErrMsg)
	} else {
		if checkPassword(password, teacher.Password) {
			token := utils.GenerateToken(teacher.ID, teacher.RoleID, number)
			res["token"] = token
			return res, nil
		} else {
			return nil, errno.NewErrno(errno.PasswordIncorrectErrCode, errno.PasswordIncorrectErrMsg)
		}
	}
}

// CheckPassword 检验密码
func checkPassword(passwordInput, passwordSave string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(passwordSave), []byte(passwordInput))
	return err == nil
}
