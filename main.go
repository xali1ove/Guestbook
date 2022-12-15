package main

import (
	"log"
	"net/http"
)

// mi proveryaem git
//all errors
func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

//функция-обработчик
func viewHandler(writer http.ResponseWriter, request *http.Request) {
	placeholder := []byte("signature list goes here")
	_, err := writer.Write(placeholder)
	check(err)
}

func main() {
	http.HandleFunc("/guestbook", viewHandler)
	err := http.ListenAndServe("localhost:8080", nil)
	log.Fatal(err)
}
