package main

import (
	_ "errors"
	"fmt"
	"os"
	"unicode"

	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
)

func IsRusByUnicode(str string) bool {
	for _, r := range str {
		if unicode.Is(unicode.Cyrillic, r) {
			return true
		}
	}
	return false
}

func brickification(message string) (brickWord string, err error) {
	var checkRussian bool = IsRusByUnicode(message)
	if checkRussian == true {
		return message, nil
	} else {
		return "Пожалуйста, на русском!", nil
	}
}

func main() {
	botToken := os.Getenv("TOKEN")
	bot, err := telego.NewBot(botToken, telego.WithDefaultDebugLogger())
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	updates, _ := bot.UpdatesViaLongPolling(nil)

	defer bot.StopLongPolling()

	for update := range updates {
		brickMessage, _ := brickification(update.Message.Text)

		if update.Message != nil {
			chatID := update.Message.Chat.ID
			sentMessage, _ := bot.SendMessage(
				tu.Message(
					tu.ID(chatID),
					brickMessage,
				),
			)
			fmt.Printf("Sent Message: %v\n", sentMessage)

		}

	}
}
