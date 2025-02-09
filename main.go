package main

import (
	_ "errors"
	"fmt"
	"os"
	"slices"
	"unicode"

	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
)

func contains(s [10]string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}

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
	if checkRussian == false {
		return "Пожалуйста, на русском!", nil
	}
	var outMessage string = ""
	vowel := []string{"а", "у", "о", "и", "э", "ы", "я", "ю", "е", "ё"}

	for i := 0; i < len(message); i++ {
		if slices.Contains(vowel, string(message[i])) == true {
			outMessage = outMessage + string(message[i]) + "с" + string(message[i])
		} else {
			outMessage = outMessage + string(message[i])
		}
	}
	return outMessage, nil
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
