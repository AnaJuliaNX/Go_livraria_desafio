package biblioteca

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/AnaJuliaNX/desafio2/banco"
	"github.com/AnaJuliaNX/desafio2/dados"
)

// Função com finalidade de cadastrar um livro novo no banco de dados
func AdiconarUmLivro(w http.ResponseWriter, r *http.Request) {

	//Faço a leitura de todo o conteúdo do corpo, ou seja, os dados escritos pelo funcionário da biblioteca
	corpoDaRequisicao, erro := io.ReadAll(r.Body)
	if erro != nil {
		TratandoErros(w, "Erro ao ler o conteúdo do corpo", 422)
		return
	}

	//Essa função faz com que eu converta de json para struct, ou seja, volto para os padrões da linguagem de go
	var livro dados.Livro
	if erro = json.Unmarshal(corpoDaRequisicao, &livro); erro != nil {
		TratandoErros(w, "Erro ao converter de json para struct", 422)
		return
	}

	//Executo a função que vai fazer a conexão com o banco (mais informações no arquivo "comandosBancoErro")
	db, erro := banco.ConectarNoBanco()
	if erro != nil {
		TratandoErros(w, "Erro ao se conectar com o banco de dados", 422)
		return
	}
	defer db.Close()

	//Prepraro e digo onde vou salvar os dados no banco
	statement, erro := db.Prepare("Insert into livro_cadastrado(titulo, autor, estoque) values (?, ?, ?)")
	if erro != nil {
		TratandoErros(w, "Erro ao criar o statement", 422)
		return
	}
	defer statement.Close()

	//Executo a solicitação feita acima para salvar os dados do novo livro cadastrado
	inserir, erro := statement.Exec(livro.Titulo, livro.Autor, livro.Estoque)
	if erro != nil {
		TratandoErros(w, "Erro ao executar o statment", 422)
		return
	}

	//Vai retorna um novo Id onde salvei esse livro e por meio dele que vou pode pesquisar o livro depois
	_, erro = inserir.LastInsertId()
	if erro != nil {
		TratandoErros(w, "Erro ao obter o ID inserido", 422)
		return
	}

	//Se nao houve nenhum erro durante a execução do código exibo essa mensagem no final
	TratandoErros(w, "Livro adicionado com sucesso", 200)
	return
}
