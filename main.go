package main

import (
	"flag"
	"fmt"
)

func main() {
	//wrirte your api key here
	apiKey := ""
	//write your login and password to data base here
	login := "login"
	password := "password"

	db := connectionOfDataBase(login, password)

	ingredientsFlag := flag.String("ingredients", "", "List of ingredients separated by commas")
	numberOfRecipesFlag := flag.Int("numberOfRecipes", 1, "Max number of recipes to display")

	flag.Parse()

	if *ingredientsFlag == "" {
		fmt.Println("Write list ingredients using flag --ingredients=")
		fmt.Println("Write max number of recipes to display using flag --numberOfRecipes=")
		fmt.Println("Example: ./main --ingredients=apple,banana --numberOfRecipes=2")
		return
	}

	exist := isHistoryExists(db, *ingredientsFlag, *numberOfRecipesFlag)
	if exist == false {
		findByIngredients(apiKey, *ingredientsFlag, *numberOfRecipesFlag, db)
	} else {
		// TODO: get data from database
		getDetailsRecipe(db, *ingredientsFlag, *numberOfRecipesFlag)
	}
	defer db.Close()

}
