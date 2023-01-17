package main

import (
	"log"
	"net/http"
	"html/template"
	"os"
	"bufio"
)

type Guestbook struct {
	SignatureCount int
	Signatures []string
}
//all errors
func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
//чтение файла
func getSrings(fileName string) []string {
	var lines []string
	file, err := os.Open(fileName)  //открытие файла
	if os.IsNotExist(err) {         //при возвращ ошибки указывающей на несуществ файл
		return nil
	}
	check(err)                      //для другой ошибки
	defer file.Close()              //закрытие файла
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	check(scanner.Err())             //сообщить о любых ошибках обработки файла
	return lines
}
//функция-обработчик
func viewHandler(writer http.ResponseWriter, request *http.Request) {
	signatures := getSrings("signatures.txt")
	html, err := template.ParseFiles("view.html") //содержимое html используется для создания нового значени teamplate
	check(err)
	guestbook := Guestbook{
		SignatureCount: len(signatures),
	}
	err = html.Execute(writer, nil) //содержимое шаблона записывается в ResponseWriter
	check(err)
}

func main() {
	http.HandleFunc("/guestbook", viewHandler)
	err := http.ListenAndServe("localhost:8080", nil)
	log.Fatal(err)
}

