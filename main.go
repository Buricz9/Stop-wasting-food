package main

import (
	"flag"
	"fmt"
)

func main() {

	//wrirte your api key here
	apiKey := "apiKey"

	ingredientsFlag := flag.String("ingredients", "", "List of ingredients separated by commas")
	numberOfRecipesFlag := flag.Int("numberOfRecipes", 1, "Max number of recipes to display")

	flag.Parse()

	if *ingredientsFlag == "" {
		fmt.Println("Write list ingredients using flag --ingredients=")
		fmt.Println("Write max number of recipes to display using flag --numberOfRecipes=")
		fmt.Println("Example: ./main --ingredients=apple,banana --numberOfRecipes=2")
		return
	}

	findByIngredients(apiKey, *ingredientsFlag, *numberOfRecipesFlag)
}
