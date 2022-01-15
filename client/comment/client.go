package comment

import (
	"asyncservice/conf"
	"asyncservice/global"
	"asyncservice/rpc/comment/pb"
	"context"
	"github.com/micro/go-micro"
)

var (
	client comment_service.CommentServerService
)

func InitClient(config *conf.Conf) {
	service := micro.NewService(micro.Name(
		config.Grpc.Name),
	)
	client = comment_service.NewCommentServerService(
		config.Grpc.Name,
		service.Client(),
	)
}

//func GetCommentList(ctx context.Context, articleID, cursor, offset int64) error {
//	return nil
//}

func PublishComment(ctx context.Context, commentID, pCommentID, articleID, uid, replyUID int64, content string) error {
	_, err := client.PublishComment(ctx, toPublishCommentRequest(commentID, pCommentID, articleID, uid, replyUID, content))
	if err != nil {
		global.ExcLog.Printf("ctx %v PublishComment commentid %v pcommentid %v articleid %v uid %v replyuid %v content %v err %v", ctx, commentID, pCommentID, articleID, uid, replyUID, content, err)
	}
	return err
}

func DeleteComment(ctx context.Context, commentID, pCommentID, articleID, uid, replyUID int64, content string) error {
	_, err := client.DeleteComment(ctx, toPublishCommentRequest(commentID, pCommentID, articleID, uid, replyUID, content))
	if err != nil {
		global.ExcLog.Printf("ctx %v DeleteComment commentid %v pcommentid %v articleid %v uid %v replyuid %v content %v err %v", ctx, commentID, pCommentID, articleID, uid, replyUID, content, err)
	}
	return err
}

func GetCommentCount(ctx context.Context, articleIDs []int64) (map[int64]int64, error) {
	res, err := client.GetCommentCount(ctx, toCountRequest(articleIDs))
	if err != nil || res == nil {
		global.ExcLog.Printf("ctx %v GetCommentCount articleids %v err %v", ctx, articleIDs, err)
	}
	return res.LikeCount, nil
}

func GetLikeCount(ctx context.Context, articleIDs []int64) (map[int64]int64, error) {
	res, err := client.GetLikeCount(ctx, toCountRequest(articleIDs))
	if err != nil || res == nil {
		global.ExcLog.Printf("ctx %v GetLikeCount articleids %v err %v", ctx, articleIDs, err)
	}
	return res.LikeCount, nil
}

func GetLikeState(ctx context.Context, articleIDs []int64, uid int64) (map[int64]bool, error) {
	res, err := client.GetLikeState(ctx, toLikeStateRequest(articleIDs, uid))
	if err != nil || res == nil {
		global.ExcLog.Printf("ctx %v GetLikeState articleids %v err %v", ctx, articleIDs, err)
	}
	return res.Ok, nil
}

func LikePoint(ctx context.Context, articleID, uid int64) error {
	_, err := client.LikePoint(ctx, toLikePointRequest(articleID, uid))
	if err != nil {
		global.ExcLog.Printf("ctx %v LikePoint articleid %v uid %v err %v", ctx, articleID, uid, err)
	}
	return err
}

func CancelLike(ctx context.Context, articleID, uid int64) error {
	_, err := client.CancelLike(ctx, toLikePointRequest(articleID, uid))
	if err != nil {
		global.ExcLog.Printf("ctx %v CancelLike articleid %v uid %v err %v", ctx, articleID, uid, err)
	}
	return err
}
