package biblioteca

import (
	"encoding/json"
	"net/http"
)

// Essa função serve para listar todos os livros previamente cadastrados que estão salvos no banco de dados
func ListarOsLivros(w http.ResponseWriter, r *http.Request) {

	//Chamo a minha função que vai realizar toda a parte de buscar os livros
	livros, erro := BuscandoOSLivros()
	if erro != nil {
		TratandoErros(w, erro.Error(), 422)
		return
	}

	//Transformo os dados recebidos no formato struct para json, facilita o entendimento em outras linguagens
	erro = json.NewEncoder(w).Encode(livros)
	if erro != nil {
		TratandoErros(w, "Erro ao converter para json", 422)
		return
	}

	//Se não houve nenhum erro durante a execução do código exibo essa mensagem no final
	TratandoErros(w, "Livros buscascados com sucesso", 200)
	return
}
