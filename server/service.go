package main

import (
	"database/sql"
	"fmt"
	"log"
)

func Connect() *sql.DB {
	psinfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", "postgres", "5432", "postgres", "123", "avito")
	database, err := sql.Open("postgres", psinfo)
	if err != nil {
		log.Fatal(err)
	}
	return database
}

func InitDB(db *sql.DB) {
	db.Exec("DROP TABLE IF EXISTS userss")
	db.Exec("DROP TABLE IF EXISTS segmentlist")
	db.Exec("CREATE TABLE userss (user_id SERIAL PRIMARY KEY, segments INTEGER[])")
	db.Exec("CREATE TABLE segmentlist (segment_id serial primary key , slug varchar(128))")
}

func First(a int, er error) int {
	return a
}
