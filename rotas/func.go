package rotas

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func myFuncHome(w http.ResponseWriter, r *http.Request) {
	//código aqui...
}

func myLogin(w http.ResponseWriter, r *http.Request) {
	//código aqui...
}

//myContent imprime o conteudo recebido da url pelo paramentro id
func myContent(w http.ResponseWriter, r *http.Request) {
	var params = mux.Vars(r)
	id := params["id"]
	fmt.Println("O valor do id é: ", id)
}
