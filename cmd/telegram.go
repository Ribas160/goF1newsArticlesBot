package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (a *App) SubmitNewArticles(articles Articles) {
	bot := a.initBot()

	bot.Debug = true

	for i := 0; i < len(articles.Articles); i++ {
		channerlId, err := strconv.Atoi(os.Getenv("CHANNEL_ID"))
		if err != nil {
			a.ErrorLog.Fatal(err)
		}

		pubDate := strings.Replace(articles.Articles[i].PubDate, " +0300", "", -1)
		msg := tgbotapi.NewMessage(int64(channerlId), fmt.Sprintf("%s\n%s", pubDate, articles.Articles[i].Link))

		_, err = bot.Send(msg)
		if err != nil {
			a.ErrorLog.Fatal(err)
		}
	}
}

func (a *App) initBot() *tgbotapi.BotAPI {
	bot, err := tgbotapi.NewBotAPI(os.Getenv("BOT_TOKEN"))
	if err != nil {
		a.ErrorLog.Fatal(err)
	}

	return bot
}
