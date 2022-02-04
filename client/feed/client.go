package feed

import (
	"asyncservice/conf"
	"asyncservice/global"
	"asyncservice/rpc/feed/pb"
	"asyncservice/util/constant"
	"context"
	"github.com/micro/go-micro"
)

var (
	client feed_service.FeedServerService
)

func InitClient(config *conf.Conf) {
	service := micro.NewService(micro.Name(
		config.Grpc.Name),
	)
	client = feed_service.NewFeedServerService(
		config.Grpc.Name,
		service.Client(),
	)
}

func FollowAfterFeed(ctx context.Context, uid, toUID int64) error {
	_, err := client.FollowAfterFeed(ctx, toActionFeedRequest(uid, toUID))
	if err != nil {
		global.ExcLog.Printf("ctx %v FollowAfterFeed uid %v touid %v err %v", ctx, uid, toUID, err)
	}
	return err
}

func UnfollowAfterFeed(ctx context.Context, uid, toUID int64) error {
	_, err := client.UnfollowAfterFeed(ctx, toActionFeedRequest(uid, toUID))
	if err != nil {
		global.ExcLog.Printf("ctx %v UnfollowAfterFeed uid %v touid %v err %v", ctx, uid, toUID, err)
	}
	return err
}

func PushFollowFeed(ctx context.Context, uids []int64, uid, articleID int64) error {
	stream, err := client.PushFollowFeed(ctx)
	if err != nil {
		global.ExcLog.Printf("ctx %v PushFollowFeed getstream uids %v uid %v articleid %v err %v", ctx, uids, uid, articleID, err)
		return err
	}
	for i := 0; i < len(uids); i += constant.BatchSize {
		left := i
		right := i + constant.BatchSize
		if right > len(uids) {
			right = len(uids)
		}
		err = stream.Send(toPushFollowFeedRequest(uids[left:right], uid, articleID))
		if err != nil {
			global.ExcLog.Printf("ctx %v PushFollowFeed send uids %v uid %v articleid %v err %v", ctx, uids[left:right], uid, articleID, err)
			continue
		}
	}
	return nil
}

func GetFollowFeed(ctx context.Context, uid, cursor, offset int64) ([]int64, int64, error) {
	res, err := client.GetFollowFeed(ctx, toGetFeedRequest(uid, cursor, offset))
	if err != nil {
		global.ExcLog.Printf("ctx %v GetFollowFeed uid %v cursor %v err %v", ctx, uid, cursor)
		return nil, 0, err
	}
	return res.ArticleIds, res.NextCursor, nil
}

func PushSelfFeed(ctx context.Context, uid, articleID int64) error {
	_, err := client.PushSelfFeed(ctx, toPushSelfFeedRequest(uid, articleID))
	if err != nil {
		global.ExcLog.Printf("ctx %v PushSelfFeed uid %v articleid %v err %v", ctx, uid, articleID, err)
	}
	return err
}

func GetSelfFeed(ctx context.Context, uid, cursor, offset int64) ([]int64, int64, error) {
	res, err := client.GetSelfFeed(ctx, toGetFeedRequest(uid, cursor, offset))
	if err != nil {
		global.ExcLog.Printf("ctx %v GetSelfFeed uid %v cursor %v err %v", ctx, uid, cursor, err)
		return nil, 0, err
	}
	return res.ArticleIds, res.NextCursor, nil
}

func Refresh(ctx context.Context, uid int64) error {
	_, err := client.Refresh(ctx, toRefreshRequest(uid))
	if err != nil {
		global.ExcLog.Printf("ctx %v Refresh uid %v err %v", ctx, err)
	}
	return err
}
