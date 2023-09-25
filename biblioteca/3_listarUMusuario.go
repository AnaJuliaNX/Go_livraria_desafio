package biblioteca

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Essa função busca pelo ID de um usuário especifico previamente salvo solicitado pelo funcionário
// e exibe os dados dele na tela
func ListarUMUsuario(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)

	//Trasformo o parametro de string para int
	ID, erro := strconv.ParseUint(parametros["id"], 10, 32)
	if erro != nil {
		TratandoErros(w, "Erro ao converter o parametro para inteiro", 422)
		return
	}

	//Executo a função de buscar um usuário pelo ID dele (mais informações no arquivo "comandosOutros")
	usuariobuscado, erro := buscandoUMUsuario(int(ID))
	if erro != nil {
		TratandoErros(w, "Erro ao buscar o ID do usuario", 422)
		return
	}

	//Altero os dados recebidos de struct para json, facilitado para outras linguagens
	erro = json.NewEncoder(w).Encode(usuariobuscado)
	if erro != nil {
		TratandoErros(w, "Erro ao converter para json", 422)
		return
	}

	//Se não houve nenhum erro durante a execução do código exibo essa mensagem no final
	TratandoErros(w, "Usuário buscado com sucesso", 200)
	return
}
