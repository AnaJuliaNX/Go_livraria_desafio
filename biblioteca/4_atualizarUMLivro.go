package biblioteca

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/AnaJuliaNX/desafio2/banco"
	"github.com/AnaJuliaNX/desafio2/dados"
	"github.com/gorilla/mux"
)

func AtualizarUMLivro(w http.ResponseWriter, r *http.Request) {
	parametro := mux.Vars(r)

	ID, erro := strconv.ParseUint(parametro["id"], 10, 32)
	if erro != nil {
		TratandoErros(w, "Erro ao converter o parametro para inteiro", 422)
		return
	}

	corpoRequisicao, erro := io.ReadAll(r.Body)
	if erro != nil {
		TratandoErros(w, "Erro ao ler o conteúdo do corpo", 422)
		return
	}

	var livro dados.Livro
	erro = json.Unmarshal(corpoRequisicao, &livro)
	if erro != nil {
		TratandoErros(w, "Erro ao converter de json para struct", 422)
		return
	}

	db, erro := banco.ConectarNoBanco()
	if erro != nil {
		TratandoErros(w, "Erro ao se conectar com o banco de dados", 422)
		return
	}
	defer db.Close()

	statement, erro := db.Prepare("Update livro_cadastrado set titulo = ?, autor = ? where id = ?")
	if erro != nil {
		TratandoErros(w, "Erro ao criar o statement", 422)
		return
	}
	defer statement.Close()

	if _, erro := statement.Exec(livro.Titulo, livro.Autor, ID); erro != nil {
		fmt.Println(erro, 1)
		TratandoErros(w, "Erro ao atualizar o livro", 422)
		return
	}

	TratandoErros(w, "Livro atualizado com sucesso", 200)
	return
}
