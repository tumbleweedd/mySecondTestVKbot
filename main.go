package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/tumbleweedd/mySecondTestVKbot/pkg/models"
	"io"
	"log"
	"net/http"
	"strconv"
)

const botToken = "6279047527:AAEtZGFDHKToNXk4tORbWqft-f0Krm-Jswc"
const botApi = "https://api.telegram.org/bot"
const botUrl = botApi + botToken

func main() {
	offset := 0
	for {
		updates, err := getUpdates(botUrl, offset)
		if err != nil {
			log.Println("Err: ", err.Error())
		}
		for _, update := range updates {
			err = sendMessage(botUrl, update)
			offset = update.UpdateId + 1
		}

		fmt.Println(updates)
	}
}

func getUpdates(botUrl string, offset int) ([]models.Update, error) {
	resp, err := http.Get(botUrl + "/getUpdates" + "?offset=" + strconv.Itoa(offset))
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var restResponse models.RestResponse
	err = json.Unmarshal(body, &restResponse)
	if err != nil {
		return nil, err
	}

	return restResponse.Result, nil
}

func sendMessage(botUrl string, update models.Update) error {
	var botMessage models.BotMessage
	botMessage.ChatId = update.Message.Chat.ChatId

	message := models.BotMessage{
		Text:        "Привет!",
		ChatId:      update.Message.Chat.ChatId,
		ReplyMarkup: getKeyboardMarkup(update.Message.Text),
	}

	body, err := json.Marshal(message)
	if err != nil {
		return err
	}

	resp, err := http.Post(botUrl+"/sendMessage", "application/json", bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("error StatusCode %d", resp.StatusCode)
	}

	return nil
}

func getKeyboardMarkup(text string) *models.ReplyKeyboardMarkup {
	var buttons [][]models.KeyboardButton

	switch text {
	case "Кнопка 1":
		buttons = [][]models.KeyboardButton{
			{
				{Text: "Подкнопка раз"},
				{Text: "Подкнопка два"},
			},
		}
	case "Кнопка 2":
		buttons = [][]models.KeyboardButton{
			{
				{Text: "Подкнопка раз"},
				{Text: "Подкнопка два"},
			},
		}
	case "Кнопка 3":
		buttons = [][]models.KeyboardButton{
			{
				{Text: "Подкнопка раз"},
				{Text: "Подкнопка два"},
			},
		}
	case "Кнопка 4":
		buttons = [][]models.KeyboardButton{
			{
				{Text: "Подкнопка раз"},
				{Text: "Подкнопка два"},
			},
		}
	default:
		buttons = [][]models.KeyboardButton{}
	}

	return &models.ReplyKeyboardMarkup{
		Keyboard:        buttons,
		ResizeKeyboard:  true,
		OneTimeKeyboard: true,
	}
}
