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
	recipes  []Recipe
	database []Recipe
	random   *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))
)

type DayPlan struct {
	Date time.Time
	Recipes []Recipe
}

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
	rawDate:= r.FormValue("startdate")
	date, _ := time.Parse("2006-01-02", rawDate)
	
	
	// sum number of recipes
	numRecipes := 0
	numMeals := 0
	if breakfast {
		numRecipes += days
		numMeals++
	}
	if lunch {
		numRecipes += days
		numMeals++
	}
	if dinner {
		numRecipes += days
		numMeals++
	}
	if dessert {
		numRecipes += days
		numMeals++
	}
	dayPlans := make([]DayPlan, days)
	
	// limit the size to that of the DB
	if len(database) < numRecipes {
		numRecipes = len(database)
	}
	recipes = make([]Recipe, 0, 0)

	// filter out repeats
	var filters []func(Recipe) bool
	filters = append(filters, func(rec Recipe) bool {
		return !containsRecipe(recipes, rec)
	})

	// breakfast
	for i := 0; i < days; i++ {
		dayPlans[i].Date = date
		date = date.Add(time.Duration(24) * time.Hour)
		if breakfast {
			// get random recipe
			mealfilter := append(filters, func(rec Recipe) bool {
				return contains(rec.Attributes, "breakfast")
			})
			recipe := getRecipe(mealfilter)
			recipes = append(recipes, recipe)
			dayPlans[i].Recipes = append(dayPlans[i].Recipes, recipe)
		}

		if lunch {
			// get random recipe
			mealfilter := append(filters, func(rec Recipe) bool {
				return contains(rec.Attributes, "lunch")
			})
			recipe := getRecipe(mealfilter)
			recipes = append(recipes, recipe)
			dayPlans[i].Recipes = append(dayPlans[i].Recipes, recipe)
		}

		if dinner {
			// get random recipe
			mealfilter := append(filters, func(rec Recipe) bool {
				return contains(rec.Attributes, "dinner")
			})
			recipe := getRecipe(mealfilter)
			recipes = append(recipes, recipe)
			dayPlans[i].Recipes = append(dayPlans[i].Recipes, recipe)
		}

		if dessert {
			// get random recipe
			mealfilter := append(filters, func(rec Recipe) bool {
				return contains(rec.Attributes, "dessert")
			})
			recipe := getRecipe(mealfilter)
			recipes = append(recipes, recipe)
			dayPlans[i].Recipes = append(dayPlans[i].Recipes, recipe)
		}
	}

	t, _ := template.ParseFiles("recipes.html")
	t.Execute(w, dayPlans)
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

	http.HandleFunc("/", handler)
	http.HandleFunc("/recipes", recipesHandler)
	http.HandleFunc("/groceries", groceriesHandler)
	http.ListenAndServe(":8080", nil)
}
