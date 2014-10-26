package plugin

// import (
// 	"github.com/daneharrigan/hipchat"
// )

var Plugins map[string]PluginComponent = make(map[string]PluginComponent)

type OutgoingMessage struct {
	RoomId  string
	Message string
}

type IncomingMessage struct {
	FromUser        string
	FromRoom        string
	FromMentionName string
	Body            string
}

type PluginComponent struct {
	HandleMessage func(message IncomingMessage) *OutgoingMessage
}

func NewPlugin(pattern string, handler func(message IncomingMessage) *OutgoingMessage) *PluginComponent {
	p := PluginComponent{handler}
	Plugins[pattern] = p
	return &p
}
