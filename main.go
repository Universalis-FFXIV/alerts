package main

import (
	"flag"
	"log"
	"os"
	"strconv"

	"github.com/Universalis-FFXIV/alerts/service/discord"
	"github.com/Universalis-FFXIV/alerts/service/email"
	"github.com/gofiber/fiber/v2"
)

type GenericResponse struct {
	Message string `json:"message"`
}

func main() {
	port := flag.Uint64("p", 7584, "service binding port")
	flag.Parse()

	_, err := discord.New(os.Getenv("UNIVERSALIS_ALERTS_DISCORD_TOKEN"))
	if err != nil {
		log.Fatalln(err)
	}

	_ = email.New("", os.Getenv("UNIVERSALIS_MAILGUN_KEY"))

	app := fiber.New()

	app.Listen(":" + strconv.FormatUint(*port, 10))
}
