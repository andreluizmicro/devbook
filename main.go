package main

import (
	"api/src/config"
	"api/src/routes"
	"fmt"
	"log"
	"net/http"
)

func main() {
	config.Bootstrap()

	fmt.Println(config.ConnectionString)

	fmt.Println("Rodando API")
	r := routes.Generate()

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), r))
}
