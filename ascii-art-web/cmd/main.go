package main

import (
	"ascii-art-web/handlers"
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handlers.HomePage)
	http.HandleFunc("/generate", handlers.GeneratePage)
	fmt.Println("Server is running on :http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
