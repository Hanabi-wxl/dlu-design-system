// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package db

import (
	"context"
	"database/sql"

	"gorm.io/gorm"

	"gorm.io/gen"

	"gorm.io/plugin/dbresolver"
)

var (
	Q               = new(Query)
	Appoint         *appoint
	Class           *class
	College         *college
	Log             *log
	Major           *major
	Message         *message
	Notice          *notice
	Plan            *plan
	Role            *role
	School          *school
	Selection       *selection
	Student         *student
	Subject         *subject
	SubjectMajor    *subjectMajor
	SubjectOrigin   *subjectOrigin
	SubjectProgress *subjectProgress
	SubjectType     *subjectType
	Task            *task
	Teacher         *teacher
	TeacherDegree   *teacherDegree
	TeacherOffice   *teacherOffice
	TeacherTitle    *teacherTitle
)

func SetDefault(db *gorm.DB, opts ...gen.DOOption) {
	*Q = *Use(db, opts...)
	Appoint = &Q.Appoint
	Class = &Q.Class
	College = &Q.College
	Log = &Q.Log
	Major = &Q.Major
	Message = &Q.Message
	Notice = &Q.Notice
	Plan = &Q.Plan
	Role = &Q.Role
	School = &Q.School
	Selection = &Q.Selection
	Student = &Q.Student
	Subject = &Q.Subject
	SubjectMajor = &Q.SubjectMajor
	SubjectOrigin = &Q.SubjectOrigin
	SubjectProgress = &Q.SubjectProgress
	SubjectType = &Q.SubjectType
	Task = &Q.Task
	Teacher = &Q.Teacher
	TeacherDegree = &Q.TeacherDegree
	TeacherOffice = &Q.TeacherOffice
	TeacherTitle = &Q.TeacherTitle
}

func Use(db *gorm.DB, opts ...gen.DOOption) *Query {
	return &Query{
		db:              db,
		Appoint:         newAppoint(db, opts...),
		Class:           newClass(db, opts...),
		College:         newCollege(db, opts...),
		Log:             newLog(db, opts...),
		Major:           newMajor(db, opts...),
		Message:         newMessage(db, opts...),
		Notice:          newNotice(db, opts...),
		Plan:            newPlan(db, opts...),
		Role:            newRole(db, opts...),
		School:          newSchool(db, opts...),
		Selection:       newSelection(db, opts...),
		Student:         newStudent(db, opts...),
		Subject:         newSubject(db, opts...),
		SubjectMajor:    newSubjectMajor(db, opts...),
		SubjectOrigin:   newSubjectOrigin(db, opts...),
		SubjectProgress: newSubjectProgress(db, opts...),
		SubjectType:     newSubjectType(db, opts...),
		Task:            newTask(db, opts...),
		Teacher:         newTeacher(db, opts...),
		TeacherDegree:   newTeacherDegree(db, opts...),
		TeacherOffice:   newTeacherOffice(db, opts...),
		TeacherTitle:    newTeacherTitle(db, opts...),
	}
}

type Query struct {
	db *gorm.DB

	Appoint         appoint
	Class           class
	College         college
	Log             log
	Major           major
	Message         message
	Notice          notice
	Plan            plan
	Role            role
	School          school
	Selection       selection
	Student         student
	Subject         subject
	SubjectMajor    subjectMajor
	SubjectOrigin   subjectOrigin
	SubjectProgress subjectProgress
	SubjectType     subjectType
	Task            task
	Teacher         teacher
	TeacherDegree   teacherDegree
	TeacherOffice   teacherOffice
	TeacherTitle    teacherTitle
}

func (q *Query) Available() bool { return q.db != nil }

func (q *Query) clone(db *gorm.DB) *Query {
	return &Query{
		db:              db,
		Appoint:         q.Appoint.clone(db),
		Class:           q.Class.clone(db),
		College:         q.College.clone(db),
		Log:             q.Log.clone(db),
		Major:           q.Major.clone(db),
		Message:         q.Message.clone(db),
		Notice:          q.Notice.clone(db),
		Plan:            q.Plan.clone(db),
		Role:            q.Role.clone(db),
		School:          q.School.clone(db),
		Selection:       q.Selection.clone(db),
		Student:         q.Student.clone(db),
		Subject:         q.Subject.clone(db),
		SubjectMajor:    q.SubjectMajor.clone(db),
		SubjectOrigin:   q.SubjectOrigin.clone(db),
		SubjectProgress: q.SubjectProgress.clone(db),
		SubjectType:     q.SubjectType.clone(db),
		Task:            q.Task.clone(db),
		Teacher:         q.Teacher.clone(db),
		TeacherDegree:   q.TeacherDegree.clone(db),
		TeacherOffice:   q.TeacherOffice.clone(db),
		TeacherTitle:    q.TeacherTitle.clone(db),
	}
}

