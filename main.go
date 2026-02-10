package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"reminder/utils"
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

}
