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

func newUserFriend(db *gorm.DB, opts ...gen.DOOption) userFriend {
	_userFriend := userFriend{}

	_userFriend.userFriendDo.UseDB(db, opts...)
	_userFriend.userFriendDo.UseModel(&model.UserFriend{})

	tableName := _userFriend.userFriendDo.TableName()
	_userFriend.ALL = field.NewAsterisk(tableName)
	_userFriend.ID = field.NewUint32(tableName, "id")
	_userFriend.UserID = field.NewUint64(tableName, "user_id")
	_userFriend.FriendID = field.NewUint64(tableName, "friend_id")
	_userFriend.CreatedAt = field.NewInt64(tableName, "created_at")
	_userFriend.UpdatedAt = field.NewInt64(tableName, "updated_at")
	_userFriend.DeletedAt = field.NewInt64(tableName, "deleted_at")

	_userFriend.fillFieldMap()

	return _userFriend
}

type userFriend struct {
	userFriendDo

	ALL       field.Asterisk
	ID        field.Uint32
	UserID    field.Uint64 // 用户ID
	FriendID  field.Uint64 // 好友ID
	CreatedAt field.Int64
	UpdatedAt field.Int64
	DeletedAt field.Int64

	fieldMap map[string]field.Expr
}

func (u userFriend) Table(newTableName string) *userFriend {
	u.userFriendDo.UseTable(newTableName)
	return u.updateTableName(newTableName)
}

func (u userFriend) As(alias string) *userFriend {
	u.userFriendDo.DO = *(u.userFriendDo.As(alias).(*gen.DO))
	return u.updateTableName(alias)
}

func (u *userFriend) updateTableName(table string) *userFriend {
	u.ALL = field.NewAsterisk(table)
	u.ID = field.NewUint32(table, "id")
	u.UserID = field.NewUint64(table, "user_id")
	u.FriendID = field.NewUint64(table, "friend_id")
	u.CreatedAt = field.NewInt64(table, "created_at")
	u.UpdatedAt = field.NewInt64(table, "updated_at")
	u.DeletedAt = field.NewInt64(table, "deleted_at")

	u.fillFieldMap()

	return u
}

func (u *userFriend) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := u.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (u *userFriend) fillFieldMap() {
	u.fieldMap = make(map[string]field.Expr, 6)
	u.fieldMap["id"] = u.ID
	u.fieldMap["user_id"] = u.UserID
	u.fieldMap["friend_id"] = u.FriendID
	u.fieldMap["created_at"] = u.CreatedAt
	u.fieldMap["updated_at"] = u.UpdatedAt
	u.fieldMap["deleted_at"] = u.DeletedAt
}

func (u userFriend) clone(db *gorm.DB) userFriend {
	u.userFriendDo.ReplaceConnPool(db.Statement.ConnPool)
	return u
}

func (u userFriend) replaceDB(db *gorm.DB) userFriend {
	u.userFriendDo.ReplaceDB(db)
	return u
}

type userFriendDo struct{ gen.DO }

