// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"github.com/hubogle/chatcode-server/internal/dal/model"
)

func newRoomBasic(db *gorm.DB, opts ...gen.DOOption) roomBasic {
	_roomBasic := roomBasic{}

	_roomBasic.roomBasicDo.UseDB(db, opts...)
	_roomBasic.roomBasicDo.UseModel(&model.RoomBasic{})

	tableName := _roomBasic.roomBasicDo.TableName()
	_roomBasic.ALL = field.NewAsterisk(tableName)
	_roomBasic.ID = field.NewUint32(tableName, "id")
	_roomBasic.UID = field.NewString(tableName, "uid")
	_roomBasic.UserUID = field.NewUint32(tableName, "user_uid")
	_roomBasic.Name = field.NewString(tableName, "name")
	_roomBasic.Info = field.NewString(tableName, "info")
	_roomBasic.Salt = field.NewString(tableName, "salt")
	_roomBasic.CreatedAt = field.NewInt64(tableName, "created_at")
	_roomBasic.UpdatedAt = field.NewInt64(tableName, "updated_at")
	_roomBasic.DeletedAt = field.NewInt64(tableName, "deleted_at")

	_roomBasic.fillFieldMap()

	return _roomBasic
}

type roomBasic struct {
	roomBasicDo

	ALL       field.Asterisk
	ID        field.Uint32
	UID       field.String // 房间唯一标识 ID
	UserUID   field.Uint32 // 创建者ID
	Name      field.String // 房间名称
	Info      field.String // 房间简介
	Salt      field.String // 聊天室密码盐
	CreatedAt field.Int64
	UpdatedAt field.Int64
	DeletedAt field.Int64

	fieldMap map[string]field.Expr
}

func (r roomBasic) Table(newTableName string) *roomBasic {
	r.roomBasicDo.UseTable(newTableName)
	return r.updateTableName(newTableName)
}

func (r roomBasic) As(alias string) *roomBasic {
	r.roomBasicDo.DO = *(r.roomBasicDo.As(alias).(*gen.DO))
	return r.updateTableName(alias)
}

func (r *roomBasic) updateTableName(table string) *roomBasic {
	r.ALL = field.NewAsterisk(table)
	r.ID = field.NewUint32(table, "id")
	r.UID = field.NewString(table, "uid")
	r.UserUID = field.NewUint32(table, "user_uid")
	r.Name = field.NewString(table, "name")
	r.Info = field.NewString(table, "info")
	r.Salt = field.NewString(table, "salt")
	r.CreatedAt = field.NewInt64(table, "created_at")
	r.UpdatedAt = field.NewInt64(table, "updated_at")
	r.DeletedAt = field.NewInt64(table, "deleted_at")

	r.fillFieldMap()

	return r
}

func (r *roomBasic) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := r.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (r *roomBasic) fillFieldMap() {
	r.fieldMap = make(map[string]field.Expr, 9)
	r.fieldMap["id"] = r.ID
	r.fieldMap["uid"] = r.UID
	r.fieldMap["user_uid"] = r.UserUID
	r.fieldMap["name"] = r.Name
	r.fieldMap["info"] = r.Info
	r.fieldMap["salt"] = r.Salt
	r.fieldMap["created_at"] = r.CreatedAt
	r.fieldMap["updated_at"] = r.UpdatedAt
	r.fieldMap["deleted_at"] = r.DeletedAt
}

func (r roomBasic) clone(db *gorm.DB) roomBasic {
	r.roomBasicDo.ReplaceConnPool(db.Statement.ConnPool)
	return r
}

func (r roomBasic) replaceDB(db *gorm.DB) roomBasic {
	r.roomBasicDo.ReplaceDB(db)
	return r
}

type roomBasicDo struct{ gen.DO }

