package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

//B ALUMNO OMIT
type alumne struct {
	legajo           int
	nombre, apellido string
}

//E ALUMNO OMIT

//B CREATEDATABASE OMIT
func createDatabase() {
	//B OPEN OMIT
	db,err := sql.Open("postgres", "user=lucifer host=localhost dbname=postgres sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	//E OPEN OMIT

	_, err = db.Exec(`create database guarani`)
	if err != nil {
		log.Fatal(err)
	}
}

//E CREATEDATABASE OMIT

func main() {
	createDatabase()

	db, err := sql.Open("postgres", "user=lucifer host=localhost dbname=guarani sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	//B CREATE TABLE OMIT
	_, err = db.Exec(`create table alumne (legajo int, nombre text, apellido text)`)
	if err != nil {
		log.Fatal(err)
	}
	//E CREATE TABLE OMIT

	//B INSERT INTO TABLE OMIT
	_, err = db.Exec(`insert into alumne values (1, 'Cristina', 'Kirchner');
	                  insert into alumne values (2, 'Juan Domingo', 'Perón');`)

	if err != nil {
		log.Fatal(err)
	}
	//E INSERT INTO TABLE OMIT

	//B SELECT OMIT
	rows, err := db.Query(`select * from alumne`)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var a alumne
	for rows.Next() {
		if err := rows.Scan(&a.legajo, &a.nombre, &a.apellido); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%v %v %v\n", a.legajo, a.nombre, a.apellido)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
	//E SELECT OMIT
}
