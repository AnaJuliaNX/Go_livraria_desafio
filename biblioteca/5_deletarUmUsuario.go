package biblioteca

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func DeletarUmUsuario(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)

	ID, erro := strconv.ParseUint(parametros["id"], 10, 32)
	if erro != nil {
		TratandoErros(w, "Erro ao converter o parametro para inteiro", 422)
		return
	}

	db, erro := ConectandoNoBanco()
	if erro != nil {
		TratandoErros(w, "Erro ao se conectar com o banco", 422)
		return
	}
	defer db.Close()

	statement, erro := db.Prepare("delete from usuario where id = ?")
	if erro != nil {
		TratandoErros(w, "Erro ao criar o statemen", 422)
		return
	}
	defer statement.Close()

	if _, erro := statement.Exec(ID); erro != nil {
		TratandoErros(w, "Erro ao executar o statement e deletar o usuário", 422)
		return
	}

	TratandoErros(w, "Usuário deletado com sucesso", 200)
	return
}
