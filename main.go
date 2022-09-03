package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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

	opt, _ := getInput("Choose option (a - Add an ingredient/amount, i - Add Instructions, s - Save Cocktail)\n", reader)

	switch opt {
	case "a":
		ingredient, _ := getInput("Ingredient: ", reader)
		amount, _ := getInput("Amount(mls): ", reader)

		// To return amount in an int(not the default string)
		a, err := strconv.ParseUint(amount, 0, 64)

		// error will be nil if everything is good
		if err != nil {
			fmt.Println("Amount must be a whole number")
			promptOptions(c)
		}

		c.addIngredients(ingredient, a)

		fmt.Printf("Added: %v %vmls \n", ingredient, amount)
		promptOptions(c)
	case "i":
		instructions, _ := getInput("Instructions to create the cocktail: ", reader)

		c.addIntructions(instructions)

		promptOptions(c)
	case "s":
		c.save()
		fmt.Println("Saved file:", c.name)
	default:
		fmt.Println("You did not choose a valid option")
		promptOptions(c)
	}
}

func main() {
	myCocktail := createCocktail()
	promptOptions(myCocktail)
}
