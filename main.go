package main

import (
	"api/src/routes"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Rodando API")
	r := routes.Generate()

	log.Fatal(http.ListenAndServe(":9001", r))
}
