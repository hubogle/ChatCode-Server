// Code generated by ginctl. DO NOT EDIT.
package room

// RoomCreateResp 创建房间
type RoomCreateResp struct {
	// RoomId 房间id
	RoomId uint64 `binding:"required" json:"room_id"`
}

// RoomCreateReq defines model for RoomCreateReq.
type RoomCreateReq struct {
	// Info 房间信息
	Info *string `json:"info,omitempty"`

	// Name 房间名称
	Name string `binding:"required" json:"name"`

	// Salt 房间密码
	Salt *string `json:"salt,omitempty"`
}

// RoomJoinReq defines model for RoomJoinReq.
type RoomJoinReq struct {
	// RoomId 房间id
	RoomId uint64 `binding:"required" json:"room_id"`
}

// RoomGetReq room get
type RoomGetReq struct {
	// Id room id
	Id int `binding:"required" uri:"id"`
}
