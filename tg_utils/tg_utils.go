package tg_utils

import (
	"context"
	"reminder/handlers"

	"github.com/go-telegram/bot"
)

func Commands(ctx context.Context, b *bot.Bot, h *handlers.HandlerModule) {
	b.RegisterHandler(bot.HandlerTypeMessageText, "/start", bot.MatchTypeExact, h.Start)
	b.RegisterHandler(bot.HandlerTypeMessageText, "/remind", bot.MatchTypePrefix, h.CreateReminder)
	b.RegisterHandler(bot.HandlerTypeMessageText, "", bot.MatchTypePrefix, h.Empty)
	b.Start(ctx)
}
