package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Olá Mundo"))
}

func usuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Carregando Usuários"))
}

func main() {

	http.HandleFunc("/home", home)

	http.HandleFunc("/usuarios", usuario)

	log.Fatal(http.ListenAndServe(":5000", nil))

}
