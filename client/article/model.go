package article

import (
	"asyncservice/rpc/article/pb"
	"time"
)

type Article struct {
	ArticleID   int64     `json:"article_id"`
	UID         int64     `json:"uid"`
	Content     string    `json:"content"`
	TopicID     int64     `json:"topic_id"`
	VisibleType int32     `json:"visible_type"`
	Ctime       time.Time `json:"ctime"`
}

type Topic struct {
	TopicID   int64  `json:"topic_id"`
	TopicName string `json:"topic_name"`
}

func toArticleInfo(articleID, uid, topicID int64, content string, vType int32) *article_service.ArticleInfo {
	return &article_service.ArticleInfo{
		ArticleId:   articleID,
		Uid:         uid,
		Content:     content,
		TopicId:     topicID,
		VisibleType: vType,
	}
}

func toPublishArticleRequest(articleID, uid, topicID int64, content string, vType int32) *article_service.PublishArticleRequest {
	return &article_service.PublishArticleRequest{
		ArticleInfo: toArticleInfo(articleID, uid, topicID, content, vType),
	}
}

func toArticleRequest(articleID int64) *article_service.ArticleRequest {
	return &article_service.ArticleRequest{
		ArticleId: articleID,
	}
}

func toArticleBatchRequest(articleIDs []int64) *article_service.ArticleBatchRequest {
	return &article_service.ArticleBatchRequest{
		ArticleIds: articleIDs,
	}
}

func toTopicRequest(topicID int64) *article_service.TopicRequest {
	return &article_service.TopicRequest{
		TopicId: topicID,
	}
}

func toTopicBatchRequest(topicIDs []int64) *article_service.TopicBatchRequest {
	return &article_service.TopicBatchRequest{
		TopicIds: topicIDs,
	}
}

func toArticle(article *article_service.ArticleInfo) *Article {
	return &Article{
		ArticleID:   article.ArticleId,
		UID:         article.Uid,
		Content:     article.Content,
		TopicID:     article.TopicId,
		VisibleType: article.VisibleType,
		Ctime:       time.Unix(article.Ctime, 0),
	}
}

func toTopic(topic *article_service.TopicInfo) *Topic {
	return &Topic{
		TopicID:   topic.TopicId,
		TopicName: topic.TopicName,
	}
}

func toPushFollowFeedRequest(uid, articleID int64, uids []int64) *article_service.PushFollowFeedRequest {
	return &article_service.PushFollowFeedRequest{
		Uid:       uid,
		ArticleId: articleID,
		Uids:      uids,
	}
}
