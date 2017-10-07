package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("home.html")
	t.Execute(w, "foo")
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("world!")
}

func recipesHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("recipes.html")

	for _, r := range readRecipes(recipeDirectory) {
		t.Execute(w, r)
	}
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/recipes", recipesHandler)
	http.ListenAndServe(":8080", nil)
}
