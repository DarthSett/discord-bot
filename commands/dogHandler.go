package commands

import (
	"encoding/json"
	"fmt"
	"github.com/DiscordBot/util"
	"github.com/bwmarrin/discordgo"
	"io/ioutil"
	"log"
	"net/http"
)

func Dog (s *discordgo.Session, m *discordgo.MessageCreate) {
	println(1)
	res, err := http.Get("https://api.thedogapi.com/v1/images/search")
	util.FailOnError(err,"Failed to get a reply from thedogapi")
	if res.StatusCode != 200 {
		log.Fatalf("Status code error: %v", res.StatusCode)
	}
	var i []map[string]interface{}
	body, err := ioutil.ReadAll(res.Body)
	util.FailOnError(err,"failed to read the response body")
	err = json.Unmarshal(body, &i)
	util.FailOnError(err,"failed to unmarshal the json body")
	url := fmt.Sprint(i[0]["url"])
	res.Body.Close()
	println(2, ": Url:", url)


	x,err := s.ChannelMessageSendComplex(m.ChannelID,&discordgo.MessageSend{
		Embed: &discordgo.MessageEmbed{
			Image: &discordgo.MessageEmbedImage{
				URL: url,
			},
		},
	})
	
	
	//x,err := s.ChannelMessageSend(m.ChannelID,url)
	util.FailOnError(err,"Failed to send dog pic to channel")
	println(x.Type)


}
