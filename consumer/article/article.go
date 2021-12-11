package article

import (
	"asyncservice/client/article"
	"asyncservice/client/kafka"
	"asyncservice/client/social"
	"asyncservice/util/constant"
	"context"
)

func PublishArticle(publish *kafka.KafkaMessage) {
	uid := publish.Info.UID
	articleID := publish.Info.ArticleID
	baseCtx := context.Background()
	ctx, cancel := context.WithTimeout(baseCtx, constant.RPCTimeOut)
	defer cancel()
	uids, err := social.GetFollowerAll(ctx, uid)
	if err != nil {
		return
	}
	// fixme 需过滤非活跃或非在线人数

	ctx, cancel = context.WithTimeout(baseCtx, constant.RPCTimeOut)
	defer cancel()
	_ = article.PushFollowFeed(ctx, uid, articleID, uids)
}
