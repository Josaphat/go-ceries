package main

import (
    "net/http"
    "html/template"
    "math/rand"
    "reflect"
)

type Recipe struct {
    Name string
}

var (
    recipes []Recipe
    database []Recipe
)

func getField(r *Recipe, field string) string {
    ref := reflect.ValueOf(r)
    f := reflect.Indirect(ref).FieldByName(field)
    return string(f.String())
}

func contains(s []string, e string) bool {
    for _, a := range s {
        if a == e {
            return true
        }
    }
    return false
}

/*
 * `excludes` maps a Recipe attribute name to the attribute value
 * e.g. {"Name": "Chicken Parm"} indicates that meals named "Chicken Parm"
 * should be excluded
 */
func getRecipe(excludes map[string][]string) Recipe {

    var recipe Recipe

    // Apply filters
    var subDb []Recipe
    var doAdd bool
    for _, r := range database {
        doAdd = true
        for k, val := range excludes {
            if contains(val, getField(&r, k)) {
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
    database = append(database, Recipe{Name:"Chicken Marsala"})
    database = append(database, Recipe{Name:"Chicken Parm"})
    database = append(database, Recipe{Name:"Chicken Vindaloo"})
    database = append(database, Recipe{Name:"Chicken Wings"})
    database = append(database, Recipe{Name:"Chicken Salad Sandwich"})
    database = append(database, Recipe{Name:"Chicken Fingers"})

    excludes := make(map[string][]string)
    for i:= 0; i < 5; i++ {
        recipe := getRecipe(excludes)
        excludes["Name"] = append(excludes["Name"], recipe.Name)
        recipes = append(recipes, recipe)
    }

    http.HandleFunc("/", handler)
    http.HandleFunc("/recipes", recipesHandler)
    http.ListenAndServe(":8080", nil)
}
