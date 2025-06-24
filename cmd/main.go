package main

import (
	"appsconcept/internal/api"
	"appsconcept/internal/repository/mysql"
	"appsconcept/internal/service"
	"log"
	"net/http"
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
