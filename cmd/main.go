package main

import (
	"asyncservice/client/article"
	"asyncservice/client/kafka"
	"asyncservice/client/social"
	"asyncservice/client/user"
	"asyncservice/conf"
	"asyncservice/consumer"
	"asyncservice/util/concurrent"
	"net/http"
)

var (
	asyncConf   *conf.Conf
	articleConf *conf.Conf
	commentConf *conf.Conf
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

	article.InitClient(articleConf)
	social.InitClient(socialConf)
	user.InitClient(userConf)

	concurrent.Go(func() {
		consumer.InitConsumer(asyncConf.Kafka.Addr, kafka.TestTopic)
	})

	_ = http.ListenAndServe(asyncConf.Grpc.Addr, nil)
}