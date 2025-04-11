package main

import (
	"awesomeProject3/handler"
	"log"
	"net/http"
)

func main() {

	handler.Router()

	log.Println("http://localhost:10000")
	if err := http.ListenAndServe(":10000", nil); err != nil {
		log.Fatal(err)
	}
}
