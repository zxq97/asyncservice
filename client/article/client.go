package article

import (
	"asyncservice/conf"
	"asyncservice/global"
	"asyncservice/rpc/article/pb"
	"context"
	"github.com/micro/go-micro"
)

var (
	client article_service.ArticleServerService
)

func InitClient(config *conf.Conf) {
	service := micro.NewService(micro.Name(
		config.Grpc.Name),
	)
	client = article_service.NewArticleServerService(
		config.Grpc.Name,
		service.Client(),
	)
}

//ChangeVisibleType(ctx context.Context, in *VisibleTypeRequest, opts ...client.CallOption) (*VisibleTypeResponse, error)

func GetArticle(ctx context.Context, articleID int64) (*Article, error) {
	res, err := client.GetArticle(ctx, toArticleRequest(articleID))
	if err != nil {
		global.ExcLog.Printf("ctx %v GetArticle articleid %v err %v", ctx, articleID, err)
		return nil, err
	}
	return toArticle(res.ArticleInfo), nil
}

func GetBatchArticle(ctx context.Context, articleIDs []int64) (map[int64]*Article, error) {
	res, err := client.GetBatchArticle(ctx, toArticleBatchRequest(articleIDs))
	if err != nil {
		global.ExcLog.Printf("ctx %v GetBatchArticle articleid %v err %v", ctx, articleIDs, err)
		return nil, err
	}
	articleMap := make(map[int64]*Article, len(articleIDs))
	for k, v := range res.ArticleInfos {
		articleMap[k] = toArticle(v)
	}
	return articleMap, nil
}

func GetTopic(ctx context.Context, topicID int64) (*Topic, error) {
	res, err := client.GetTopic(ctx, toTopicRequest(topicID))
	if err != nil {
		global.ExcLog.Printf("ctx %v GetTopic topicid %v err %v", ctx, topicID, err)
		return nil, err
	}
	return toTopic(res.TopicInfo), nil
}

func GetBatchTopic(ctx context.Context, topicIDs []int64) (map[int64]*Topic, error) {
	res, err := client.GetBatchTopic(ctx, toTopicBatchRequest(topicIDs))
	if err != nil {
		global.ExcLog.Printf("ctx %v GetBatchTopic topicids %v err %v", ctx, topicIDs, err)
		return nil, err
	}
	topicMap := make(map[int64]*Topic, len(topicIDs))
	for k, v := range res.TopicInfos {
		topicMap[k] = toTopic(v)
	}
	return topicMap, nil
}

func PublishArticle(ctx context.Context, articleID, uid, topicID int64, content string, vType int32) error {
	_, err := client.PublishArticle(ctx, toPublishArticleRequest(articleID, uid, topicID, content, vType))
	if err != nil {
		global.ExcLog.Printf("ctx %v PublishArticle articleid %v err %v", ctx, articleID, err)
	}
	return err
}

func DeleteArticle(ctx context.Context, articleID int64) error {
	_, err := client.DeleteArticle(ctx, toArticleRequest(articleID))
	if err != nil {
		global.ExcLog.Printf("ctx %v DeleteArticle articleid %v err %v", ctx, articleID, err)
	}
	return err
}

func PushInBox(ctx context.Context, uid, articleID int64) error {
	_, err := client.PushInBox(ctx, toPushInBoxRequest(uid, articleID))
	if err != nil {
		global.ExcLog.Printf("ctx %v PushInBox uid %v articleid %v err %v", ctx, uid, articleID, err)
	}
	return err
}

func GetInBox(ctx context.Context, uid, cursor, offset int64) ([]int64, int64, bool, error) {
	res, err := client.GetInBox(ctx, toGetInBoxRequest(uid, cursor, offset))
	if err != nil {
		global.ExcLog.Printf("ctx %v GetInBox uid %v cursor %v offset %v err %v", ctx, uid, cursor, offset, err)
		return nil, 0, false, err
	}
	return res.ArticleIds, res.NextCursor, res.HasMore, nil
}