func (q *Query) ReadDB() *Query {
	return q.ReplaceDB(q.db.Clauses(dbresolver.Read))
}

func (q *Query) WriteDB() *Query {
	return q.ReplaceDB(q.db.Clauses(dbresolver.Write))
}

func (q *Query) ReplaceDB(db *gorm.DB) *Query {
	return &Query{
		db:              db,
		Appoint:         q.Appoint.replaceDB(db),
		Class:           q.Class.replaceDB(db),
		College:         q.College.replaceDB(db),
		Log:             q.Log.replaceDB(db),
		Major:           q.Major.replaceDB(db),
		Message:         q.Message.replaceDB(db),
		Notice:          q.Notice.replaceDB(db),
		Plan:            q.Plan.replaceDB(db),
		Role:            q.Role.replaceDB(db),
		School:          q.School.replaceDB(db),
		Selection:       q.Selection.replaceDB(db),
		Student:         q.Student.replaceDB(db),
		Subject:         q.Subject.replaceDB(db),
		SubjectMajor:    q.SubjectMajor.replaceDB(db),
		SubjectOrigin:   q.SubjectOrigin.replaceDB(db),
		SubjectProgress: q.SubjectProgress.replaceDB(db),
		SubjectType:     q.SubjectType.replaceDB(db),
		Task:            q.Task.replaceDB(db),
		Teacher:         q.Teacher.replaceDB(db),
		TeacherDegree:   q.TeacherDegree.replaceDB(db),
		TeacherOffice:   q.TeacherOffice.replaceDB(db),
		TeacherTitle:    q.TeacherTitle.replaceDB(db),
	}
}

type queryCtx struct {
	Appoint         IAppointDo
	Class           IClassDo
	College         ICollegeDo
	Log             ILogDo
	Major           IMajorDo
	Message         IMessageDo
	Notice          INoticeDo
	Plan            IPlanDo
	Role            IRoleDo
	School          ISchoolDo
	Selection       ISelectionDo
	Student         IStudentDo
	Subject         ISubjectDo
	SubjectMajor    ISubjectMajorDo
	SubjectOrigin   ISubjectOriginDo
	SubjectProgress ISubjectProgressDo
	SubjectType     ISubjectTypeDo
	Task            ITaskDo
	Teacher         ITeacherDo
	TeacherDegree   ITeacherDegreeDo
	TeacherOffice   ITeacherOfficeDo
	TeacherTitle    ITeacherTitleDo
}

func (q *Query) WithContext(ctx context.Context) *queryCtx {
	return &queryCtx{
		Appoint:         q.Appoint.WithContext(ctx),
		Class:           q.Class.WithContext(ctx),
		College:         q.College.WithContext(ctx),
		Log:             q.Log.WithContext(ctx),
		Major:           q.Major.WithContext(ctx),
		Message:         q.Message.WithContext(ctx),
		Notice:          q.Notice.WithContext(ctx),
		Plan:            q.Plan.WithContext(ctx),
		Role:            q.Role.WithContext(ctx),
		School:          q.School.WithContext(ctx),
		Selection:       q.Selection.WithContext(ctx),
		Student:         q.Student.WithContext(ctx),
		Subject:         q.Subject.WithContext(ctx),
		SubjectMajor:    q.SubjectMajor.WithContext(ctx),
		SubjectOrigin:   q.SubjectOrigin.WithContext(ctx),
		SubjectProgress: q.SubjectProgress.WithContext(ctx),
		SubjectType:     q.SubjectType.WithContext(ctx),
		Task:            q.Task.WithContext(ctx),
		Teacher:         q.Teacher.WithContext(ctx),
		TeacherDegree:   q.TeacherDegree.WithContext(ctx),
		TeacherOffice:   q.TeacherOffice.WithContext(ctx),
		TeacherTitle:    q.TeacherTitle.WithContext(ctx),
	}
}

func (q *Query) Transaction(fc func(tx *Query) error, opts ...*sql.TxOptions) error {
	return q.db.Transaction(func(tx *gorm.DB) error { return fc(q.clone(tx)) }, opts...)
}

func (q *Query) Begin(opts ...*sql.TxOptions) *QueryTx {
	return &QueryTx{q.clone(q.db.Begin(opts...))}
}

type QueryTx struct{ *Query }

func (q *QueryTx) Commit() error {
	return q.db.Commit().Error
}

func (q *QueryTx) Rollback() error {
	return q.db.Rollback().Error
}

func (q *QueryTx) SavePoint(name string) error {
	return q.db.SavePoint(name).Error
}

func (q *QueryTx) RollbackTo(name string) error {
	return q.db.RollbackTo(name).Error
}