type IRoomBasicDo interface {
	gen.SubQuery
	Debug() IRoomBasicDo
	WithContext(ctx context.Context) IRoomBasicDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IRoomBasicDo
	WriteDB() IRoomBasicDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IRoomBasicDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IRoomBasicDo
	Not(conds ...gen.Condition) IRoomBasicDo
	Or(conds ...gen.Condition) IRoomBasicDo
	Select(conds ...field.Expr) IRoomBasicDo
	Where(conds ...gen.Condition) IRoomBasicDo
	Order(conds ...field.Expr) IRoomBasicDo
	Distinct(cols ...field.Expr) IRoomBasicDo
	Omit(cols ...field.Expr) IRoomBasicDo
	Join(table schema.Tabler, on ...field.Expr) IRoomBasicDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IRoomBasicDo
	RightJoin(table schema.Tabler, on ...field.Expr) IRoomBasicDo
	Group(cols ...field.Expr) IRoomBasicDo
	Having(conds ...gen.Condition) IRoomBasicDo
	Limit(limit int) IRoomBasicDo
	Offset(offset int) IRoomBasicDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IRoomBasicDo
	Unscoped() IRoomBasicDo
	Create(values ...*model.RoomBasic) error
	CreateInBatches(values []*model.RoomBasic, batchSize int) error
	Save(values ...*model.RoomBasic) error
	First() (*model.RoomBasic, error)
	Take() (*model.RoomBasic, error)
	Last() (*model.RoomBasic, error)
	Find() ([]*model.RoomBasic, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.RoomBasic, err error)
	FindInBatches(result *[]*model.RoomBasic, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.RoomBasic) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IRoomBasicDo
	Assign(attrs ...field.AssignExpr) IRoomBasicDo
	Joins(fields ...field.RelationField) IRoomBasicDo
	Preload(fields ...field.RelationField) IRoomBasicDo
	FirstOrInit() (*model.RoomBasic, error)
	FirstOrCreate() (*model.RoomBasic, error)
	FindByPage(offset int, limit int) (result []*model.RoomBasic, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IRoomBasicDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (r roomBasicDo) Debug() IRoomBasicDo {
	return r.withDO(r.DO.Debug())
}

func (r roomBasicDo) WithContext(ctx context.Context) IRoomBasicDo {
	return r.withDO(r.DO.WithContext(ctx))
}

func (r roomBasicDo) ReadDB() IRoomBasicDo {
	return r.Clauses(dbresolver.Read)
}

func (r roomBasicDo) WriteDB() IRoomBasicDo {
	return r.Clauses(dbresolver.Write)
}

func (r roomBasicDo) Session(config *gorm.Session) IRoomBasicDo {
	return r.withDO(r.DO.Session(config))
}

func (r roomBasicDo) Clauses(conds ...clause.Expression) IRoomBasicDo {
	return r.withDO(r.DO.Clauses(conds...))
}

func (r roomBasicDo) Returning(value interface{}, columns ...string) IRoomBasicDo {
	return r.withDO(r.DO.Returning(value, columns...))
}

func (r roomBasicDo) Not(conds ...gen.Condition) IRoomBasicDo {
	return r.withDO(r.DO.Not(conds...))
}

func (r roomBasicDo) Or(conds ...gen.Condition) IRoomBasicDo {
	return r.withDO(r.DO.Or(conds...))
}

func (r roomBasicDo) Select(conds ...field.Expr) IRoomBasicDo {
	return r.withDO(r.DO.Select(conds...))
}

func (r roomBasicDo) Where(conds ...gen.Condition) IRoomBasicDo {
	return r.withDO(r.DO.Where(conds...))
}

func (r roomBasicDo) Order(conds ...field.Expr) IRoomBasicDo {
	return r.withDO(r.DO.Order(conds...))
}

func (r roomBasicDo) Distinct(cols ...field.Expr) IRoomBasicDo {
	return r.withDO(r.DO.Distinct(cols...))
}

func (r roomBasicDo) Omit(cols ...field.Expr) IRoomBasicDo {
	return r.withDO(r.DO.Omit(cols...))
}

func (r roomBasicDo) Join(table schema.Tabler, on ...field.Expr) IRoomBasicDo {
	return r.withDO(r.DO.Join(table, on...))
}

func (r roomBasicDo) LeftJoin(table schema.Tabler, on ...field.Expr) IRoomBasicDo {
	return r.withDO(r.DO.LeftJoin(table, on...))
}

func (r roomBasicDo) RightJoin(table schema.Tabler, on ...field.Expr) IRoomBasicDo {
	return r.withDO(r.DO.RightJoin(table, on...))
}

func (r roomBasicDo) Group(cols ...field.Expr) IRoomBasicDo {
	return r.withDO(r.DO.Group(cols...))
}

func (r roomBasicDo) Having(conds ...gen.Condition) IRoomBasicDo {
	return r.withDO(r.DO.Having(conds...))
}

func (r roomBasicDo) Limit(limit int) IRoomBasicDo {
	return r.withDO(r.DO.Limit(limit))
}

func (r roomBasicDo) Offset(offset int) IRoomBasicDo {
	return r.withDO(r.DO.Offset(offset))
}

func (r roomBasicDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IRoomBasicDo {
	return r.withDO(r.DO.Scopes(funcs...))
}

func (r roomBasicDo) Unscoped() IRoomBasicDo {
	return r.withDO(r.DO.Unscoped())
}

func (r roomBasicDo) Create(values ...*model.RoomBasic) error {
	if len(values) == 0 {
		return nil
	}
	return r.DO.Create(values)
}

func (r roomBasicDo) CreateInBatches(values []*model.RoomBasic, batchSize int) error {
	return r.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (r roomBasicDo) Save(values ...*model.RoomBasic) error {
	if len(values) == 0 {
		return nil
	}
	return r.DO.Save(values)
}

func (r roomBasicDo) First() (*model.RoomBasic, error) {
	if result, err := r.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.RoomBasic), nil
	}
}

func (r roomBasicDo) Take() (*model.RoomBasic, error) {
	if result, err := r.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.RoomBasic), nil
	}
}

