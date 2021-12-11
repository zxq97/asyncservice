package user

import (
	"asyncservice/conf"
	"asyncservice/rpc/user/pb"
	"context"
	"github.com/micro/go-micro"
	"log"
)

var (
	client user_service.UserServerService
)

func InitClient(config *conf.Conf) {
	service := micro.NewService(micro.Name(
		config.Grpc.Name),
	)
	client = user_service.NewUserServerService(
		config.Grpc.Name,
		service.Client(),
	)
}

//GetHistoryBrowse(ctx context.Context, in *FeedListRequest, opts ...client.CallOption) (*FeedListResponse, error)
//GetBlackList(ctx context.Context, in *BlackListRequest, opts ...client.CallOption) (*FeedListResponse, error)
//GetCollectionList(ctx context.Context, in *FeedListRequest, opts ...client.CallOption) (*FeedListResponse, error)
//Black(ctx context.Context, in *BlackRequest, opts ...client.CallOption) (*BlackResponse, error)
//CancelBlack(ctx context.Context, in *CancelBlackRequest, opts ...client.CallOption) (*BlackResponse, error)
//Collection(ctx context.Context, in *CollectionRequest, opts ...client.CallOption) (*CollectionResponse, error)
//CancelCollection(ctx context.Context, in *CancelCollectionRequest, opts ...client.CallOption) (*CollectionResponse, error)
//AddBrowse(ctx context.Context, in *AddBrowseRequest, opts ...client.CallOption) (*AddBrowseResponse, error)

func GetUserinfo(ctx context.Context, uid int64) (*User, error) {
	res, err := client.GetUserinfo(ctx, toUserinfoRequest(uid))
	if err != nil {
		log.Printf("ctx %v GetUserinfo uid %v err %v", ctx, uid, err)
		return nil, err
	}
	user := toUser(res.Userinfo)
	return user, nil
}

func GetBatchUserinfo(ctx context.Context, uids []int64) (map[int64]*User, error) {
	res, err := client.GetBatchUserinfo(ctx, toBatchUserinfoRequest(uids))
	if err != nil {
		log.Printf("ctx %v GetBatchUserinfo uids %v err %v", ctx, uids, err)
		return nil, err
	}
	userMap := make(map[int64]*User, len(res.Userinfos))
	for k, v := range res.Userinfos {
		userMap[k] = toUser(v)
	}
	return userMap, nil
}

func CreateUser(ctx context.Context, uid int64, nickname, introduction string, gender int32) error {
	_, err := client.CreateUser(ctx, toCreateUserRequest(uid, nickname, introduction, gender))
	if err != nil {
		log.Printf("ctx %v CreateUser uid %v nickname %v introduction %v gender %v err %v", ctx, uid, nickname, introduction, gender, err)
		return err
	}
	return nil
}
