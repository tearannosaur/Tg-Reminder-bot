package utils

import (
	"log"
	"os"
	"time"

	"github.com/go-telegram/bot"
	"github.com/joho/godotenv"
)

func TgConnection() (*bot.Bot, error) {
	if err := godotenv.Load(); err != nil {
		return nil, err
	}
	tgKey := os.Getenv("TgBotToket")
	var b *bot.Bot
	for {
		time.Sleep(5 * time.Second)
		bot, err := bot.New(tgKey)
		if err == nil {
			b = bot
			break
		}
		log.Println(err)
	}
	return b, nil
}
