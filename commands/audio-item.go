package commands

import (
	"errors"
	"github.com/bwmarrin/discordgo"
)

type audioItem struct {
	opusData chan []byte
	dead     bool // set to true if needs to be cleaned up
}

func (a *audioItem) OpusChan() chan []byte {
	return a.opusData
}

func (a *audioItem) IsClosed() bool {
	return a.dead
}

func (a *audioItem) Cleanup() {
	a.dead = true
	<-a.opusData
}

func getVoiceChannelID(s *discordgo.Session, m *discordgo.MessageCreate) (string, error) {
	guild, err := s.Guild(m.GuildID)
	if err != nil {
		return "", err
	}

	for _, vc := range guild.VoiceStates {
		if vc.UserID == m.Author.ID {
			return vc.ChannelID, nil
		}
	}

	return "", errors.New("voice channel not found")
}