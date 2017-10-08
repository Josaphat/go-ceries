package main

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Ingredient struct {
	Name     string
	Quantity float32
	Unit     string
}

type Recipe struct {
	Title       string
	Subtitle    string
	Time        string
	Servings    int
	Calories    int
	Ingredients []Ingredient
	Steps       []string
	Picture     string
	Attributes  []string
	Id          int
}

var recipeDirectory = "./recipes"

/**
 * Remove element from recipes array.
 */
func remove(s []Recipe, i int) []Recipe {
	s[len(s)-1], s[i] = s[i], s[len(s)-1]
	return s[:len(s)-1]
}

/**
 * Reads the recipes from the given dirctory and returns them as an array.
 */
func readRecipes(dir string) []Recipe {
	// get files in folder
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		panic(err)
	}

	// array of recipes
	recipes := make([]Recipe, len(files), len(files))

	for i, file := range files {
		// open file
		source, err := ioutil.ReadFile(dir + "/" + file.Name())
		if err != nil {
			panic(err)
		}

		// unmarshal data
		err = yaml.Unmarshal(source, &recipes[i])
		if err != nil {
			panic(err)
		}
		recipes[i].Id = i

		if len(recipes[i].Picture) == 0 {
			recipes[i].Picture = "/images/noimage.png"
		}
	}

	return recipes
}
