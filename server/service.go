package main

import (
	"database/sql"
	"fmt"
	"github.com/lib/pq"
	"log"
	"strconv"
	"strings"
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

func ArrayToPQ(array pq.Int32Array) string {
	str := ""
	for _, el := range array {
		str += strconv.Itoa(int(el)) + ","
	}
	return str[:len(str)-1]
}

func Int32ArrayToPQ(array []int32) pq.Int32Array {
	var res pq.Int32Array
	for _, el := range array {
		res = append(res, el)
	}
	return res
}

func First(a int, er error) int {
	return a
}

func PQtoArray(str string) []int32 {
	splitted := strings.Split(str, ",")
	var res []int32
	for _, el := range splitted {
		res = append(res, int32(First(strconv.Atoi(el))))
	}
	return res
}
