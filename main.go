package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "<host>"
	port     = "<port>"
	user     = "<user>"
	password = "<password>"
	dbname   = "<db_name>"
)

type member struct {
	id   int
	name string
	age  int
}

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	QueryMem(db)

	member := member{name: "Carlos", age: 27}
	AddMember(db, member.name, member.age)

}

func QueryMem(db *sql.DB) {
	rows, err := db.Query("SELECT * FROM mem")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		member := member{}
		err = rows.Scan(&member.name, &member.age, &member.id)
		if err != nil {
			panic(err)
		}
		fmt.Println(member.id, member.name, member.age)
	}

	err = rows.Err()
	if err != nil {
		panic(err)
	}
}

func AddMember(db *sql.DB, name string, age int) {
	if CheckIfExists(db, name, age) {
		fmt.Printf("The member has already been inserted, name: %s, age: %d\n", name, age)
		return
	}
	statement, err := db.Prepare("INSERT INTO mem (name, age) VALUES ($1, $2)")
	if err != nil {
		panic(err)
	}
	defer statement.Close()

	_, err = statement.Exec(name, age)
	if err != nil {
		panic(err)
	}

	fmt.Println("successfully inserted")
	QueryMem(db)
}

func CheckIfExists(db *sql.DB, name string, age int) bool {
	var count int
	statement, err := db.Prepare("SELECT COUNT(*) FROM mem WHERE name=$1 AND age=$2")
	if err != nil {
		panic(err)
	}
	defer statement.Close()

	err = statement.QueryRow(name, age).Scan(&count)
	if err != nil {
		panic(err)
	}
	fmt.Println()
	fmt.Println("Check IF Exists, Count: ", count)

	return count > 0
}
