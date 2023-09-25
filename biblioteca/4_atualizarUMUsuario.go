package biblioteca

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/AnaJuliaNX/desafio2/dados"
	"github.com/gorilla/mux"
)

// Essa função vai atualizar/alterar os dados do usuário previamente cadastrado
// Caso precise, por exemplo, atualizar o nome dele uso essa função
func AtualizarUMUsuario(w http.ResponseWriter, r *http.Request) {
	parametro := mux.Vars(r)

	//Converto o parametro que recebo como string para int
	ID, erro := strconv.ParseUint(parametro["id"], 10, 32)
	if erro != nil {
		TratandoErros(w, "Erro ao converter o parametro para inteiro", 422)
		return
	}

	//Leio todos os dados digitados
	corpoRequisicao, erro := io.ReadAll(r.Body)
	if erro != nil {
		TratandoErros(w, "Erro ao ler o conteúdo do corpo", 422)
		return
	}

	//Crio uma variavel apontando para o pacote dados onde tem a struct usuário
	//Altero os dados de struct para json, facilitando para outras linguagens
	var usuario dados.Usuario
	erro = json.Unmarshal(corpoRequisicao, &usuario)
	if erro != nil {
		TratandoErros(w, "Erro ao converter de json para struct", 422)
		return
	}

	//Abro a conexão com o banco para pode fazer a alteração dos dados e eles serem salvos
	db, erro := ConectandoNoBanco()
	if erro != nil {
		TratandoErros(w, "Erro ao se conectar com o banco de dados", 422)
		return
	}
	defer db.Close()

	//Crio um statement onde vou especificar quais campos da tabela vou preencher
	statement, erro := db.Prepare("Update usuario set nome = ? where id = ?")
	if erro != nil {
		TratandoErros(w, "Erro ao criar o statement", 422)
		return
	}
	defer statement.Close()

	//Executo o statement e faço a atualização/alteração dos dados do usuário
	if _, erro := statement.Exec(usuario.Nome, ID); erro != nil {
		TratandoErros(w, "Erro ao atualizar o usuário", 422)
		return
	}

	//Se não houver nenhum erro durante a execução do código exibo essa mensagem no final
	TratandoErros(w, "Usuário atualizado com sucesso", 200)
}
