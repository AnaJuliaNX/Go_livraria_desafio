package biblioteca

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Essa função serve para deletar do meu banco de dados um usuário previamente cadastrado
func DeletarUmUsuario(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)

	ID, erro := strconv.ParseUint(parametros["id"], 10, 32)
	if erro != nil {
		TratandoErros(w, "Erro ao converter o parametro para inteiro", 422)
		return
	}

	//Executo a função que vai fazer a conexão com o banco (mais informações no arquivo "comandosBancoErro")
	db, erro := ConectandoNoBanco()
	if erro != nil {
		TratandoErros(w, "Erro ao se conectar com o banco", 422)
		return
	}
	defer db.Close()

	//Crio os statement que vai excluir um usuário especificado pelo ID
	statement, erro := db.Prepare("delete from usuario where id = ?")
	if erro != nil {
		TratandoErros(w, "Erro ao criar o statemen", 422)
		return
	}
	defer statement.Close()

	//Executo o statement que vai deletar o usuário e seus dados do banco
	if _, erro := statement.Exec(ID); erro != nil {
		TratandoErros(w, "Erro ao executar o statement e deletar o usuário", 422)
		return
	}

	//Se não houve nenhum erro durante a execução do código exibo essa mensagem no final
	TratandoErros(w, "Usuário deletado com sucesso", 200)
	return
}
