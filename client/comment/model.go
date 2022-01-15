package comment

import (
	"asyncservice/rpc/comment/pb"
)

func toCommentInfo(commentID, pCommentID, articleID, uid, replyUID int64, content string) *comment_service.CommentInfo {
	return &comment_service.CommentInfo{
		CommentId:  commentID,
		PCommentId: pCommentID,
		ArticleId:  articleID,
		Uid:        uid,
		ReplyUid:   replyUID,
		Content:    content,
	}
}

func toPublishCommentRequest(commentID, pCommentID, articleID, uid, replyUID int64, content string) *comment_service.PublishCommentRequest {
	return &comment_service.PublishCommentRequest{
		CommentInfo: toCommentInfo(commentID, pCommentID, articleID, uid, replyUID, content),
	}
}

func toCountRequest(articleIDs []int64) *comment_service.GetCountRequest {
	return &comment_service.GetCountRequest{
		ArticleIds: articleIDs,
	}
}

func toLikeStateRequest(articleIDs []int64, uid int64) *comment_service.GetLikeStateRequest {
	return &comment_service.GetLikeStateRequest{
		ArticleIds: articleIDs,
		Uid:        uid,
	}
}

func toLikeInfo(articleID, uid int64) *comment_service.LikeInfo {
	return &comment_service.LikeInfo{
		ArticleId: articleID,
		Uid:       uid,
	}
}

func toLikePointRequest(articleID, uid int64) *comment_service.LikePointRequest {
	return &comment_service.LikePointRequest{
		LikeInfo: toLikeInfo(articleID, uid),
	}
}
