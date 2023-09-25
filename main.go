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

	//Rota para cadastrar um livro novo
	router.HandleFunc("/livros", biblioteca.AdiconarUmLivro).Methods(http.MethodPost)
	//Rota para listar todos os livros já cadastrados
	router.HandleFunc("/livros", biblioteca.ListarOsLivros).Methods(http.MethodGet)
	//Rota para listar um livro, especificado pelo ID
	router.HandleFunc("/livros/{id}", biblioteca.ListarUMLivro).Methods(http.MethodGet)
	//Rota para atualizar um livro, especificado pelo ID
	router.HandleFunc("/livros/{id}", biblioteca.AtualizarUMLivro).Methods(http.MethodPut)
	//Rota para deletar um livro, especificado pelo ID
	router.HandleFunc("/livros/{id}", biblioteca.DeletarUMLivro).Methods(http.MethodDelete)

	//Rota para cadastrar um usuário novo
	router.HandleFunc("/usuarios", biblioteca.AdicionarUsuario).Methods(http.MethodPost)
	//Rota para listar todos os usuários já cadastrados
	router.HandleFunc("/usuarios", biblioteca.ListarOsUsuarios).Methods(http.MethodGet)
	//Rota para listar um usuário, especificado pelo ID
	router.HandleFunc("/usuarios/{id}", biblioteca.ListarUMUsuario).Methods(http.MethodGet)
	//Rota para atualizar um usuário, especificado pelo ID
	router.HandleFunc("/usuarios/{id}", biblioteca.AtualizarUMUsuario).Methods(http.MethodPut)
	//Rota para deletar um usuário, especificado pelo ID
	router.HandleFunc("/usuarios/{id}", biblioteca.DeletarUmUsuario).Methods(http.MethodDelete)

	//Rota para emprestar um livro, especificando o ID do usuário e do livro
	router.HandleFunc("/usuario/{usuario_id}/emprestar/livro/{livro_id}", biblioteca.Emprestando).Methods(http.MethodPut)
	//Rota para devolver um livro, especificando o ID do usuário e do livro
	router.HandleFunc("/usuario/{usuario_id}/devolver/livro/{livro_id}", biblioteca.Devolvendo).Methods(http.MethodPut)

	//Apenas um aviso de que o servidor está funcionando certo
	fmt.Println("Executando na porta 2000")
	//Especificando a porta
	log.Fatal(http.ListenAndServe(":2000", router))
}
