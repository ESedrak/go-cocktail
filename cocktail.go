package main

import (
	"fmt"
	"os"
)

type cocktail struct {
	name         string
	ingredients  map[string]uint64
	instructions string
}

// make new cocktails
func oneCocktail(name string) cocktail {
	c := cocktail{
		name:         name,
		ingredients:  map[string]uint64{},
		instructions: "",
	}

	return c
}

// Format the cocktail
func (c *cocktail) format() string {
	fs := "Cocktail list \n"

	// list ingredients
	for key, value := range c.ingredients {
		fs += fmt.Sprintf("Ingredient: %-20v - Amount(mls): %v \n", key, value)
	}

	// add instructions
	fs += fmt.Sprintf("Instructions: %v\n", c.instructions)

	return fs
}

// add ingredients
func (c *cocktail) addIngredients(name string, amount uint64) {
	c.ingredients[name] = amount
}

// add instructions
func (c *cocktail) addIntructions(instructions string) {
	c.instructions += instructions + "\n"
}

// save cocktail - need to be saved in byte slice format
func (c *cocktail) save() {
	data := []byte(c.format())

	err := os.WriteFile("cocktails/"+c.name+".txt", data, 0644)

	if err != nil {
		panic(err)
	}

}
