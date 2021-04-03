package main

import (
	"encoding/json"
	"flag"
	"log"
	"os"
	"strconv"

	"github.com/Universalis-FFXIV/alerts/service/discord"
	"github.com/Universalis-FFXIV/alerts/service/email"
	"github.com/gofiber/fiber/v2"
)

type DiscordNotificationInfo struct {
	UserID string `json:"uid"`
	Body   string `json:"body"`
}

type EmailNotificationInfo struct {
	Address string `json:"address"`
	Body    string `json:"body"`
}

func main() {
	port := flag.Uint64("p", 7584, "service binding port")
	flag.Parse()

	// Instantiate services
	discordClient, err := discord.New(os.Getenv("UNIVERSALIS_ALERTS_DISCORD_TOKEN"))
	if err != nil {
		log.Fatalln(err)
	}

	emailClient := email.New("", os.Getenv("UNIVERSALIS_MAILGUN_KEY"))

	// Configure router
	app := fiber.New()

	discordRouter := app.Group("/discord")
	discordRouter.Post("/send", func(ctx *fiber.Ctx) error {
		dni := &DiscordNotificationInfo{}
		if err := json.Unmarshal(ctx.Body(), dni); err != nil {
			log.Println(err)
			return err
		}

		return discordClient.SendNotification(dni.UserID, dni.Body)
	})

	emailRouter := app.Group("/email")
	emailRouter.Post("/send", func(ctx *fiber.Ctx) error {
		eni := &EmailNotificationInfo{}
		if err := json.Unmarshal(ctx.Body(), eni); err != nil {
			log.Println(err)
			return err
		}

		return emailClient.SendNotification(eni.Address, eni.Body)
	})

	app.Listen(":" + strconv.FormatUint(*port, 10))
}
