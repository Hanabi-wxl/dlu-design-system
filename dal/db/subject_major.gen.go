// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package db

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"github.com/Hanabi-wxl/dlu-design-system/dal/model"
)

func newSubjectMajor(db *gorm.DB, opts ...gen.DOOption) subjectMajor {
	_subjectMajor := subjectMajor{}

	_subjectMajor.subjectMajorDo.UseDB(db, opts...)
	_subjectMajor.subjectMajorDo.UseModel(&model.SubjectMajor{})

	tableName := _subjectMajor.subjectMajorDo.TableName()
	_subjectMajor.ALL = field.NewAsterisk(tableName)
	_subjectMajor.ID = field.NewInt64(tableName, "id")
	_subjectMajor.SubjectID = field.NewInt64(tableName, "subject_id")
	_subjectMajor.MajorID = field.NewInt64(tableName, "major_id")
	_subjectMajor.IsDelete = field.NewField(tableName, "is_delete")

	_subjectMajor.fillFieldMap()

	return _subjectMajor
}

type subjectMajor struct {
	subjectMajorDo

	ALL       field.Asterisk
	ID        field.Int64 // id
	SubjectID field.Int64 // 题目id
	MajorID   field.Int64 // 专业id
	IsDelete  field.Field // 存在标志

	fieldMap map[string]field.Expr
}

func (s subjectMajor) Table(newTableName string) *subjectMajor {
	s.subjectMajorDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s subjectMajor) As(alias string) *subjectMajor {
	s.subjectMajorDo.DO = *(s.subjectMajorDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *subjectMajor) updateTableName(table string) *subjectMajor {
	s.ALL = field.NewAsterisk(table)
	s.ID = field.NewInt64(table, "id")
	s.SubjectID = field.NewInt64(table, "subject_id")
	s.MajorID = field.NewInt64(table, "major_id")
	s.IsDelete = field.NewField(table, "is_delete")

	s.fillFieldMap()

	return s
}

func (s *subjectMajor) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *subjectMajor) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 4)
	s.fieldMap["id"] = s.ID
	s.fieldMap["subject_id"] = s.SubjectID
	s.fieldMap["major_id"] = s.MajorID
	s.fieldMap["is_delete"] = s.IsDelete
}

func (s subjectMajor) clone(db *gorm.DB) subjectMajor {
	s.subjectMajorDo.ReplaceConnPool(db.Statement.ConnPool)
	return s
}

func (s subjectMajor) replaceDB(db *gorm.DB) subjectMajor {
	s.subjectMajorDo.ReplaceDB(db)
	return s
}

type subjectMajorDo struct{ gen.DO }

