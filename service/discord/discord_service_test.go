package discord_test

import (
	"log"
	"os"
	"testing"

	"github.com/Universalis-FFXIV/alerts/model"
	"github.com/Universalis-FFXIV/alerts/service/discord"
)

// TestSend tests the entire SendNotification method of the DiscordService.
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
			Reasons: []string{
				"Prices_PricePerUnit < 9000",
				"History_Added",
				"History_IsHQ = true",
				"Prices_TownID = 4",
				"Prices_PricePerUnit < 9000",
				"History_Added",
				"History_IsHQ = true",
				"Prices_TownID = 4",
				"Prices_PricePerUnit < 9000",
				"History_Added",
				"History_IsHQ = true",
				"Prices_TownID = 4",
				"Prices_PricePerUnit < 9000",
				"History_Added",
				"History_IsHQ = true",
				"Prices_TownID = 4",
			},
		},
	}

	err = discordClient.SendNotification(ni.TargetUser, &ni.Notification)
	if err != nil {
		t.Fatal(err)
	}
}
