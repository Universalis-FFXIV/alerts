package discord

import (
	"github.com/bwmarrin/discordgo"
)

type Discord interface {
	SendNotification(uid uint64, text string) error
}

type DiscordService struct {
	client *discordgo.Session
}

func New(token string) (Discord, error) {
	client, err := discordgo.New("Bot " + token)
	if err != nil {
		return nil, err
	}

	d := &DiscordService{client: client}
	return d, nil
}

func (d *DiscordService) SendNotification(uid uint64, text string) error {
	return nil
}
