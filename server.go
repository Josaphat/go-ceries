package main

import (
    "fmt"
    "net/http"
    "html/template"
)

type Recipe struct {
    Name string
}

func getRecipes() []Recipe {
    //r1 := Recipe{Name:"Chicken parm"}
    var recipes []Recipe
    recipes = append(recipes, Recipe{Name:"Chicken Parm"})
    recipes = append(recipes, Recipe{Name:"Chicken Marsala"})
    return recipes
}

func handler(w http.ResponseWriter, r *http.Request) {
    t, _ := template.ParseFiles("home.html")
    t.Execute(w, "foo")
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Println("world!")
}

func recipesHandler(w http.ResponseWriter, r *http.Request) {
    t, _ := template.ParseFiles("recipes.html")

    for _, r := range getRecipes() {
        t.Execute(w, r)
    }
}

func main() {
    http.HandleFunc("/", handler)
    http.HandleFunc("/hello", helloHandler)
    http.HandleFunc("/recipes", recipesHandler)
    http.ListenAndServe(":8080", nil)
}
