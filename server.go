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
    var recipes []Recipe

    // TODO: get recipes here
    recipes = append(recipes, Recipe{Name:"Chicken Parm"})
    recipes = append(recipes, Recipe{Name:"Chicken Marsala"})

    return recipes
}

func handler(w http.ResponseWriter, r *http.Request) {
    t, _ := template.ParseFiles("home.html")
    t.Execute(w, "foo")
}

func recipesHandler(w http.ResponseWriter, r *http.Request) {
    t, _ := template.ParseFiles("recipes.html")

    for i, r := range getRecipes() {
        t.Execute(w, map[string]interface{}{"Recipe":r, "Index":i})
    }
}

func main() {
    http.HandleFunc("/", handler)
    http.HandleFunc("/recipes", recipesHandler)
    http.ListenAndServe(":8080", nil)
}
