package main

import (
	"database/sql"
	"fmt"
	"log"
	"sort"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

type History struct {
	Ingridients string
	Number      int
}

func isHistoryExists(db *sql.DB, ingredients string, number int) bool {
	allHistory := getAllHistory(db)

	areIngredientsEqual := func(ing1, ing2 string) bool {
		ingSlice1 := strings.Split(ing1, ",")
		ingSlice2 := strings.Split(ing2, ",")
		sort.Strings(ingSlice1)
		sort.Strings(ingSlice2)
		return strings.Join(ingSlice1, ",") == strings.Join(ingSlice2, ",")
	}

	for _, history := range allHistory {
		if areIngredientsEqual(history.Ingridients, ingredients) && history.Number == number {
			fmt.Println("Exist in DB")
			return true
		}
	}
	fmt.Println("Does not exist in DB")
	return false
}

func getAllHistory(db *sql.DB) []History {
	rows, err := db.Query("SELECT historyOfIngredients, historyOfNumber FROM history_of_inputs")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	data := []History{}
	var ing string
	var num int

	for rows.Next() {
		err := rows.Scan(&ing, &num)
		if err != nil {
			log.Fatal(err)
		}
		data = append(data, History{ing, num})
	}

	return data
}

func getDetailsRecipe(db *sql.DB, ingredients string, number int) {
	rows, err := db.Query("SELECT id FROM history_of_inputs where historyOfIngredients = ? AND historyOfNumber = ?", ingredients, number)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var pk int
	if rows.Next() {
		err = rows.Scan(&pk)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		fmt.Println("No history found for the specified ingredients and number.")
		return
	}

	rows, err = db.Query("SELECT id, title, carbs, proteins, calories FROM recipes WHERE history_id = ?", pk)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var title string
		var carbs, proteins, calories float64
		err := rows.Scan(&pk, &title, &carbs, &proteins, &calories)
		if err != nil {
			log.Fatal(err)
		}

		recipe := Recipe{
			Title:         title,
			Carbohydrates: carbs,
			Protein:       proteins,
			Calories:      calories,
		}

		rows, err := db.Query("SELECT ingredient_name FROM missingingredients WHERE recipe_id = ?", pk)
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()

		for rows.Next() {
			var ingredientName string
			err := rows.Scan(&ingredientName)
			if err != nil {
				log.Fatal(err)
			}
			recipe.MissedIngredients = append(recipe.MissedIngredients, Ingredient{Name: ingredientName})
		}

		rows, err = db.Query("SELECT ingredient_name FROM availableingredients WHERE recipe_id = ?", pk)
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()

		for rows.Next() {
			var ingredientName string
			err := rows.Scan(&ingredientName)
			if err != nil {
				log.Fatal(err)
			}
			recipe.UsedIngredients = append(recipe.UsedIngredients, Ingredient{Name: ingredientName})
		}

		displayRecipeInfo(recipe)
	}
}
