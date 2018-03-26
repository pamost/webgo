package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"net/url"
	"strings"
	"text/template"
)

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() // Анализирует переданные параметры url, затем анализирует пакет ответа для тела POST (тела запроса)
	// внимание: без вызова метода ParseForm последующие данные не будут получены
	fmt.Println(r.Form) // печатает информацию на сервере
	fmt.Println("Путь: ", r.URL.Path)
	fmt.Println("Схема: ", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("Ключ: ", k)
		fmt.Println("Значение: ", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Привет Pavel!") // пишет данные в ответ
}

func login(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Метод:", r.Method) // получаем информацию о методе запроса
	if r.Method == "GET" {
		t, _ := template.ParseFiles("login.html")
		t.Execute(w, nil)
	} else {

		fmt.Println(r.FormValue("username")) // Go вызовет r.ParseForm автоматически, вернет первое из значений

		r.ParseForm()
		// логическая часть процесса входа
		//fmt.Println("Пользователь:", r.Form["username"])
		//fmt.Println("Пароль:", r.Form["password"])

		// возвращает версию с заменой потенциально опасных символов на их escape-последовательности.
		fmt.Println("Имя пользователя:", template.HTMLEscapeString(r.Form.Get("username"))) // печатает на стороне сервера
		fmt.Println("Пароль:", template.HTMLEscapeString(r.Form.Get("password")))

		//отправляет в w версию с заменой потенциально опасных символов на их escape-последовательности.
		template.HTMLEscape(w, []byte(r.Form.Get("username"))) // отправляет клиенту

		v := url.Values{}

		v.Set("name", "Ava")
		v.Add("name", "Pavel")
		v.Add("friend", "Jess")
		v.Add("friend", "Sarah")
		v.Add("friend", "Zoe")
		// v.Encode() == "name=Ava&friend=Jess&friend=Sarah&friend=Zoe"
		fmt.Println(v.Get("name"))   // безопаснее использовать Get т.к. возращает пустое значение если пусто
		fmt.Println(v.Get("friend")) // только одно первое значение
		fmt.Println(v["friend"])     // карта значений
		fmt.Println(v["name"])       // карта значений

	}
}

func main() {
	http.HandleFunc("/", sayhelloName) // устанавливаем правила маршрутизатора
	http.HandleFunc("/login", login)
	err := http.ListenAndServe(":9090", nil) // устанавливаем порт для прослушивания
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
