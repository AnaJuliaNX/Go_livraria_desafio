package biblioteca

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/AnaJuliaNX/desafio2/banco"
	"github.com/AnaJuliaNX/desafio2/dados"
	"github.com/gorilla/mux"
)

func DevolverLivro(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)

	ID, erro := strconv.ParseUint(parametros["id"], 10, 32)
	if erro != nil {
		TratandoErros(w, "Erro ao converter o parametro para inteiro", 422)
		return
	}

	db, erro := banco.ConectarNoBanco()
	if erro != nil {
		TratandoErros(w, "Erro ao se conectar com o banco de dados", 422)
	}

	linhas, erro := db.Query("select * from livro_cadastrado where id = ?", ID)
	if erro != nil {
		TratandoErros(w, "Erro ao buscar o livro", 422)
		return
	}

	var livro dados.Livro
	if linhas.Next() {
		if erro := linhas.Scan(&livro.ID, &livro.Titulo, &livro.Autor, &livro.Estoque); erro != nil {
			TratandoErros(w, "Erro ao escaner o livro", 422)
			return
		}
	}

	corpoRequisicao, erro := io.ReadAll(r.Body)
	if erro != nil {
		TratandoErros(w, "Erro ao ler os dados da requisição", 422)
		return
	}

	var dadosRequest dados.RequestDevolverLivro
	erro = json.Unmarshal(corpoRequisicao, &dadosRequest)
	if erro != nil {
		TratandoErros(w, "Erro ao pegar dados inseridos", 422)
		return
	}

	if dadosRequest.Devolvido <= 0 {
		TratandoErros(w, "Quantidade inserida não pode ser menor do que zero", 422)
		return
	}

	livro.Estoque = livro.Estoque + dadosRequest.Devolvido

	erro = AlterarEstoque(int(ID), livro.Estoque)
	if erro != nil {
		TratandoErros(w, erro.Error(), 422)
		return
	}
}
