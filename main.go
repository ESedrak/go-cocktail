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

	switch opt {
	case "a":
		fmt.Println("You chose a")
	case "i":
		fmt.Println("You chose i")
	case "s":
		fmt.Println("You chose s")
	default:
		fmt.Println("You did not choose a valid option")
		promptOptions(c)
	}
}

func main() {
	myCocktail := createCocktail()
	promptOptions(myCocktail)
}
