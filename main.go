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

	input, err := r.ReadString('\n')

	return strings.TrimSpace(input), err
}

// create a reader - for the terminal
func createCocktail() cocktail {
	reader := bufio.NewReader(os.Stdin)

	name, _ := getInput("Cocktail name: ", reader)

	c := oneCocktail(name)
	fmt.Println("Created a cocktail: ", c.name)

	return c
}

func promptOptions(c cocktail) {
	reader := bufio.NewReader(os.Stdin)

	opt, _ := getInput("Choose option (a - add an ingredient, i - add instructions, s - save bill)", reader)
	fmt.Println(opt)
}

func main() {
	myCocktail := createCocktail()
	promptOptions(myCocktail)

	fmt.Println(myCocktail)
}
