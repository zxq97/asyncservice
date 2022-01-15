package main

import (
	"asyncservice/client/article"
	"asyncservice/client/comment"
	"asyncservice/client/kafka"
	"asyncservice/client/online"
	"asyncservice/client/remind"
	"asyncservice/client/social"
	"asyncservice/client/user"
	"asyncservice/conf"
	"asyncservice/consumer"
	"asyncservice/global"
	"asyncservice/util/concurrent"
	"net/http"
)

var (
	asyncConf   *conf.Conf
	articleConf *conf.Conf
	commentConf *conf.Conf
	onlineConf  *conf.Conf
	remindConf  *conf.Conf
	socialConf  *conf.Conf
	userConf    *conf.Conf
	err         error
)

func main() {
	asyncConf, err = conf.LoadYaml(conf.ASyncConfPath)
	if err != nil {
		panic(err)
	}
	articleConf, err = conf.LoadYaml(conf.ArticleConfPath)
	if err != nil {
		panic(err)
	}
	commentConf, err = conf.LoadYaml(conf.CommentConfPath)
	if err != nil {
		panic(err)
	}
	onlineConf, err = conf.LoadYaml(conf.OnlineConfPath)
	if err != nil {
		panic(err)
	}
	remindConf, err = conf.LoadYaml(conf.RemindConfPath)
	if err != nil {
		panic(err)
	}
	socialConf, err = conf.LoadYaml(conf.SocialConfPath)
	if err != nil {
		panic(err)
	}
	userConf, err = conf.LoadYaml(conf.UserConfPath)
	if err != nil {
		panic(err)
	}

	global.InfoLog, err = conf.InitLog(asyncConf.InfoLog.Path)
	if err != nil {
		panic(err)
	}
	global.ExcLog, err = conf.InitLog(asyncConf.ExcLog.Path)
	if err != nil {
		panic(err)
	}

	article.InitClient(articleConf)
	comment.InitClient(commentConf)
	online.InitClient(onlineConf)
	remind.InitClient(remindConf)
	social.InitClient(socialConf)
	user.InitClient(userConf)

	concurrent.Go(func() {
		consumer.InitConsumer(asyncConf.Kafka.Addr, kafka.UserActionTopic)
	})

	_ = http.ListenAndServe(asyncConf.Grpc.Addr, nil)
}
