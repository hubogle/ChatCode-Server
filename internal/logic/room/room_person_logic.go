package room

import (
	"context"

	"github.com/hubogle/chatcode-server/internal/types"
	"github.com/hubogle/chatcode-server/internal/types/room"
)

// RoomPerson room person
// GET /api/v1/room/:id/person
func (l *logic) RoomPerson(ctx context.Context, req *room.RoomPersonReq) (resp room.GetRoomPersonResp, err error) {
	userBasicList, err := l.repo.GetUserBasicByRoomID(ctx, req.Id)
	if err != nil {
		return room.GetRoomPersonResp{}, err
	}

	resp.List = make([]types.ChatItemInfo, 0, len(userBasicList))
	for _, userBasic := range userBasicList {
		resp.List = append(resp.List, types.ChatItemInfo{
			Name: userBasic.Account,
			Uid:  userBasic.UID,
			Type: 1,
		})
	}
	return
}
