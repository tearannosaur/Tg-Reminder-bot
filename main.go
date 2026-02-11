package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"reminder/handlers"
	"reminder/repository"
	"reminder/utils"

	"github.com/go-telegram/bot"
)

func main() {
	db, err := utils.DbConnection()
	if err != nil {
		log.Println("Не удалось подключиться к базе данных:", err)
		return
	}
	fmt.Println("Успешное подключение к базе данных")
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()
	b, err := utils.TgConnection()
	if err != nil {
		log.Println("Не удалось подключиться к телеграмм боту", err)
		return
	}
	repository.ModuleInit(db)
	b.RegisterHandler(bot.HandlerTypeMessageText, "/start", bot.MatchTypeExact, handlers.Start)
	b.Start(ctx)
}
