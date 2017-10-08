package main

import (
	"fmt"
	"html/template"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

var (
	recipes []Recipe
	excludes []Recipe
	database []Recipe
	filters []func(Recipe) bool
	random   *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))
)

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func containsRecipe(r []Recipe, e Recipe) bool {
	for _, a := range r {
		if a.Title == e.Title {
			return true
		}
	}
	return false
}

func recipeByTitle(title string) (int, Recipe) {
    var retI int
    var retRec Recipe
    for i, r := range recipes {
        if r.Title == title {
            retRec = r
            retI = i
            break
        }
    }
    return retI, retRec
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
	// read in form
	days, _ := strconv.Atoi(r.FormValue("days"))
	breakfast := r.FormValue("breakfast") == "on"
	lunch := r.FormValue("lunch") == "on"
	dinner := r.FormValue("dinner") == "on"
	dessert := r.FormValue("dessert") == "on"

	// sum number of recipes
	numRecipes := 0
	if breakfast {
		numRecipes += days
	}
	if lunch {
		numRecipes += days
	}
	if dinner {
		numRecipes += days
	}
	if dessert {
		numRecipes += days
	}
	// limit the size to that of the DB
	if len(database) < numRecipes {
		numRecipes = len(database)
	}
	recipes = make([]Recipe, 0, 0)
	excludes = make([]Recipe, 0, 0)


	// breakfast
	for i := 0; (i < days) && breakfast; i++ {
		// get random recipe
		mealfilter := append(filters, func(rec Recipe) bool {
			return contains(rec.Attributes, "breakfast")
		})
		recipe := getRecipe(mealfilter)
		recipes = append(recipes, recipe)
		excludes = append(excludes, recipe)
	}
	// lunch
	for i := 0; (i < days) && lunch; i++ {
		// get random recipe
		mealfilter := append(filters, func(rec Recipe) bool {
			return contains(rec.Attributes, "lunch")
		})
		recipe := getRecipe(mealfilter)
		recipes = append(recipes, recipe)
		excludes = append(excludes, recipe)
	}
	// dinner
	for i := 0; (i < days) && dinner; i++ {
		// get random recipe
		mealfilter := append(filters, func(rec Recipe) bool {
			return contains(rec.Attributes, "dinner")
		})
		recipe := getRecipe(mealfilter)
		recipes = append(recipes, recipe)
		excludes = append(excludes, recipe)
	}
	// dessert
	for i := 0; (i < days) && dessert; i++ {
		// get random recipe
		mealfilter := append(filters, func(rec Recipe) bool {
			return contains(rec.Attributes, "dessert")
		})
		recipe := getRecipe(mealfilter)
		recipes = append(recipes, recipe)
		excludes = append(excludes, recipe)
	}

	t, _ := template.ParseFiles("recipes.html")
	t.Execute(w, recipes)
}

func replaceHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Println("replaced!")
	title := r.FormValue("recipe")
    fmt.Println(title)

    i, rec := recipeByTitle(title)
    fmt.Println(rec)
    fmt.Println(i)

    excludes = append(excludes, rec)

	// breakfast
	if contains(rec.Attributes, "breakfast") {
		// get random recipe
		mealfilter := append(filters, func(rec Recipe) bool {
			return contains(rec.Attributes, "breakfast")
		})
		recipe := getRecipe(mealfilter)
        recipes[i] = recipe
		excludes = append(excludes, recipe)
	}
	// lunch
	if contains(rec.Attributes, "lunch") {
		// get random recipe
		mealfilter := append(filters, func(rec Recipe) bool {
			return contains(rec.Attributes, "lunch")
		})
		recipe := getRecipe(mealfilter)
        recipes[i] = recipe
		excludes = append(excludes, recipe)
	}
	// dinner
	if contains(rec.Attributes, "dinner") {
		// get random recipe
		mealfilter := append(filters, func(rec Recipe) bool {
			return contains(rec.Attributes, "dinner")
		})
		recipe := getRecipe(mealfilter)
        recipes[i] = recipe
		excludes = append(excludes, recipe)
	}
	// dessert
	if contains(rec.Attributes, "dessert") {
		// get random recipe
		mealfilter := append(filters, func(rec Recipe) bool {
			return contains(rec.Attributes, "dessert")
		})
		recipe := getRecipe(mealfilter)
        recipes[i] = recipe
		excludes = append(excludes, recipe)
	}

	t, _ := template.ParseFiles("recipes.html")
	t.Execute(w, recipes)
}

func groceriesHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("groceries.html")

	ingredients := make(map[string]Ingredient)
	var names []string

	for _, r := range recipes {
		for _, i := range r.Ingredients {
			if !contains(names, i.Name) {
				names = append(names, i.Name)
				ingredients[i.Name] = i
			} else {
				ingr := ingredients[i.Name]
				ingr.Quantity += i.Quantity
				ingredients[i.Name] = ingr
			}
		}
	}

	t.Execute(w, ingredients)
}

func main() {
	fmt.Println("this is a test")

	database = readRecipes(recipeDirectory)
    filters = append(filters, func(rec Recipe) bool {
		return !containsRecipe(excludes, rec)
    })

	http.HandleFunc("/", handler)
	http.HandleFunc("/recipes", recipesHandler)
	http.HandleFunc("/replace", replaceHandler)
	http.HandleFunc("/groceries", groceriesHandler)
	http.ListenAndServe(":8080", nil)
}
