package main

import (
    "net/http"
    "html/template"
    "math/rand"
)

var (
    recipes []Recipe
    database []Recipe
)

func contains(s []string, e string) bool {
    for _, a := range s {
        if a == e {
            return true
        }
    }
    return false
}

func getRecipe(filters []func(Recipe)bool) Recipe {

    var recipe Recipe

    // Apply filters
    var subDb []Recipe
    var doAdd bool
    for _, r := range database {
        doAdd = true
        for _, f := range filters {
            if !f(r) {
                doAdd = false
                break
            }
        }

        if doAdd == true {
            subDb = append(subDb, r)
        }
    }

    // TODO: get recipe here
    recipe = subDb[rand.Intn(len(subDb))]
    // END TODO

    return recipe
}

func handler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("home.html")
	t.Execute(w, "foo")
}

func recipesHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("recipes.html")

    for i, r := range recipes {
        t.Execute(w, map[string]interface{}{"Recipe":r, "Index":i})
    }
}

func main() {
    database = readRecipes(recipeDirectory)

    var filters []func(Recipe)bool
    var names []string
    for i:= 0; i < 5; i++ {
        recipe := getRecipe(filters)
        names = append(names, recipe.Title)
        filters = append(filters, func(r Recipe) bool {
            return !contains(names, r.Title)
        })
        recipes = append(recipes, recipe)
    }

    http.HandleFunc("/", handler)
    http.HandleFunc("/recipes", recipesHandler)
    http.ListenAndServe(":8080", nil)
}
