package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "user:user@tcp(localhost:13306)/todo_db")
	if err != nil {
		log.Fatal(err)
	}
	// main関数の最後にDBを閉じる
	defer db.Close()

	// データの挿入を行うSQLステートメントを準備
	stmt, err := db.Prepare("INSERT INTO todos (task) VALUES (?)")
	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close()

	tasks := []string{
		"Task 1",
		"Task 2",
		"Task 3",
	}

	// rangeを使用してtasksスライスをループ
	for _, task := range tasks {
		// DBインサート
		_, err := stmt.Exec(task)
		if err != nil {
			log.Fatal(err)
		}
	}

	fmt.Println("Seeder executed successfully")
}
