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

func newMajor(db *gorm.DB, opts ...gen.DOOption) major {
	_major := major{}

	_major.majorDo.UseDB(db, opts...)
	_major.majorDo.UseModel(&model.Major{})

	tableName := _major.majorDo.TableName()
	_major.ALL = field.NewAsterisk(tableName)
	_major.ID = field.NewInt64(tableName, "id")
	_major.Name = field.NewString(tableName, "name")
	_major.CollegeID = field.NewInt64(tableName, "college_id")
	_major.IsDelete = field.NewField(tableName, "is_delete")

	_major.fillFieldMap()

	return _major
}

type major struct {
	majorDo

	ALL       field.Asterisk
	ID        field.Int64  // id
	Name      field.String // 专业名称
	CollegeID field.Int64  // 学院
	IsDelete  field.Field  // 是否有效

	fieldMap map[string]field.Expr
}

func (m major) Table(newTableName string) *major {
	m.majorDo.UseTable(newTableName)
	return m.updateTableName(newTableName)
}

func (m major) As(alias string) *major {
	m.majorDo.DO = *(m.majorDo.As(alias).(*gen.DO))
	return m.updateTableName(alias)
}

func (m *major) updateTableName(table string) *major {
	m.ALL = field.NewAsterisk(table)
	m.ID = field.NewInt64(table, "id")
	m.Name = field.NewString(table, "name")
	m.CollegeID = field.NewInt64(table, "college_id")
	m.IsDelete = field.NewField(table, "is_delete")

	m.fillFieldMap()

	return m
}

func (m *major) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := m.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (m *major) fillFieldMap() {
	m.fieldMap = make(map[string]field.Expr, 4)
	m.fieldMap["id"] = m.ID
	m.fieldMap["name"] = m.Name
	m.fieldMap["college_id"] = m.CollegeID
	m.fieldMap["is_delete"] = m.IsDelete
}

func (m major) clone(db *gorm.DB) major {
	m.majorDo.ReplaceConnPool(db.Statement.ConnPool)
	return m
}

func (m major) replaceDB(db *gorm.DB) major {
	m.majorDo.ReplaceDB(db)
	return m
}

type majorDo struct{ gen.DO }

