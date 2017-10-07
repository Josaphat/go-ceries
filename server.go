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

func helloHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Println("world!")
}

func main() {
    http.HandleFunc("/", handler)
    http.HandleFunc("/hello", helloHandler)
    http.ListenAndServe(":8080", nil)
}
