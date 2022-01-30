package article

import (
	"asyncservice/client/kafka"
	"asyncservice/client/online"
	"asyncservice/client/social"
	"asyncservice/util/constant"
	"context"
)

func PublishArticle(publish *kafka.KafkaMessage) {
	uid := publish.Info.UID
	//articleID := publish.Info.ArticleID
	baseCtx := context.Background()
	ctx, cancel := context.WithTimeout(baseCtx, constant.RPCTimeOut)
	defer cancel()
	uids, err := social.GetFollowerAll(ctx, uid)
	if err != nil {
		return
	}

	if len(uids) > constant.FollowCountLimit {
		ctx, cancel = context.WithTimeout(baseCtx, constant.RPCTimeOut)
		defer cancel()
		onlineUIDs, _, err := online.GetOnlineAll(ctx)
		if err != nil {
			return
		}
		uids = union(uids, onlineUIDs)
	}

	//ctx, cancel = context.WithTimeout(baseCtx, constant.RPCTimeOut)
	//defer cancel()
	//_ = article.PushFollowFeed(ctx, uid, articleID, uids)
}

func union(a, b []int64) []int64 {
	c := make([]int64, 0, len(a))
	m := make(map[int64]struct{}, len(b))
	for _, v := range b {
		m[v] = struct{}{}
	}
	for _, v := range a {
		if _, ok := m[v]; ok {
			c = append(c, v)
		}
	}
	return c
}
