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

// Função com finalidade de cadastrar um livro novo no banco de dados
// w: write, ou seja, o que vou escrever pro usuário
// r: request, ou seja, o que vou receber do usuário/postman, etc
func AdiconarUmLivro(w http.ResponseWriter, r *http.Request) {

	//Faço a leitura de todo o conteúdo do corpo, ou seja, os dados escritos pelo funcionário da biblioteca
	corpoDaRequisicao, erro := io.ReadAll(r.Body)
	if erro != nil {
		TratandoErros(w, "Erro ao ler o conteúdo do corpo", 422)
		return
	}

	//Essa função faz com que eu converta de json para struct, ou seja, volto para os padrões de go
	var body map[string]interface{}
	erro = json.Unmarshal(corpoDaRequisicao, &body)
	if erro != nil {
		fmt.Println(erro)
		TratandoErros(w, "Erro ao converter de json para struct", 422)
		return
	}
	//Verifica se o campo titulo foi preenchido, exibe uma mensagem de erro caso esteja vazio
	if body["titulo"] == nil {
		TratandoErros(w, "O campo titulo é obrigatório", 422)
		return
	}
	//Verifico se o que foi digitado é uma string, se não for exibo a mensagem de erro
	if reflect.TypeOf(body["titulo"]).Kind() != reflect.String {
		TratandoErros(w, "Titulo inválido", 422)
		return
	}
	//Verifica se o titulo foi deixado em branco, se foi exibo a mensagem de erro
	if body["titulo"].(string) == "" {
		TratandoErros(w, "O campo titulo é obrigatório", 422)
		return
	}
	//Verifica se o campo autor foi preenchido, exibe uma mensagem de erro caso esteja vazio
	if body["autor"] == nil || body["autor"] == "" {
		TratandoErros(w, "O campo autor é obrigatório", 422)
		return
	}
	//verifica se o titulo ou o autor atingiram o limite de caracteres, se sim exibe a mensagem de erro
	limiteDeCaracter := 20
	if len(body["titulo"].(string)) > limiteDeCaracter || len(body["autor"].(string)) > limiteDeCaracter {
		TratandoErros(w, "Limite de caracteres atingido", 422)
		return
	}
	//Verifica se o campo estoque foi preenchido, se não exibe a mensagem de erro
	if body["estoque"] == nil {
		TratandoErros(w, "O campo estoque é obrigatório", 422)
		return
	}
	//Verifica se o que estou inserindo no estoque é um número, se não for ele exibe a msg de erro
	if reflect.TypeOf(body["estoque"]).Kind() != reflect.Float64 {
		TratandoErros(w, "No campo de estoque são aceitos apenas números", 422)
		return
	}
	//Verifico se o estoque inserido é zero, se for exibo a mensagem de erro
	if body["estoque"].(float64) == 0 {
		TratandoErros(w, "O estoque não pode ser zero", 422)
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

	var livro dados.Livro
	//Executo a solicitação feita acima para salvar os dados do novo livro cadastrado
	inserir, erro := statement.Exec(livro.Titulo, livro.Autor, livro.Estoque)
	if erro != nil {
		TratandoErros(w, "Erro ao executar o statment", 422)
		return
	}

	//Vai retorna o Id do livro que acabei de adicionar e por meio dele que vou poder pesquisar o livro depois
	_, erro = inserir.LastInsertId()
	if erro != nil {
		TratandoErros(w, "Erro ao obter o ID inserido", 422)
		return
	}

	//Se nao houve nenhum erro durante a execução do código exibo essa mensagem no final
	TratandoErros(w, "Livro adicionado com sucesso", 200)
}
