package main

import (
	"fmt"
	"log"

	cone "github.com/DiegoSantosWS/encurtador-url/connection"
	"github.com/DiegoSantosWS/encurtador-url/helpers"
)

func main() {
	fmt.Println("Iniciando servidor...")
	err := cone.Connection()
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	fmt.Println("Seja bem vindo!")
	// for example, server receive token string in request header.
	tokenstring := helpers.TokenGenerate()

}
