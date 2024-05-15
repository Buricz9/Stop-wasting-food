package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

type Ingredient struct {
	Name string `json:"name"`
}

type Nutrients struct {
	Name   string  `json:"name"`
	Amount float64 `json:"amount"`
}
type Recipe struct {
	Title             string       `json:"title"`
	MissedIngredients []Ingredient `json:"missedIngredients"`
	UsedIngredients   []Ingredient `json:"usedIngredients"`
	Calories          float64
	Carbohydrates     float64
	Protein           float64
}

type Response struct {
	Results []struct {
		Nutrition struct {
			Nutrients []Nutrients `json:"nutrients"`
		} `json:"nutrition"`
	} `json:"results"`
}

func findByIngredients(apiKey string, ingredients string, numb int, db *sql.DB) {

	client := http.DefaultClient
	url := "https://api.spoonacular.com/recipes/findByIngredients?ingredients=" + ingredients + "&number=" + strconv.Itoa(numb) + "&ranking=2&apiKey=" + apiKey
	res, err := client.Get(url)
	if err != nil {
		log.Println("Error sending request: findByIngredients", err)
		os.Exit(1)
	}
	defer res.Body.Close()

	var recipes []Recipe
	err = json.NewDecoder(res.Body).Decode(&recipes)

	if err != nil {
		log.Println("Error decoding response JSONaaaa:", err)
		os.Exit(1)
	}

	for i, recipe := range recipes {
		recipes[i] = findByTitle(apiKey, &recipe)
	}

	id := insertHistory(db, ingredients, numb)
	for _, recipe := range recipes {
		insertDetailsHistory(db, ingredients, numb, &recipe, id)
	}
	for _, recipe := range recipes {
		displayRecipeInfo(recipe)
	}
}

func findByTitle(apiKey string, recip *Recipe) Recipe {
	url := "https://api.spoonacular.com/recipes/complexSearch?titleMatch=" + recip.Title + "&addRecipeNutrition=true&apiKey=" + apiKey

	res, err := http.Get(url)
	if err != nil {
		fmt.Println("Błąd podczas pobierania danych:", err)
		os.Exit(1)
	}
	defer res.Body.Close()

	var nut Response
	if err := json.NewDecoder(res.Body).Decode(&nut); err != nil {
		fmt.Println("Błąd podczas dekodowania danych JSON:", err)
		os.Exit(1)
	}

	recip.Calories = nut.Results[0].Nutrition.Nutrients[0].Amount
	recip.Carbohydrates = nut.Results[0].Nutrition.Nutrients[1].Amount
	recip.Protein = nut.Results[0].Nutrition.Nutrients[2].Amount
	return *recip

}

func displayRecipeInfo(recipe Recipe) {
	fmt.Println("--------------------")
	fmt.Println("Title:", recipe.Title)
	fmt.Println("Missed ingredients:")
	for _, ingredient := range recipe.MissedIngredients {
		fmt.Println("  -", ingredient.Name)
	}
	fmt.Println("Used ingredients:")
	for _, ingredient := range recipe.UsedIngredients {
		fmt.Println("  -", ingredient.Name)
	}
	fmt.Println("Carbohydrates:", recipe.Carbohydrates)
	fmt.Println("Protein:", recipe.Protein)
	fmt.Println("Calories:", recipe.Calories)

}
