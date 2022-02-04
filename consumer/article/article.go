package article

import (
	"asyncservice/client/feed"
	"asyncservice/client/kafka"
	"asyncservice/client/social"
	"asyncservice/global"
	"asyncservice/util/concurrent"
	"context"
)

func PublishArticle(ctx context.Context, publish *kafka.KafkaMessage) {
	uid := publish.Info.UID
	articleID := publish.Info.ArticleID
	var (
		uids []int64
		err  error
	)
	wg := concurrent.NewWaitGroup()
	wg.Run(func() {
		err = feed.PushSelfFeed(ctx, uid, articleID)
		if err != nil {
			global.ExcLog.Printf("ctx %v publisharticle pushselffeed uid %v articleid %v err %v", ctx, uid, articleID)
		}
	})
	wg.Run(func() {
		uids, err = social.GetFollowerAll(ctx, uid)
		if err != nil {
			global.ExcLog.Printf("ctx %v publisharticle getfollowerall uid %v err %v", ctx, uid, err)
		}
	})
	wg.Wait()

	//if len(uids) > constant.FollowCountLimit {
	//	onlineUIDs, _, err := online.GetOnlineAll(ctx)
	//	if err != nil {
	//		return
	//	}
	//	uids = union(uids, onlineUIDs)
	//}
	// todo 之后需要改成 推拉结合的

	err = feed.PushFollowFeed(ctx, uids, uid, articleID)
	if err != nil {
		global.ExcLog.Printf("ctx %v publisharticle pushfollowfeed uids %v uid %v articleid %v err %v", ctx, uids, uid, err)
	}

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
