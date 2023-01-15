package main

import (
	"fmt"
	"net/http"
)

func main() {
	port := ":8080"
	http.Handle("/", http.FileServer(http.Dir("./public")))
	http.HandleFunc("/request", requestHandler)
	fmt.Println("Listening port", port)
	http.ListenAndServe(port, nil)

}

func requestHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "POST":
		r.ParseMultipartForm(0)
		message := r.FormValue("message")
		fmt.Println("Message from Client: ", message)
	default:
		fmt.Fprintf(w, "only post")
	}
}
