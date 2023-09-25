package biblioteca

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/AnaJuliaNX/desafio2/banco"
	"github.com/AnaJuliaNX/desafio2/dados"
	"github.com/gorilla/mux"
)

// Uso essa função para atualizar ou alterar os dados de um livro previamente cadastrado
// Caso precise, por exemplo, alterar o nome dele ou corrigir um erro de digitação na hora que foi cadastrado
func AtualizarUMLivro(w http.ResponseWriter, r *http.Request) {
	parametro := mux.Vars(r)

	//Altero o parametro de string para int
	ID, erro := strconv.ParseUint(parametro["id"], 10, 32)
	if erro != nil {
		TratandoErros(w, "Erro ao converter o parametro para inteiro", 422)
		return
	}

	//Leio todo o corpo digitado
	corpoRequisicao, erro := io.ReadAll(r.Body)
	if erro != nil {
		TratandoErros(w, "Erro ao ler o conteúdo do corpo", 422)
		return
	}

	//Crio uma variavel apontando para o pacote dados onde tem a struct livro
	//Altero os dados de lá de struct para json, facilita para outras linguagens
	var livro dados.Livro
	erro = json.Unmarshal(corpoRequisicao, &livro)
	if erro != nil {
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

	//Crio um statement especificando a tabela do banco e quais dados dela vou alterar
	statement, erro := db.Prepare("Update livro_cadastrado set titulo = ?, autor = ? where id = ?")
	if erro != nil {
		TratandoErros(w, "Erro ao criar o statement", 422)
		return
	}
	defer statement.Close()

	//Executo o statement e faço a atualização/alteração dos dados
	if _, erro := statement.Exec(livro.Titulo, livro.Autor, ID); erro != nil {
		fmt.Println(erro, 1)
		TratandoErros(w, "Erro ao atualizar o livro", 422)
		return
	}

	//Se não houver nenhum erro durante todo o código exibo essa mensagem no final
	TratandoErros(w, "Livro atualizado com sucesso", 200)
	return
}
