package defaultPlugins

import (
	"fmt"
	"github.com/asteroidteeth/protochat/plugin"
)

var Matcher = "you there?"

func handler(msg plugin.IncomingMessage) *plugin.OutgoingMessage {
	return &plugin.OutgoingMessage{msg.FromRoom, fmt.Sprintf("@%s yep!", msg.FromMentionName)}
}

var YepPlugin = plugin.NewPlugin(Matcher, handler)