type IUserFriendDo interface {
	gen.SubQuery
	Debug() IUserFriendDo
	WithContext(ctx context.Context) IUserFriendDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IUserFriendDo
	WriteDB() IUserFriendDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IUserFriendDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IUserFriendDo
	Not(conds ...gen.Condition) IUserFriendDo
	Or(conds ...gen.Condition) IUserFriendDo
	Select(conds ...field.Expr) IUserFriendDo
	Where(conds ...gen.Condition) IUserFriendDo
	Order(conds ...field.Expr) IUserFriendDo
	Distinct(cols ...field.Expr) IUserFriendDo
	Omit(cols ...field.Expr) IUserFriendDo
	Join(table schema.Tabler, on ...field.Expr) IUserFriendDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IUserFriendDo
	RightJoin(table schema.Tabler, on ...field.Expr) IUserFriendDo
	Group(cols ...field.Expr) IUserFriendDo
	Having(conds ...gen.Condition) IUserFriendDo
	Limit(limit int) IUserFriendDo
	Offset(offset int) IUserFriendDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IUserFriendDo
	Unscoped() IUserFriendDo
	Create(values ...*model.UserFriend) error
	CreateInBatches(values []*model.UserFriend, batchSize int) error
	Save(values ...*model.UserFriend) error
	First() (*model.UserFriend, error)
	Take() (*model.UserFriend, error)
	Last() (*model.UserFriend, error)
	Find() ([]*model.UserFriend, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.UserFriend, err error)
	FindInBatches(result *[]*model.UserFriend, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.UserFriend) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IUserFriendDo
	Assign(attrs ...field.AssignExpr) IUserFriendDo
	Joins(fields ...field.RelationField) IUserFriendDo
	Preload(fields ...field.RelationField) IUserFriendDo
	FirstOrInit() (*model.UserFriend, error)
	FirstOrCreate() (*model.UserFriend, error)
	FindByPage(offset int, limit int) (result []*model.UserFriend, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IUserFriendDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (u userFriendDo) Debug() IUserFriendDo {
	return u.withDO(u.DO.Debug())
}

func (u userFriendDo) WithContext(ctx context.Context) IUserFriendDo {
	return u.withDO(u.DO.WithContext(ctx))
}

func (u userFriendDo) ReadDB() IUserFriendDo {
	return u.Clauses(dbresolver.Read)
}

func (u userFriendDo) WriteDB() IUserFriendDo {
	return u.Clauses(dbresolver.Write)
}

func (u userFriendDo) Session(config *gorm.Session) IUserFriendDo {
	return u.withDO(u.DO.Session(config))
}

func (u userFriendDo) Clauses(conds ...clause.Expression) IUserFriendDo {
	return u.withDO(u.DO.Clauses(conds...))
}

func (u userFriendDo) Returning(value interface{}, columns ...string) IUserFriendDo {
	return u.withDO(u.DO.Returning(value, columns...))
}

func (u userFriendDo) Not(conds ...gen.Condition) IUserFriendDo {
	return u.withDO(u.DO.Not(conds...))
}

func (u userFriendDo) Or(conds ...gen.Condition) IUserFriendDo {
	return u.withDO(u.DO.Or(conds...))
}

func (u userFriendDo) Select(conds ...field.Expr) IUserFriendDo {
	return u.withDO(u.DO.Select(conds...))
}

func (u userFriendDo) Where(conds ...gen.Condition) IUserFriendDo {
	return u.withDO(u.DO.Where(conds...))
}

func (u userFriendDo) Order(conds ...field.Expr) IUserFriendDo {
	return u.withDO(u.DO.Order(conds...))
}

func (u userFriendDo) Distinct(cols ...field.Expr) IUserFriendDo {
	return u.withDO(u.DO.Distinct(cols...))
}

func (u userFriendDo) Omit(cols ...field.Expr) IUserFriendDo {
	return u.withDO(u.DO.Omit(cols...))
}

func (u userFriendDo) Join(table schema.Tabler, on ...field.Expr) IUserFriendDo {
	return u.withDO(u.DO.Join(table, on...))
}

func (u userFriendDo) LeftJoin(table schema.Tabler, on ...field.Expr) IUserFriendDo {
	return u.withDO(u.DO.LeftJoin(table, on...))
}

func (u userFriendDo) RightJoin(table schema.Tabler, on ...field.Expr) IUserFriendDo {
	return u.withDO(u.DO.RightJoin(table, on...))
}

func (u userFriendDo) Group(cols ...field.Expr) IUserFriendDo {
	return u.withDO(u.DO.Group(cols...))
}

func (u userFriendDo) Having(conds ...gen.Condition) IUserFriendDo {
	return u.withDO(u.DO.Having(conds...))
}

func (u userFriendDo) Limit(limit int) IUserFriendDo {
	return u.withDO(u.DO.Limit(limit))
}

func (u userFriendDo) Offset(offset int) IUserFriendDo {
	return u.withDO(u.DO.Offset(offset))
}

func (u userFriendDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IUserFriendDo {
	return u.withDO(u.DO.Scopes(funcs...))
}

func (u userFriendDo) Unscoped() IUserFriendDo {
	return u.withDO(u.DO.Unscoped())
}

func (u userFriendDo) Create(values ...*model.UserFriend) error {
	if len(values) == 0 {
		return nil
	}
	return u.DO.Create(values)
}

func (u userFriendDo) CreateInBatches(values []*model.UserFriend, batchSize int) error {
	return u.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (u userFriendDo) Save(values ...*model.UserFriend) error {
	if len(values) == 0 {
		return nil
	}
	return u.DO.Save(values)
}

func (u userFriendDo) First() (*model.UserFriend, error) {
	if result, err := u.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserFriend), nil
	}
}

func (u userFriendDo) Take() (*model.UserFriend, error) {
	if result, err := u.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserFriend), nil
	}
}

func (u userFriendDo) Last() (*model.UserFriend, error) {
	if result, err := u.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserFriend), nil
	}
}

func (u userFriendDo) Find() ([]*model.UserFriend, error) {
	result, err := u.DO.Find()
	return result.([]*model.UserFriend), err
}

func (u userFriendDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.UserFriend, err error) {
	buf := make([]*model.UserFriend, 0, batchSize)
	err = u.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (u userFriendDo) FindInBatches(result *[]*model.UserFriend, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return u.DO.FindInBatches(result, batchSize, fc)
}

func (u userFriendDo) Attrs(attrs ...field.AssignExpr) IUserFriendDo {
	return u.withDO(u.DO.Attrs(attrs...))
}

func (u userFriendDo) Assign(attrs ...field.AssignExpr) IUserFriendDo {
	return u.withDO(u.DO.Assign(attrs...))
}

func (u userFriendDo) Joins(fields ...field.RelationField) IUserFriendDo {
	for _, _f := range fields {
		u = *u.withDO(u.DO.Joins(_f))
	}
	return &u
}

func (u userFriendDo) Preload(fields ...field.RelationField) IUserFriendDo {
	for _, _f := range fields {
		u = *u.withDO(u.DO.Preload(_f))
	}
	return &u
}

func (u userFriendDo) FirstOrInit() (*model.UserFriend, error) {
	if result, err := u.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserFriend), nil
	}
}

func (u userFriendDo) FirstOrCreate() (*model.UserFriend, error) {
	if result, err := u.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserFriend), nil
	}
}

func (u userFriendDo) FindByPage(offset int, limit int) (result []*model.UserFriend, count int64, err error) {
	result, err = u.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = u.Offset(-1).Limit(-1).Count()
	return
}

func (u userFriendDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = u.Count()
	if err != nil {
		return
	}

	err = u.Offset(offset).Limit(limit).Scan(result)
	return
}

func (u userFriendDo) Scan(result interface{}) (err error) {
	return u.DO.Scan(result)
}

func (u userFriendDo) Delete(models ...*model.UserFriend) (result gen.ResultInfo, err error) {
	return u.DO.Delete(models)
}

func (u *userFriendDo) withDO(do gen.Dao) *userFriendDo {
	u.DO = *do.(*gen.DO)
	return u
}
