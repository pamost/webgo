package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()       // разбор аргументов, необходимо вызвать самостоятельно
	fmt.Println(r.Form) // печать данных формы на стороне сервера
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello Pavel!") // отправка данных на клиент
}

func main() {
	http.HandleFunc("/", sayhelloName)       // устанавливаем обработчик
	err := http.ListenAndServe(":9090", nil) // устанавливаем порт, который будем слушать
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
