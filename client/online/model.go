package online

import "asyncservice/rpc/online/pb"

func toOnlineRequest(uid int64) *online_service.OnlineRequest {
	return &online_service.OnlineRequest{
		Uid: uid,
	}
}
