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

	if err = d.client.Open(); err != nil {
		return nil, err
	}

	return d, nil
}

func (d *DiscordService) SendNotification(uid string, text string) error {
	user, _ := d.client.UserChannelCreate(uid)
	_, err := d.client.ChannelMessageSend(user.ID, text)
	return err
}
