package biblioteca

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

// Essa função serve para listar todos os livros previamente cadastrados que estão salvos no banco de dados
// Eles serão listados de 15 em 15 separados por páginas
func ListarOsLivros(w http.ResponseWriter, r *http.Request) {

	search := r.URL.Query().Get("search")

	//Explicação completa no arquivo "informacoes"
	pagAtual := r.URL.Query().Get("page")
	//Converto o parametro de string pra int
	page, erro := strconv.Atoi(pagAtual)
	if erro != nil {
		page = 1 //se não tiver nenhum parametro ou ele estiver em branco concluo que é a pag 1
	}
	limit := 15 //limito a quantidade de itens que serão buscados e exibidos
	offset := limit * (page - 1)

	//Chamo a minha função que vai realizar toda a parte de buscar os livros até o limite permitido
	livros, erro := BuscandoOSLivros(search, offset)
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
	linhas, erro := db.Query("select count(id) from livro_cadastrado where (titulo like ? or autor like ?)", "%"+search+"%", "%"+search+"%")
	if erro != nil {
		fmt.Println(erro)
		TratandoErros(w, "Erro ao fazer a contagem dos livros", 422)
		return
	}
	defer linhas.Close()

	//Faço o scan do total de livros
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
	response.Meta.Current_page = int64(page) //em cada página ele vai exibir em qual página está

	//Se não houve nenhum erro durante a execução até aqui finalizo transformando os dados recebeidos em json
	erro = json.NewEncoder(w).Encode(response)
	if erro != nil {
		TratandoErros(w, "Erro ao converter para json", 422)
		return
	}

}
