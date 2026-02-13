package handlers

import (
	"context"
	"fmt"
	"log"
	m "reminder/models"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func (h *HandlerModule) Start(ctx context.Context, b *bot.Bot, update *models.Update) {
	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   fmt.Sprintf("Привет, %s я бот Напоминалкин", update.Message.Chat.Username),
	})
	user := m.Users{
		Id:        int(update.Message.From.ID),
		NickName:  update.Message.From.Username,
		FirstName: update.Message.From.FirstName,
		LastName:  update.Message.From.LastName,
	}
	err := h.repo.SaveUsers(user)
	if err != nil {
		log.Println("Не удалось сохранить пользователя в бд:", err)
		return
	}
	log.Println("Пользователь сохранен или существует в бд:", user)
}

func (h *HandlerModule) Empty(ctx context.Context, b *bot.Bot, update *models.Update) {
	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   "Cписок моих команд:\n /remind - создать напоминание, формат ввода: 02-01-2006 15:04:05 Сходить за хлебом",
	})
}
