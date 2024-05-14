package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type Ingredient struct {
	Name string `json:"name"`
}

type Recipe struct {
	Title             string       `json:"title"`
	MissedIngredients []Ingredient `json:"missedIngredients"`
	UsedIngredients   []Ingredient `json:"usedIngredients"`
}

func findByIngredients(apiKey string) {
	url := fmt.Sprintf("https://api.spoonacular.com/recipes/findByIngredients?ingredients=apples,flour,sugar&ranking=2&apiKey=%s", apiKey)
	res, err := http.Get(url)
	if err != nil {
		fmt.Println("Błąd podczas wysyłania zapytania:findByIngredients", err)
		os.Exit(1)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	var recipes []Recipe
	err = json.Unmarshal(body, &recipes)
	if err != nil {
		panic(err)
	}

	for _, recipe := range recipes {
		fmt.Println("Title:", recipe.Title)
		fmt.Println("Missed Ingredients:")
		for _, missedIngredient := range recipe.MissedIngredients {
			fmt.Println("  - Name:", missedIngredient.Name)
		}
		fmt.Println("Used Ingredients:")
		for _, usedIngredient := range recipe.UsedIngredients {
			fmt.Println("  - Name:", usedIngredient.Name)
		}
		fmt.Println()
	}
}

// func findByTitle(apiKey string, title string, recipe Recipe) {
// 	recipeURL := fmt.Sprintf("https://api.spoonacular.com/recipes/complexSearch?titleMatch=%s&addRecipeNutrition=true&apiKey=%s", title, apiKey)
// 	resp, err := http.Get(recipeURL)
// 	if err != nil {
// 		fmt.Println("Błąd podczas wysyłania zapytania:", err)
// 		return
// 	}
// 	defer resp.Body.Close()

// 	body, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		fmt.Println("Błąd odczytu odpowiedzi:", err)
// 		return
// 	}

// 	fmt.Println(string(body))

// }
