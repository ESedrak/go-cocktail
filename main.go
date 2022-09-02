package main

import (
	"fmt"
)

func main() {
	newCocktail := oneCocktail("Margarita")

	newCocktail.addIngredients("gin", 30)

	fmt.Println(newCocktail.format())
}
