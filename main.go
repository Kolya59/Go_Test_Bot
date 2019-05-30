package main

import (
	"github.com/Syfaro/telegram-bot-api"
	"log"
)

func main() {
	bot, err := tgbotapi.NewBotAPI("804292110:AAGnUrMivE__LB98hv7uUTNG7ATnpkgubDY")
		// Проверка наличия ошибок при подключении
		if err != nil {
			log.Panic(err)
		}

		bot.Debug = true
		log.Printf("Authorizaed on account %s", bot.Self.UserName)

	// Инициализация канала, который будет сожержать обновления от API
	var ucfg tgbotapi.UpdateConfig = tgbotapi.NewUpdate(0)
		ucfg.Timeout = 60
		upd, _ := bot.GetUpdatesChan(ucfg)

	// Читаем обновления из канала
	for {
		select {
		case update:= <-upd:
				// Пользователь, написавший боту
				UserName := update.Message.From.UserName

				// ID чата
				ChatID := update.Message.Chat.ID

				Text := update.Message.Text

				log.Printf("[%s] %d %s", UserName, ChatID, Text)

				// Ответ
				reply := Text

				// Создаем сообщение
				msg := tgbotapi.NewMessage(ChatID, reply)

				// Отправляем сообщение
				bot.Send(msg)

		}
	}

}
