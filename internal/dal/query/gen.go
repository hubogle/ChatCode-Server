// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"
	"database/sql"

	"gorm.io/gorm"

	"gorm.io/gen"

	"gorm.io/plugin/dbresolver"
)

var (
	Q            = new(Query)
	Group        *group
	MessageBasic *messageBasic
	RoomBasic    *roomBasic
	UserBasic    *userBasic
	UserFriend   *userFriend
	UserRoom     *userRoom
)

func SetDefault(db *gorm.DB, opts ...gen.DOOption) {
	*Q = *Use(db, opts...)
	Group = &Q.Group
	MessageBasic = &Q.MessageBasic
	RoomBasic = &Q.RoomBasic
	UserBasic = &Q.UserBasic
	UserFriend = &Q.UserFriend
	UserRoom = &Q.UserRoom
}

func Use(db *gorm.DB, opts ...gen.DOOption) *Query {
	return &Query{
		db:           db,
		Group:        newGroup(db, opts...),
		MessageBasic: newMessageBasic(db, opts...),
		RoomBasic:    newRoomBasic(db, opts...),
		UserBasic:    newUserBasic(db, opts...),
		UserFriend:   newUserFriend(db, opts...),
		UserRoom:     newUserRoom(db, opts...),
	}
}

type Query struct {
	db *gorm.DB

	Group        group
	MessageBasic messageBasic
	RoomBasic    roomBasic
	UserBasic    userBasic
	UserFriend   userFriend
	UserRoom     userRoom
}

func (q *Query) Available() bool { return q.db != nil }

func (q *Query) clone(db *gorm.DB) *Query {
	return &Query{
		db:           db,
		Group:        q.Group.clone(db),
		MessageBasic: q.MessageBasic.clone(db),
		RoomBasic:    q.RoomBasic.clone(db),
		UserBasic:    q.UserBasic.clone(db),
		UserFriend:   q.UserFriend.clone(db),
		UserRoom:     q.UserRoom.clone(db),
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
		db:           db,
		Group:        q.Group.replaceDB(db),
		MessageBasic: q.MessageBasic.replaceDB(db),
		RoomBasic:    q.RoomBasic.replaceDB(db),
		UserBasic:    q.UserBasic.replaceDB(db),
		UserFriend:   q.UserFriend.replaceDB(db),
		UserRoom:     q.UserRoom.replaceDB(db),
	}
}

type queryCtx struct {
	Group        IGroupDo
	MessageBasic IMessageBasicDo
	RoomBasic    IRoomBasicDo
	UserBasic    IUserBasicDo
	UserFriend   IUserFriendDo
	UserRoom     IUserRoomDo
}

func (q *Query) WithContext(ctx context.Context) *queryCtx {
	return &queryCtx{
		Group:        q.Group.WithContext(ctx),
		MessageBasic: q.MessageBasic.WithContext(ctx),
		RoomBasic:    q.RoomBasic.WithContext(ctx),
		UserBasic:    q.UserBasic.WithContext(ctx),
		UserFriend:   q.UserFriend.WithContext(ctx),
		UserRoom:     q.UserRoom.WithContext(ctx),
	}
}

func (q *Query) Transaction(fc func(tx *Query) error, opts ...*sql.TxOptions) error {
	return q.db.Transaction(func(tx *gorm.DB) error { return fc(q.clone(tx)) }, opts...)
}

func (q *Query) Begin(opts ...*sql.TxOptions) *QueryTx {
	tx := q.db.Begin(opts...)
	return &QueryTx{Query: q.clone(tx), Error: tx.Error}
}

type QueryTx struct {
	*Query
	Error error
}

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
