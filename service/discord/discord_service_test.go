package discord_test

import (
	"log"
	"os"
	"testing"

	"github.com/Universalis-FFXIV/alerts/model"
	"github.com/Universalis-FFXIV/alerts/service/discord"
)

func TestSend(t *testing.T) {
	discordClient, err := discord.New(os.Getenv("UNIVERSALIS_ALERTS_DISCORD_TOKEN"))
	if err != nil {
		log.Fatalln(err)
	}

	ni := &model.NotificationInfo{
		TargetUser: os.Getenv("UNIALERTS_TEST_DISCORD_ID"),
		Notification: model.Notification{
			ItemName: "Purpure Bead",
			PageURL:  "https://universalis.app/market/29959",
			Reasons:  []string{"Prices_PricePerUnit < 9000", "History_Added"},
		},
	}

	err = discordClient.SendNotification(ni.TargetUser, &ni.Notification)
	if err != nil {
		t.Fatal(err)
	}
}
