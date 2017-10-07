package main

import (
	"fmt"
	//	"log"
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
	Attributes  []string
}

func main() {
	fmt.Println("test")

	// vars
	var rec1 Recipe

	// TO-DO open all files in folder

	// open generic filename
	filename := "recipes/PennePastaAndBeefBolognese.yaml"
	source, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	fmt.Println(len(source))

	// unmarshal data
	err = yaml.Unmarshal(source, &rec1)
	if err != nil {
		//log.Fatalf("error: %v", err)
		panic(err)
	}

	// print confirmation file
	fmt.Printf("--- Recipe:\n%v\n\n", rec1)
}
