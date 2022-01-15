package article

import (
	"asyncservice/conf"
	"asyncservice/global"
	"asyncservice/rpc/article/pb"
	"asyncservice/util/constant"
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

func PushFollowFeed(ctx context.Context, uid, articleID int64, uids []int64) error {
	stream, err := client.PushFollowFeed(ctx)
	defer stream.Close()
	if err != nil {
		global.ExcLog.Printf("ctx %v PushFollowFeed uid %v articleid %v uids %v err %v", ctx, uids, articleID, uids, err)
		return err
	}
	for i := 0; i < len(uids); i += constant.BatchSize {
		left := i
		right := i + constant.BatchSize
		if right > len(uids) {
			right = len(uids)
		}
		tus := uids[left:right]
		err = stream.Send(toPushFollowFeedRequest(uid, articleID, tus))
		if err != nil {
			global.ExcLog.Printf("ctx %v PushFollowFeed streamsend uid %v articleid %v uids %v err %v", ctx, uids, articleID, uids, err)
			return err
		}
	}
	return nil
}
