package utils

import (
	"os"

	"github.com/go-telegram/bot"
	"github.com/joho/godotenv"
)

func TgConnection() (*bot.Bot, error) {
	if err := godotenv.Load(); err != nil {
		return nil, err
	}
	tgKey := os.Getenv("TgBotToket")
	b, err := bot.New(tgKey)
	if err != nil {
		return nil, err
	}
	return b, nil
}
