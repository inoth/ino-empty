package logger

import (
	"defaultProject/config"
	"defaultProject/logger/hooks"
	"defaultProject/util"
	"github.com/nsqio/go-nsq"
	"github.com/sirupsen/logrus"
)

type LogrusConfig struct {
	Hooks []logrus.Hook
}

func (l *LogrusConfig) AddHook(hook ...logrus.Hook) {
	l.Hooks = hook
}

func (l *LogrusConfig) Init() error {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	for _, hook := range l.Hooks {
		logrus.AddHook(hook)
	}
	return nil
}

func DefualtNsqHook(topicName string,level logrus.Level) logrus.Hook {
	cfg := nsq.NewConfig()
	client, err := nsq.NewProducer(config.Cfg.GetString("Nsq.Host"), cfg)
	util.Must(err)

	defer client.Stop()
	hook, err := hooks.NewAsyncNsqHook(client, topicName, level)
	util.Must(err)
	return hook
}
