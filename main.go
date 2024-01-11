package main

import (
	"database/sql"
	"encoding/json"
	"fmt"

	_ "github.com/lib/pq"
)

type Student struct {
	ID            int
	Name          string
	Country       string
	Age           int
	Course        int
	Faculty       string
	IsScholarship bool
}

func main() {

	reqString := []byte(`
	{
		"NAME": "Mehmet Hasimoglu",
		"COUNTRY": "Turkiye",
		"AGE": 23,
		"COURSE": 4,
		"FACULTY": "Marine knowledges",
		"IS_SCHOLARSHIP": true
	}`)

	connection := "user=admin password=123 dbname=homework sslmode=disable"
	db, err := sql.Open("postgres", connection)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	var student Student

	if err := json.Unmarshal(reqString, &student); err != nil {
		panic(err)
	}

	var respSt Student

	rowSt := db.QueryRow(`INSERT INTO students(NAME, COUNTRY, AGE, COURSE, FACULTY, IS_SCHOLARSHIP) VALUES($1, $2, $3, $4, $5, $6) RETURNING ID, NAME, COUNTRY, AGE, COURSE, FACULTY, IS_SCHOLARSHIP;`,
		student.Name,
		student.Country,
		student.Age,
		student.Course,
		student.Faculty,
		student.IsScholarship)

	if err := rowSt.Scan(&respSt.ID, &respSt.Name, &respSt.Country, &respSt.Age, &respSt.Course, &respSt.Faculty, &respSt.IsScholarship); err != nil {
		panic(err)
	}

	fmt.Println("Student is successfuly registred with ID ", respSt.ID)

	

}
