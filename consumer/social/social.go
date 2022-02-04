package social

import (
	"asyncservice/client/feed"
	"asyncservice/client/kafka"
	"asyncservice/global"
	"context"
)

func Follow(ctx context.Context, follow *kafka.KafkaMessage) {
	uid := follow.Info.UID
	toUID := follow.Info.ToUID
	err := feed.FollowAfterFeed(ctx, uid, toUID)
	if err != nil {
		global.ExcLog.Printf("ctx %v Follow uid %v touid %v err %v", ctx, uid, toUID, err)
	}
}

func Unfollow(ctx context.Context, unfollow *kafka.KafkaMessage) {
	uid := unfollow.Info.UID
	toUID := unfollow.Info.ToUID
	err := feed.UnfollowAfterFeed(ctx, uid, toUID)
	if err != nil {
		global.ExcLog.Printf("ctx %v Unfollow uid %v touid %v err %v", ctx, uid, toUID, err)
	}
}
