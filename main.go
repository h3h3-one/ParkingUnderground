package main

import (
	"log"
	"net/http"
)

func main() {
	log.Println("Запуск программы.")
	http.ListenAndServe(":8080", nil)
}
