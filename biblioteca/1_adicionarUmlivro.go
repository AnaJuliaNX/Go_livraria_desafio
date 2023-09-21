package biblioteca

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/AnaJuliaNX/desafio2/banco"
	"github.com/AnaJuliaNX/desafio2/dados"
)

func AdiconarUmLivro(w http.ResponseWriter, r *http.Request) {
	corpoDaRequisicao, erro := io.ReadAll(r.Body)
	if erro != nil {
		TratandoErros(w, "Erro ao ler o conteúdo do corpo", 422)
		return
	}

	var livro dados.Livro
	if erro = json.Unmarshal(corpoDaRequisicao, &livro); erro != nil {
		TratandoErros(w, "Erro ao converter de json para struct", 422)
		return
	}

	db, erro := banco.ConectarNoBanco()
	if erro != nil {
		TratandoErros(w, "Erro ao se conectar com o banco de dados", 422)
		return
	}
	defer db.Close()

	statement, erro := db.Prepare("Insert into livro_cadastrado(titulo, autor, estoque) values (?, ?, ?)")
	if erro != nil {
		TratandoErros(w, "Erro ao criar o statement", 422)
		return
	}
	defer statement.Close()

	inserir, erro := statement.Exec(livro.Titulo, livro.Autor, livro.Estoque)
	if erro != nil {
		TratandoErros(w, "Erro ao executar o statment", 422)
		return
	}

	_, erro = inserir.LastInsertId()
	if erro != nil {
		TratandoErros(w, "Erro ao obter o ID inserido", 422)
		return
	}

	TratandoErros(w, "Livro adicionado com sucesso", 200)
	return
}
