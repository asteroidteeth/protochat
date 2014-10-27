package plugin

import (
	"github.com/texttheater/golang-levenshtein/levenshtein"
)

var Plugins map[string]PluginComponent = make(map[string]PluginComponent)

func GetHandler(message string) PluginComponent {
	var bestHandler PluginComponent
	levDist := 9999
	for pattern, handler := range Plugins {
		thisDist := levenshtein.DistanceForStrings([]rune(message), []rune(pattern), levenshtein.DefaultOptions)
		if thisDist < levDist {
			levDist = thisDist
			bestHandler = handler
		}
	}
	return bestHandler
}

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
