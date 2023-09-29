package biblioteca

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Essa função busca pelo ID de um livro especifico previamente salvo solicitado pelo funcionário
// e exibe os dados dele na tela
func ListarUMLivro(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)

	//Converto o meu parametro de string para int
	ID, erro := strconv.ParseInt(parametros["id"], 10, 32)
	if erro != nil {
		TratandoErros(w, "Erro ao converter o parametro para int", 422)
		return
	}

	//Função de buscar um livro pelo ID selecionado (mais detalhes sobre no aquivo "comandosParaLivros")
	livroencontrado, erro := BuscandoUMLivro(int(ID))
	if erro != nil {
		TratandoErros(w, "Erro ao converter o parametro para int", 422)
		return
	}

	if livroencontrado.ID == 0 {
		TratandoErros(w, "Livro não cadastrado", 404)
		return
	}

	erro = json.NewEncoder(w).Encode(livroencontrado)
	if erro != nil {
		TratandoErros(w, "Erro ao converter para json", 422)
		return
	}
}
