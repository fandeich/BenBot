package main

import (
	tgbotapi "gopkg.in/telegram-bot-api.v4"
	"math/rand"
	"time"
)

func telegramBot() {
	bot, err := tgbotapi.NewBotAPI("5247616780:AAF4iFBQCMuMLQRwS7gVXSwboPlihXcI3Qg")
	if err != nil {
		panic(err)
	}
	rand.Seed(time.Now().UnixNano())
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		if update.Message.Text != "" {
			switch update.Message.Text {
			case "/start":
				msg := tgbotapi.NewAudioUpload(update.Message.Chat.ID, "voice/Ben.ogg")
				_, err = bot.Send(msg)
				if err != nil {
					panic(err)
				}
			default:
				if rand.Intn(100) < 50 {
					msg := tgbotapi.NewAudioUpload(update.Message.Chat.ID, "voice/No.ogg")
					_, err = bot.Send(msg)
					if err != nil {
						panic(err)
					}
				} else {
					msg := tgbotapi.NewAudioUpload(update.Message.Chat.ID, "voice/Yes.ogg")
					_, err = bot.Send(msg)
					if err != nil {
						panic(err)
					}
				}
			}
		}
	}
}

func main() {
	telegramBot()
}
