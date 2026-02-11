package handlers

import (
	"context"
	"fmt"
	"log"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func (h *HandlerModule) Start(ctx context.Context, b *bot.Bot, update *models.Update) {
	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   fmt.Sprintf("Привет, %s я бот Напоминалкин", update.Message.Chat.Username),
	})
	log.Println(update.Message.Chat.ID, update.Message.Chat.FirstName, update.Message.Chat.LastName)
}

func (h *HandlerModule) Empty(ctx context.Context, b *bot.Bot, update *models.Update) {
	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   "Вот список моих команд:\n",
	})
}
