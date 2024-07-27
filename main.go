package main

import (
	g "xabbo.b7c.io/goearth"
	"xabbo.b7c.io/goearth/shockwave/out"
)

var ext = g.NewExt(g.ExtInfo{
	Title:       "Shoutify",
	Description: "Are you tired of your chat messages being unreadable? Want to make sure your voice is heard loud and clear? Introducing Shoutify - the ultimate chat extension that transforms all your say-chats into attention-grabbing shouts!",
	Version:     "1.0.0",
	Author:      "Nanobyte",
})

func main() {
	ext.Intercept(out.CHAT).With(handleChat)
	ext.Run()
}

func handleChat(e *g.Intercept) {
	e.Packet.Header = ext.Headers().Get(out.SHOUT)
}