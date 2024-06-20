package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

//http.Handler - обработчики формируют ответы на запросы.
//http.ServeMux - мультиплексор, muxer или router, все это одно и то же
//его функции заключаются в анализе пути(сопоставление указанного
//пути со списком зарегистрированных шаблонов)
//и выборе соответствующего обработчика.
//Для поддержки сложной маршрутизации лучше воспользоваться сторонними библиотеками
//gorilla/mux и go-chi/chi, эти библиотеки позволяют реализовать промежуточную обработку
//без всяких проблем
//http.Server - задача сервера слушать входящие соединения и перенаправлять запросы
//по правильному обработчику.
//Промежуточное ПО (middleware) - это по которое находится между сервером и клиентов,
//обрабатывая запросы и ответы. Это способ инкапсулировать общие функции, которые используются
//на нескольких маршрутах, проверка авторизации, шифрование, сжатие, обработка заголвков

func main() {
	router := http.NewServeMux()

	router.HandleFunc("/hello-gophers1", helloHandler)
	router.HandleFunc("/hello-gophers2", jsonHandler)
	router.HandleFunc("/template", indexHandle)
	//router.Handle("/auth", RequireAuthentication)

	wrappedRouter := loggingMiddleware(router)

	// http-сервер
	http.ListenAndServe(":80", wrappedRouter)

	// https-сервер, но для его использования следует получить сертификат и закрытый ключ
	//http.ListenAndServeTLS(":443", "cert.pem", "key.pem", nil)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello ")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Cophers!"))
}

func jsonHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "Hello, Gophers!"}`))
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Receiver request", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
		log.Println("Finished handling request")
	})
}

func TestHelloHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/hello-gophers1", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(helloHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := "Hello, Gophers!"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

var header = `
{{define "header"}}
<h1>Header</h1>
<hr>
{{end}}`

var tmpl = `
{{template "header" .}}
<h1>
<a href="{{.Link}}">{{.Text}}</a>
</h1>`

type page struct {
	Text string
	Link string
}

func indexHandle(w http.ResponseWriter, r *http.Request) {
	t, _ := template.New("page").Parse(header)
	t.Parse(tmpl)
	p := page{
		Text: "github",
		Link: "https://github.com/",
	}
	t.Execute(w, p)
}
