package main

import (
	router "ShorUrl/Router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	r := router.Router()
	fmt.Println("Server is running")
	log.Fatal(http.ListenAndServe(":8999", r))
}
