package main

import (
	"crud/servidor"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// CRUD - CREATE - READ - UPDATE - DELETE

	router := mux.NewRouter()
	router.HandleFunc("/usuarios", servidor.CriarUsuario).Methods("POST") //pode usar "POST" ou http.MethodPost

	fmt.Println("Escutando na porta 5000")

	log.Fatal(http.ListenAndServe(":5000", router))

}
