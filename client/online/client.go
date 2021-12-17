package online

import (
	"asyncservice/conf"
	"asyncservice/rpc/online/pb"
	"context"
	"github.com/micro/go-micro"
	"log"
)

var (
	client online_service.OnlineServerService
)

func InitClient(config *conf.Conf) {
	service := micro.NewService(micro.Name(
		config.Grpc.Name),
	)
	client = online_service.NewOnlineServerService(
		config.Grpc.Name,
		service.Client(),
	)
}

func StartUp(ctx context.Context, uid int64) error {
	_, err := client.StartUp(ctx, toOnlineRequest(uid))
	if err != nil {
		log.Printf("ctx %v StartUp uid %v err %v", ctx, uid, err)
	}
	return err
}

func Shutdown(ctx context.Context, uid int64) error {
	_, err := client.StartUp(ctx, toOnlineRequest(uid))
	if err != nil {
		log.Printf("ctx %v Shutdown uid %v err %v", ctx, uid, err)
	}
	return err
}

func GetOnlineAll(ctx context.Context) ([]int64, int64, error) {
	res, err := client.GetOnlineAll(ctx, &online_service.EmptyRequest{})
	if err != nil {
		log.Printf("ctx %v GetOnlineAll err %v", ctx, err)
		return nil, 0, err
	}
	return res.Uids, res.Count, nil
}
