package main

type cocktail struct {
	name         string
	ingredients  map[string]int64
	instructions string
}

// make new cocktails
func newCocktail(name string) cocktail {
	c := cocktail{
		name:         name,
		ingredients:  map[string]int64{},
		instructions: "",
	}

	return c
}
