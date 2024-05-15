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

func printFormattedHistory(data [][]interface{}) {
	for _, d := range data {
		ingredients := d[0].([]string)
		number := d[1].(int)
		fmt.Printf("[%v] %v\n", ingredients, number)
	}
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
			fmt.Println("History already exists in DB")
			return true
		}
	}
	fmt.Println("History does not exist in DB")
	return false
}

func insertHistory(db *sql.DB, ingredients string, number int) int {

	query := `INSERT INTO history_of_inputs (historyOfIngredients, historyOfNumber) 
		VALUES (?, ?) RETURNING id`

	var pk int
	err := db.QueryRow(query, ingredients, number).Scan(&pk)
	if err != nil {
		log.Fatal(err)
	}
	return pk
}
