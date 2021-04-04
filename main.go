package main

import (
	"encoding/json"
	"flag"
	"log"
	"os"
	"strconv"

	"github.com/Universalis-FFXIV/alerts/model"
	"github.com/Universalis-FFXIV/alerts/service/common"
	"github.com/Universalis-FFXIV/alerts/service/discord"
	"github.com/gofiber/fiber/v2"
)

func main() {
	port := flag.Uint64("p", 7584, "service binding port")
	flag.Parse()

	// Instantiate services
	discordClient, err := discord.New(os.Getenv("UNIVERSALIS_ALERTS_DISCORD_TOKEN"))
	if err != nil {
		log.Fatalln(err)
	}

	/*
		emailClient, err := email.New("", os.Getenv("UNIVERSALIS_MAILGUN_KEY"))
		if err != nil {
			log.Fatalln(err)
		}
	*/

	services := map[string]common.NotificationService{
		"discord": discordClient,
		//"email":   emailClient,
	}

	// Configure router
	app := fiber.New()

	for serviceName, service := range services {
		router := app.Group("/" + serviceName)
		router.Post("/send", func(ctx *fiber.Ctx) error {
			ni := &model.NotificationInfo{}
			if err := json.Unmarshal(ctx.Body(), ni); err != nil {
				log.Println(err)
				return err
			}

			return service.SendNotification(ni.TargetUser, &ni.Notification)
		})
	}

	app.Listen(":" + strconv.FormatUint(*port, 10))
}
