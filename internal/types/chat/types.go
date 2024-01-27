// Code generated by ginctl. DO NOT EDIT.
package chat

import (
	"github.com/hubogle/chatcode-server/internal/types"
)

// ChatListResp 会话列表
type ChatListResp struct {
	List []types.ChatItemInfo `binding:"required" json:"list"`
}

// ChatCreateReq defines model for ChatCreateReq.
type ChatCreateReq struct {
	// Salt 会话密码
	Salt *string `json:"salt,omitempty"`

	// Type 会话类型
	Type int `binding:"required,oneof=1 2" json:"type"`

	// Uid 会话id
	Uid uint64 `binding:"required" json:"uid"`
}
