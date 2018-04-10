package rotas

import (
	"net/http"

	"github.com/gorilla/mux"
)

func Routers() {
	r := mux.NewRouter()

	fs := http.FileServer(http.Dir("assets/"))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))

	r.HandleFunc("/", myFuncHome).Methods("GET") //PRESSUPONDO QUE ESSA VAI SER A
	//ROTA QUE RENDERISA A TEMPLATE DE LOGIN

	r.HandleFunc("/login", myLogin).Methods("POST") //RECEBE UM POST COM OS DADOS DO FORM

	r.HandleFunc("/content/{id}", myContent)

	http.ListenAndServe(":12345", r)
}
