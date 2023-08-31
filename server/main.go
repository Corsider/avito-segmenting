package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"log"
)

var DB *sql.DB

func main() {
	fmt.Println("Connecting to Postgres...")
	DB = Connect()
	defer DB.Close()

	err := DB.Ping()
	if err != nil {
		log.Fatal(err)
	}

	InitDB(DB)

	r := gin.Default()
	InitRouters(r)

	err = r.Run("0.0.0.0:8080")
}
