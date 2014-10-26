package plugin

import (
	"testing"
)

func TestPluginRegistration(t *testing.T) {
	underTest := NewPlugin("a", func(msg IncomingMessage) *OutgoingMessage {
		return nil
	})
	if len(Plugins) != 1 {
		t.Error("Expected 1 plugin to be registered.")
	}
	for key, p := range Plugins {
		if key != "a" && &p != underTest {
			t.Errorf("Unexpected plugin for key %s", key)
		}
	}
}
