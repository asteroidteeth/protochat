package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path"
	"strings"

	"github.com/daneharrigan/hipchat"
	"github.com/rakyll/globalconf"
	"github.com/texttheater/golang-levenshtein/levenshtein"

	"github.com/asteroidteeth/protochat/plugin"
	_ "github.com/asteroidteeth/protochat/plugins"
)

var (
	configPath        = path.Join(os.Getenv("GOPATH"), "src/github.com/asteroidteeth/protochat/config.ini")
	config, configErr = globalconf.NewWithOptions(&globalconf.Options{
		Filename: configPath,
	})

	username = flag.String("hipchat_username", "", "Hipchat bot username")
	password = flag.String("hipchat_password", "", "Hipchat bot password")
	c        *hipchat.Client
	users    []*hipchat.User
)

const (
	host     = "chat.hipchat.com:5222"
	fullname = "ProtoTest Bot"
)

func main() {
	if configErr != nil {
		fmt.Printf("Error loading config! \"%s\"\n", configErr.Error())
	}
	config.Parse()
	log.Printf("Username is %s", *username)
	log.Printf("Password is %s", *password)
	// if appKey == nil {
	// 	fmt.Println("Hipchat app key not specified...")
	// 	os.Exit(1)
	// }

	// xmpp.DefaultConfig = tls.Config{
	// 	ServerName: "chat.hipchat.com",
	// 	//InsecureSkipVerify: false,
	// }
	// client, err := xmpp.NewClient(host, username, password)
	// if err != nil {
	// 	fmt.Printf("Error starting XMPP client: %s\n", err.Error())
	// 	os.Exit(1)
	// }
	// client.Status(xmpp.DoNotDisturb, "Go away, I'm a reclusive bot")

	doTheUnexpected()
}

func doTheUnexpected() {
	c, err := hipchat.NewClient(*username, *password, "bot")
	if err != nil {
		log.Fatalf("Failed to start hipchat client: %s", err.Error())
	}

	c.Status("chat")
	users = c.Users()

	rooms := c.Rooms()

	for _, room := range rooms {
		fmt.Println(room.Name)
		fmt.Println(room.Id)
		c.Join(room.Id, fullname)
	}
	for message := range c.Messages() {
		log.Println(message.From)
		if strings.HasPrefix(message.Body, "@ProtoBot") {
			room, userFullName := parseFrom(message.From)
			mName := mentionName(userFullName)
			msgData := strings.TrimPrefix(message.Body, "@ProtoBot ")

			var bestHandler plugin.PluginComponent
			levDist := 9999
			for pattern, handler := range plugin.Plugins {
				thisDist := levenshtein.DistanceForStrings([]rune(msgData), []rune(pattern), levenshtein.DefaultOptions)
				if thisDist < levDist {
					bestHandler = handler
				}
			}

			outgoing := bestHandler.HandleMessage(plugin.IncomingMessage{userFullName, room, mName, message.Body})
			if outgoing != nil {
				c.Say(outgoing.RoomId, fullname, outgoing.Message)
			}
		}
	}

}

func parseFrom(from string) (room, user string) {
	result := strings.Split(from, "/")
	return result[0], result[1]
}

func mentionName(fullname string) string {
	for _, user := range users {
		if user.Name == fullname {
			return user.MentionName
		}
	}
	return ""
}
