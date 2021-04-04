package discord

import (
	"github.com/Universalis-FFXIV/alerts/model"
	"github.com/Universalis-FFXIV/alerts/service/common"
	"github.com/bwmarrin/discordgo"
)

type discordService struct {
	client *discordgo.Session
}

// New creates a new Discord-backed NotificationService.
func New(token string) (common.NotificationService, error) {
	client, err := discordgo.New("Bot " + token)
	if err != nil {
		return nil, err
	}

	d := &discordService{client: client}

	if err = d.client.Open(); err != nil {
		return nil, err
	}

	return d, nil
}

func (d *discordService) SendNotification(uid string, notification *model.Notification) error {
	user, _ := d.client.UserChannelCreate(uid)

	embed := &discordgo.MessageEmbed{
		URL:         notification.PageURL,
		Title:       "Alert triggered for " + notification.ItemName,
		Description: notification.Body,
		Color:       0xBD983A,
	}

	_, err := d.client.ChannelMessageSendEmbed(user.ID, embed)
	return err
}
