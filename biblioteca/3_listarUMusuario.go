package biblioteca

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/AnaJuliaNX/desafio2/dados"
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
	//Caso na busca não encontre nenhum usuário exibo isso, usuários não cadastrados tem id 0
	if usuariobuscado.ID == 0 {
		TratandoErros(w, "Usuário não encontrado", 404)
	} else {

		//Deixo dentro de um dsata tudo que será exibido, forma padronizada
		var dadosUsuario dados.DataUsuario
		dadosUsuario.Data = usuariobuscado
		//Altero os dados recebidos de struct para json, facilitado para outras linguagens
		erro = json.NewEncoder(w).Encode(dadosUsuario)
		if erro != nil {
			TratandoErros(w, "Erro ao converter para json", 422)
			return
		}
	}

}
