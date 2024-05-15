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
	Results           []struct {
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

}

func findByTitle(apiKey string, recip *Recipe) Recipe {
	url := "https://api.spoonacular.com/recipes/complexSearch?titleMatch=" + recip.Title + "&addRecipeNutrition=true&apiKey=" + apiKey

	res, err := http.Get(url)
	if err != nil {
		fmt.Println("Błąd podczas pobierania danych:", err)
		os.Exit(1)
	}
	defer res.Body.Close()

	if err := json.NewDecoder(res.Body).Decode(&recip); err != nil {
		fmt.Println("Błąd podczas dekodowania danych JSON:", err)
		os.Exit(1)
	}

	if len(recip.Results) == 0 {
		fmt.Println("Brak przepisów w odpowiedzi.")
		os.Exit(1)
	}
	return *recip

}

func displayRecipeInfo(recip Recipe) {
	fmt.Println("\n-----------------")

	fmt.Println("\nTytuł przepisu:", recip.Title)

	fmt.Println("\nMissed Ingredients:")
	for _, missedIngredient := range recip.MissedIngredients {
		fmt.Println("  - Name:", missedIngredient.Name)
	}
	fmt.Println("\nUsed Ingredients:")
	for _, usedIngredient := range recip.UsedIngredients {
		fmt.Println("  - Name:", usedIngredient.Name)
	}

	fmt.Println("\nInformacje o składnikach odżywczych:")
	for _, nu := range recip.Results[0].Nutrition.Nutrients {
		if nu.Name == "Calories" || nu.Name == "Carbohydrates" || nu.Name == "Protein" {
			fmt.Printf(" - %s: %.2f\n", nu.Name, nu.Amount)
		}
	}
}
