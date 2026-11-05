package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type CryptoResponse struct {
	Bitcoin struct {
		USD float64 `json:"usd"`
	} `json:"bitcoin"`
}

func main() {
	// Put your secret token from BotFather between the quotes below
	bot, err := tgbotapi.NewBotAPI("8986914406:AAHo1ukEuIAiyuL6b_GAY76GoAcMdvVtlcE")
	if err != nil {
		log.Panic("Failed to initialize bot: ", err)
	}

	bot.Debug = true
	fmt.Printf("Authorized on account %s\n", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		chatID := update.Message.Chat.ID
		userText := update.Message.Text
		var replyText string

		switch {
		case strings.HasPrefix(userText, "/start"):
			replyText = "Hello! I am your Go multi-utility bot. Try typing /crypto or /ping!"

		case strings.HasPrefix(userText, "/ping"):
			replyText = "Pong! 🏓 Your bot is fully functional."

		case strings.HasPrefix(userText, "/crypto"):
			price, err := getBitcoinPrice()
			if err != nil {
				replyText = "Sorry, I couldn't fetch the crypto price right now."
			} else {
				replyText = fmt.Sprintf("🪙 The current price of Bitcoin is $%.2f USD.", price)
			}
		case strings.HasPrefix(userText, "/coinflip"):
			if update.Message.MessageID%2 == 0 {
				replyText = "🪙 Heads!"
			} else {
				replyText = "🪙 Tails!"
			}

		default:
			replyText = "I don't recognize that command. Try /crypto or /ping!"
		}

		msg := tgbotapi.NewMessage(chatID, replyText)
		_, err = bot.Send(msg)
		if err != nil {
			log.Println("Error sending message: ", err)
		}
	}
}

func getBitcoinPrice() (float64, error) {
	resp, err := http.Get("https://coingecko.com")
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	var data CryptoResponse
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return 0, err
	}

	return data.Bitcoin.USD, nil
}
