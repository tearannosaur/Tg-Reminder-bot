package handlers

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func (h *HandlerModule) CreateReminder(ctx context.Context, b *bot.Bot, update *models.Update) {
	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   "Устанавливаем напоминание...",
	})

	layout := "02-01-2006 15:04:05"
	text := update.Message.Text
	textSplit := strings.Split(text, " ")
	if len(textSplit) < 3 {
		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: update.Message.Chat.ID,
			Text:   "Не удалось установить напоминание..Проверьте ввода формат даты и времени.",
		})
		return
	}

	dirtyDate := textSplit[1] + " " + textSplit[2]
	date, err := time.ParseInLocation(layout, dirtyDate, time.Local)
	if err != nil {
		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: update.Message.Chat.ID,
			Text:   "Не удалось установить напоминание..Проверьте ввода формат даты и времени.",
		})
		return
	}
	if time.Now().After(date) {
		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: update.Message.Chat.ID,
			Text:   "Введите корректную дату и время.",
		})
		return
	}

	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   "Напоминание установленно!",
	})

	reminderText := strings.Join(textSplit[3:], " ")

	go func() {
		for {
			time.Sleep(time.Until(date))
			b.SendMessage(ctx, &bot.SendMessageParams{
				ChatID: update.Message.Chat.ID,
				Text:   fmt.Sprintf("Напоминание:\n%s", reminderText),
			})
			return
		}
	}()
}
