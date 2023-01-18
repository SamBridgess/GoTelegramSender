package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	tgbotapi "gopkg.in/telegram-bot-api.v4"
	"io"
	"net/http"
	"os"
	"strconv"
)

type Config struct {
	TelegramBotToken string
}

var bot *tgbotapi.BotAPI
var usersMap map[string]int64

func main() {
	usersMap = map[string]int64{}
	readUsers()
	//for key, value := range mp {
	//	fmt.Println(key, " ", value)
	//}

	file, _ := os.Open("config.json")
	decoder := json.NewDecoder(file)
	configuration := Config{}
	decoder.Decode(&configuration)

	bot, _ = tgbotapi.NewBotAPI(configuration.TelegramBotToken)
	bot.Debug = true

	go botListener(bot)

	port := ":8080"
	http.Handle("/", http.FileServer(http.Dir("./public")))
	http.HandleFunc("/request", frontRequestHandler)

	fmt.Println("Listening port", port)
	http.ListenAndServe(port, nil)
}
func botListener(bot *tgbotapi.BotAPI) {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 120
	updates, _ := bot.GetUpdatesChan(u)
	for update := range updates {
		if update.Message == nil {
			continue
		}
		username := update.Message.Chat.UserName
		chatID := update.Message.Chat.ID
		message := update.Message.Text

		usersMap[username] = chatID
		writeUsers()
		fmt.Println(username, chatID, message)
	}
}
func sendTGMessage(bot *tgbotapi.BotAPI, username string, text string) bool {

	delete(usersMap, "1")
	msg := tgbotapi.NewMessage(usersMap[username], text)
	msg.ParseMode = "Markdown"
	_, err := bot.Send(msg)
	if err != nil {
		return false
	}
	return true
}
func frontRequestHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		r.ParseMultipartForm(0)
		username := r.FormValue("username")
		message := r.FormValue("message")
		fmt.Println("Front request :")
		fmt.Println("username: ", username)
		fmt.Println("message: ", message)
		resp := make(map[string]string)

		if sendTGMessage(bot, username, message) {
			resp["status"] = "success"
			resp["message"] = "Message sent successfully"
			jsonResp, _ := json.Marshal(resp)
			w.Write(jsonResp)
		} else {
			resp["status"] = "error"
			resp["message"] = "There is no such user in database"
			jsonResp, _ := json.Marshal(resp)
			w.Write(jsonResp)
		}
	default:
		fmt.Fprintf(w, "only post")
	}
}
func readUsers() {
	f, _ := os.Open("users.csv")
	r := csv.NewReader(f)

	for {
		line, err := r.Read()

		if err == io.EOF {
			fmt.Println("User base loaded")
			break
		}
		if err != nil {
			fmt.Println("An error occurred while reading csv")
		}

		usersMap[line[0]], _ = strconv.ParseInt(line[1], 10, 64)
	}
}
func writeUsers() {
	f, _ := os.Create("users.csv")
	w := csv.NewWriter(f)

	for key, value := range usersMap {
		err := w.Write([]string{fmt.Sprintf("%v", key), fmt.Sprintf("%v", value)})
		if err != nil {
			fmt.Println("An error occurred while reading csv")
		}
	}
	w.Flush()
}
