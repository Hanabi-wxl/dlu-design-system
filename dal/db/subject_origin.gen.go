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

func newSubjectOrigin(db *gorm.DB, opts ...gen.DOOption) subjectOrigin {
	_subjectOrigin := subjectOrigin{}

	_subjectOrigin.subjectOriginDo.UseDB(db, opts...)
	_subjectOrigin.subjectOriginDo.UseModel(&model.SubjectOrigin{})

	tableName := _subjectOrigin.subjectOriginDo.TableName()
	_subjectOrigin.ALL = field.NewAsterisk(tableName)
	_subjectOrigin.ID = field.NewInt64(tableName, "id")
	_subjectOrigin.Name = field.NewString(tableName, "name")
	_subjectOrigin.IsDelete = field.NewField(tableName, "is_delete")

	_subjectOrigin.fillFieldMap()

	return _subjectOrigin
}

type subjectOrigin struct {
	subjectOriginDo

	ALL      field.Asterisk
	ID       field.Int64  // id
	Name     field.String // 名称
	IsDelete field.Field  // 存在标志

	fieldMap map[string]field.Expr
}

func (s subjectOrigin) Table(newTableName string) *subjectOrigin {
	s.subjectOriginDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s subjectOrigin) As(alias string) *subjectOrigin {
	s.subjectOriginDo.DO = *(s.subjectOriginDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *subjectOrigin) updateTableName(table string) *subjectOrigin {
	s.ALL = field.NewAsterisk(table)
	s.ID = field.NewInt64(table, "id")
	s.Name = field.NewString(table, "name")
	s.IsDelete = field.NewField(table, "is_delete")

	s.fillFieldMap()

	return s
}

func (s *subjectOrigin) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *subjectOrigin) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 3)
	s.fieldMap["id"] = s.ID
	s.fieldMap["name"] = s.Name
	s.fieldMap["is_delete"] = s.IsDelete
}

func (s subjectOrigin) clone(db *gorm.DB) subjectOrigin {
	s.subjectOriginDo.ReplaceConnPool(db.Statement.ConnPool)
	return s
}

func (s subjectOrigin) replaceDB(db *gorm.DB) subjectOrigin {
	s.subjectOriginDo.ReplaceDB(db)
	return s
}

type subjectOriginDo struct{ gen.DO }

