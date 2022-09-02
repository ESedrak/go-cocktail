package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// helper function
func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)

	input, _ := reader.ReadString('\n')
}

// create a reader - for the terminal
func createCocktail() cocktail {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Cocktail name: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	c := oneCocktail(name)
	fmt.Println("Created a cocktail: ", c.name)

	return c
}

func main() {
	myCocktail := createCocktail()

	fmt.Println(myCocktail)
}
