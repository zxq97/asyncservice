package social

import (
	"asyncservice/client/kafka"
	"asyncservice/client/user"
	"asyncservice/global"
	"context"
)

func Follow(follow *kafka.KafkaMessage) {
	uid := follow.Info.UID
	touid := follow.Info.ToUID
	userMap, err := user.GetBatchUserinfo(context.Background(), []int64{uid, touid})
	global.InfoLog.Printf("uid %v touid %v usermap %v err %v", uid, touid, userMap, err)
}

func Unfollow(unfollow *kafka.KafkaMessage) {

}