type ISubjectMajorDo interface {
	gen.SubQuery
	Debug() ISubjectMajorDo
	WithContext(ctx context.Context) ISubjectMajorDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ISubjectMajorDo
	WriteDB() ISubjectMajorDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ISubjectMajorDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ISubjectMajorDo
	Not(conds ...gen.Condition) ISubjectMajorDo
	Or(conds ...gen.Condition) ISubjectMajorDo
	Select(conds ...field.Expr) ISubjectMajorDo
	Where(conds ...gen.Condition) ISubjectMajorDo
	Order(conds ...field.Expr) ISubjectMajorDo
	Distinct(cols ...field.Expr) ISubjectMajorDo
	Omit(cols ...field.Expr) ISubjectMajorDo
	Join(table schema.Tabler, on ...field.Expr) ISubjectMajorDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ISubjectMajorDo
	RightJoin(table schema.Tabler, on ...field.Expr) ISubjectMajorDo
	Group(cols ...field.Expr) ISubjectMajorDo
	Having(conds ...gen.Condition) ISubjectMajorDo
	Limit(limit int) ISubjectMajorDo
	Offset(offset int) ISubjectMajorDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ISubjectMajorDo
	Unscoped() ISubjectMajorDo
	Create(values ...*model.SubjectMajor) error
	CreateInBatches(values []*model.SubjectMajor, batchSize int) error
	Save(values ...*model.SubjectMajor) error
	First() (*model.SubjectMajor, error)
	Take() (*model.SubjectMajor, error)
	Last() (*model.SubjectMajor, error)
	Find() ([]*model.SubjectMajor, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.SubjectMajor, err error)
	FindInBatches(result *[]*model.SubjectMajor, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.SubjectMajor) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ISubjectMajorDo
	Assign(attrs ...field.AssignExpr) ISubjectMajorDo
	Joins(fields ...field.RelationField) ISubjectMajorDo
	Preload(fields ...field.RelationField) ISubjectMajorDo
	FirstOrInit() (*model.SubjectMajor, error)
	FirstOrCreate() (*model.SubjectMajor, error)
	FindByPage(offset int, limit int) (result []*model.SubjectMajor, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ISubjectMajorDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (s subjectMajorDo) Debug() ISubjectMajorDo {
	return s.withDO(s.DO.Debug())
}

func (s subjectMajorDo) WithContext(ctx context.Context) ISubjectMajorDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s subjectMajorDo) ReadDB() ISubjectMajorDo {
	return s.Clauses(dbresolver.Read)
}

func (s subjectMajorDo) WriteDB() ISubjectMajorDo {
	return s.Clauses(dbresolver.Write)
}

func (s subjectMajorDo) Session(config *gorm.Session) ISubjectMajorDo {
	return s.withDO(s.DO.Session(config))
}

func (s subjectMajorDo) Clauses(conds ...clause.Expression) ISubjectMajorDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s subjectMajorDo) Returning(value interface{}, columns ...string) ISubjectMajorDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s subjectMajorDo) Not(conds ...gen.Condition) ISubjectMajorDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s subjectMajorDo) Or(conds ...gen.Condition) ISubjectMajorDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s subjectMajorDo) Select(conds ...field.Expr) ISubjectMajorDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s subjectMajorDo) Where(conds ...gen.Condition) ISubjectMajorDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s subjectMajorDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) ISubjectMajorDo {
	return s.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (s subjectMajorDo) Order(conds ...field.Expr) ISubjectMajorDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s subjectMajorDo) Distinct(cols ...field.Expr) ISubjectMajorDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s subjectMajorDo) Omit(cols ...field.Expr) ISubjectMajorDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s subjectMajorDo) Join(table schema.Tabler, on ...field.Expr) ISubjectMajorDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s subjectMajorDo) LeftJoin(table schema.Tabler, on ...field.Expr) ISubjectMajorDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s subjectMajorDo) RightJoin(table schema.Tabler, on ...field.Expr) ISubjectMajorDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s subjectMajorDo) Group(cols ...field.Expr) ISubjectMajorDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s subjectMajorDo) Having(conds ...gen.Condition) ISubjectMajorDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s subjectMajorDo) Limit(limit int) ISubjectMajorDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s subjectMajorDo) Offset(offset int) ISubjectMajorDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s subjectMajorDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ISubjectMajorDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s subjectMajorDo) Unscoped() ISubjectMajorDo {
	return s.withDO(s.DO.Unscoped())
}

func (s subjectMajorDo) Create(values ...*model.SubjectMajor) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s subjectMajorDo) CreateInBatches(values []*model.SubjectMajor, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s subjectMajorDo) Save(values ...*model.SubjectMajor) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s subjectMajorDo) First() (*model.SubjectMajor, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.SubjectMajor), nil
	}
}

func (s subjectMajorDo) Take() (*model.SubjectMajor, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.SubjectMajor), nil
	}
}

func (s subjectMajorDo) Last() (*model.SubjectMajor, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.SubjectMajor), nil
	}
}

func (s subjectMajorDo) Find() ([]*model.SubjectMajor, error) {
	result, err := s.DO.Find()
	return result.([]*model.SubjectMajor), err
}

func (s subjectMajorDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.SubjectMajor, err error) {
	buf := make([]*model.SubjectMajor, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s subjectMajorDo) FindInBatches(result *[]*model.SubjectMajor, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s subjectMajorDo) Attrs(attrs ...field.AssignExpr) ISubjectMajorDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s subjectMajorDo) Assign(attrs ...field.AssignExpr) ISubjectMajorDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s subjectMajorDo) Joins(fields ...field.RelationField) ISubjectMajorDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s subjectMajorDo) Preload(fields ...field.RelationField) ISubjectMajorDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s subjectMajorDo) FirstOrInit() (*model.SubjectMajor, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.SubjectMajor), nil
	}
}

func (s subjectMajorDo) FirstOrCreate() (*model.SubjectMajor, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.SubjectMajor), nil
	}
}

func (s subjectMajorDo) FindByPage(offset int, limit int) (result []*model.SubjectMajor, count int64, err error) {
	result, err = s.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = s.Offset(-1).Limit(-1).Count()
	return
}

func (s subjectMajorDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s subjectMajorDo) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s subjectMajorDo) Delete(models ...*model.SubjectMajor) (result gen.ResultInfo, err error) {
	return s.DO.Delete(models)
}

func (s *subjectMajorDo) withDO(do gen.Dao) *subjectMajorDo {
	s.DO = *do.(*gen.DO)
	return s
}