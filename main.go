package main

import (
	"flag"
	"log"
	"os"
	"strconv"

	"github.com/bwmarrin/discordgo"
	"github.com/gofiber/fiber/v2"
)

type GenericResponse struct {
	Message string `json:"message"`
}

func main() {
	port := flag.Uint64("p", 7584, "service binding port")
	flag.Parse()

	discord, err := discordgo.New("Bot " + os.Getenv("UNIVERSALIS_ALERTS_DISCORD_TOKEN"))
	if err != nil {
		log.Fatalln(err)
	}

	app := fiber.New()

	app.Get("/", func(ctx *fiber.Ctx) error {
		user, err := discord.User("@me")
		if err != nil {
			log.Fatalln(err)
		}

		res := &GenericResponse{Message: "Logged in as " + user.Username + "#" + user.Discriminator}

		return ctx.JSON(res)
	})

	app.Listen(":" + strconv.FormatUint(*port, 10))
}
