package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	fmt.Printf("Serving in http://localhost:%s", port)
	http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
}