type IMajorDo interface {
	gen.SubQuery
	Debug() IMajorDo
	WithContext(ctx context.Context) IMajorDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IMajorDo
	WriteDB() IMajorDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IMajorDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IMajorDo
	Not(conds ...gen.Condition) IMajorDo
	Or(conds ...gen.Condition) IMajorDo
	Select(conds ...field.Expr) IMajorDo
	Where(conds ...gen.Condition) IMajorDo
	Order(conds ...field.Expr) IMajorDo
	Distinct(cols ...field.Expr) IMajorDo
	Omit(cols ...field.Expr) IMajorDo
	Join(table schema.Tabler, on ...field.Expr) IMajorDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IMajorDo
	RightJoin(table schema.Tabler, on ...field.Expr) IMajorDo
	Group(cols ...field.Expr) IMajorDo
	Having(conds ...gen.Condition) IMajorDo
	Limit(limit int) IMajorDo
	Offset(offset int) IMajorDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IMajorDo
	Unscoped() IMajorDo
	Create(values ...*model.Major) error
	CreateInBatches(values []*model.Major, batchSize int) error
	Save(values ...*model.Major) error
	First() (*model.Major, error)
	Take() (*model.Major, error)
	Last() (*model.Major, error)
	Find() ([]*model.Major, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Major, err error)
	FindInBatches(result *[]*model.Major, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.Major) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IMajorDo
	Assign(attrs ...field.AssignExpr) IMajorDo
	Joins(fields ...field.RelationField) IMajorDo
	Preload(fields ...field.RelationField) IMajorDo
	FirstOrInit() (*model.Major, error)
	FirstOrCreate() (*model.Major, error)
	FindByPage(offset int, limit int) (result []*model.Major, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IMajorDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (m majorDo) Debug() IMajorDo {
	return m.withDO(m.DO.Debug())
}

func (m majorDo) WithContext(ctx context.Context) IMajorDo {
	return m.withDO(m.DO.WithContext(ctx))
}

func (m majorDo) ReadDB() IMajorDo {
	return m.Clauses(dbresolver.Read)
}

func (m majorDo) WriteDB() IMajorDo {
	return m.Clauses(dbresolver.Write)
}

func (m majorDo) Session(config *gorm.Session) IMajorDo {
	return m.withDO(m.DO.Session(config))
}

func (m majorDo) Clauses(conds ...clause.Expression) IMajorDo {
	return m.withDO(m.DO.Clauses(conds...))
}

func (m majorDo) Returning(value interface{}, columns ...string) IMajorDo {
	return m.withDO(m.DO.Returning(value, columns...))
}

func (m majorDo) Not(conds ...gen.Condition) IMajorDo {
	return m.withDO(m.DO.Not(conds...))
}

func (m majorDo) Or(conds ...gen.Condition) IMajorDo {
	return m.withDO(m.DO.Or(conds...))
}

func (m majorDo) Select(conds ...field.Expr) IMajorDo {
	return m.withDO(m.DO.Select(conds...))
}

func (m majorDo) Where(conds ...gen.Condition) IMajorDo {
	return m.withDO(m.DO.Where(conds...))
}

func (m majorDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) IMajorDo {
	return m.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (m majorDo) Order(conds ...field.Expr) IMajorDo {
	return m.withDO(m.DO.Order(conds...))
}

func (m majorDo) Distinct(cols ...field.Expr) IMajorDo {
	return m.withDO(m.DO.Distinct(cols...))
}

func (m majorDo) Omit(cols ...field.Expr) IMajorDo {
	return m.withDO(m.DO.Omit(cols...))
}

func (m majorDo) Join(table schema.Tabler, on ...field.Expr) IMajorDo {
	return m.withDO(m.DO.Join(table, on...))
}

func (m majorDo) LeftJoin(table schema.Tabler, on ...field.Expr) IMajorDo {
	return m.withDO(m.DO.LeftJoin(table, on...))
}

func (m majorDo) RightJoin(table schema.Tabler, on ...field.Expr) IMajorDo {
	return m.withDO(m.DO.RightJoin(table, on...))
}

func (m majorDo) Group(cols ...field.Expr) IMajorDo {
	return m.withDO(m.DO.Group(cols...))
}

func (m majorDo) Having(conds ...gen.Condition) IMajorDo {
	return m.withDO(m.DO.Having(conds...))
}

func (m majorDo) Limit(limit int) IMajorDo {
	return m.withDO(m.DO.Limit(limit))
}

func (m majorDo) Offset(offset int) IMajorDo {
	return m.withDO(m.DO.Offset(offset))
}

func (m majorDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IMajorDo {
	return m.withDO(m.DO.Scopes(funcs...))
}

func (m majorDo) Unscoped() IMajorDo {
	return m.withDO(m.DO.Unscoped())
}

func (m majorDo) Create(values ...*model.Major) error {
	if len(values) == 0 {
		return nil
	}
	return m.DO.Create(values)
}

func (m majorDo) CreateInBatches(values []*model.Major, batchSize int) error {
	return m.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (m majorDo) Save(values ...*model.Major) error {
	if len(values) == 0 {
		return nil
	}
	return m.DO.Save(values)
}

func (m majorDo) First() (*model.Major, error) {
	if result, err := m.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Major), nil
	}
}

func (m majorDo) Take() (*model.Major, error) {
	if result, err := m.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Major), nil
	}
}

func (m majorDo) Last() (*model.Major, error) {
	if result, err := m.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Major), nil
	}
}

func (m majorDo) Find() ([]*model.Major, error) {
	result, err := m.DO.Find()
	return result.([]*model.Major), err
}

func (m majorDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Major, err error) {
	buf := make([]*model.Major, 0, batchSize)
	err = m.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (m majorDo) FindInBatches(result *[]*model.Major, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return m.DO.FindInBatches(result, batchSize, fc)
}

func (m majorDo) Attrs(attrs ...field.AssignExpr) IMajorDo {
	return m.withDO(m.DO.Attrs(attrs...))
}

func (m majorDo) Assign(attrs ...field.AssignExpr) IMajorDo {
	return m.withDO(m.DO.Assign(attrs...))
}

func (m majorDo) Joins(fields ...field.RelationField) IMajorDo {
	for _, _f := range fields {
		m = *m.withDO(m.DO.Joins(_f))
	}
	return &m
}

func (m majorDo) Preload(fields ...field.RelationField) IMajorDo {
	for _, _f := range fields {
		m = *m.withDO(m.DO.Preload(_f))
	}
	return &m
}

func (m majorDo) FirstOrInit() (*model.Major, error) {
	if result, err := m.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Major), nil
	}
}

func (m majorDo) FirstOrCreate() (*model.Major, error) {
	if result, err := m.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Major), nil
	}
}

func (m majorDo) FindByPage(offset int, limit int) (result []*model.Major, count int64, err error) {
	result, err = m.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = m.Offset(-1).Limit(-1).Count()
	return
}

func (m majorDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = m.Count()
	if err != nil {
		return
	}

	err = m.Offset(offset).Limit(limit).Scan(result)
	return
}

func (m majorDo) Scan(result interface{}) (err error) {
	return m.DO.Scan(result)
}

func (m majorDo) Delete(models ...*model.Major) (result gen.ResultInfo, err error) {
	return m.DO.Delete(models)
}

func (m *majorDo) withDO(do gen.Dao) *majorDo {
	m.DO = *do.(*gen.DO)
	return m
}
