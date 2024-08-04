package main

import (
	"fmt"
	"log"

	spritedb "github.com/AyeshGK/spritedb/src/spritedb"
)

func main() {
	user1 := map[string]string{
		"name": "joe",
		"city": "beijing",
		"age":  "18",
	}

	user2 := map[string]string{
		"name": "jane",
		"city": "shanghai",
		"age":  "19",
	}

	db, err := spritedb.NewDB()
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.CreateCollection("users")
	if err != nil {
		log.Fatal(err)
	}

	query1 := db.NewQueryBuilder().Collection("users").Insert(user1).Build()

	id1, err := query1.Insert()
	if err != nil {
		// handle error
	}
	fmt.Println("inserting user 1 id: ", id1)

	query2 := db.NewQueryBuilder().Collection("users").Insert(user2).Build()

	id2, err := query2.Insert()
	if err != nil {
		// handle error
	}

	fmt.Println("inserting user 2 id: ", id2)

	fmt.Println("inserting done")

	query3 := db.NewQueryBuilder().Collection("users").Select("age", "name").Build()

	results, err := query3.Select()
	if err != nil {
		// handle error
	}
	fmt.Println("selecting done")

	fmt.Println(results)
	db.Close()
}
