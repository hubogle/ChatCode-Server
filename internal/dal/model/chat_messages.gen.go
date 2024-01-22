// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameChatMessage = "chat_messages"

// ChatMessage mapped from table <chat_messages>
type ChatMessage struct {
	ID        uint64 `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	UserID    int32  `gorm:"column:user_id;not null" json:"user_id"` // 用户ID
	RoomID    int32  `gorm:"column:room_id;not null" json:"room_id"` // 房间ID
	Content   string `gorm:"column:content;not null" json:"content"` // 聊天内容
	CreatedAt int64  `gorm:"column:created_at;not null" json:"created_at"`
	UpdatedAt int64  `gorm:"column:updated_at;not null" json:"updated_at"`
	DeletedAt *int64 `gorm:"column:deleted_at" json:"deleted_at"`
}

// IsEmpty determines whether the structure is empty
func (m *ChatMessage) IsEmpty() bool {
	if m == nil {
		return true
	}
	return m.ID == 0
}

// GetID get the ID of the database table
func (m *ChatMessage) GetID() int64 {
	return int64(m.ID)
}

// TableName ChatMessage's table name
func (*ChatMessage) TableName() string {
	return TableNameChatMessage
}
