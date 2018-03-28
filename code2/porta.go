package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	fmt.Println("Iniciando...")
	http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
}
