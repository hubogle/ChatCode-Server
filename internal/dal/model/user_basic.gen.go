// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameUserBasic = "user_basic"

// UserBasic mapped from table <user_basic>
type UserBasic struct {
	ID        uint64  `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	UID       uint64  `gorm:"column:uid;not null" json:"uid"`           // 用户唯一标识 ID
	Account   string  `gorm:"column:account;not null" json:"account"`   // 用户名
	Password  string  `gorm:"column:password;not null" json:"password"` // 密码
	Nickname  string  `gorm:"column:nickname;not null" json:"nickname"` // 昵称
	Email     *string `gorm:"column:email" json:"email"`                // 邮箱
	CreatedAt int64   `gorm:"column:created_at;not null" json:"created_at"`
	UpdatedAt int64   `gorm:"column:updated_at;not null" json:"updated_at"`
	DeletedAt *int64  `gorm:"column:deleted_at" json:"deleted_at"`
}

// IsEmpty determines whether the structure is empty
func (m *UserBasic) IsEmpty() bool {
	if m == nil {
		return true
	}
	return m.ID == 0
}

// GetID get the ID of the database table
func (m *UserBasic) GetID() int64 {
	return int64(m.ID)
}

// TableName UserBasic's table name
func (*UserBasic) TableName() string {
	return TableNameUserBasic
}
