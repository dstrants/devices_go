package slack

import (
	"github.com/slack-go/slack"

	config "devices/lib/config"
)

var cnf config.Config = config.LoadConfig()

// Sends a slack plain text message
func SendSimpleMessage(msg string) error {
	api := slack.New(cnf.Slack.Token)

	_, _, err := api.PostMessage(
		cnf.Slack.Channel,
		slack.MsgOptionText(msg, false),
		slack.MsgOptionAsUser(true),
	)
	return err
}
