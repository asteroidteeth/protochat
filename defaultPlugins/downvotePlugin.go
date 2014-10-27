package defaultPlugins

import (
	"math/rand"
	"time"

	"github.com/asteroidteeth/protochat/plugin"
)

var rng = rand.New(rand.NewSource(int64(time.Now().Nanosecond())))

var downvoteLinks = []string{
	"http://www.dstaylor.me/wp-content/uploads/2012/02/downvote.png",
	"http://www.reactiongifs.us/wp-content/uploads/2013/02/downvote_dodgeball.gif",
	"http://cdn.gifbay.com/2013/04/anchorman_downvote-44171.gif",
	"http://stream1.gifsoup.com/view6/4098973/rain-downvote-o.gif",
}

var i = 0

func nextDownvote() string {
	return downvoteLinks[rng.Int()%len(downvoteLinks)]
}

var dvMatcher = "downvote %s"

func dvHandler(msg plugin.IncomingMessage) *plugin.OutgoingMessage {
	return &plugin.OutgoingMessage{msg.FromRoom, nextDownvote()}
}

var dvPlugin = plugin.NewPlugin(dvMatcher, dvHandler)
