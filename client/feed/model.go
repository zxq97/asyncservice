package feed

import (
	"asyncservice/rpc/feed/pb"
)

func toRefreshRequest(uid int64) *feed_service.ReFreshRequest {
	return &feed_service.ReFreshRequest{
		Uid: uid,
	}
}

func toGetFeedRequest(uid, cursor, offset int64) *feed_service.GetFeedRequest {
	return &feed_service.GetFeedRequest{
		Uid:    uid,
		Cursor: cursor,
		Offset: offset,
	}
}

func toPushSelfFeedRequest(uid, articleID int64) *feed_service.PushSelfFeedRequest {
	return &feed_service.PushSelfFeedRequest{
		Uid:       uid,
		ArticleId: articleID,
	}
}

func toPushFollowFeedRequest(uids []int64, uid, articleID int64) *feed_service.PushFollowFeedRequest {
	return &feed_service.PushFollowFeedRequest{
		Uids:      uids,
		Uid:       uid,
		ArticleId: articleID,
	}
}

func toActionFeedRequest(uid, toUID int64) *feed_service.ActionFeedRequest {
	return &feed_service.ActionFeedRequest{
		Uid:   uid,
		ToUid: toUID,
	}
}
