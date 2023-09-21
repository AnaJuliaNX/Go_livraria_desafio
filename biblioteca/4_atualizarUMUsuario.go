package biblioteca

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/AnaJuliaNX/desafio2/dados"
	"github.com/gorilla/mux"
)

func AtualizarUMUsuario(w http.ResponseWriter, r *http.Request) {
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

	var usuario dados.Usuario
	erro = json.Unmarshal(corpoRequisicao, &usuario)
	if erro != nil {
		TratandoErros(w, "Erro ao converter de json para struct", 422)
		return
	}

	db, erro := ConectandoNoBanco()
	if erro != nil {
		TratandoErros(w, "Erro ao se conectar com o banco de dados", 422)
		return
	}
	defer db.Close()

	statement, erro := db.Prepare("Update usuario set nome = ? where id = ?")
	if erro != nil {
		TratandoErros(w, "Erro ao criar o statement", 422)
		return
	}
	defer statement.Close()

	if _, erro := statement.Exec(usuario.Nome, ID); erro != nil {
		TratandoErros(w, "Erro ao atualizar o usuário", 422)
		return
	}

	TratandoErros(w, "Usuário atualizado com sucesso", 200)
}
