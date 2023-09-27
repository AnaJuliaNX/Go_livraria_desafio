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

	//Executo a função que vai fazer a conexão com o banco de dados (mais informações em "ComandosBancoeErro")
	db, erro := ConectandoNoBanco()
	if erro != nil {
		TratandoErros(w, "Erro ao se conectar no banco de dados ", 422)
		return
	}
	defer db.Close()

	//Selecioso da tabela livro_cadastrado os Id de todos os livros e somo eles
	linhas, erro := db.Query("select count(id) from livro_cadastrado")
	if erro != nil {
		TratandoErros(w, "Erro ao fazer a contagem dos livros", 422)
		return
	}
	defer linhas.Close()

	//Faço o scan do total de livros encontrados pelo id
	var totalDeLivros int64
	if linhas.Next() {
		erro := linhas.Scan(&totalDeLivros)
		if erro != nil {
			TratandoErros(w, "Erro ao fazer o scan", 422)
			return
		}
	}

	//Se o total de livros encontrado for igual a 0 exibo o erro
	if totalDeLivros == 0 {
		TratandoErros(w, "Nenhum livro encontrado", 404)
		return
	}

	//Executo a função Paginacao (mais informações em "comandosBancoeErro")
	response := Paginacao(totalDeLivros, livros)

	//Se não houve nenhum erro durante a execução do código exibo essa mensagem no final
	//Transformo os dados recebidos no formato struct para json, facilita o entendimento em outras linguagens
	erro = json.NewEncoder(w).Encode(response)
	if erro != nil {
		TratandoErros(w, "Erro ao converter para json", 422)
		return
	}
}
