package biblioteca

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"reflect"

	"github.com/AnaJuliaNX/desafio2/banco"
	"github.com/AnaJuliaNX/desafio2/dados"
)

// Função para fazer a verificação dos dados inseridos
func verificaDados() {
	var body map[string]interface{}

}

// Função com finalidade de cadastrar um livro novo no banco de dados
// w: write, ou seja o que vou escrever pro usuário
// r: request, ou seja, o que vou receber do usuário/postman, etc
func AdiconarUmLivro(w http.ResponseWriter, r *http.Request) {

	//Faço a leitura de todo o conteúdo do corpo, ou seja, os dados escritos pelo funcionário da biblioteca
	corpoDaRequisicao, erro := io.ReadAll(r.Body)
	if erro != nil {
		TratandoErros(w, "Erro ao ler o conteúdo do corpo", 422)
		return
	}

	var body map[string]interface{}
	//Verifica se os campos autor e titulo foram preenchidos, exibe uma mensagem de erro caso estejam vazios
	if body["titulo"] == "" || body["autor"] == "" {
		fmt.Println(body["titulo"])
		TratandoErros(w, "Os campos de titulo e/ou autor não podem estar em branco", 422)
		return
	}

	//Limita o tanto de caracteres que podem ser inseridos e exibe uma mensagem de erro caso ultrapasse
	//limiteDeCaracter := 20

	titulo, ok := body["titulo"].(string)
	if !ok {
		fmt.Println(len(titulo))
		TratandoErros(w, "O campo titulo deve ser preenchido com caracteres", 422)
		return
	}
	autor, ok := body["autor"].(string)
	if !ok {
		TratandoErros(w, "No campo autor é permitido apenas caracteres", 422)
		return
	}

	fmt.Println(len(titulo), len(autor))
	// if len(titulo) > limiteDeCaracter || len(autor) > limiteDeCaracter {
	// 	TratandoErros(w, "Limite de caracateres superior a vinte (20) no campo de titulo e/ou autor", 422)
	// 	return
	// }

	//Verifica se o que estou inserindo no estoque é um número, se não for ele exibe a msg de erro
	if reflect.TypeOf(body["estoque"]).Kind() != reflect.Float64 {
		TratandoErros(w, "No campo de estoque são aceitos apenas números", 422)
		return
	}

	//Verifico se o estoque inserido é zero ou nulo, se for exibo a mensagem de erro
	if body["estoque"] == 0 {
		TratandoErros(w, "O estoque não pode ser zero ou nulo", 422)
		return
	}

	//Essa função faz com que eu converta de json para struct, ou seja, volto para os padrões da linguagem de go
	var livro dados.Livro
	erro = json.Unmarshal(corpoDaRequisicao, &body)
	if erro != nil {
		fmt.Println(erro)
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

	//Vai retorna o Id do livro que acabei de adicionar e por meio dele que vou pode pesquisar o livro depois
	_, erro = inserir.LastInsertId()
	if erro != nil {
		TratandoErros(w, "Erro ao obter o ID inserido", 422)
		return
	}

	//Se nao houve nenhum erro durante a execução do código exibo essa mensagem no final
	TratandoErros(w, "Livro adicionado com sucesso", 200)
	return
}
