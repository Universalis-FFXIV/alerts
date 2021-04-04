package discord

import (
	"bytes"
	_ "embed"
	"text/template"

	"github.com/Universalis-FFXIV/alerts/model"
	"github.com/Universalis-FFXIV/alerts/service/common"
	"github.com/bwmarrin/discordgo"
)

//go:embed embed_template.md
var embedTemplate string

type discordService struct {
	client *discordgo.Session
	et     *template.Template
}

// New creates a new Discord-backed NotificationService.
func New(token string) (common.NotificationService, error) {
	client, err := discordgo.New("Bot " + token)
	if err != nil {
		return nil, err
	}

	et, err := template.New("universalis_discord_template").Parse(embedTemplate)
	if err != nil {
		return nil, err
	}

	d := &discordService{client: client, et: et}

	if err = d.client.Open(); err != nil {
		return nil, err
	}

	return d, nil
}

func (d *discordService) SendNotification(uid string, notification *model.Notification) error {
	user, _ := d.client.UserChannelCreate(uid)

	var description bytes.Buffer
	err := d.et.Execute(&description, notification)
	if err != nil {
		return err
	}

	embed := &discordgo.MessageEmbed{
		URL:         notification.PageURL,
		Title:       "Alert triggered for " + notification.ItemName,
		Description: description.String(),
		Color:       0xBD983A,
	}

	_, err = d.client.ChannelMessageSendEmbed(user.ID, embed)
	return err
}
