package main

import (
	"crypto/md5"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() // Анализирует переданные параметры url, затем анализирует пакет ответа для тела POST (тела запроса)
	// внимание: без вызова метода ParseForm последующие данные не будут получены
	fmt.Println("sayhelloName")
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
	fmt.Println("method:", r.Method) // получаем метод запроса
	fmt.Println("login")
	if r.Method == "GET" {
		curtime := time.Now().Unix()
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(curtime, 12))
		token := fmt.Sprintf("%x", h.Sum(nil))

		t, _ := template.ParseFiles("login.html")
		t.Execute(w, token)
	} else {
		// запрос данных о входе
		r.ParseForm()
		token := r.Form.Get("token")
		if token != "" {
			// проверяем валидность токена
		} else {
			// если нет токена, возвращаем ошибку
		}
		fmt.Println("email length:", len(r.Form["email"][0]))
		fmt.Println("email:", template.HTMLEscapeString(r.Form.Get("email"))) // печатаем на стороне сервера
		fmt.Println("password:", template.HTMLEscapeString(r.Form.Get("password")))
		template.HTMLEscape(w, []byte(r.Form.Get("email"))) // отвечаем клиенту
	}
}

// обработка закачки
func upload(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Метод:", r.Method)
	fmt.Println("upload")
	if r.Method == "GET" {
		crutime := time.Now().Unix()
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(crutime, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))

		t, _ := template.ParseFiles("upload.html")
		t.Execute(w, token)
	} else {
		r.ParseMultipartForm(32 << 20)
		file, handler, err := r.FormFile("uploadfile")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()
		fmt.Fprintf(w, "%v", handler.Header)
		f, err := os.OpenFile("/test/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer f.Close()
		io.Copy(f, file)
	}
}

func main() {
	http.HandleFunc("/", sayhelloName) // устанавливаем правила маршрутизатора
	http.HandleFunc("/login", login)
	http.HandleFunc("/upload", upload)
	err := http.ListenAndServe(":9090", nil) // устанавливаем порт для прослушивания
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
