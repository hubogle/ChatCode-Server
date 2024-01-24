// Code generated by ginctl. DO NOT EDIT.
package chat

// ChatListReq chat list
type ChatListReq struct {
	// RoomUid room uid
	RoomUid string `binding:"required" form:"room_uid"`

	// Page page
	Page int `binding:"required" form:"page"`

	// PageSize page size
	PageSize int `binding:"required" form:"page_size"`
}

// ChatPrivateReq defines parameters for ChatPrivate.
type ChatPrivateReq struct {
	// Content chat content
	Content string `binding:"required" json:"content"`

	// ToUserId to user id
	ToUserId int `binding:"required" json:"to_user_id"`
}

// ChatRoomReq defines parameters for ChatRoom.
type ChatRoomReq struct {
	// Content chat content
	Content string `binding:"required" json:"content"`

	// RoomId room id
	RoomId int `binding:"required" json:"room_id"`
}
