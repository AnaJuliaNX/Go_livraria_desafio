package biblioteca

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Essa função busca pelo ID de um livro especifico previamente salvo solicitado pelo funcionário
// e exibe os dados dele na tela
func ListarUMLivro(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)

	//Converto o meu parametro de string para int
	ID, erro := strconv.ParseUint(parametros["id"], 10, 32)
	if erro != nil {
		fmt.Println(erro, 1)
		TratandoErros(w, "Erro ao converter o parametro para int", 422)
		return
	}

	//Função de buscar um livro pelo ID selecionado (mais detalhes sobre no aquivo "comandosParaLivros")
	livroencontrado, erro := BuscandoUMLivro(int(ID))
	if erro != nil {
		fmt.Println(erro, 1)
		TratandoErros(w, "Erro ao converter o parametro para int", 422)
		return
	}

	//Transformo os dados recebidos em struct para json, assim facilita o entendimento em outras linguagens
	erro = json.NewEncoder(w).Encode(livroencontrado)
	if erro != nil {
		TratandoErros(w, "Erro ao converter para json", 422)
		return
	}

	//Se não houver nenhum erro durante toda a execução do código exibo essa mensagem no final
	TratandoErros(w, "Livro buscado com sucesso", 200)
	return
}
