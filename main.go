package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"reminder/handlers"
	"reminder/repository"
	t "reminder/tg_utils"
	"reminder/utils"
	"time"
)

func main() {
	db, err := utils.DbConnection()
	if err != nil {
		log.Println("Не удалось подключиться к базе данных:", err)
		return
	}
	log.Println("TimeZone:", time.Local, time.Now())
	repo := repository.ModuleInit(db)
	h := handlers.HandlerModuleInit(repo)
	fmt.Println("Успешное подключение к базе данных")
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()
	b, err := utils.TgConnection()
	if err != nil {
		log.Println("Не удалось подключиться к телеграмм боту", err)
		return
	}
	t.Commands(ctx, b, h)

}
