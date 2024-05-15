package main

import (
	"database/sql"
	"fmt"
	"log"

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
	defer db.Close()

	// err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully connected to MariaDB")
	return db
	// history := History{"jabłko,banan", 2}
	// pk := insertHistory(db, history)
	// insertHistory(db, History{"jabłko,banan,jajko", 2})
	// insertHistory(db, History{"jabłko,banan,jajko,chleb", 3})

	// fmt.Println("Inserted history with primary key:", pk)

	// getFromHistory(db, 8)
	// getAllHistory(db)
}

func getAllHistory(db *sql.DB, ingredients string, number int) {
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
	fmt.Println("All history from DB:", data)
}

// func searchDB(db *sql.DB, ingredients string, number int) {
// 	var historyOfIngredients string
// 	var historyOfNumber int
// 	query := `SELECT historyOfIngredients, historyOfNumber FROM history_of_inputs WHERE historyOfIngredients = '` + ingredients + `' AND historyOfNumber = ` + fmt.Sprint(number) + `;`
// 	err := db.QueryRow(query).Scan(&historyOfIngredients, &historyOfNumber)
// 	if err != nil {
// 		if err != sql.ErrNoRows {
// 			log.Fatal("No rows found in data base: ")
// 		}
// 		log.Fatal(err)
// 	}
// 	fmt.Println("History from DB:", historyOfIngredients, historyOfNumber)
// }

func insertHistory(db *sql.DB, history History) int {
	query := `INSERT INTO history_of_inputs (historyOfIngredients, historyOfNumber) 
		VALUES (?, ?) RETURNING id`

	var pk int
	err := db.QueryRow(query, history.Ingridients, history.Number).Scan(&pk)
	if err != nil {
		log.Fatal(err)
	}
	return pk
}
