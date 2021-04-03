package discord

import (
	"github.com/Universalis-FFXIV/alerts/service/common"
	"github.com/bwmarrin/discordgo"
)

type DiscordService struct {
	client *discordgo.Session
}

func New(token string) (common.NotificationService, error) {
	client, err := discordgo.New("Bot " + token)
	if err != nil {
		return nil, err
	}

	d := &DiscordService{client: client}
	return d, nil
}

func (d *DiscordService) SendNotification(uid string, text string) error {
	_, err := d.client.ChannelMessageSend(uid, text)
	return err
}
