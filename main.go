package main

import (
	"fmt"
	"log"
	"reminder/utils"
)

func main() {
	db, err := utils.DbConnection()
	if err != nil {
		log.Println("Не удалось подключиться к базе данных:", err)
		return
	}
	fmt.Println("Успешное подключение к базе данных")
	_ = db

}
