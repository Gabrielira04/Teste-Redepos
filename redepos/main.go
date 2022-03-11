package main

import (
	"crud/servidor"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/Usuario", servidor.CriarUsuario).Methods(http.MethodPost)
	router.HandleFunc("/Usuario", servidor.AtualizarUsuario).Methods(http.MethodPut)
	router.HandleFunc("/Usuario/{Cpf}", servidor.ApagarUsuario).Methods(http.MethodDelete)
	router.HandleFunc("/Usuario", servidor.LocalizarUsuarios).Methods(http.MethodGet)
	router.HandleFunc("/Usuario/{Cpf}", servidor.LocalizarUsuario).Methods(http.MethodGet)

	fmt.Println("rodando na porta 5000")
	log.Fatal(http.ListenAndServe(":5000", router))

}
