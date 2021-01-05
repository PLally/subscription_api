package destinations

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/plally/subscription_api/subscription"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func init() {
	subscription.SetDestinationHandler("discord", &DiscordDestinationHandler{})
}

type DiscordDestinationHandler struct{}

func (d *DiscordDestinationHandler) GetType() string {
	return "discord"
}

func (d *DiscordDestinationHandler) Dispatch(id string, item subscription.SubscriptionItem) error {
	session, err := discordgo.New(viper.GetString("discord_authorization"))
	if err != nil {
		return err
	}

	log.Infof("[Discord Destination]: %v, %v (%v)", id, item.Title, item.Url)
	message := discordMessage{
		Content: "",
		Embed: embed{
			Title:       item.Title,
			Description: fmt.Sprintf("Author: %v\n%v", item.Author, item.Description),
			Url:         item.Url,
			Footer: embedFooter{
				Text: item.Type + " : " + item.Tags,
			},
			Image: embedImage{
				Url: item.Image,
			},
		},
	}
	channelUrl := fmt.Sprintf("https://discord.com/api/channels/%v/messages", id)
	_, err = session.Request("POST", channelUrl, message)
	return err
}
