package biblioteca

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/AnaJuliaNX/desafio2/dados"
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

	//Trato o erro caso a busca me dê um livro não cadastrado, livros não cadastrados apresentam id 0
	if livroencontrado.ID == 0 {
		TratandoErros(w, "Livro não cadastrado", 404)
		return
	}
	//Coloco todos os dados dentro do data para que fique padronizado
	var dadosLivro dados.DataLivros
	dadosLivro.Data = livroencontrado

	//Transformo em struct para que fique legivel
	erro = json.NewEncoder(w).Encode(dadosLivro)
	if erro != nil {
		TratandoErros(w, "Erro ao converter para json", 422)
		return
	}
}
