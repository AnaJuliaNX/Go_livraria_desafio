package biblioteca

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Essa função serve para deletar do meu banco de dados um livro previamente cadastrado
func DeletarUMLivro(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)

	//Converto o parametro de string para int
	ID, erro := strconv.ParseUint(parametros["id"], 10, 32)
	if erro != nil {
		TratandoErros(w, "Erro ao converter o parametro para inteiro", 422)
	}

	//Executo o comando que faz a conexão com o banco (mais informações no arquivo "comandosBancoErro")
	db, erro := ConectandoNoBanco()
	if erro != nil {
		TratandoErros(w, "Erro ao se conectar com o banco de dados", 422)
		return
	}
	defer db.Close()

	//Crio o statement que vai excluir o livro especificado pelo Id
	statement, erro := db.Prepare("delete from livro_cadastrado where id = ?")
	if erro != nil {
		TratandoErros(w, "Erro ao criar o statement", 422)
		return
	}
	defer statement.Close()

	//Executo o statement e excluo o livro
	if _, erro := statement.Exec(ID); erro != nil {
		TratandoErros(w, "Erro ao executar o statement e deletar o livro", 422)
		return
	}

	//Se não houver nenhum erro durante a execução do código exibo essa mensagem no final
	TratandoErros(w, "Livro deletado com sucesso", 200)
	return
}
