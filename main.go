package main

import (
	"github.com/DiscordBot/bot"
)



func main() {

	bot.Start()
	<- make (chan struct{})
}



