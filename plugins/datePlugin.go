package defaultPlugins

import (
	"fmt"
	"time"

	"github.com/asteroidteeth/protochat/plugin"
)

var dateMatcher = "what's the time?"

func datehandler(msg plugin.IncomingMessage) *plugin.OutgoingMessage {
	now := time.Now()
	mName := msg.FromMentionName
	return &plugin.OutgoingMessage{
		msg.FromRoom,
		fmt.Sprintf("@%s it's %d:%d on %s, %s %d %d", mName,
			now.Hour(),
			now.Minute(),
			now.Weekday().String(),
			now.Month().String(),
			now.Day(),
			now.Year()),
	}
}

var DatePlugin = plugin.NewPlugin(dateMatcher, datehandler)
