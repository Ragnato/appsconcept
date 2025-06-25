package main

import (
	"log"
	"net/http"

	"appsconcept/internal/api"
	"appsconcept/internal/repository/mysql"
	"appsconcept/internal/service"
)

func main() {
	mysqlConn, err := mysql.NewMySQL()
	if err != nil {
		log.Fatalf("Error connecting to MySQL: %v", err)
	}

	fizzbuzzRepo := mysql.NewFizzBuzzRepo(mysqlConn.DB)

	fizzbuzzService := service.NewFizzBuzzService(fizzbuzzRepo)

	handler := api.NewHandler(fizzbuzzService)

	http.HandleFunc("/fizzbuzz", handler.FizzBuzz)
	http.HandleFunc("/stats", handler.Stats)

	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
