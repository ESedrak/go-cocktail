package main

import "fmt"

type cocktail struct {
	name         string
	ingredients  map[string]int64
	instructions string
}

// make new cocktails
func oneCocktail(name string) cocktail {
	c := cocktail{
		name:         name,
		ingredients:  map[string]int64{},
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

	return fs
}

// add ingredients
func (c *cocktail) addIngredients(name string, amount int64) {
	c.ingredients[name] = amount
}

// add instructions
func (c *cocktail) addIntructions(instructions string) {
	c.instructions = instructions
}
