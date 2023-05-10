package telegram

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

func Run(botUrl string) {
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

	if update.Message.Text == "Назад" {
		botMessage.Text = "Выберите одну из кнопок"
		botMessage.ReplyMarkup = &models.ReplyKeyboardMarkup{
			Keyboard: [][]models.KeyboardButton{
				{
					{Text: "Кнопка 1"},
					{Text: "Кнопка 2"},
				},
				{
					{Text: "Кнопка 3"},
					{Text: "Кнопка 4"},
				},
			},
			ResizeKeyboard:  true,
			OneTimeKeyboard: true,
		}
	} else {
		botMessage.Text = "Выберите подкнопку"
		botMessage.ReplyMarkup = getKeyboardMarkup(update.Message.Text)
	}

	body, err := json.Marshal(botMessage)
	if err != nil {
		return err
	}

	resp, err := http.Post(botUrl+"/sendMessage", "application/json", bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code %d", resp.StatusCode)
	}

	return nil
}

func getKeyboardMarkup(text string) *models.ReplyKeyboardMarkup {
	var buttons [][]models.KeyboardButton

	switch text {
	case "Кнопка 1":
		buttons = [][]models.KeyboardButton{
			{
				{Text: "Подкнопка 1-1"},
				{Text: "Подкнопка 1-2"},
			},
			{
				{Text: "Назад"},
			},
		}
	case "Кнопка 2":
		buttons = [][]models.KeyboardButton{
			{
				{Text: "Подкнопка 2-1"},
				{Text: "Подкнопка 2-2"},
			},
			{
				{Text: "Назад"},
			},
		}
	case "Кнопка 3":
		buttons = [][]models.KeyboardButton{
			{
				{Text: "Подкнопка 3-1"},
				{Text: "Подкнопка 3-2"},
			},
			{
				{Text: "Назад"},
			},
		}
	case "Кнопка 4":
		buttons = [][]models.KeyboardButton{
			{
				{Text: "Подкнопка 4-1"},
				{Text: "Подкнопка 4-2"},
			},
			{
				{Text: "Назад"},
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
