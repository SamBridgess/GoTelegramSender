package main

import (
	"embed"
	"encoding/csv"
	"encoding/json"
	"fmt"
	tgbotapi "gopkg.in/telegram-bot-api.v4"
	"io"
	"io/fs"
	"net/http"
	"os"
	"path"
	"runtime"
	"strconv"
)

type Config struct {
	TelegramBotToken string
}

//go:embed public
var frontend embed.FS
var configuration = Config{}
var _, filename, _, _ = runtime.Caller(0)
var bot *tgbotapi.BotAPI
var usersMap map[string]int64

func main() {
	port := ":8080"
	usersMap = map[string]int64{}
	readUsers()
	readConfig()

	bot, _ = tgbotapi.NewBotAPI(configuration.TelegramBotToken)
	go botListener(bot)

	stripped, _ := fs.Sub(frontend, "public")
	http.Handle("/", http.FileServer(http.FS(stripped)))
	http.HandleFunc("/request", frontRequestHandler)

	fmt.Println("Listening port", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		fmt.Println("Unable to start server")
	}
}
func botListener(bot *tgbotapi.BotAPI) {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 30
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
	msg := tgbotapi.NewMessage(usersMap[username], text)
	msg.ParseMode = "Markdown"
	_, err := bot.Send(msg)
	return err == nil
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
		} else {
			resp["status"] = "error"
			resp["message"] = "There is no such user in database"
		}
		jsonResp, _ := json.Marshal(resp)
		_, err := w.Write(jsonResp)
		if err != nil {
			fmt.Println("Couldn't send response")
		}
		fmt.Println("Server response: " + resp["message"] + "\n")
	default:
		fmt.Fprintf(w, "Server only supports POST")
	}
}
func readUsers() {
	f, _ := os.Open(path.Join(path.Dir(filename), "users.csv"))
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
	f, _ := os.Create(path.Join(path.Dir(filename), "users.csv"))
	w := csv.NewWriter(f)

	for key, value := range usersMap {
		err := w.Write([]string{fmt.Sprintf("%v", key), fmt.Sprintf("%v", value)})
		if err != nil {
			fmt.Println("An error occurred while writing to csv")
		}
	}
	w.Flush()
	fmt.Println("User info saved")
}
func readConfig() {
	f, _ := os.Open(path.Join(path.Dir(filename), "config.json"))

	decoder := json.NewDecoder(f)
	err := decoder.Decode(&configuration)
	if err != nil {
		fmt.Println("An error occurred while reading config")
	}
	fmt.Println("Config loaded")
}
