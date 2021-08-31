package bot

import (
	"github.com/DiscordBot/commands"
	"github.com/DiscordBot/util"
	"github.com/bwmarrin/discordgo"
	"strings"
)
var BotId string
const Token string = "NjIyNDQ0ODkwMTQ3NzgyNjk2.XX0HAg.ugZaHCFfCDmQ6WKA-5Iz3jKRPFE"


func Start() {
	dg,err := discordgo.New("Bot " + Token)
	util.FailOnError(err,"Failed to create a new discord session")
	u,err := dg.User("@me")
	util.FailOnError(err,"Failed to fetch user id")
	BotId = u.ID
	dg.AddHandler(MessageHandler)
	err = dg.Open()
	util.FailOnError(err,"Error opening WebSocket to Discord")

	println("Bot is running")
}

func MessageHandler (s *discordgo.Session,m *discordgo.MessageCreate) {
	println("Guild Id: ",m.GuildID)
	if m.Author.ID == BotId {
		return
	}
	if m.Author.Bot {
		println(m.Content)
	}
	if m.Content == "ping" {
		_,err := s.ChannelMessageSend(m.ChannelID,"pong")
		util.FailOnError(err,"Message not Sent")
	}
	if m.Content == "!cat" {
		commands.Cat(s,m)
	}
	if m.Content == "!dog" {
		commands.Dog(s,m)
	}
	if strings.HasPrefix(m.Content,"#") {
		commands.Music(s,m)
	}
}