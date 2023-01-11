package main

import (
	"log"
	"net/http"
	"html/template"
)

//all errors
func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

//функция-обработчик
func viewHandler(writer http.ResponseWriter, request *http.Request) {
	html, err := template.ParseFiles("view.html") //содержимое html используется для создания нового значени teamplate
	check(err)
	err = html.Execute(writer, nil) //содержимое шаблона записывается в ResponseWriter
	check(err)
}

func main() {
	http.HandleFunc("/guestbook", viewHandler)
	err := http.ListenAndServe("localhost:8080", nil)
	log.Fatal(err)
}
