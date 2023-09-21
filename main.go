package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/AnaJuliaNX/desafio2/biblioteca"
	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/livros", biblioteca.AdiconarUmLivro).Methods(http.MethodPost)
	router.HandleFunc("/livros", biblioteca.ListarOsLivros).Methods(http.MethodGet)
	router.HandleFunc("/livros/{id}", biblioteca.ListarUMLivro).Methods(http.MethodGet)
	router.HandleFunc("/livros/{id}", biblioteca.AtualizarUMLivro).Methods(http.MethodPut)
	router.HandleFunc("/livros/{id}", biblioteca.DeletarUMLivro).Methods(http.MethodDelete)

	router.HandleFunc("/usuario/{usuario_id}/emprestar/livro/{livro_id}", biblioteca.Emprestando).Methods(http.MethodPut)
	router.HandleFunc("/livros/{id}/devolver", biblioteca.DevolverLivro).Methods(http.MethodPut)

	router.HandleFunc("/usuarios", biblioteca.AdicionarUsuario).Methods(http.MethodPost)
	router.HandleFunc("/usuarios", biblioteca.ListarOsUsuarios).Methods(http.MethodGet)
	router.HandleFunc("/usuarios/{id}", biblioteca.ListarUMUsuario).Methods(http.MethodGet)
	router.HandleFunc("/usuarios/{id}", biblioteca.AtualizarUMUsuario).Methods(http.MethodPut)
	router.HandleFunc("/usuarios/{id}", biblioteca.DeletarUmUsuario).Methods(http.MethodDelete)

	fmt.Println("Executando na porta 2000")
	log.Fatal(http.ListenAndServe(":2000", router))
}