func (r roomBasicDo) Last() (*model.RoomBasic, error) {
	if result, err := r.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.RoomBasic), nil
	}
}

func (r roomBasicDo) Find() ([]*model.RoomBasic, error) {
	result, err := r.DO.Find()
	return result.([]*model.RoomBasic), err
}

func (r roomBasicDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.RoomBasic, err error) {
	buf := make([]*model.RoomBasic, 0, batchSize)
	err = r.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (r roomBasicDo) FindInBatches(result *[]*model.RoomBasic, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return r.DO.FindInBatches(result, batchSize, fc)
}

func (r roomBasicDo) Attrs(attrs ...field.AssignExpr) IRoomBasicDo {
	return r.withDO(r.DO.Attrs(attrs...))
}

func (r roomBasicDo) Assign(attrs ...field.AssignExpr) IRoomBasicDo {
	return r.withDO(r.DO.Assign(attrs...))
}

func (r roomBasicDo) Joins(fields ...field.RelationField) IRoomBasicDo {
	for _, _f := range fields {
		r = *r.withDO(r.DO.Joins(_f))
	}
	return &r
}

func (r roomBasicDo) Preload(fields ...field.RelationField) IRoomBasicDo {
	for _, _f := range fields {
		r = *r.withDO(r.DO.Preload(_f))
	}
	return &r
}

func (r roomBasicDo) FirstOrInit() (*model.RoomBasic, error) {
	if result, err := r.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.RoomBasic), nil
	}
}

func (r roomBasicDo) FirstOrCreate() (*model.RoomBasic, error) {
	if result, err := r.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.RoomBasic), nil
	}
}

func (r roomBasicDo) FindByPage(offset int, limit int) (result []*model.RoomBasic, count int64, err error) {
	result, err = r.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = r.Offset(-1).Limit(-1).Count()
	return
}

func (r roomBasicDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = r.Count()
	if err != nil {
		return
	}

	err = r.Offset(offset).Limit(limit).Scan(result)
	return
}

func (r roomBasicDo) Scan(result interface{}) (err error) {
	return r.DO.Scan(result)
}

func (r roomBasicDo) Delete(models ...*model.RoomBasic) (result gen.ResultInfo, err error) {
	return r.DO.Delete(models)
}

func (r *roomBasicDo) withDO(do gen.Dao) *roomBasicDo {
	r.DO = *do.(*gen.DO)
	return r
}
