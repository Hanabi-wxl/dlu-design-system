package service

import (
	"github.com/Hanabi-wxl/dlu-design-system/dal/db"
	"github.com/Hanabi-wxl/dlu-design-system/pkg/consts"
	"github.com/Hanabi-wxl/dlu-design-system/pkg/errno"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"sync"
	"time"
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
	//teacher, err = db.Teacher.Where(db.Teacher.Number.Eq(number)).Count()
	//student, err = db.Student.Where(db.Student.Number.Eq(number)).Count()
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
		Select(db.Student.ID, db.Student.Number, db.Student.Password).First(); err != nil {
		return nil, errno.NewErrno(errno.DbErrorCode, err.Error())
	} else {
		if checkPassword(password, student.Password) {
			token := generateToken(student.ID, 1)
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
		Select(db.Teacher.ID, db.Teacher.Number, db.Teacher.Password, db.Teacher.RoleID).First(); err != nil {
		return nil, errno.NewErrno(errno.DbErrorCode, err.Error())
	} else {
		if checkPassword(password, teacher.Password) {
			token := generateToken(teacher.ID, teacher.RoleID)
			res["token"] = token
			return res, nil
		} else {
			return nil, errno.NewErrno(errno.PasswordIncorrectErrCode, errno.PasswordIncorrectErrMsg)
		}
	}
}

type Claims struct {
	UserId int64
	RoleId int64
	jwt.StandardClaims
}

// generateToken 根据username生成一个token
func generateToken(id, roleId int64) string {
	token := newToken(id, roleId)
	return token
}

// newToken 根据信息创建token
func newToken(id, roleId int64) string {
	expiresTime := time.Now().Unix() + consts.DefaultExpiresTime
	userClaims := Claims{
		UserId: id,
		RoleId: roleId,
		StandardClaims: jwt.StandardClaims{
			Audience:  "dlu-ie-college",
			ExpiresAt: expiresTime,
			IssuedAt:  time.Now().Unix(),
			Issuer:    "dlu-design-system",
			NotBefore: time.Now().Unix(),
			Subject:   "token",
		},
	}
	var jwtSecret = []byte(consts.JwtSecret)
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, userClaims)
	if token, err := tokenClaims.SignedString(jwtSecret); err == nil {
		token = "Bearer " + token
		return token
	} else {
		println("generate token fail\n")
		return "fail"
	}
}

// CheckPassword 检验密码
func checkPassword(passwordInput, passwordSave string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(passwordSave), []byte(passwordInput))
	return err == nil
}
