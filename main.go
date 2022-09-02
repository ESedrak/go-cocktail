package main

import (
	"fmt"
)

func main() {
	newCocktail := oneCocktail("Margarita")

	fmt.Println(newCocktail.format())
}
