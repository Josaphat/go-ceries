package main

import (
    "fmt"
    "net/http"
    "html/template"
)

func handler(w http.ResponseWriter, r *http.Request) {
    t, _ := template.ParseFiles("client.html")
    t.Execute(w, "foo")
}

func fooHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Println("foo!")
}

func main() {
    http.HandleFunc("/", handler)
    http.HandleFunc("/foo", fooHandler)
    http.ListenAndServe(":8080", nil)
}
