package biblioteca

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func DeletarUMLivro(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)

	ID, erro := strconv.ParseUint(parametros["id"], 10, 32)
	if erro != nil {
		TratandoErros(w, "Erro ao converter o parametro para inteiro", 422)
	}

	db, erro := ConectandoNoBanco()
	if erro != nil {
		TratandoErros(w, "Erro ao se conectar com o banco de dados", 422)
		return
	}
	defer db.Close()

	statement, erro := db.Prepare("delete from livro_cadastrado where id = ?")
	if erro != nil {
		TratandoErros(w, "Erro ao criar o statement", 422)
		return
	}
	defer statement.Close()

	if _, erro := statement.Exec(ID); erro != nil {
		TratandoErros(w, "Erro ao executar o statement e deletar o livro", 422)
		return
	}

	TratandoErros(w, "Livro deletado com sucesso", 200)
	return
}
