package main

import (
	"log"
	"myapp/internal/api"
	"net/http"
)

const port = "8080"

func main() {
	// dbURL := os.Getenv("DB_URL")
	// //db, err := sql.Open("postgres", dbURL)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	mux := http.NewServeMux()
	server := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}
	api.RegisterRoutes(mux)
	println("Server running on port:", port)
	log.Fatal(server.ListenAndServe())
}