type ISubjectOriginDo interface {
	gen.SubQuery
	Debug() ISubjectOriginDo
	WithContext(ctx context.Context) ISubjectOriginDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ISubjectOriginDo
	WriteDB() ISubjectOriginDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ISubjectOriginDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ISubjectOriginDo
	Not(conds ...gen.Condition) ISubjectOriginDo
	Or(conds ...gen.Condition) ISubjectOriginDo
	Select(conds ...field.Expr) ISubjectOriginDo
	Where(conds ...gen.Condition) ISubjectOriginDo
	Order(conds ...field.Expr) ISubjectOriginDo
	Distinct(cols ...field.Expr) ISubjectOriginDo
	Omit(cols ...field.Expr) ISubjectOriginDo
	Join(table schema.Tabler, on ...field.Expr) ISubjectOriginDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ISubjectOriginDo
	RightJoin(table schema.Tabler, on ...field.Expr) ISubjectOriginDo
	Group(cols ...field.Expr) ISubjectOriginDo
	Having(conds ...gen.Condition) ISubjectOriginDo
	Limit(limit int) ISubjectOriginDo
	Offset(offset int) ISubjectOriginDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ISubjectOriginDo
	Unscoped() ISubjectOriginDo
	Create(values ...*model.SubjectOrigin) error
	CreateInBatches(values []*model.SubjectOrigin, batchSize int) error
	Save(values ...*model.SubjectOrigin) error
	First() (*model.SubjectOrigin, error)
	Take() (*model.SubjectOrigin, error)
	Last() (*model.SubjectOrigin, error)
	Find() ([]*model.SubjectOrigin, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.SubjectOrigin, err error)
	FindInBatches(result *[]*model.SubjectOrigin, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.SubjectOrigin) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ISubjectOriginDo
	Assign(attrs ...field.AssignExpr) ISubjectOriginDo
	Joins(fields ...field.RelationField) ISubjectOriginDo
	Preload(fields ...field.RelationField) ISubjectOriginDo
	FirstOrInit() (*model.SubjectOrigin, error)
	FirstOrCreate() (*model.SubjectOrigin, error)
	FindByPage(offset int, limit int) (result []*model.SubjectOrigin, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ISubjectOriginDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (s subjectOriginDo) Debug() ISubjectOriginDo {
	return s.withDO(s.DO.Debug())
}

func (s subjectOriginDo) WithContext(ctx context.Context) ISubjectOriginDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s subjectOriginDo) ReadDB() ISubjectOriginDo {
	return s.Clauses(dbresolver.Read)
}

func (s subjectOriginDo) WriteDB() ISubjectOriginDo {
	return s.Clauses(dbresolver.Write)
}

func (s subjectOriginDo) Session(config *gorm.Session) ISubjectOriginDo {
	return s.withDO(s.DO.Session(config))
}

func (s subjectOriginDo) Clauses(conds ...clause.Expression) ISubjectOriginDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s subjectOriginDo) Returning(value interface{}, columns ...string) ISubjectOriginDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s subjectOriginDo) Not(conds ...gen.Condition) ISubjectOriginDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s subjectOriginDo) Or(conds ...gen.Condition) ISubjectOriginDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s subjectOriginDo) Select(conds ...field.Expr) ISubjectOriginDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s subjectOriginDo) Where(conds ...gen.Condition) ISubjectOriginDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s subjectOriginDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) ISubjectOriginDo {
	return s.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (s subjectOriginDo) Order(conds ...field.Expr) ISubjectOriginDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s subjectOriginDo) Distinct(cols ...field.Expr) ISubjectOriginDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s subjectOriginDo) Omit(cols ...field.Expr) ISubjectOriginDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s subjectOriginDo) Join(table schema.Tabler, on ...field.Expr) ISubjectOriginDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s subjectOriginDo) LeftJoin(table schema.Tabler, on ...field.Expr) ISubjectOriginDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s subjectOriginDo) RightJoin(table schema.Tabler, on ...field.Expr) ISubjectOriginDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s subjectOriginDo) Group(cols ...field.Expr) ISubjectOriginDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s subjectOriginDo) Having(conds ...gen.Condition) ISubjectOriginDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s subjectOriginDo) Limit(limit int) ISubjectOriginDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s subjectOriginDo) Offset(offset int) ISubjectOriginDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s subjectOriginDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ISubjectOriginDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s subjectOriginDo) Unscoped() ISubjectOriginDo {
	return s.withDO(s.DO.Unscoped())
}

func (s subjectOriginDo) Create(values ...*model.SubjectOrigin) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s subjectOriginDo) CreateInBatches(values []*model.SubjectOrigin, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s subjectOriginDo) Save(values ...*model.SubjectOrigin) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s subjectOriginDo) First() (*model.SubjectOrigin, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.SubjectOrigin), nil
	}
}

func (s subjectOriginDo) Take() (*model.SubjectOrigin, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.SubjectOrigin), nil
	}
}

func (s subjectOriginDo) Last() (*model.SubjectOrigin, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.SubjectOrigin), nil
	}
}

func (s subjectOriginDo) Find() ([]*model.SubjectOrigin, error) {
	result, err := s.DO.Find()
	return result.([]*model.SubjectOrigin), err
}

func (s subjectOriginDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.SubjectOrigin, err error) {
	buf := make([]*model.SubjectOrigin, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s subjectOriginDo) FindInBatches(result *[]*model.SubjectOrigin, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s subjectOriginDo) Attrs(attrs ...field.AssignExpr) ISubjectOriginDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s subjectOriginDo) Assign(attrs ...field.AssignExpr) ISubjectOriginDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s subjectOriginDo) Joins(fields ...field.RelationField) ISubjectOriginDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s subjectOriginDo) Preload(fields ...field.RelationField) ISubjectOriginDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s subjectOriginDo) FirstOrInit() (*model.SubjectOrigin, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.SubjectOrigin), nil
	}
}

func (s subjectOriginDo) FirstOrCreate() (*model.SubjectOrigin, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.SubjectOrigin), nil
	}
}

func (s subjectOriginDo) FindByPage(offset int, limit int) (result []*model.SubjectOrigin, count int64, err error) {
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

func (s subjectOriginDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s subjectOriginDo) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s subjectOriginDo) Delete(models ...*model.SubjectOrigin) (result gen.ResultInfo, err error) {
	return s.DO.Delete(models)
}

func (s *subjectOriginDo) withDO(do gen.Dao) *subjectOriginDo {
	s.DO = *do.(*gen.DO)
	return s
}
