package destinations

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/plally/subscription_api/subscription"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"net/http"
)

func init() {
	subscription.SetDestinationHandler("discord", &DiscordDestinationHandler{})
}

type DiscordDestinationHandler struct{}

func (d *DiscordDestinationHandler) GetType() string {
	return "discord"
}

func (d *DiscordDestinationHandler) Dispatch(id string, item subscription.SubscriptionItem) error {
	log.Infof("[Discord Destination]: %v, %v (%v)", id, item.Title, item.Url)
	message := discordMessage{
		Content: "",
		Embed: embed{
			Title: "Subscription Item",
			Description: fmt.Sprintf("Author: %v\n%v", item.Author, item.Description),
			Url: item.Url,
			Footer: embedFooter{
				Text: item.Type +" : "+item.Tags,
			},
			Image: embedImage{
				Url: item.Image,
			},
		},
	}
	channelUrl := fmt.Sprintf("https://discord.com/api/channels/%v/messages", id)
	data, err := json.Marshal(message)
	req, err := http.NewRequest("POST", channelUrl, bytes.NewBuffer(data))
	if err != nil { return err }
	req.Header.Set("Authorization", viper.GetString("discord_authorization"))
	req.Header.Set("Content-Type", "application/json")

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		return errors.New("non success status code "+resp.Status)
	}
	return err
}
