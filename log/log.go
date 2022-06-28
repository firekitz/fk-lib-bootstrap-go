package fkLog

import (
	fkBootstrap "github.com/firekitz/fk-lib-bootstrap-go"
	"github.com/firekitz/fk-lib-bootstrap-go/slack"
	"github.com/sirupsen/logrus"
)

func E(msg string) {
	defer func() {
		logrus.Error(msg)
	}()
	payload := slack.Payload{
		Text:     msg,
		Username: fkBootstrap.SERVICE_NAME,
		//IconEmoji: "zzz",
	}
	payload.SendSlack(slack.SLACK_CHANNEL_EMERG)
}

func W(msg string) {
	defer func() {
		logrus.Warn(msg)
	}()
	payload := slack.Payload{
		Text:     msg,
		Username: fkBootstrap.SERVICE_NAME,
		//IconEmoji: "zzz",
	}
	payload.SendSlack(slack.SLACK_CHANNEL_WARN)
}

func I(msg string) {
	defer func() {
		logrus.Info(msg)
	}()
	payload := slack.Payload{
		Text:     msg,
		Username: fkBootstrap.SERVICE_NAME,
		//IconEmoji: "zzz",
	}
	payload.SendSlack(slack.SLACK_CHANNEL_INFO)
}

func D(msg string) {
	defer func() {
		logrus.Debug(msg)
	}()
	payload := slack.Payload{
		Text:     msg,
		Username: fkBootstrap.SERVICE_NAME,
		//IconEmoji: "zzz",
	}
	payload.SendSlack(slack.SLACK_CHANNEL_DEBUG)
}
