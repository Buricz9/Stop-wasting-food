package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func connectionOfDataBase(login string, password string) *sql.DB {
	db, err := sql.Open("mysql", ""+login+":"+password+"@tcp(localhost:3306)/recipedatabase")
	if err != nil {
		log.Fatal(err)
	}
	// defer db.Close()

	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully connected to MariaDB")
	return db

}

func insertDetailsHistory(db *sql.DB, ingredients string, number int, recipe *Recipe, id int) {

	var pk int
	query := `INSERT INTO recipes (history_id,title,carbs,proteins,calories) VALUES (?,?,?,?,?) RETURNING id`
	err := db.QueryRow(query, id, recipe.Title, recipe.Carbohydrates, recipe.Protein, recipe.Calories).Scan(&pk)
	if err != nil {
		fmt.Println("Error inserting history into DB ----------2")
		log.Fatal(err)
	}

	for _, ingredient := range recipe.MissedIngredients {
		query = `INSERT INTO missingingredients (recipe_id, ingredient_name) VALUES (?, ?)`
		_, err := db.Exec(query, pk, ingredient.Name)
		if err != nil {
			fmt.Println("Error inserting history into DB ----------3")
			log.Fatal(err)
		}
	}

	for _, ingredient := range recipe.UsedIngredients {
		query = `INSERT INTO availableingredients (recipe_id, ingredient_name) VALUES (?, ?)`
		_, err := db.Exec(query, pk, ingredient.Name)
		if err != nil {
			fmt.Println("Error inserting history into DB ----------4")

			log.Fatal(err)
		}
	}
}

func insertHistory(db *sql.DB, ingredients string, number int) int {
	var pk int
	query := `INSERT INTO history_of_inputs (historyOfIngredients, historyOfNumber) VALUES (?, ?) RETURNING id`
	err := db.QueryRow(query, ingredients, number).Scan(&pk)
	if err != nil {
		log.Fatal(err)
	}
	return pk
}
