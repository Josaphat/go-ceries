package main

import (
	"html/template"
	"math/rand"
	"net/http"
)

var (
	recipes  []Recipe
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

func getRecipe(filters []func(Recipe) bool) Recipe {

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

	recipe = subDb[rand.Intn(len(subDb))]

	return recipe
}

func handler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("home.html")
	t.Execute(w, "foo")
}

func recipesHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("recipes.html")
	t.Execute(w, recipes)
}

func main() {
	database = readRecipes(recipeDirectory)

	var filters []func(Recipe) bool
	for i := 0; i < 5; i++ {
		recipe := getRecipe(filters)
		filters = append(filters, func(r Recipe) bool {
			return recipe.Title != r.Title
		})
		recipes = append(recipes, recipe)
	}

	http.HandleFunc("/", handler)
	http.HandleFunc("/recipes", recipesHandler)
	http.ListenAndServe(":8080", nil)
}
