package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	registerRoutes()
	fmt.Println("Server running on :8080")
	fmt.Println("hello github!")
	log.Fatal(http.ListenAndServe(":8080", nil))
}