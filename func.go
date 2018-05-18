package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

func main() {
	key, err := ioutil.ReadFile("secret.str")
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	key = []byte(key)
	tokenstring := generateToke(key)
	token, _ := jwt.Parse(tokenstring, func(token *jwt.Token) (interface{}, error) {
		return []byte(key), nil
	})
	if token.Valid == true {
		fmt.Println(tokenstring)
	}
}
func generateToke(secretKey []byte) string {

	claims := &jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Hour * time.Duration(1)).Unix(),
		Issuer:    "APId2js898ilsje6272g726g072gso",
		IssuedAt:  time.Now().Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims) //Incorporar informações do usuário ao token
	tokenstring, err := token.SignedString(secretKey)          // token -> string. Only server knows this secret (wsitesb).
	if err != nil {
		log.Fatalln(err)
	}
	return tokenstring
}
